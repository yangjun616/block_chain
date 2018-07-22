package gotest

import (
	"fmt"
	"encoding/json"
)

func main() {
	p := &Product{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	stu :=student{}
	stu.Name="xiaoming"
	byte,_:=json.Marshal(stu)
	fmt.Println(string(byte))

}

type Product struct {
	Name      string
	ProductID int64
	Number    int
	Price     float64
	IsOnSale  bool
}

type student struct {
	Name string
	Number    int

}