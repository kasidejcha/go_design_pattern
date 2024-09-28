package main

import "fmt"

// Component Interface: Defines the method to get the cost
type Coffee interface {
	Cost() int
	Description() string
}

// Concrete Component: SimpleCoffee
type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() int {
	return 5
}

func (c *SimpleCoffee) Description() string {
	return "Simple Coffee"
}


// Decorator: Base Decorator
type CoffeeDecorator struct {
	coffee Coffee
}

func (d *CoffeeDecorator) Cost() int {
	return d.coffee.Cost()
}

func (d *CoffeeDecorator) Description() string {
	return d.coffee.Description()
}


// Concrete Decorator: MilkDecorator
type MilkDecorator struct {
	CoffeeDecorator
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
	return &MilkDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (m *MilkDecorator) Cost() int {
	return m.coffee.Cost() + 2 // Add cost of milk
}

func (m *MilkDecorator) Description() string {
	return m.coffee.Description() + ", Milk"
}


// Concrete Decorator: SugarDecorator
type SugarDecorator struct {
	CoffeeDecorator
}

func NewSugarDecorator(coffee Coffee) *SugarDecorator {
	return &SugarDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (s *SugarDecorator) Cost() int {
	return s.coffee.Cost() + 1 // Add cost of sugar
}

func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + ", Sugar"
}


func main() {
	// Create a simple coffee
	simpleCoffee := &SimpleCoffee{}
	fmt.Printf("%s: $%d\n", simpleCoffee.Description(), simpleCoffee.Cost())

	// Add milk to the coffee
	milkCoffee := NewMilkDecorator(simpleCoffee)
	fmt.Printf("%s: $%d\n", milkCoffee.Description(), milkCoffee.Cost())

	// Add sugar to the coffee with milk
	sugarMilkCoffee := NewSugarDecorator(milkCoffee)
	fmt.Printf("%s: $%d\n", sugarMilkCoffee.Description(), sugarMilkCoffee.Cost())
}
