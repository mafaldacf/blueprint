package assurance

import (
	"context"
)

type AssuranceType struct {
	Index int64
	Name  string
	Price float64
}

var TRAFFIC_ACCIDENT = AssuranceType{1, "Traffic Accident Assurance", 3.0}
var ALL_ASSURANCES = []AssuranceType{TRAFFIC_ACCIDENT}

func getAssuranceType(ctx context.Context, index int64) (AssuranceType, error) {
	if index == TRAFFIC_ACCIDENT.Index {
		return TRAFFIC_ACCIDENT, nil
	}
	// FIXME: ADD SUPPORT FOR ORIGINAL CODE
	/* return AssuranceType{}, errors.New(fmt.Sprintf("Assurance with index %d does not exist", index)) */
	return AssuranceType{}, nil
}

type Assurance struct {
	ID      string
	OrderID string
	AT      AssuranceType
}
