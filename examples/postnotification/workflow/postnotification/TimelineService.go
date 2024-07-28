package postnotification

import (
	"context"
	"strconv"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type TimelineService interface {
	ReadTimeline(ctx context.Context, reqID int64) (Post, Analytics, error)
}

type TimelineServiceImpl struct {
	storage_service StorageService
	timeline_cache  backend.Cache
}

func NewTimelineServiceImpl(ctx context.Context, storage_service StorageService, timeline_cache backend.Cache) (TimelineService, error) {
	n := &TimelineServiceImpl{storage_service: storage_service, timeline_cache: timeline_cache}
	return n, nil
}

func (s *TimelineServiceImpl) ReadTimeline(ctx context.Context, reqID int64) (Post, Analytics, error) {
	var timeline Timeline
	reqIDStr := strconv.FormatInt(reqID, 10)
	s.timeline_cache.Get(ctx, reqIDStr, &timeline)
	post, analytics, err := s.storage_service.ReadPostNoSQL(ctx, reqID, timeline.PostID)
	return post, analytics, err
}
