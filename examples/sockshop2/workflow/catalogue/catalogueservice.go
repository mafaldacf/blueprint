// Package catalogue implements the SockShop catalogue microservice
package catalogue

import (
	"context"
	"strings"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type CatalogueService interface {
	List(ctx context.Context, tags []string, order string, pageNum, pageSize int) ([]Sock, error)
	Count(ctx context.Context, tags []string) (int, error)
	Get(ctx context.Context, id string) (Sock, error)
	Tags(ctx context.Context) ([]string, error)
	AddTags(ctx context.Context, tags []string) error
	AddSock(ctx context.Context, sock Sock) (string, error)
	DeleteSock(ctx context.Context, id string) error
}

type catalogueImpl struct {
	catalogue_db backend.NoSQLDatabase
}

func NewCatalogueService(ctx context.Context, catalogue_db backend.NoSQLDatabase) (CatalogueService, error) {
	c := &catalogueImpl{catalogue_db: catalogue_db}
	return c, nil
}

// List implements CatalogueService.
func (s *catalogueImpl) List(ctx context.Context, tags []string, order string, pageNum int, pageSize int) ([]Sock, error) {
	collection, _ := s.catalogue_db.GetCollection(ctx, "catalogue", "catalogue")
	var allSocks []Sock
	var socks []Sock
	filter := bson.D{}
	result, _ := collection.FindMany(ctx, filter)
	result.All(ctx, &allSocks)
	/* for _, sock := range allSocks {
		for _, tag := range tags {
			if slices.Contains(sock.Tags, tag) {
				socks = append(socks, sock)
				break
			}
		}
	} */
	return socks, nil
}

// Count implements CatalogueService.
func (s *catalogueImpl) Count(ctx context.Context, tags []string) (int, error) {
	collection, _ := s.catalogue_db.GetCollection(ctx, "catalogue", "catalogue")
	var socks []Sock
	filter := bson.D{}
	result, _ := collection.FindMany(ctx, filter)
	result.All(ctx, &socks)
	return len(socks), nil
}

func (s *catalogueImpl) Get(ctx context.Context, id string) (Sock, error) {
	collection, _ := s.catalogue_db.GetCollection(ctx, "catalogue", "catalogue")
	query := bson.D{{Key: "id", Value: id}}
	var sock Sock
	result, _ := collection.FindOne(ctx, query)
	result.One(ctx, &sock)
	sock.ImageURL = []string{sock.ImageURL_1, sock.ImageURL_2}
	sock.Tags = strings.Split(sock.TagString, ",")
	return sock, nil
}

// Tags implements CatalogueService.
func (s *catalogueImpl) Tags(ctx context.Context) ([]string, error) {
	collection, _ := s.catalogue_db.GetCollection(ctx, "catalogue", "catalogue")
	var socks []Sock
	projection := bson.D{{Key: "tags", Value: true}}
	filter := bson.D{}
	result, _ := collection.FindMany(ctx, filter, projection)
	result.All(ctx, &socks)
	var tags []string
	/* for _, sock := range socks {
		tags = append(tags, sock.Tags...)
	} */
	return tags, nil
}

// AddTags implements CatalogueService.
func (s *catalogueImpl) AddTags(ctx context.Context, tags []string) error {
	collection, _ := s.catalogue_db.GetCollection(ctx, "catalogue", "catalogue")
	update := bson.D{{Key: "tags", Value: tags}}
	filter := bson.D{}
	collection.Upsert(ctx, filter, update)
	return nil
}

// AddSock implements CatalogueService.
func (s *catalogueImpl) AddSock(ctx context.Context, sock Sock) (string, error) {
	collection, _ := s.catalogue_db.GetCollection(ctx, "catalogue", "catalogue")
	collection.InsertOne(ctx, sock)
	return sock.ID, nil
}

// DeleteSock implements CatalogueService.
func (s *catalogueImpl) DeleteSock(ctx context.Context, id string) error {
	collection, _ := s.catalogue_db.GetCollection(ctx, "catalogue", "catalogue")
	query := bson.D{{Key: "id", Value: id}}
	collection.DeleteOne(ctx, query)
	return nil
}
