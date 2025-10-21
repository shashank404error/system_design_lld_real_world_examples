// https://workat.tech/machine-coding/practice/splitwise-problem-0kp2yneec2q2
package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	EQUAL   = "EQUAL"
	EXACT   = "EXACT"
	PERCENT = "PERCENT"
)

type user struct {
	userId       string
	name         string
	email        string
	mobileNumber string
	transactions []transaction
}

func newUser(userId, name, email, mobileNumber string) *user {
	return &user{
		userId:       userId,
		name:         name,
		email:        email,
		mobileNumber: mobileNumber,
	}
}

func (d *user) addTransaction(transactionType string, userIDs []string, amount float64, exactShare []float64, percentageShare []int) (transaction, error) {
	var transaction transaction
	var err error
	var splits []float64
	switch transactionType {
	case EQUAL:
		log.Println("[addTransaction] Recording an equal expense")
		transaction, err = createEqualShareTransaction(d.userId, userIDs, amount)
		split := amount / float64(len(userIDs))
		for _, _ = range userIDs {
			splits = append(splits, split)
		}
		break
	case PERCENT:
		log.Println("[addTransaction] Recording a percentage expense")
		transaction, err = createPercentageShareTransaction(d.userId, userIDs, amount)
		for _, p := range percentageShare {
			splits = append(splits, (float64(p) * amount * 0.01))
		}
		break
	case EXACT:
		log.Println("[addTransaction] Recording a percentage expense")
		transaction, err = createExactShareTransactionTransaction(d.userId, userIDs, amount)
		for _, e := range exactShare {
			splits = append(splits, e)
		}
		break
	}

	// common logic
	if err == nil {
		transaction.addUserIds(userIDs)
		transaction.addSplits(splits)
		transaction.updateSplit()
		d.transactions = append(d.transactions, transaction)
		return transaction, nil
	}

	log.Println("failed to add transaction")
	return nil, errors.New("failed to add transaction")
}

func (d *user) showBalance() {
	totalHeOwes := 0.0
	totalOwedToHim := 0.0

	if len(d.transactions) == 0 {
		log.Println("No balances for user:", d.userId)
	}

	for _, t := range d.transactions {
		if t.getWhoPaid() == d.userId {
			for k, v := range t.getUserToSplitAmount() {
				if k != d.userId {
					log.Println(k, "owes", d.userId, ":", v)
					totalOwedToHim = totalOwedToHim + v
				}
			}
		} else {
			for k, v := range t.getUserToSplitAmount() {
				if k == d.userId {
					log.Println(d.userId, "owes", k, ":", v)
					totalHeOwes = totalHeOwes + v
				}
			}
		}
	}
}

type transaction interface {
	getWhoPaid() string
	getUserToSplitAmount() map[string]float64
	updateSplit()
	addUserIds([]string)
	addSplits([]float64)
}

type baseTransaction struct {
	transactionType   string
	amount            float64
	userToAmountSplit map[string]float64
	id                string
	whoPaid           string // userid of who paid
	userIds           []string
	splits            []float64
}

func (d *baseTransaction) getWhoPaid() string {
	return d.whoPaid
}

func (d *baseTransaction) getUserToSplitAmount() map[string]float64 {
	return d.userToAmountSplit
}

func (d *baseTransaction) updateSplit() {
	userToAmountSplit := make(map[string]float64)
	for index, u := range d.userIds {
		userToAmountSplit[u] = d.splits[index]
	}
	d.userToAmountSplit = userToAmountSplit
}

func (d *baseTransaction) addUserIds(userIds []string) {
	d.userIds = userIds
}

func (d *baseTransaction) addSplits(splits []float64) {
	d.splits = splits
}

type equalShareTransaction struct {
	baseTransaction
}

func createEqualShareTransaction(paidByUserId string, userIDs []string, amount float64) (transaction, error) {
	transaction := &equalShareTransaction{}
	baseTransaction := baseTransaction{}
	baseTransaction.amount = amount
	baseTransaction.id = time.Now().String()
	baseTransaction.whoPaid = paidByUserId
	baseTransaction.transactionType = EQUAL
	transaction.baseTransaction = baseTransaction
	log.Println("[createEqualShareTransaction] splitting", amount, "among", len(userIDs), "users equally")
	return transaction, nil
}

type percentageShareTransaction struct {
	baseTransaction
}

