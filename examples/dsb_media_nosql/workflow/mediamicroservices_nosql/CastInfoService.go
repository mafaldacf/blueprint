package mediamicroservices_nosql

import (
	"context"
	"errors"
	"fmt"

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
	ReadCastInfos(ctx context.Context, reqID int64, castInfoIDs []string) ([]CastInfo, error)
}

type CastInfoServiceImpl struct {
	database backend.NoSQLDatabase
	cache    backend.Cache
}

func NewCastInfoServiceImpl(ctx context.Context, database backend.NoSQLDatabase, cache backend.Cache) (CastInfoService, error) {
	s := &CastInfoServiceImpl{database: database, cache: cache}
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

func (s *CastInfoServiceImpl) ReadCastInfos(ctx context.Context, reqID int64, castInfoIDs []string) ([]CastInfo, error) {
	if len(castInfoIDs) == 0 {
		return nil, nil
	}

	// check duplicates
	unique_cast_info_ids := make(map[string]bool)
	for _, pid := range castInfoIDs {
		unique_cast_info_ids[pid] = true
	}
	if len(unique_cast_info_ids) != len(castInfoIDs) {
		return []CastInfo{}, errors.New("cast_info_ids are duplicated")
	}

	// get cast infos from memcached
	var cache_keys []string
	for _, id := range castInfoIDs {
		cache_keys = append(cache_keys, id)
	}

	cast_infos_cached := make([]CastInfo, len(cache_keys))
	var cached_values []interface{}
	for idx := range cast_infos_cached {
		cached_values = append(cached_values, &cast_infos_cached[idx])
	}

	err := s.cache.Mget(ctx, castInfoIDs, cached_values)
	if err != nil {
		return nil, err
	}

	// filter cast info ids not in memcached
	cast_info_ids_not_cached := unique_cast_info_ids
	for _, cast_info := range cast_infos_cached {
		delete(cast_info_ids_not_cached, cast_info.CastInfoID)
	}

	ret_cast_infos := cast_infos_cached

	// find the rest in MongoDB
	if len(cast_info_ids_not_cached) > 0 {
		var new_cast_infos []CastInfo
		collection, err := s.database.GetCollection(ctx, "cast_info_db", "cast")
		if err != nil {
			return nil, err
		}
		var ids []string
		for id := range cast_info_ids_not_cached {
			ids = append(ids, id)
		}
		query := bson.D{
			{Key: "CastInfoID", Value: bson.D{
				{Key: "$in", Value: ids},
			}},
		}
		cursor, err := collection.FindMany(ctx, query)
		if err != nil {
			return nil, err
		}
		err = cursor.All(ctx, &new_cast_infos)
		if err != nil {
			return nil, err
		}

		// append to return values
		ret_cast_infos = append(ret_cast_infos, new_cast_infos...)
		
		// update cast-info to memcached
		for _, new_cast_info := range new_cast_infos {
			s.cache.Put(ctx, new_cast_info.CastInfoID, new_cast_info)
		}

	}

	if len(ret_cast_infos) != len(castInfoIDs) {
		return nil, fmt.Errorf("cast-info-service return set incomplete")
	}

	return ret_cast_infos, err
}
