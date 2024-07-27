package postnotification

import (
	"context"
	"sync"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification/common"
)

type AnalyticsService interface {
	Run(ctx context.Context) error
	ReadAnalytics(ctx context.Context, postID int64) (Analytics, error)
}

type AnalyticsServiceImpl struct {
	analytics_queue backend.Queue
	analytics_db    backend.NoSQLDatabase
	num_workers     int
}

func NewAnalyticsServiceImpl(ctx context.Context, analytics_db backend.NoSQLDatabase, analytics_queue backend.Queue) (AnalyticsService, error) {
	n := &AnalyticsServiceImpl{analytics_db: analytics_db, analytics_queue: analytics_queue, num_workers: 4}
	return n, nil
}

func (a *AnalyticsServiceImpl) ReadAnalytics(ctx context.Context, postID int64) (Analytics, error) {
	var analytics Analytics
	collection, err := a.analytics_db.GetCollection(ctx, "analytics_db", "analytics_collection")
	if err != nil {
		return analytics, err
	}
	query := bson.D{{Key: "postid", Value: postID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return analytics, err
	}
	res, err := result.One(ctx, &analytics)
	if !res || err != nil {
		return analytics, err
	}
	return analytics, nil
}

func (a *AnalyticsServiceImpl) handleMessage(ctx context.Context, message TriggerAnalyticsMessage) error {
	postID, err := common.StringToInt64(message.PostID)
	if err != nil {
		return err
	}

	collection, err := a.analytics_db.GetCollection(ctx, "analytics_db", "analytics_collection")
	if err != nil {
		return err
	}

	analytics := Analytics{
		PostID: postID,
	}
	return collection.InsertOne(ctx, analytics)
}

func (n *AnalyticsServiceImpl) workerThread(ctx context.Context) error {
	var forever chan struct{}
	go func() {
		var event TriggerAnalyticsMessage
		n.analytics_queue.Pop(ctx, &event)
		n.handleMessage(ctx, event)
	}()
	<-forever
	return nil
}

func (n *AnalyticsServiceImpl) Run(ctx context.Context) error {
	backend.GetLogger().Info(ctx, "initializing %d workers", n.num_workers)
	var wg sync.WaitGroup
	wg.Add(n.num_workers)
	for i := 1; i <= n.num_workers; i++ {
		go func(i int) {
			defer wg.Done()
			err := n.workerThread(ctx)
			if err != nil {
				backend.GetLogger().Error(ctx, "error in worker thread: %s", err.Error())
				panic(err)
			}
		}(i)
	}
	wg.Wait()
	backend.GetLogger().Info(ctx, "joining %d workers", n.num_workers)
	return nil
}
