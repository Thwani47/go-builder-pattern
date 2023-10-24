package main

type Response struct {
	User    User
	Account Account
	Address Address
}

type User struct {
	Email string
}

type Account struct {
	Balance float64
}

type Address struct {
	City string
}
