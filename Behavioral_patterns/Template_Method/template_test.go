package t_test

import (
	"DESIGN_PATTERNS/Behavioral_patterns/Template_Method/templateMethod"
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	smsOTP := &templateMethod.Sms{}
	o := templateMethod.Otp{
		IOtp: smsOTP,
	}
	err := o.GenAndSendOTP(4)
	if err != nil {
		t.Error(err)
	}

	fmt.Println()
	emailOTP := &templateMethod.Email{}
	o = templateMethod.Otp{
		IOtp: emailOTP,
	}
	err = o.GenAndSendOTP(4)
	if err != nil {
		t.Error(err)
	}
}