func createPercentageShareTransaction(paidByUserId string, userIDs []string, amount float64) (transaction, error) {
	transaction := &percentageShareTransaction{}
	baseTransaction := baseTransaction{}
	baseTransaction.amount = amount
	baseTransaction.id = time.Now().String()
	baseTransaction.whoPaid = paidByUserId
	baseTransaction.transactionType = PERCENT
	transaction.baseTransaction = baseTransaction
	log.Println("[createPercentageShareTransaction] splitting", amount, "among", len(userIDs), "users by percentage")
	return transaction, nil
}

type exactShareTransaction struct {
	baseTransaction
}

func createExactShareTransactionTransaction(paidByUserId string, userIDs []string, amount float64) (transaction, error) {
	transaction := &exactShareTransaction{}
	baseTransaction := baseTransaction{}
	baseTransaction.amount = amount
	baseTransaction.id = time.Now().String()
	baseTransaction.whoPaid = paidByUserId
	baseTransaction.transactionType = EXACT
	transaction.baseTransaction = baseTransaction
	log.Println("[createPercentageShareTransaction] splitting", amount, "among", len(userIDs), "users exactly")
	return transaction, nil
}

type expenseTracker struct {
	users        []*user
	transactions []transaction
}

func (d *expenseTracker) showBalance() {

}

func (d *expenseTracker) getUser(userId string) (*user, error) {
	for _, u := range d.users {
		if u.userId == userId {
			return u, nil
		}
	}
	log.Println("user with userId:", userId, "not found")
	return nil, errors.New("user not found")
}

func (d *expenseTracker) addTransaction(transaction transaction) {
	log.Println("Added transaction to ledger")
	d.transactions = append(d.transactions, transaction)
}

func main() {
	expenseTracker := &expenseTracker{}
	expenseTracker.users = append(expenseTracker.users, newUser("u1", "shashank", "shashank@gmail.com", "9340212623"))
	expenseTracker.users = append(expenseTracker.users, newUser("u2", "prakash", "prakash@gmail.com", "8828232123"))
	expenseTracker.users = append(expenseTracker.users, newUser("u3", "sharma", "sharma@gmail.com", "9340212623"))
	expenseTracker.users = append(expenseTracker.users, newUser("u4", "golu", "golu@gmail.com", "8240212623"))

	cmds := []string{
		"EXPENSE u1 1000 4 u1 u2 u3 u4 EQUAL",
		"EXPENSE u4 1200 4 u1 u2 u3 u4 PERCENT 40 20 20 20",
		"EXPENSE u1 1250 2 u2 u3 EXACT 370 880",
		"SHOW u4",
		"SHOW u1",
		"SHOW u3",
	}

	for _, cmd := range cmds {
		log.Println("executing:", cmd)
		cmdSplit := strings.Fields(cmd)
		// cmd := "SHOW"
		switch cmdSplit[0] {
		case "SHOW":
			if len(cmdSplit) == 1 {
				expenseTracker.showBalance()
				continue
			}
			if len(cmdSplit) > 1 {
				user, err := expenseTracker.getUser(cmdSplit[1])
				if err == nil {
					user.showBalance()
				}
				expenseTracker.showBalance()
				continue
			}
		case "EXPENSE":
			amountStr := cmdSplit[2]
			amount, _ := strconv.ParseFloat(amountStr, 64)
			totalUsersStr := cmdSplit[3]
			totalUsers, _ := strconv.Atoi(totalUsersStr)
			var users []string
			for i := 0; i < totalUsers; i++ {
				users = append(users, cmdSplit[4+i])
			}
			user, err := expenseTracker.getUser(cmdSplit[1])
			if err == nil {
				transactionType := cmdSplit[totalUsers+4]
				var transaction transaction
				var err error
				switch transactionType {
				case EQUAL:
					transaction, err = user.addTransaction(transactionType, users, amount, []float64{}, []int{})
					break
				case PERCENT:
					var percentageShare []int
					for i := 0; i < totalUsers; i++ {
						s, _ := strconv.Atoi(cmdSplit[totalUsers+5+i])
						percentageShare = append(percentageShare, s)
					}
					transaction, err = user.addTransaction(transactionType, users, amount, []float64{}, percentageShare)
					break
				case EXACT:
					var exactShare []float64
					for i := 0; i < totalUsers; i++ {
						s, _ := strconv.ParseFloat(cmdSplit[totalUsers+5+i], 64)
						exactShare = append(exactShare, s)
					}
					transaction, err = user.addTransaction(transactionType, users, amount, exactShare, []int{})
					break
				}
				if err == nil {
					expenseTracker.addTransaction(transaction)
				}

			}
			log.Println("failed adding expense")
		}
	}
}
