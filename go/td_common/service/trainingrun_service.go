package service

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/model"
)

type ITrainingrunService interface {
	FindAllTrainingRuns(ctx context.Context) ([]*model.TrainingRun, error)
	FindAllTrainingRunsForModelname(ctx context.Context, modelName string) ([]*model.TrainingRun, error)
	FindAllByModelId(ctx context.Context, modelId string) ([]*model.TrainingRun, error)
	DeleteAllByModelId(ctx context.Context, id string) error
	DeleteByID(ctx context.Context, id string) error
}

type TraingingrunService struct {
	trainingrunRepo model.TrainingRunRepository
}

func NewTraingingrunService(trainingrunRepo model.TrainingRunRepository) *TraingingrunService {
	return &TraingingrunService{
		trainingrunRepo: trainingrunRepo,
	}
}

func (s *TraingingrunService) FindAllTrainingRuns(ctx context.Context) ([]*model.TrainingRun, error) {
	return s.trainingrunRepo.FindAll(ctx)
}

func (s *TraingingrunService) FindAllTrainingRunsForModelname(ctx context.Context, modelName string) ([]*model.TrainingRun, error) {
	return s.trainingrunRepo.FindByModelName(ctx, modelName)
}

func (s *TraingingrunService) FindAllByModelId(ctx context.Context, modelId string) ([]*model.TrainingRun, error) {
	return s.trainingrunRepo.FindByModelID(ctx, modelId)
}

func (s *TraingingrunService) DeleteAllByModelId(ctx context.Context, id string) error {
	return s.trainingrunRepo.DeleteAllByModelID(ctx, id)
}

func (s *TraingingrunService) DeleteByID(ctx context.Context, id string) error {
	return s.trainingrunRepo.DeleteByID(ctx, id)
}
