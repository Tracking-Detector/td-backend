package job

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"io"
	"sync"

	"strings"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/extractor"
	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_exporter/converter"

	"github.com/robertkrimen/otto"
	log "github.com/sirupsen/logrus"
)

type IExportJob interface {
	Execute(exporter *model.Exporter, reducer string, dataset string) *model.ExportMetrics
}

type InternalExportJob struct {
	Wg             sync.WaitGroup
	requestRepo    model.RequestRepository
	storageService service.IStorageService
}

func NewInternalExportJob(requestRepo model.RequestRepository, storageService service.IStorageService) *InternalExportJob {
	return &InternalExportJob{
		requestRepo:    requestRepo,
		storageService: storageService,
		Wg:             sync.WaitGroup{},
	}
}

func (j *InternalExportJob) Execute(exporter *model.Exporter, reducer string, dataset string) *model.ExportMetrics {
	ctx := context.TODO()
	extractor, err := j.getCorrectExporter(exporter)
	if err != nil {
		return &model.ExportMetrics{
			Error: err.Error(),
		}
	}

	pr, pw := io.Pipe()
	gzipWriter := gzip.NewWriter(pw)

	resultChannel, _ := j.requestRepo.StreamByDataset(ctx, dataset)

	var wg sync.WaitGroup
	wg.Add(1)
	tracker := 0
	nonTracker := 0
	total := 0
	// Concurrently handle writing to gzip writer
	go func() {
		defer pw.Close()
		defer gzipWriter.Close()
		defer wg.Done()

		for requestData := range resultChannel {
			reduced := converter.ConvertRequestModel(requestData, converter.ReduceType(reducer))
			encoded, encodedErr := extractor.Encode(*reduced)

			if encodedErr != nil {
				continue
			}
			arr, err := json.Marshal(encoded)
			if err != nil {
				log.WithFields(log.Fields{
					"service": "InternalExportJob",
					"error":   err.Error(),
				}).Error("Could not convert int[] to string.")
				continue
			}

			data := strings.Trim(string(arr), "[]") + "\n"

			if _, err := gzipWriter.Write([]byte(data)); err != nil {
				log.WithFields(log.Fields{
					"service": "InternalExportJob",
					"error":   err.Error(),
				}).Error("Failed to write to gzip writer.")
				break // Break the loop if there's an error writing to the gzip writer
			}
			if reduced.Tracker {
				tracker = tracker + 1
			} else {
				nonTracker = nonTracker + 1
			}
			total = total + 1
		}
	}()
	err = j.storageService.PutObject(ctx, config.EnvExportBucketName(), exporter.Name+"_"+reducer+"_"+dataset+".csv.gz", pr, -1, "application/gzip")
	wg.Wait()
	if err != nil {
		return &model.ExportMetrics{
			Error: err.Error(),
		}
	}
	return &model.ExportMetrics{
		Total:      total,
		Tracker:    tracker,
		NonTracker: nonTracker,
	}
}

func (j *InternalExportJob) getCorrectExporter(exporter *model.Exporter) (*extractor.Extractor, error) {
	for _, ext := range extractor.EXTRACTORS {
		if ext.GetName() == exporter.Name {
			return &ext, nil
		}
	}
	return nil, errors.New("In Services Extractor could not be found.")
}

type ExternalExportJob struct {
	requestRepo    model.RequestRepository
	storageService service.IStorageService
}

func NewExternalExportJob(requestRepo model.RequestRepository, storageService service.IStorageService) *ExternalExportJob {
	return &ExternalExportJob{
		requestRepo:    requestRepo,
		storageService: storageService,
	}
}

func (j *ExternalExportJob) Execute(exporter *model.Exporter, reducer string, dataset string) *model.ExportMetrics {
	ctx := context.TODO()
	vm := otto.New()
	obj, err := j.storageService.GetObject(ctx, config.EnvExtractorBucketName(), *exporter.ExportScriptLocation)
	_, err = vm.Run(obj)
	if err != nil {
		return &model.ExportMetrics{
			Error: err.Error(),
		}
	}

	pr, pw := io.Pipe()
	gzipWriter := gzip.NewWriter(pw)

	resultChannel, _ := j.requestRepo.StreamByDataset(ctx, dataset)

	var wg sync.WaitGroup
	wg.Add(1)
	tracker := 0
	nonTracker := 0
	total := 0
	// Concurrently handle writing to gzip writer
	go func() {
		defer pw.Close()
		defer gzipWriter.Close()
		defer wg.Done()

		for requestData := range resultChannel {
			reduced := converter.ConvertRequestModel(requestData, converter.ReduceType(reducer))
			result, err := vm.Call("extract", nil, reduced)
			if err != nil {
				continue
			}
			encoded, err := result.Export()
			if err != nil {
				return
			}
			arr, err := json.Marshal(encoded)
			if err != nil {
				log.WithFields(log.Fields{
					"service": "InternalExportJob",
					"error":   err.Error(),
				}).Error("Could not convert int[] to string.")
				continue
			}

			data := strings.Trim(string(arr), "[]") + "\n"
			if _, err := gzipWriter.Write([]byte(data)); err != nil {
				log.WithFields(log.Fields{
					"service": "InternalExportJob",
					"error":   err.Error(),
				}).Error("Failed to write to gzip writer.")
				break // Break the loop if there's an error writing to the gzip writer
			}
			if reduced.Tracker {
				tracker = tracker + 1
			} else {
				nonTracker = nonTracker + 1
			}
			total = total + 1
		}
	}()

	err = j.storageService.PutObject(ctx, config.EnvExportBucketName(), exporter.Name+"_"+reducer+"_"+dataset+".csv.gz", pr, -1, "application/gzip")
	if err != nil {
		return &model.ExportMetrics{
			Error: err.Error(),
		}
	}
	wg.Wait()
	return &model.ExportMetrics{
		Total:      total,
		Tracker:    tracker,
		NonTracker: nonTracker,
	}
}
