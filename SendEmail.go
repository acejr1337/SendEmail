import {
	"gopkg.in/gomail.v2"
	"fmt"
}

func main() {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "from@gmail.com")
	mail.SetHeader("To", "to@gmail.com")
	mail.SetHeader("Subject", "Subject")
	mail.SetBody("text/plain", "text goes here")
	dialer := gomail.NewPlainDialer("smtp.gmail.com", 587, "from@gmail.com", "password")
	if err := dialer.DialAndSend(mail);
		err != nil {
			panic(err)
		}
		fmt.Print("Email has been successfully sent")
}