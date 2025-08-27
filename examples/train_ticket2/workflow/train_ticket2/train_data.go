// package train implements ts-train-service from the original TrainTicket application
package train_ticket2

type TrainType struct {
	ID           string
	Name         string
	EconomyClass int
	ComfortClass int
	AvgSpeed     int64
}
