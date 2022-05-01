package t_test

import (
	"DESIGN_PATTERNS/Behavioral_patterns/Chain_Of_Responsibility/chain"
	"testing"
)

func TestChain(t *testing.T) {
	cashier := &chain.Cashier{} // 最后一步

	//Set next for medical department
	medical := &chain.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &chain.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &chain.Reception{}
	reception.SetNext(doctor)

	patient := &chain.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
