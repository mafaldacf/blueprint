package mediamicroservices_sql

import (
	"context"
	"os"
	"strings"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type MovieInfo struct {
	MovieID string `bson:"_id"`
	Title   string `bson:"title"`
	Casts   string `bson:"casts"`
}

type MovieInfoService interface {
	WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error)
	ReadMovieInfo(ctx context.Context, reqID int64, movieID string) (MovieInfo, error)
}

type MovieInfoServiceImpl struct {
	movieInfoDB backend.RelationalDB
}

func NewMovieInfoServiceImpl(ctx context.Context, movieIdDB backend.RelationalDB) (MovieInfoService, error) {
	m := &MovieInfoServiceImpl{movieInfoDB: movieIdDB}
	return m, m.createTables(ctx)
}

func (m *MovieInfoServiceImpl) WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error) {
	movieInfo := MovieInfo{
		MovieID: movieID,
		Title:   title,
		Casts:   casts,
	}
	_ , err := m.movieInfoDB.Exec(ctx, "INSERT INTO movieinfo(movieid, title, casts) VALUES (?, ?);", movieID, title, casts)
	return movieInfo, err
}

func (m *MovieInfoServiceImpl) ReadMovieInfo(ctx context.Context, reqID int64, movieID string) (MovieInfo, error) {
	var movieInfo MovieInfo
	err := m.movieInfoDB.Select(ctx, &movieInfo, "SELECT * FROM movieinfo WHERE movieid = ?", movieID)
	return movieInfo, err
}

func (m *MovieInfoServiceImpl) createTables(ctx context.Context) error {
	sqlBytes, err := os.ReadFile("database/movieinfo.sql")
	if err != nil {
		return err
	}
	sqlStatements := strings.Split(string(sqlBytes), ";")
	for _, stmt := range sqlStatements {
		_, err := m.movieInfoDB.Exec(ctx, stmt)
		if err != nil {
			return err
		}
	}
	return nil
}
