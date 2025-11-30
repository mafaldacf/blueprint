package mediamicroservices

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"
)


type UniqueIdService interface {
	UploadUniqueId(ctx context.Context, reqID int64) error
}

type UniqueIdServiceImpl struct {
	counter              int64
	currentTimestamp     int64
	machineID            string
	composeReviewService ComposeReviewService
}

func NewUniqueIdServiceImpl(ctx context.Context, composeReviewService ComposeReviewService) (UniqueIdService, error) {
	s := &UniqueIdServiceImpl{counter: 0, currentTimestamp: -1, machineID: GetMachineID(), composeReviewService: composeReviewService}
	return s, nil
}

func (s *UniqueIdServiceImpl) UploadUniqueId(ctx context.Context, reqID int64) error {
	reviewID, err := s.GenerateUniqueId()
	if err != nil {
		return err
	}

	return s.composeReviewService.UploadUniqueId(ctx, reqID, reviewID)
}

func (s *UniqueIdServiceImpl) GenerateUniqueId() (int64, error) {
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

func (s *UniqueIdServiceImpl) GetCounter(timestamp int64) int64 {
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
