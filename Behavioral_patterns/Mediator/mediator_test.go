package t_test

import (
	"DESIGN_PATTERNS/Behavioral_patterns/Mediator/mediator"
	"testing"
)

func TestMediator(t *testing.T) {
	stationManager := mediator.NewStationManger()

	passengerTrain := &mediator.PassengerTrain{
		Mediator: stationManager,
	}
	freightTrain := &mediator.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
