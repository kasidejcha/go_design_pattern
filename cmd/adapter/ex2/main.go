package main

import "fmt"

// Adaptee Interface: This represents different types of sockets
type Socket interface {
	PlugIn()
}

// Concrete Adaptee 1: EuropeanSocket implements the Socket interface
type EuropeanSocket struct{}

func (e *EuropeanSocket) PlugIn() {
	fmt.Println("Plugged into European socket: 220V")
}

// Concrete Adaptee 2: USASocket implements the Socket interface
type USASocket struct{}

func (u *USASocket) PlugIn() {
	fmt.Println("Plugged into USA socket: 110V")
}

// Adapter: This adapts any Socket to the Charger interface
type SocketAdapter struct {
	socket Socket
}

// Adapter's method implements the Charger interface
func (a *SocketAdapter) Charge() {
	fmt.Println("Adapter converting socket power...")
	a.socket.PlugIn()
}


// Client code: Uses the Charger interface
func main() {
	// Create an instance of EuropeanSocket
	europeanSocket := &EuropeanSocket{}

	// Create an adapter for the EuropeanSocket that fits the Charger interface
	adapter1 := &SocketAdapter{
		socket: europeanSocket,
	}

	// Use the adapter as a Charger
	adapter1.Charge()

	// Create an instance of USASocket
	usaSocket := &USASocket{}

	// Create an adapter for the USASocket
	adapter2 := &SocketAdapter{
		socket: usaSocket,
	}

	// Use the adapter as a Charger
	adapter2.Charge()

}
