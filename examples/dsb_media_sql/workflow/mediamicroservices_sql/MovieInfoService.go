package mediamicroservices_sql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type MovieInfo struct {
	movieid string `bson:"_id"`
	title   string `bson:"title"`
	castid  string `bson:"cast_id"`
	plotid  string `bson:"plot_id"`
}

type MovieInfoService interface {
	WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, castID string, plotID string) (MovieInfo, error)
	ReadMovieInfo(ctx context.Context, reqID int64, movieID string) (MovieInfo, error)
}

type MovieInfoServiceImpl struct {
	movieInfoDB backend.RelationalDB
}

func NewMovieInfoServiceImpl(ctx context.Context, movieInfoDB backend.RelationalDB) (MovieInfoService, error) {
	m := &MovieInfoServiceImpl{movieInfoDB: movieInfoDB}
	return m, nil //, m.createTables(ctx)
}

func (m *MovieInfoServiceImpl) WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, castID string, plotID string) (MovieInfo, error) {
	movieInfo := MovieInfo{
		movieid: movieID,
		title:   title,
		castid:  castID,
		plotid:  plotID,
	}
	_, err := m.movieInfoDB.Exec(ctx, "INSERT INTO movieinfo(movieid, title, castid, plotid) VALUES (?, ?, ?, ?);", movieID, title, castID, plotID)
	return movieInfo, err
}

func (m *MovieInfoServiceImpl) ReadMovieInfo(ctx context.Context, reqID int64, movieID string) (MovieInfo, error) {
	var movieInfo MovieInfo
	err := m.movieInfoDB.Select(ctx, &movieInfo, "SELECT * FROM movieinfo WHERE movieid = ?", movieID)
	return movieInfo, err
}

/* func (m *MovieInfoServiceImpl) createTables(ctx context.Context) error {
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
} */
