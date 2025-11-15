package mediamicroservices_sql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type CastInfo struct {
	castinfoid string `bson:"_id"`
	name       string `bson:"name"`
	gender     string `bson:"gender"`
	intro      string `bson:"intro"`
}

type CastInfoService interface {
	WriteCastInfo(ctx context.Context, reqID int64, castInfoID string, name string, gender string, intro string) (CastInfo, error)
	ReadCastInfo(ctx context.Context, reqID int64, castInfoID string) (CastInfo, error)
}

type CastInfoServiceImpl struct {
	CastInfoDB backend.RelationalDB
}

func NewCastInfoServiceImpl(ctx context.Context, CastInfoDB backend.RelationalDB) (CastInfoService, error) {
	m := &CastInfoServiceImpl{CastInfoDB: CastInfoDB}
	return m, nil
}

func (m *CastInfoServiceImpl) WriteCastInfo(ctx context.Context, reqID int64, castInfoID string, name string, gender string, intro string) (CastInfo, error) {
	CastInfo := CastInfo{
		castinfoid: castInfoID,
		name:       name,
		gender:     gender,
		intro:      intro,
	}
	_, err := m.CastInfoDB.Exec(ctx, "INSERT INTO castinfo(castinfoid, name, gender, intro) VALUES (?, ?, ?, ?);", castInfoID, name, gender, intro)
	return CastInfo, err
}

func (m *CastInfoServiceImpl) ReadCastInfo(ctx context.Context, reqID int64, castInfoID string) (CastInfo, error) {
	var CastInfo CastInfo
	err := m.CastInfoDB.Select(ctx, &CastInfo, "SELECT * FROM castinfo WHERE castinfoid = ?", castInfoID)
	return CastInfo, err
}
