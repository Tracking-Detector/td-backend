package service

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/model"
)

type IExportRunService interface {
	Save(ctx context.Context, exportRun *model.ExportRun) (*model.ExportRun, error)
	GetAll(ctx context.Context) ([]*model.ExportRun, error)
	GetByExporterID(ctx context.Context, exporterId string) ([]*model.ExportRun, error)
	ExistByExporterIDAndRecducer(ctx context.Context, exporterId, reducer string) (bool, error)
	GetByID(ctx context.Context, id string) (*model.ExportRun, error)
}

type ExportRunService struct {
	exportRunRepository model.ExportRunRepository
}

func NewExportRunService(exportRunRepository model.ExportRunRepository) *ExportRunService {
	return &ExportRunService{
		exportRunRepository: exportRunRepository,
	}
}

func (s *ExportRunService) Save(ctx context.Context, exportRun *model.ExportRun) (*model.ExportRun, error) {
	return s.exportRunRepository.Save(ctx, exportRun)
}

func (s *ExportRunService) GetAll(ctx context.Context) ([]*model.ExportRun, error) {
	return s.exportRunRepository.FindAll(ctx)
}

func (s *ExportRunService) GetByExporterID(ctx context.Context, exporterId string) ([]*model.ExportRun, error) {
	return s.exportRunRepository.FindByExporterID(ctx, exporterId)
}

func (s *ExportRunService) GetByID(ctx context.Context, id string) (*model.ExportRun, error) {
	return s.exportRunRepository.FindByID(ctx, id)
}

func (s *ExportRunService) ExistByExporterIDAndRecducer(ctx context.Context, exporterId, reducer string) (bool, error) {
	return s.exportRunRepository.ExistByExporterIDAndRecducer(ctx, exporterId, reducer)
}
