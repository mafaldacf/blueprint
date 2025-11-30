// package train implements ts-train-service from the original TrainTicket application
package trainticket

type TrainType struct {
	ID           string
	Name         string
	EconomyClass int
	ComfortClass int
	AvgSpeed     int64
}
