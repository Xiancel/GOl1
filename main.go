package main

func main() {
	age := 16
	isAdult := age >= 18
	hasPerm := isLoggedIn && hasRole
	canAccess := isAdmin || hasPerm
	isRestricted := !canAccess
}
