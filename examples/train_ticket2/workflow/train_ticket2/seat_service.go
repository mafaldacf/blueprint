package train_ticket2

import (
	"context"
	"math/rand"
)

type SeatService interface {
	DistributeSeat(ctx context.Context, seatRequest SeatRequest) (Ticket, error)
}

type SeatServiceImpl struct {
	orderService OrderService
}

func NewSeatServiceImpl(ctx context.Context, orderService OrderService) (SeatService, error) {
	return &SeatServiceImpl{orderService: orderService}, nil
}

func (s *SeatServiceImpl) DistributeSeat(ctx context.Context, seatRequest SeatRequest) (Ticket, error) {
	leftTicketInfo, err := s.orderService.GetTicketListByDateAndTripID(ctx, seatRequest)
	if err != nil {
		return Ticket{}, nil
	}

	// assign seats
	stationList := seatRequest.Stations
	seatTotalNum := seatRequest.TotalNum
	startStation := seatRequest.StartStation

	var ticket Ticket
	ticket.StartStation = startStation
	ticket.DestStation = seatRequest.DestStation

	// assign new tickets
	rangeNum := seatTotalNum
	seat := rand.Intn(seatTotalNum) + 1

	if len(leftTicketInfo.SoldTickets) > 0 {
		soldTickets := leftTicketInfo.SoldTickets
		// give priority to tickets already sold
		for _, soldTicket := range soldTickets {
			soldTicketDestStation := soldTicket.DestStation
			// tickets can be allocated if the sold ticket's end station is before the start station of the request
			if indexOf(stationList, soldTicketDestStation) < indexOf(stationList, startStation) {
				ticket.SeatNo = soldTicket.SeatNo
				return Ticket{}, nil
			}
		}
		for isContained(soldTickets, seat) {
			seat = rand.Intn(rangeNum) + 1
		}
	}
	ticket.SeatNo = seat
	return ticket, nil
}
