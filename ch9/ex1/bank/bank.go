package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var refunds = make(chan int) 
var refundResult = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Refund(amount int) { refunds <- amount }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-refunds:
			if balance >= amount {
				balance -= amount
				refundResult <- true
			} else {
				refundResult <- false
			}
		case balances <- balance:
		}
	}
}

func Withdraw(amount int) bool {
	Refund(amount)
	return <- refundResult

}

func init() {
	go teller() // start the monitor goroutine
}

//!-
