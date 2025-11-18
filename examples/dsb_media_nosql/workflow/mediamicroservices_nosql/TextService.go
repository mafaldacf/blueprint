package mediamicroservices_nosql

import (
	"context"
)

type Text struct {
	MovieID string `bson:"_id"`
	Title   string
	CastID  string
	PlotID  string
}

type TextService interface {
	UploadText(ctx context.Context, reqID int64, text string) error
}

type TextServiceImpl struct {
	composeReviewService ComposeReviewService
}

func NewTextServiceImpl(ctx context.Context, composeReviewService ComposeReviewService) (TextService, error) {
	s := &TextServiceImpl{composeReviewService: composeReviewService}
	return s, nil
}

func (s *TextServiceImpl) UploadText(ctx context.Context, reqID int64, text string) error {
	return s.composeReviewService.UploadText(ctx, reqID, text)
}
