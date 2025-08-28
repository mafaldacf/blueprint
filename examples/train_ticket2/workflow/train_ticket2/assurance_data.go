package train_ticket2

import (
	"context"
	"errors"
	"fmt"
)

type AssuranceType struct {
	Index int
	Name  string
	Price float64
}

var TRAFFIC_ACCIDENT = AssuranceType{1, "Traffic Accident Assurance", 3.0}
var ALL_ASSURANCES = []AssuranceType{TRAFFIC_ACCIDENT}

func getAssuranceType(ctx context.Context, index int) (AssuranceType, error) {
	if index == TRAFFIC_ACCIDENT.Index {
		return TRAFFIC_ACCIDENT, nil
	}
	return AssuranceType{}, errors.New(fmt.Sprintf("Assurance with index %d does not exist", index))
}

type Assurance struct {
	ID      string
	OrderID string
	AT      AssuranceType
}
