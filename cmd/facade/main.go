package main

import "fmt"

// Subsystem: DVD Player
type DvdPlayer struct{}

func (d *DvdPlayer) On() {
	fmt.Println("DVD Player is On")
}

func (d *DvdPlayer) Play(movie string) {
	fmt.Printf("Playing movie: %s\n", movie)
}

func (d *DvdPlayer) Off() {
	fmt.Println("DVD Player is Off")
}

// Subsystem: Amplifier
type Amplifier struct{}

func (a *Amplifier) On() {
	fmt.Println("Amplifier is On")
}

func (a *Amplifier) SetVolume(volume int) {
	fmt.Printf("Setting amplifier volume to %d\n", volume)
}

func (a *Amplifier) Off() {
	fmt.Println("Amplifier is Off")
}

// Subsystem: Projector
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Projector is On")
}

func (p *Projector) Off() {
	fmt.Println("Projector is Off")
}


// Facade: Home Theater Facade
type HomeTheaterFacade struct {
	dvdPlayer  *DvdPlayer
	amplifier  *Amplifier
	projector  *Projector
}

func NewHomeTheaterFacade(dvd *DvdPlayer, amp *Amplifier, proj *Projector) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		dvdPlayer:  dvd,
		amplifier:  amp,
		projector:  proj,
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Starting movie...")
	h.projector.On()
	h.amplifier.On()
	h.amplifier.SetVolume(10)
	h.dvdPlayer.On()
	h.dvdPlayer.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting down movie theater...")
	h.dvdPlayer.Off()
	h.amplifier.Off()
	h.projector.Off()
}


func main() {
	// Subsystem components
	dvd := &DvdPlayer{}
	amp := &Amplifier{}
	projector := &Projector{}

	// Facade
	homeTheater := NewHomeTheaterFacade(dvd, amp, projector)

	// Use the facade to control the subsystem
	homeTheater.WatchMovie("Inception")
	fmt.Println()
	homeTheater.EndMovie()
}
