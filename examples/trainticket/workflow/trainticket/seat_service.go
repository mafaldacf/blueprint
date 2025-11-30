package trainticket

import (
	"context"
	"math/rand"
	"strconv"
	"strings"
)

type SeatService interface {
	DistributeSeat(ctx context.Context, seatRequest SeatRequest) (Ticket, error)
	GetLeftTicketOfInterval(ctx context.Context, seatRequest SeatRequest) (int, error)
}

type SeatServiceImpl struct {
	orderService  OrderService
	configService ConfigService
}

func NewSeatServiceImpl(ctx context.Context, orderService OrderService, configService ConfigService) (SeatService, error) {
	return &SeatServiceImpl{orderService: orderService, configService: configService}, nil
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

func (s *SeatServiceImpl) GetLeftTicketOfInterval(ctx context.Context, seatRequest SeatRequest) (int, error) {
	numOfLeftTicket := 0
	var leftTicketInfo LeftTicketInfo

	trainNumber := seatRequest.TrainNumber
	if strings.HasPrefix(trainNumber, "G") || strings.HasPrefix(trainNumber, "D") {
		var err error
		leftTicketInfo, err = s.orderService.GetTicketListByDateAndTripID(ctx, seatRequest)
		if err != nil {
			return -1, nil
		}
	} else {
		// request to order-other-service
	}

	// count the seats remaining in certain sections
	stationList := seatRequest.Stations
	seatTotalNum := seatRequest.TotalNum
	solidTicketSize := 0
	if len(leftTicketInfo.SoldTickets) > 0 {
		startStation := seatRequest.StartStation
		soldTickets := leftTicketInfo.SoldTickets
		solidTicketSize = len(soldTickets)
		for _, soldTicket := range soldTickets {
			soldTicketDestStation := soldTicket.DestStation
			if indexOf(stationList, soldTicketDestStation) < indexOf(stationList, startStation) {
				numOfLeftTicket++
			}
		}
	}

	// count the unsold tickets
	config, err := s.configService.Find(ctx, "DirectTicketAllocationProportion")
	if err != nil {
		return -1, err
	}
	direstPart, err := strconv.ParseFloat(config.Value, 64)
	if err != nil {
		return -1, err
	}

	if stationList[0] == seatRequest.StartStation && stationList[len(stationList)-1] == seatRequest.DestStation {
		// do nothing
	} else {
		direstPart = 1.0 - direstPart
	}

	unusedNum := (seatTotalNum * int(direstPart)) - solidTicketSize
	numOfLeftTicket += unusedNum
	return numOfLeftTicket, nil
}
