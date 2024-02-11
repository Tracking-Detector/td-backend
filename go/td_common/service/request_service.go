package service

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/model"
)

type IRequestService interface {
	GetRequestById(ctx context.Context, id string) (*model.RequestData, error)
	InsertManyRequests(ctx context.Context, requests []*model.RequestData) error
	SaveRequest(ctx context.Context, request *model.RequestData) (*model.RequestData, error)
	StreamAll(ctx context.Context) (<-chan *model.RequestData, <-chan error)
	GetPagedRequestsFilterdByUrl(ctx context.Context, url string, page, pageSize int) ([]*model.RequestData, error)
	CountDocumentsForUrlFilter(ctx context.Context, url string) (int64, error)
}

type RequestService struct {
	requestRepo model.RequestRepository
}

func NewRequestService(requestRepo model.RequestRepository) *RequestService {
	return &RequestService{
		requestRepo: requestRepo,
	}
}

func (s *RequestService) StreamAll(ctx context.Context) (<-chan *model.RequestData, <-chan error) {
	return s.requestRepo.StreamAll(ctx)
}

func (s *RequestService) GetRequestById(ctx context.Context, id string) (*model.RequestData, error) {
	return s.requestRepo.FindByID(ctx, id)
}

func (s *RequestService) InsertManyRequests(ctx context.Context, requests []*model.RequestData) error {
	_, err := s.requestRepo.SaveAll(ctx, requests)
	return err
}

func (s *RequestService) SaveRequest(ctx context.Context, request *model.RequestData) (*model.RequestData, error) {
	return s.requestRepo.Save(ctx, request)
}

func (s *RequestService) GetPagedRequestsFilterdByUrl(ctx context.Context, url string, page, pageSize int) ([]*model.RequestData, error) {
	return s.requestRepo.FindAllByUrlLikePaged(ctx, url, page, pageSize)
}

func (s *RequestService) CountDocumentsForUrlFilter(ctx context.Context, url string) (int64, error) {
	return s.requestRepo.CountByUrlLike(ctx, url)
}
