package main

import (
	"fmt"
	"strconv"
)

type user interface {
	getId() int
	printDetail()
}

type appUser struct {
	id    int
	name  string
	email string
}

func (d *appUser) getId() int {
	return d.id
}

func (d *appUser) printDetail() {
	fmt.Println("[UserDetail] userId:", d.id, "UserName:", d.name, "email:", d.email)
}

func addUser(id int, name, email string) user {
	d := &appUser{}
	d.id = id
	d.name = name
	d.email = email
	return d
}

type instrument interface {
	getId() int
	printDetail()
}

type paymentInstrument struct {
	id   int
	name string
}

type card struct {
	paymentInstrument
}

func (d *card) getId() int {
	return d.id
}

func (d *card) printDetail() {
	fmt.Println("[InstrumentDetail] instrument id:", d.id, "instrumentName:", d.name)
}

func setCardInstrument(id int, name string) instrument {
	return &card{
		paymentInstrument{
			id:   id,
			name: name,
		},
	}
}

type bank struct {
	paymentInstrument
}

func (d *bank) getId() int {
	return d.id
}

func (d *bank) printDetail() {
	fmt.Println("[InstrumentDetail] instrument id:", d.id, "instrumentName:", d.name)
}

func setBankInstrument(id int, name string) instrument {
	return &bank{
		paymentInstrument{
			id:   id,
			name: name,
		},
	}
}

type transaction interface {
	doTransaction()
}

type paymentTransaction struct {
	amount            int
	userManager       *userManager
	instrumentManager *instrumentManager
	processor         processor
	notification      notification
}

func (d *paymentTransaction) doTransaction(fromSendId, toSendId, instrumentId int) bool {
	// check if toSend exists
	if !d.userManager.checkIfUserExists(toSendId) {
		fmt.Println("[TRANSACTION_IN_PROGRESS] userId:", toSendId, "donot exists. Aborting with payment")
		return false
	}
	fmt.Println("[TRANSACTION_IN_PROGRESS] userId:", toSendId, "exists. Proceeding with payment")

	// validate payment
	if !d.instrumentManager.validatePayment(fromSendId, instrumentId) {
		fmt.Println("[TRANSACTION_IN_PROGRESS] instrumentId:", instrumentId, "donot exists for userId: ", fromSendId, ". Aborting payment")
		return false
	}
	fmt.Println("[TRANSACTION_IN_PROGRESS] instrumentId:", instrumentId, "exists for userId: ", fromSendId, ". Proceeding with payment")

	// process payment
	d.processor.processPayment()

	// notify users
	notification1 := getNotificationObj(1, fromSendId, "Amount Debited:"+strconv.Itoa(d.amount))
	notification1.notifyUser()
	notification2 := getNotificationObj(2, toSendId, "Amount Credited:"+strconv.Itoa(d.amount))
	notification2.notifyUser()
	return true

}

type processor interface {
	processPayment()
}

type paymentProcessor struct {
	id int
}

func (d *paymentProcessor) processPayment() {
	fmt.Println("[PAYMENT-PROCESSOR] Processing payment with processorId:", d.id)
}

type notification interface {
	notifyUser()
}

type paymentNotification struct {
	notificationId int
	userId         int
	message        string
}

func (d *paymentNotification) notifyUser() {
	fmt.Println("[NOTIFICATION] Notification sent to UserId:", d.userId, "msg:", d.message, "notificationId:", d.notificationId)
}

func getNotificationObj(notificationId, userId int, message string) notification {
	return &paymentNotification{
		notificationId: notificationId,
		userId:         userId,
		message:        message,
	}
}

type userManager struct {
	users []user
}

func (d *userManager) addUser(id int, name, email string) {
	fmt.Println("Adding user with ID:", id, "name:", name, "email", email)
	u := addUser(id, name, email)
	d.users = append(d.users, u)

}

func (d *userManager) getUserDetails(userId int) user {
	for _, u := range d.users {
		if u.getId() == userId {
			return u
		}
	}
	fmt.Println("user not found")
	return nil
}

func (d *userManager) checkIfUserExists(id int) bool {
	for _, u := range d.users {
		if u.getId() == id {
			return true
		}
	}
	return false
}

type instrumentManager struct {
	userToInstrumentsMap map[int][]instrument
}

func (d *instrumentManager) addInstrument(userId int, instrumentID string, instrumentType string) {
	fmt.Println("Adding", instrumentID, "which is a", instrumentType, "for user", userId)
	if len(d.userToInstrumentsMap) == 0 {
		d.userToInstrumentsMap = make(map[int][]instrument)
	}
	if instruments, ok := d.userToInstrumentsMap[userId]; ok {
		lastID := len(instruments)
		switch instrumentType {
		case "card":
			instruments = append(instruments, setCardInstrument(lastID+1, instrumentID))
		case "bank":
			instruments = append(instruments, setBankInstrument(lastID+1, instrumentID))
		}
		d.userToInstrumentsMap[userId] = instruments
	} else {
		switch instrumentType {
		case "card":
			instruments = append(instruments, setCardInstrument(1, instrumentID))
		case "bank":
			instruments = append(instruments, setBankInstrument(1, instrumentID))
		}
		d.userToInstrumentsMap[userId] = instruments
	}
}

func (d *instrumentManager) selectInstrument(userId, instrumentId int) instrument {
	if instrument, ok := d.userToInstrumentsMap[userId]; ok {
		for _, i := range instrument {
			if i.getId() == instrumentId {
				return i
			}
		}
	}
	fmt.Println("card not found")
	return nil
}

func (d *instrumentManager) validatePayment(userId, paymentId int) bool {
	if payments, ok := d.userToInstrumentsMap[userId]; ok {
		for _, p := range payments {
			if p.getId() == paymentId {
				return true
			}
		}
	}
	return false
}

func main() {
	userManager := &userManager{}
	userManager.addUser(1, "shashank", "shashank@gmail.com")
	userManager.addUser(2, "prakash", "prakash@gmail.com")

	instrumentManager := &instrumentManager{}
	instrumentManager.addInstrument(1, "SBI_Card_1", "card")
	instrumentManager.addInstrument(2, "HDFC_Card_1", "card")
	instrumentManager.addInstrument(2, "AXIS_BANK", "bank")

	fmt.Println("----------------Running inference----------------")
	sendFrom := userManager.getUserDetails(2)
	sendTo := userManager.getUserDetails(1)

	instrumentToSend := instrumentManager.selectInstrument(2, 2)
	instrumentToSend.printDetail()

	transaction := &paymentTransaction{
		amount:            100,
		userManager:       userManager,
		instrumentManager: instrumentManager,
		processor:         &paymentProcessor{id: 1},
	}

	status := transaction.doTransaction(sendFrom.getId(), sendTo.getId(), instrumentToSend.getId())
	if status {
		fmt.Println("Transaction successful")
		return
	}
	fmt.Println("Transaction failed")

}
