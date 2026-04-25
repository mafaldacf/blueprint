package mediamicroservices

import (
	"context"
)

type Text struct {
	MovieID string `bson:"MovieID"`
	Title   string `bson:"Title"`
	CastID  string `bson:"CastID"`
	PlotID  string `bson:"PlotID"`
}

type TextService interface {
	UploadNewText(ctx context.Context, reqID int64, text string) error
}

type TextServiceImpl struct {
	composeReviewService ComposeReviewService
}

func NewTextServiceImpl(ctx context.Context, composeReviewService ComposeReviewService) (TextService, error) {
	s := &TextServiceImpl{composeReviewService: composeReviewService}
	return s, nil
}

func (s *TextServiceImpl) UploadNewText(ctx context.Context, reqID int64, text string) error {
	return s.composeReviewService.UploadText(ctx, reqID, text)
}
