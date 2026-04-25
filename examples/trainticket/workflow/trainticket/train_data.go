// package train implements ts-train-service from the original TrainTicket application
package trainticket

type TrainType struct {
	ID           string `bson:"ID"`
	Name         string `bson:"Name"`
	EconomyClass int    `bson:"EconomyClass"`
	ComfortClass int    `bson:"ComfortClass"`
	AvgSpeed     int64  `bson:"AvgSpeed"`
}
