package postnotification

/* import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
) */

/* type TimelineService interface {
	ReadTimeline(ctx context.Context, reqID int64) (Post, Analytics, error)
}

type TimelineServiceImpl struct {
	storageService StorageService
	timelineCache  backend.Cache
}

func NewTimelineServiceImpl(ctx context.Context, storageService StorageService, timelineCache backend.Cache) (TimelineService, error) {
	n := &TimelineServiceImpl{storageService: storageService, timelineCache: timelineCache}
	return n, nil
}

func (s *TimelineServiceImpl) ReadTimeline(ctx context.Context, reqID int64) (Post, Analytics, error) {
	var timeline Timeline
	reqIDStr := strconv.FormatInt(reqID, 10)
	s.timelineCache.Get(ctx, reqIDStr, &timeline)
	post, analytics, err := s.storageService.ReadPostNoSQL(ctx, reqID, timeline.PostID)
	return post, analytics, err
} */
