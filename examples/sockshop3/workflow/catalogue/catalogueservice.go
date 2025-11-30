// Package catalogue implements the SockShop catalogue microservice
package catalogue

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
	// extra endpoints
	AddTags(ctx context.Context, tags []string) error
	AddSock(ctx context.Context, sock Sock) (string, error)
	DeleteSock(ctx context.Context, id string) error
}

type CatalogueServiceImpl struct {
	catalogue_db backend.RelationalDB
}

func NewCatalogueServiceImpl(ctx context.Context, catalogue_db backend.RelationalDB) (CatalogueService, error) {
	c := &CatalogueServiceImpl{catalogue_db: catalogue_db}
	return c, c.createTables(ctx)
}

var baseQuery = `SELECT sock.SockID, 
						sock.Name, 
						sock.Description, 
						sock.Price, 
						sock.Quantity, 
						sock.ImageURL1, 
						sock.ImageURL2, 
						GROUP_CONCAT(DISTINCT alltags.Name) AS tag_name 
				FROM sock 
				JOIN sock_tag allsocktags ON sock.SockID=allsocktags.SockID 
				JOIN tag alltags ON allsocktags.TagID=alltags.TagID
				JOIN sock_tag ON sock.SockID=sock_tag.SockID
				JOIN tag ON tag.TagID=sock_tag.TagID`

func (s *CatalogueServiceImpl) List(ctx context.Context, tags []string, order string, pageNum int, pageSize int) ([]Sock, error) {
	var socks []Sock
	query := baseQuery

	var args []interface{}

	for i, t := range tags {
		if i == 0 {
			query += " WHERE tag.Name=?"
			args = append(args, t)
		} else {
			query += " OR tag.Name=?"
			args = append(args, t)
		}
	}

	query += " GROUP BY sock.SockID"

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
		socks[i].ImageURL = []string{s.ImageURL1, s.ImageURL2}
		socks[i].Tags = strings.Split(s.TagString, ",")
	}

	socks = cut(socks, pageNum, pageSize)

	return socks, nil
}

