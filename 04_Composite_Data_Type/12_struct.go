package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	dilbert.Salary -= 5000
	fmt.Println(dilbert)

	position := &dilbert.Position
	*position = "California"
	fmt.Println(dilbert)

	emp := &dilbert
	emp.Address = "5th Avenue"
	fmt.Println(dilbert)
}
