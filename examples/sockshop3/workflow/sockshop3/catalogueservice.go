// Package catalogue implements the SockShop catalogue microservice
package sockshop3

import (
	"context"
	"strings"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CatalogueService interface {
	List(ctx context.Context, tags []string, order string, pageNum, pageSize int) ([]Sock, error)
	Count(ctx context.Context, tags []string) (int, error)
	Get(ctx context.Context, id string) (Sock, error)
	Tags(ctx context.Context) ([]string, error)

	// Disabled
	AddTags(ctx context.Context, tags []string) error
	AddSock(ctx context.Context, sock Sock) (string, error)
	
	// Disabled
	DeleteSock(ctx context.Context, id string) error
}

type CatalogueServiceImpl struct {
	catalogue_db backend.RelationalDB
}

func NewCatalogueServiceImpl(ctx context.Context, catalogue_db backend.RelationalDB) (CatalogueService, error) {
	c := &CatalogueServiceImpl{catalogue_db: catalogue_db}
	return c, nil
}

func (s *CatalogueServiceImpl) List(ctx context.Context, tags []string, order string, pageNum int, pageSize int) ([]Sock, error) {
	return nil, nil
	var socks []Sock
	query := `SELECT sock.ID, 
						sock.name, 
						sock.description, 
						sock.price, 
						sock.quantity, 
						sock.image_url_1, 
						sock.image_url_2, 
						GROUP_CONCAT(DISTINCT alltags.name) AS tag_name 
				FROM sock 
				JOIN sock_tag allsocktags ON sock.ID=allsocktags.ID 
				JOIN tag alltags ON allsocktags.tag_id=alltags.tag_id
				JOIN sock_tag ON sock.ID=sock_tag.ID
				JOIN tag ON tag.tag_id=sock_tag.tag_id`

	var args []interface{}

	for i, t := range tags {
		if i == 0 {
			query += " WHERE tag.name=?"
			args = append(args, t)
		} else {
			query += " OR tag.name=?"
			args = append(args, t)
		}
	}

	query += " GROUP BY sock.ID"

	if order != "" {
		query += " ORDER BY ?"
		args = append(args, order)
	}

	query += ";"

	err := s.catalogue_db.Select(ctx, &socks, query, args...)
	if err != nil {
		return []Sock{}, err
	}
	for i, s := range socks {
		socks[i].ImageURL = []string{s.ImageURL_1, s.ImageURL_2}
		socks[i].Tags = strings.Split(s.TagString, ",")
	}

	// cut
	if pageNum == 0 || pageSize == 0 {
		return nil, nil
	}
	start := (pageNum * pageSize) - pageSize
	if start > len(socks) {
		return nil, nil
	}
	end := (pageNum * pageSize)
	if end > len(socks) {
		end = len(socks)
	}
	return socks[start:end], nil
}

func (s *CatalogueServiceImpl) Count(ctx context.Context, tags []string) (int, error) {
	query := "SELECT COUNT(DISTINCT sock.ID) FROM sock JOIN sock_tag ON sock.ID=sock_tag.ID JOIN tag ON sock_tag.tag_id=tag.tag_id"

	var args []interface{}

	for i, t := range tags {
		if i == 0 {
			query += " WHERE tag.name=?"
			args = append(args, t)
		} else {
			query += " OR tag.name=?"
			args = append(args, t)
		}
	}

	query += ";"

	sel, err := s.catalogue_db.Prepare(ctx, query)

	if err != nil {
		return 0, err
	}
	defer sel.Close()

	var count int
	err = sel.QueryRow(args...).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *CatalogueServiceImpl) Get(ctx context.Context, id string) (Sock, error) {
	//TODO
	/* query := `SELECT sock.ID, 
						sock.name, 
						sock.description, 
						sock.price, 
						sock.quantity, 
						sock.image_url_1, 
						sock.image_url_2, 
						GROUP_CONCAT(DISTINCT alltags.name) AS tag_name 
				FROM sock 
				JOIN sock_tag allsocktags ON sock.ID=allsocktags.ID 
				JOIN tag alltags ON allsocktags.tag_id=alltags.tag_id
				JOIN sock_tag ON sock.ID=sock_tag.ID
				JOIN tag ON tag.tag_id=sock_tag.tag_id
				WHERE sock.ID =? GROUP BY sock.ID;` */

	var sock Sock
	err := s.catalogue_db.Select(ctx, &sock, "SELECT * FROM sock WHERE sock.ID =?", id)
	if err != nil {
		return Sock{}, errors.Wrapf(err, "CatalogueService.Get %v", id)
	}

	sock.ImageURL = []string{sock.ImageURL_1, sock.ImageURL_2}
	sock.Tags = strings.Split(sock.TagString, ",")

	return sock, nil
}

// Tags implements CatalogueService.
func (s *CatalogueServiceImpl) Tags(ctx context.Context) ([]string, error) {
	var tags []string
	err := s.catalogue_db.Select(ctx, &tags, "SELECT name FROM tag")
	return tags, err
}

// AddTags implements CatalogueService.
func (s *CatalogueServiceImpl) AddTags(ctx context.Context, tags []string) error {
	var currentTags []tag
	if err := s.catalogue_db.Select(ctx, &currentTags, "SELECT * FROM tag;"); err != nil {
		return err
	}

	tagLookup := make(map[string]int)
	for _, tag := range currentTags {
		tagLookup[tag.Name] = tag.ID
	}

	tagIds := []int{}
	for _, tagName := range tags {
		if _, tagAlreadyExists := tagLookup[tagName]; !tagAlreadyExists {
			// Insert the tag
			res, err := s.catalogue_db.Exec(ctx, "INSERT INTO tag (name) VALUES (?);", tagName)
			if err != nil {
				return err
			}
			id, err := res.LastInsertId()
			if err != nil {
				return err
			}
			tagLookup[tagName] = int(id)
		}

		tagIds = append(tagIds, tagLookup[tagName])
	}

	return nil
}

func (s *CatalogueServiceImpl) AddSock(ctx context.Context, sock Sock) (string, error) {
	// Delete any existing sock with this ID
	if sock.ID != "" {
		// DeleteSock()
		// Delete sock's tags
		_, err := s.catalogue_db.Exec(ctx, "DELETE FROM sock_tag WHERE sock_tag.ID=?;", sock.ID)
		if err != nil {
			return "", err
		}

		// Delete existing sock
		_, err = s.catalogue_db.Exec(ctx, "DELETE FROM sock WHERE sock.ID=?;", sock.ID)
		if err != nil {
			return "", err
		}
	} else {
		sock.ID = uuid.NewString()
	}

	// Add the sock
	_, err := s.catalogue_db.Exec(ctx, "INSERT INTO sock (ID, name, description, price, quantity, image_url_1, image_url_2) VALUES (?, ?, ?, ?, ?, ?, ?);",
		sock.ID, sock.Name, sock.Description, sock.Price, sock.Quantity, sock.ImageURL_1, sock.ImageURL_2)
	if err != nil {
		return "", err
	}

	// Make sure the tags are in the DB
	// addTags()
	var currentTags []tag
	if err := s.catalogue_db.Select(ctx, &currentTags, "SELECT * FROM tag;"); err != nil {
		return "", err
	}

	tagLookup := make(map[string]int)
	for _, tag := range currentTags {
		tagLookup[tag.Name] = tag.ID
	}

	tagIds := []int{}
	for _, tagName := range sock.Tags {
		if _, tagAlreadyExists := tagLookup[tagName]; !tagAlreadyExists {
			// Insert the tag
			res, err := s.catalogue_db.Exec(ctx, "INSERT INTO tag (name) VALUES (?);", tagName)
			if err != nil {
				return "", err
			}
			id, err := res.LastInsertId()
			if err != nil {
				return "", err
			}
			tagLookup[tagName] = int(id)
		}

		tagIds = append(tagIds, tagLookup[tagName])
	}

	// Add the tags to the sock
	for _, tagId := range tagIds {
		_, err = s.catalogue_db.Exec(ctx, "INSERT INTO sock_tag (ID, tag_id) VALUES (?, ?);", sock.ID, tagId)
		if err != nil {
			return "", err
		}
	}

	return sock.ID, nil
}

// DeleteSock implements CatalogueService.
func (s *CatalogueServiceImpl) DeleteSock(ctx context.Context, id string) error {
	if id == "" {
		return nil
	}

	// Delete sock's tags
	_, err := s.catalogue_db.Exec(ctx, "DELETE FROM sock_tag WHERE sock_tag.ID=?;", id)
	if err != nil {
		return err
	}

	// Delete existing sock
	_, err = s.catalogue_db.Exec(ctx, "DELETE FROM sock WHERE sock.ID=?;", id)
	if err != nil {
		return err
	}

	return nil
}
