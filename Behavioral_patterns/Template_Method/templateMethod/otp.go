package templateMethod

type iOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type Otp struct {
	IOtp iOtp
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
	otp := o.IOtp.genRandomOTP(otpLength)
	o.IOtp.saveOTPCache(otp)
	message := o.IOtp.getMessage(otp)
	err := o.IOtp.sendNotification(message)
	if err != nil {
		return err
	}
	o.IOtp.publishMetric()
	return nil
}
