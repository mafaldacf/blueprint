package shopping_app

import (
	"context"
	"sync"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type AnalyticsService interface {
	Run(ctx context.Context) error
}

type AnalyticsServiceImpl struct {
	analytics_db    backend.NoSQLDatabase
	analytics_queue backend.Queue
	num_workers     int
}

func NewAnalyticsServiceImpl(ctx context.Context, analytics_db backend.NoSQLDatabase, analytics_queue backend.Queue) (AnalyticsService, error) {
	return &AnalyticsServiceImpl{analytics_db: analytics_db, analytics_queue: analytics_queue, num_workers: 4}, nil
}

func (s *AnalyticsServiceImpl) updateAnalytics(ctx context.Context, message AnalyticsMessage) error {
	collection, _ := s.analytics_db.GetCollection(ctx, "analytics_database", "analytics_collection")
	var analytics Analytics
	query := bson.D{{Key: "userID", Value: message.UserID}}
	result, _ := collection.FindOne(ctx, query)
	result.One(ctx, &analytics)

	// dummy expr because analyzer doesn't detect the user id right now
	updatedAnalytics := Analytics{
		UserID: message.UserID,
		Categories: append(analytics.Categories, message.ProductCategory),
	}
	
	//filter := bson.D{{Key: "userID", Value: message.UserID}}
	// simulate upsert
	collection.InsertOne(ctx, updatedAnalytics)
	return nil
}

func (s *AnalyticsServiceImpl) workerThread(ctx context.Context) error {
	var forever chan struct{}
	go func() {
		var event map[string]interface{}
		s.analytics_queue.Pop(ctx, &event)
		workerMessage := AnalyticsMessage{
			UserID:          event["UserID"].(string),
			ProductCategory: event["ProductCategory"].(string),
		}
		s.updateAnalytics(ctx, workerMessage)
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
