package mediamicroservices_sql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type MovieId struct {
	movieid string `bson:"_id"`
	title   string `bson:"title"`
}

type MovieIdService interface {
	RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error)
	ReadMovieId(ctx context.Context, reqID int64, movieID string) (MovieId, error)
}

type MovieIdServiceImpl struct {
	movieIdDB backend.RelationalDB
}

func NewMovieIdServiceImpl(ctx context.Context, movieIdDB backend.RelationalDB) (MovieIdService, error) {
	m := &MovieIdServiceImpl{movieIdDB: movieIdDB}
	return m, nil //,m.createTables(ctx)
}

func (m *MovieIdServiceImpl) RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error) {
	movie := MovieId{
		movieid: movieID,
		title:   title,
	}
	_, err := m.movieIdDB.Exec(ctx, "INSERT INTO movieid(movieid, title) VALUES (?, ?);", 
		movieID, title)
	return movie, err
}

func (m *MovieIdServiceImpl) ReadMovieId(ctx context.Context, reqID int64, movieID string) (MovieId, error) {
	var movieId MovieId
	err := m.movieIdDB.Select(ctx, &movieId, "SELECT * FROM movieid WHERE movieid = ?", movieID)
	return movieId, err
}

/* func (m *MovieIdServiceImpl) createTables(ctx context.Context) error {
	sqlBytes, err := os.ReadFile("database/movieid.sql")
	if err != nil {
		return err
	}
	sqlStatements := strings.Split(string(sqlBytes), ";")
	for _, stmt := range sqlStatements {
		_, err := m.movieIdDB.Exec(ctx, stmt)
		if err != nil {
			return err
		}
	}
	return nil
} */
