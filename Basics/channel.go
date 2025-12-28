package basics

import "fmt"


// this sendEmail function is using emaillist rechive only channel & done is send only channel
func sendEmail (emailList <- chan string, done chan<- bool) {
	defer func () {done <- true}()

	for value := range emailList {
		fmt.Println("Sending Mail to : ", value)
	}
}


func SendMailChannel () {
	mailList := make(chan string, 100)
	done := make(chan bool)

	for i := 0; i < 100; i++ {
		mail := fmt.Sprintf("%d@ankan.in", i)
		mailList <- mail
	}

	go sendEmail(mailList, done)
	close(mailList)
	<- done
}