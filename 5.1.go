package main

import "fmt"

type Dish struct {
	Name  string
	Price float64
}

type Order struct {
	Dishes []Dish
	Total  float64
}

func (order *Order) AddDish(dish Dish) {
	order.Dishes = append(order.Dishes, dish)
}
func (order *Order) RemoveDish(dish Dish) {
	for i := range order.Dishes {
		if order.Dishes[i].Name == dish.Name {
			order.Dishes[i].Name = ""
			order.Dishes[i].Price = 0.0
		}
	}
}
func (order *Order) CalculateTotal() {
	order.Total = 0.0
	for _, value := range order.Dishes {
		order.Total += value.Price
	}
}

func main() {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	order.CalculateTotal()
	fmt.Println("Total:", order.Total)

	order.RemoveDish(dish1)

	order.CalculateTotal()
	fmt.Println("Total:", order.Total)
}
