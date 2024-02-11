package service

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/extractor"
	"github.com/Tracking-Detector/td-backend/go/td_common/model"
)

type IExporterService interface {
	GetAllExporter(ctx context.Context) ([]*model.Exporter, error)
	InitInCodeExports(ctx context.Context)
	IsValidExporter(ctx context.Context, identifier string) bool
	FindByID(ctx context.Context, id string) (*model.Exporter, error)
}

type ExporterService struct {
	exporterRepo model.ExporterRepository
}

func NewExporterService(extractorRepo model.ExporterRepository) *ExporterService {
	return &ExporterService{
		exporterRepo: extractorRepo,
	}
}

func (s *ExporterService) GetAllExporter(ctx context.Context) ([]*model.Exporter, error) {
	return s.exporterRepo.FindAll(ctx)
}

func (s *ExporterService) InitInCodeExports(ctx context.Context) {
	for _, ext := range extractor.EXTRACTORS {
		exporterData := model.Exporter{
			Name:        ext.GetName(),
			Description: ext.GetDescription(),
			Dimensions:  ext.GetDimensions(),
			Type:        model.IN_SERVICE,
		}
		s.exporterRepo.Save(ctx, &exporterData)
	}
}

func (s *ExporterService) IsValidExporter(ctx context.Context, exporter string) bool {
	_, err := s.exporterRepo.FindByID(ctx, exporter)
	return err == nil

}

func (s *ExporterService) FindByID(ctx context.Context, id string) (*model.Exporter, error) {
	return s.exporterRepo.FindByID(ctx, id)
}
