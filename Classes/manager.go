// See https://github.com/yulin02/go/blob/0b7bf4aabdccb18c8b4b920b88aef4efb9c99164/sample/interfaces.go for interfaces in Go!
package main

import "fmt"

// Account interface (also known as an ABC in python + abstract class in java) //
type Account interface {
    Credit(amount int) int 
    Debit(amount int) int
    GetBalance() int
}

// Business Account //
type BusinessAccount struct {
    credits, debits int    
}

func (business *BusinessAccount) Credit(amount int) {
    business.credits = business.credits + amount
}

func (business *BusinessAccount) Debit(amount int) {
    business.debits = business.debits + amount
}

func (business BusinessAccount) GetBalance() int {
    return business.credits - business.debits
}

// Checking Account //
type CheckingAccount struct {
    credits, debits int    
}

func (checking *CheckingAccount) Credit(amount int) {
    checking.credits = checking.credits + amount
}

func (checking *CheckingAccount) Debit(amount int) {
    checking.debits = checking.debits + amount
}

func (checking CheckingAccount) GetBalance() int {
    return checking.credits - checking.debits
}

// Savings Account //
type SavingsAccount struct {
    credits, debits int    
}

func (savings *SavingsAccount) Credit(amount int) {
    savings.credits = savings.credits + amount
}

func (savings *SavingsAccount) Debit(amount int) {
    savings.debits = savings.debits + amount
}

func (savings SavingsAccount) GetBalance() int {
    return savings.credits - savings.debits
}

// Main routines //
func businessAccount() {
    business := BusinessAccount{0, 50}    
    business.Credit(200)
    business.Debit(250)
    fmt.Println("Business Balance:", business.GetBalance())
}

func checkingAccount() {
    checking := CheckingAccount{100, 20}    
    checking.Credit(900)
    checking.Debit(250)
    fmt.Println("Checking Balance:", checking.GetBalance())
}

func savingsAccount() {
    savings := SavingsAccount{200, 0}    
    savings.Credit(200)
    savings.Debit(126)
    fmt.Println("Savings Balance:", savings.GetBalance())
}

func main() {
    businessAccount()
    checkingAccount()
    savingsAccount()
}
