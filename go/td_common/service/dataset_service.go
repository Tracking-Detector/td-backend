package service

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/payload"
)

type IDatasetService interface {
	Save(ctx context.Context, dataset *model.Dataset) (*model.Dataset, error)
	CreateDataset(ctx context.Context, datasetPayload *payload.CreateDatasetPayload) (*model.Dataset, error)
	SaveAll(ctx context.Context, datasets []*model.Dataset) ([]*model.Dataset, error)
	GetAllDatasets() []*model.Dataset
	ReloadCache(ctx context.Context)
	IsValidDataset(ctx context.Context, id string) bool
	GetDatasetByID(ctx context.Context, id string) (*model.Dataset, error)
	DeleteDataset(ctx context.Context, id string) error
	IsLabelValid(label string) bool
}

type DatasetService struct {
	datasetRepo  model.DatasetRepository
	requestRepo  model.RequestRepository
	datasetCache []*model.Dataset
}

func NewDatasetService(datasetRepo model.DatasetRepository, requestRepo model.RequestRepository) *DatasetService {
	service := &DatasetService{
		datasetRepo: datasetRepo,
		requestRepo: requestRepo,
	}
	service.ReloadCache(context.Background())
	return service
}

func (s *DatasetService) SaveAll(ctx context.Context, datasets []*model.Dataset) ([]*model.Dataset, error) {
	res, err := s.datasetRepo.SaveAll(ctx, datasets)
	s.ReloadCache(ctx)
	return res, err
}

func (s *DatasetService) CreateDataset(ctx context.Context, datasetPayload *payload.CreateDatasetPayload) (*model.Dataset, error) {
	dataset := &model.Dataset{
		Name:        datasetPayload.Name,
		Description: datasetPayload.Description,
		Label:       datasetPayload.Label,
	}
	return s.Save(ctx, dataset)
}

func (s *DatasetService) Save(ctx context.Context, dataset *model.Dataset) (*model.Dataset, error) {
	res, err := s.datasetRepo.Save(ctx, dataset)
	s.ReloadCache(ctx)
	return res, err
}

func (s *DatasetService) GetDatasetByID(ctx context.Context, id string) (*model.Dataset, error) {
	return s.datasetRepo.FindByID(ctx, id)
}

func (s *DatasetService) GetAllDatasets() []*model.Dataset {
	return s.datasetCache
}

func (s *DatasetService) IsValidDataset(ctx context.Context, id string) bool {
	dataset, err := s.datasetRepo.FindByID(ctx, id)
	if err != nil {
		return false
	}
	return dataset != nil
}

func (s *DatasetService) DeleteDataset(ctx context.Context, id string) error {

	dataset, err := s.datasetRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if dataset.Label != "" {
		if err := s.requestRepo.DeleteAllByLabel(ctx, dataset.Label); err != nil {
			return err
		}
	}
	err = s.datasetRepo.DeleteByID(ctx, id)

	s.ReloadCache(ctx)
	return err
}

func (s *DatasetService) ReloadCache(ctx context.Context) {
	datasets, _ := s.datasetRepo.FindAll(ctx)
	s.datasetCache = datasets
}

func (s *DatasetService) IsLabelValid(label string) bool {
	for _, dataset := range s.datasetCache {
		if dataset.Label == label {
			return true
		}
	}
	return false
}
