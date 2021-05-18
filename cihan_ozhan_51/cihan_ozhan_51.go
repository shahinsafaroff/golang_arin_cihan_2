package main

import "fmt"

func main() {
//User Data Creation VERSION 1.0
	fmt.Println("User Creation VERSION 1.0")
	user1 := &User{
		ID:			1,
		FirstName:  "Shahin",
		LastName:   "Safaroff",
		UserName:   "shahinsafaroff",
		Age: 		34,
		Pay:		&Payment{
			Salary: 355.5,
			Bonus: 45.5,
		},
	}
	//For VERSION V1 OUTPUT
	fmt.Println(user1)
/*	fmt.Println(user1.GetPayment())
	fmt.Println(user1.GetFullNAme())
	fmt.Println("Salary: ",  user1.GetPayment())*/

	//User Data Creation VERSION 2.0
	fmt.Println("User Creation VERSION 2.0")

	user2 :=NewUser()
	user2.FirstName = "Cihan"
	user2.LastName = "Ozhan"
	user2.Age = 28
	user2.UserName = "GopHer"

	user2.Pay = &Payment{Salary: 755, Bonus: 960}
	fmt.Println(user2.GetFullNAme())
	fmt.Println(user2.GetPayment())
	fmt.Println("Salary: ", user2.GetPayment())

}

//User Struct
type User struct {
	ID			int
	FirstName 	string
	LastName 	string
	UserName 	string
	Age 		int
	Pay			*Payment
}

//User constructor struct method
func NewUser() *User{
	u := new(User)
	u.Pay = NewPayment()
	return u
}

//Payment struct
type Payment struct {
	Salary 		float64
	Bonus		float64
}

//Payment constructor struct method
func NewPayment() *Payment {
	p := new(Payment)
	return p
}

//Methods creation
func (u User) GetFullNAme() string {
	return u.FirstName + " " + u.LastName
}
func (u *User) GetUserName() string {
	return u.UserName
}
func ( u *User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}