package service

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/model"
)

type IModelService interface {
	Save(ctx context.Context, model *model.Model) (*model.Model, error)
	GetAllModels(ctx context.Context) ([]*model.Model, error)
	GetModelByName(ctx context.Context, name string) (*model.Model, error)
	DeleteModelByID(ctx context.Context, id string) error
	GetModelById(ctx context.Context, id string) (*model.Model, error)
}

type ModelService struct {
	modelRepo          model.ModelRepository
	trainingrunService ITrainingrunService
}

func NewModelService(modelRepo model.ModelRepository, trainingrunService ITrainingrunService) *ModelService {
	return &ModelService{
		modelRepo:          modelRepo,
		trainingrunService: trainingrunService,
	}
}

func (s *ModelService) Save(ctx context.Context, model *model.Model) (*model.Model, error) {
	return s.modelRepo.Save(ctx, model)
}

func (s *ModelService) GetAllModels(ctx context.Context) ([]*model.Model, error) {
	return s.modelRepo.FindAll(ctx)
}

func (s *ModelService) GetModelByName(ctx context.Context, name string) (*model.Model, error) {
	return s.modelRepo.FindByName(ctx, name)
}

func (s *ModelService) GetModelById(ctx context.Context, id string) (*model.Model, error) {
	return s.modelRepo.FindByID(ctx, id)
}

func (s *ModelService) DeleteModelByID(ctx context.Context, id string) error {
	return s.modelRepo.InTransaction(ctx, func(ctx context.Context) error {
		if err := s.modelRepo.DeleteByID(ctx, id); err != nil {
			return err
		}
		return s.trainingrunService.DeleteAllByModelId(ctx, id)
	})
}
