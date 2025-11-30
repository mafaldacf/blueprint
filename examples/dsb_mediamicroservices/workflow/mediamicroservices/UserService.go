package mediamicroservices

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	UserID    int64 `bson:"_id"`
	Username  string
	FirstName string
	LastName  string
	Password  string
	Salt      string
}

type UserService interface {
	RegisterUser(ctx context.Context, reqID string, firstName string, lastName string, username string, password string) (User, error)
	RegisterUserWithId(ctx context.Context, reqID string, firstName string, lastName string, username string, password string, userID int64) (User, error)
	UploadUserWithUsername(ctx context.Context, reqID int64, username string) error
	UploadUserWithUserId(ctx context.Context, reqID int64, userID int64) error
	Login(ctx context.Context, reqID int64, username string, password string) error
}

type UserServiceImpl struct {
	counter              int64
	currentTimestamp     int64
	machineID            string
	database             backend.NoSQLDatabase
	cache                backend.Cache
	composeReviewService ComposeReviewService
}

func NewUserServiceImpl(ctx context.Context, database backend.NoSQLDatabase, cache backend.Cache, composeReviewService ComposeReviewService) (UserService, error) {
	s := &UserServiceImpl{counter: 0, currentTimestamp: -1, machineID: GetMachineID(), database: database, cache: cache, composeReviewService: composeReviewService}
	return s, nil
}

func (s *UserServiceImpl) RegisterUser(ctx context.Context, reqID string, firstName string, lastName string, username string, password string) (User, error) {
	var userID int64
	userID, err := s.GenerateUniqueId()
	if err != nil {
		return User{}, err
	}

	salt := genRandomString(32)
	passwordHashed := hashSHA256(password + salt)

	user := User{
		UserID:    userID,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Salt:      salt,
		Password:  passwordHashed,
	}

	collection, err := s.database.GetCollection(ctx, "user_db", "user")
	if err != nil {
		return User{}, err
	}
	err = collection.InsertOne(ctx, user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *UserServiceImpl) RegisterUserWithId(ctx context.Context, reqID string, firstName string, lastName string, username string, password string, userID int64) (User, error) {
	salt := genRandomString(32)
	passwordHashed := hashSHA256(password + salt)

	user := User{
		UserID:    userID,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Salt:      salt,
		Password:  passwordHashed,
	}

	collection, err := s.database.GetCollection(ctx, "user_db", "user")
	if err != nil {
		return User{}, err
	}
	err = collection.InsertOne(ctx, user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *UserServiceImpl) UploadUserWithUsername(ctx context.Context, reqID int64, username string) error {
	var userID int64
	ok, err := s.cache.Get(ctx, username+":user_id", &userID)
	if err != nil {
		return err
	}

	if !ok {
		var user User
		collection, err := s.database.GetCollection(ctx, "user_db", "user")
		if err != nil {
			return err
		}
		query := bson.D{{Key: "Username", Value: username}}
		result, err := collection.FindOne(ctx, query)
		if err != nil {
			return err
		}
		res, err := result.One(ctx, &user)
		if !res || err != nil {
			return err
		}
		userID = user.UserID
	}

	err = s.composeReviewService.UploadUserId(ctx, reqID, userID)
	if err != nil {
		return err
	}

	err = s.cache.Put(ctx, username+":user_id", userID)
	return err
}

func (s *UserServiceImpl) UploadUserWithUserId(ctx context.Context, reqID int64, userID int64) error {
	return s.composeReviewService.UploadUserId(ctx, reqID, userID)
}

func (s *UserServiceImpl) Login(ctx context.Context, reqID int64, username string, password string) error {
	var passwordHashed string
	ok1, err := s.cache.Get(ctx, username+":password", &passwordHashed)
	if err != nil {
		return err
	}

	var salt string
	ok2, err := s.cache.Get(ctx, username+":salt", &salt)
	if err != nil {
		return err
	}

	var userID int64
	ok3, err := s.cache.Get(ctx, username+":user_id", &userID)
	if err != nil {
		return err
	}

	if !ok1 || !ok2 || !ok3 {
		var user User
		collection, err := s.database.GetCollection(ctx, "user_db", "user")
		if err != nil {
			return err
		}
		query := bson.D{{Key: "Username", Value: username}}
		result, err := collection.FindOne(ctx, query)
		if err != nil {
			return err
		}
		res, err := result.One(ctx, &user)
		if !res || err != nil {
			return err
		}

		if !ok1 {
			passwordHashed = user.Password
		}
		if !ok2 {
			salt = user.Salt
		}
		if !ok3 {
			userID = user.UserID
		}
	}

	// TODO: verify password
	return nil
}

func (s *UserServiceImpl) GenerateUniqueId() (int64, error) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	idx := s.GetCounter(timestamp)
	timestamp_hex := strconv.FormatInt(timestamp, 16)
	if len(timestamp_hex) > 10 {
		timestamp_hex = timestamp_hex[:10]
	} else if len(timestamp_hex) < 10 {
		timestamp_hex = strings.Repeat("0", 10-len(timestamp_hex)) + timestamp_hex
	}
	counter_hex := strconv.FormatInt(idx, 16)
	if len(counter_hex) > 1 {
		counter_hex = counter_hex[:1]
	} else if len(counter_hex) < 1 {
		counter_hex = strings.Repeat("0", 1-len(counter_hex)) + counter_hex
	}
	log.Println(s.machineID, timestamp_hex, counter_hex)
	unique_id_str := s.machineID + timestamp_hex + counter_hex
	unique_id, err := strconv.ParseInt(unique_id_str, 16, 64)
	if err != nil {
		return 0, err
	}
	unique_id = unique_id & 0x7FFFFFFFFFFFFFFF
	return unique_id, nil
}

func (s *UserServiceImpl) GetCounter(timestamp int64) int64 {
	if s.currentTimestamp == timestamp {
		retVal := s.counter
		s.counter += 1
		return retVal
	} else {
		s.currentTimestamp = timestamp
		s.counter = 1
		return 0
	}
}

func genRandomString(n int) string {
	var alphaNum = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = alphaNum[r.Intn(len(alphaNum))]
	}
	return string(b)
}

func hashSHA256(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
