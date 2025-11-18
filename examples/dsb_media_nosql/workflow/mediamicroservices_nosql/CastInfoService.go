package mediamicroservices_nosql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type CastInfo struct {
	CastInfoID string `bson:"_id"`
	Name       string
	Gender     string
	Intro      string
}

type CastInfoService interface {
	WriteCastInfo(ctx context.Context, reqID int64, castInfoID string, name string, gender string, intro string) (CastInfo, error)
	ReadCastInfos(ctx context.Context, reqID int64, castIDs []string) ([]CastInfo, error)
}

type CastInfoServiceImpl struct {
	database backend.NoSQLDatabase
}

func NewCastInfoServiceImpl(ctx context.Context, database backend.NoSQLDatabase) (CastInfoService, error) {
	s := &CastInfoServiceImpl{database: database}
	return s, nil
}

func (s *CastInfoServiceImpl) WriteCastInfo(ctx context.Context, reqID int64, castInfoID string, name string, gender string, intro string) (CastInfo, error) {
	castInfo := CastInfo{
		CastInfoID: castInfoID,
		Name:       name,
		Gender:     gender,
		Intro:      intro,
	}

	collection, err := s.database.GetCollection(ctx, "cast_info_db", "cast")
	if err != nil {
		return CastInfo{}, err
	}
	err = collection.InsertOne(ctx, castInfo)
	if err != nil {
		return CastInfo{}, err
	}

	return castInfo, err
}

func (s *CastInfoServiceImpl) ReadCastInfos(ctx context.Context, reqID int64, castIDs []string) ([]CastInfo, error) {
	collection, err := s.database.GetCollection(ctx, "cast_info_db", "cast")
	if err != nil {
		return nil, err
	}

	var castInfos []CastInfo
	query := bson.D{
		{Key: "CastInfoID", Value: bson.D{
			{Key: "$in", Value: castIDs},
		}},
	}

	vals, err := collection.FindMany(ctx, query)
	if err != nil {
		return nil, err
	}
	err = vals.All(ctx, &castInfos)
	if err != nil {
		return nil, err
	}

	return castInfos, err
}
