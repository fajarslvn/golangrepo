package main

import (
	"fmt"
)

type LatLong struct {
	Lat  float64
	Long float64
}

type Animal interface {
	GetLocation() LatLong
	SetLocation(loc LatLong)
	CanFly() bool
	Speak() string
	GetName() string
}

type Lion struct {
	name       string
	maneLength int
	location   LatLong
}

func (l *Lion) GetLocation() LatLong {
	return l.location
}

func (l *Lion) SetLocation(loc LatLong) {
	l.location = loc
}

func (l *Lion) CanFly() bool {
	return false
}

func (l *Lion) Speak() string {
	return "Roaaar!"
}

func (l *Lion) GetManeLength() int {
	return l.maneLength
}

func (l *Lion) GetName() string {
	return l.name
}

type Pigeon struct {
	name     string
	location LatLong
}

func (p *Pigeon) GetLocation() LatLong {
	return p.location
}

func (p *Pigeon) SetLocation(loc LatLong) {
	p.location = loc
}

func (p *Pigeon) CanFly() bool {
	return true
}

func (p *Pigeon) Speak() string {
	return "KuKuk!"
}

func (p *Pigeon) GetName() string {
	return p.name
}

func MakeThemTalk(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.GetName() + " says " + animal.Speak())
	}
}

func main() {
	var myZoo []Animal
	Leo := Lion{
		"Leo",
		10,
		LatLong{10.40, 11.5},
	}
	myZoo = append(myZoo, &Leo)

	Tweetie := Pigeon{
		"Tweetie",
		LatLong{10.40, 11.5},
	}

	myZoo = append(myZoo, &Tweetie)
	MakeThemTalk(myZoo)
}
