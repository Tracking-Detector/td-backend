package testsupport

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"os"

	"github.com/Tracking-Detector/td-backend/go/td_common/model"
)

func LoadRequestJson() []*model.RequestData {
	rawJson, _ := os.ReadFile("./testdata/requests/requests.json")
	var requestDataList []*model.RequestData
	err := json.Unmarshal(rawJson, &requestDataList)
	if err != nil {
		panic(err)
	}
	return requestDataList
}

func LoadFile(path string) string {
	file, _ := os.ReadFile(path)
	return string(file)
}

func LoadGzFile(path string) string {
	// Open the gzipped file
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return ""
	}
	defer gzipReader.Close()

	// Read the decompressed content
	content, err := io.ReadAll(gzipReader)
	if err != nil {
		return ""
	}

	return string(content)
}

func Unzip(file io.ReadSeekCloser) string {
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return ""
	}
	defer gzipReader.Close()

	// Read the decompressed content
	content, err := io.ReadAll(gzipReader)
	if err != nil {
		return ""
	}

	return string(content)
}

func CreateResultsChannel(requests []*model.RequestData) <-chan *model.RequestData {
	objectsCh := make(chan *model.RequestData, len(requests))
	for _, obj := range requests {
		objectsCh <- obj
	}
	defer close(objectsCh)
	return objectsCh
}

func CreateErrorChannel(errs []error) <-chan error {
	objectsCh := make(chan error, len(errs))
	for _, obj := range errs {
		objectsCh <- obj
	}
	defer close(objectsCh)
	return objectsCh
}