func (s *CatalogueServiceImpl) Count(ctx context.Context, tags []string) (int, error) {
	query := "SELECT COUNT(DISTINCT sock.SockID) FROM sock JOIN sock_tag ON sock.SockID=sock_tag.TagID JOIN tag ON sock_tag.TagID=tag.TagID"

	var args []interface{}

	for i, t := range tags {
		if i == 0 {
			query += " WHERE tag.Name=?"
			args = append(args, t)
		} else {
			query += " OR tag.Name=?"
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
	//query := baseQuery + " WHERE sock.SockID =? GROUP BY sock.SockID;"
	var sock Sock
	err := s.catalogue_db.Select(ctx, &sock, "SELECT * FROM sock WHERE sock.SockID =? GROUP BY sock.SockID;", id)
	if err != nil {
		return Sock{}, errors.Wrapf(err, "CatalogueService.Get %v", id)
	}

	sock.ImageURL = []string{sock.ImageURL1, sock.ImageURL2}
	sock.Tags = strings.Split(sock.TagString, ",")

	return sock, nil
}

// Tags implements CatalogueService.
func (s *CatalogueServiceImpl) Tags(ctx context.Context) ([]string, error) {
	var tags []string
	err := s.catalogue_db.Select(ctx, &tags, "SELECT Name FROM tag")
	return tags, err
}

// AddTags implements CatalogueService.
func (s *CatalogueServiceImpl) AddTags(ctx context.Context, tags []string) error {
	_, err := s.addTags(ctx, tags...)
	return err
}

func (s *CatalogueServiceImpl) AddSock(ctx context.Context, sock Sock) (string, error) {
	// Delete any existing sock with this ID
	if sock.SockID != "" {
		if err := s.DeleteSock(ctx, sock.SockID); err != nil {
			return "", err
		}
	} else {
		sock.SockID = uuid.NewString()
	}

	// Add the sock
	_, err := s.catalogue_db.Exec(ctx, "INSERT INTO sock (SockID, Name, Description, Price, Quantity, ImageURL1, ImageURL2) VALUES (?, ?, ?, ?, ?, ?, ?);",
		sock.SockID, sock.Name, sock.Description, sock.Price, sock.Quantity, sock.ImageURL1, sock.ImageURL2)
	if err != nil {
		return "", err
	}

	// Make sure the tags are in the DB
	tagIds, err := s.addTags(ctx, sock.Tags...)
	if err != nil {
		return "", err
	}

	// Add the tags to the sock
	for _, tagId := range tagIds {
		_, err = s.catalogue_db.Exec(ctx, "INSERT INTO sock_tag (SockID, TagID) VALUES (?, ?);", sock.SockID, tagId)
		if err != nil {
			return "", err
		}
	}

	return sock.SockID, nil
}

// DeleteSock implements CatalogueService.
func (s *CatalogueServiceImpl) DeleteSock(ctx context.Context, id string) error {
	if id == "" {
		return nil
	}

	// Delete sock's tags
	_, err := s.catalogue_db.Exec(ctx, "DELETE FROM sock_tag WHERE sock_tag.TagID=?;", id)
	if err != nil {
		return err
	}

	// Delete existing sock
	_, err = s.catalogue_db.Exec(ctx, "DELETE FROM sock WHERE sock.SockID=?;", id)
	if err != nil {
		return err
	}

	return nil
}

func cut(socks []Sock, pageNum, pageSize int) []Sock {
	if pageNum == 0 || pageSize == 0 {
		return []Sock{} // pageNum is 1-indexed
	}
	start := (pageNum * pageSize) - pageSize
	if start > len(socks) {
		return []Sock{}
	}
	end := (pageNum * pageSize)
	if end > len(socks) {
		end = len(socks)
	}
	return socks[start:end]
}

func (s *CatalogueServiceImpl) addTags(ctx context.Context, tags ...string) ([]int, error) {
	var currentTags []tag
	if err := s.catalogue_db.Select(ctx, &currentTags, "SELECT * FROM tag;"); err != nil {
		return nil, err
	}

	tagLookup := make(map[string]int)
	for _, tag := range currentTags {
		tagLookup[tag.Name] = tag.TagID
	}

	tagIds := []int{}
	for _, tagName := range tags {
		if _, tagAlreadyExists := tagLookup[tagName]; !tagAlreadyExists {
			// Insert the tag
			res, err := s.catalogue_db.Exec(ctx, "INSERT INTO tag (Name) VALUES (?);", tagName)
			if err != nil {
				return nil, err
			}
			id, err := res.LastInsertId()
			if err != nil {
				return nil, err
			}
			tagLookup[tagName] = int(id)
		}

		tagIds = append(tagIds, tagLookup[tagName])
	}

	return tagIds, nil
}

// Creates database tables if they don't already exist
func (c *CatalogueServiceImpl) createTables(ctx context.Context) (err error) {
	if _, err = c.catalogue_db.Exec(ctx, createSockTable); err != nil {
		return errors.Wrap(err, "unable to create sock table")
	}
	if _, err = c.catalogue_db.Exec(ctx, createTagTable); err != nil {
		if _, err = c.catalogue_db.Exec(ctx, createTagTable2); err != nil {
			return errors.Wrap(err, "unable to create Tag table")
		}
	}
	if _, err = c.catalogue_db.Exec(ctx, createSockTagTable); err != nil {
		return errors.Wrap(err, "unable to create socktag table")
	}
	return nil
}

var createSockTable = `CREATE TABLE IF NOT EXISTS sock (
	SockID varchar(40) NOT NULL, 
	Name varchar(20), 
	Description varchar(200), 
	Price float, 
	Quantity int, 
	ImageURL1 varchar(40), 
	ImageURL2 varchar(40), 
	PRIMARY KEY(SockID)
);`

// AUTOINCREMENT should be AUTO_INCREMENT in mysql
var createTagTable = `CREATE TABLE IF NOT EXISTS tag (
	TagID INTEGER PRIMARY KEY AUTO_INCREMENT, 
	Name varchar(20)
);`

var createTagTable2 = `CREATE TABLE IF NOT EXISTS tag (
	TagID INTEGER PRIMARY KEY AUTOINCREMENT, 
	Name varchar(20)
);`

var createSockTagTable = `CREATE TABLE IF NOT EXISTS sock_tag (
	SockID varchar(40), 
	TagID INTEGER, 
	FOREIGN KEY (SockID) 
		REFERENCES sock(SockID), 
	FOREIGN KEY(TagID)
		REFERENCES tag(TagID)
);`
