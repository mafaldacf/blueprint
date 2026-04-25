package trainticket

import (
	"context"
	"fmt"
)

type AssuranceType struct {
	Index int     `bson:"Index"`
	Name  string  `bson:"Name"`
	Price float64 `bson:"Price"`
}

type Assurance struct {
	ID      string        `bson:"ID"`
	OrderID string        `bson:"OrderID"`
	AT      AssuranceType `bson:"AT"`
}

var TRAFFIC_ACCIDENT = AssuranceType{1, "Traffic Accident Assurance", 3.0}
var ALL_ASSURANCES = []AssuranceType{TRAFFIC_ACCIDENT}

func getAssuranceType(ctx context.Context, index int) (AssuranceType, error) {
	if index == TRAFFIC_ACCIDENT.Index {
		return TRAFFIC_ACCIDENT, nil
	}
	return AssuranceType{}, fmt.Errorf("Assurance with index %d does not exist", index)
}
