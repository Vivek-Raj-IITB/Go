package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods")
	// no inheritance in golang; No super or parent
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 20}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMain()
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email) // email is not updated

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u User) NewMain() {
	u.Email = "test@go.com"
	fmt.Println("New email of user is: ", u.Email)
}
