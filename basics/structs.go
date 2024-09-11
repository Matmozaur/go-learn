package main

import (
	"fmt"
	"time"
)

type Location struct {
	Lat float64
	Lng float64
}

func NewLocation(lat, lng float64) (Location, error) {
	if lat < -90 || lat > 90 {
		return Location{}, fmt.Errorf("invalid lat: %#v", lat)
	}
	if lng < -180 || lng > 180 {
		return Location{}, fmt.Errorf("invalid lng: %#v", lng)
	}

	loc := Location{
		Lat: lat,
		Lng: lng,
	}
	return loc, nil
}

func (l *Location) Move(lat, lng float64) {
	l.Lat = lat
	l.Lng = lng
}

// Embeddings
type Car struct {
	ID string
	Location
}

func NewCar(id string, lat, lng float64) (Car, error) {
	loc, err := NewLocation(lat, lng)
	if err != nil {
		return Car{}, err
	}

	car := Car{
		ID:       id,
		Location: loc,
	}
	return car, nil
}

// Interface
type Mover interface {
	Move(float64, float64)
}

func moveAll(items []Mover, lat, lng float64) {
	for _, item := range items {
		item.Move(lat, lng)
	}
}

// generics
func Add[T int | float64 | string](a, b T) T {
	return a + b
}

// challenge
const PerHour = 5.0

type VM struct {
	startTime time.Time
	endTime   time.Time
}

func NewVM(startTime, endTime time.Time) (VM, error) {
	if endTime.Before(startTime) {
		return VM{}, fmt.Errorf("start time: %#v later than end time: %#v", startTime, endTime)
	}
	vm := VM{
		startTime: startTime,
		endTime:   endTime,
	}
	return vm, nil
}

func (vm *VM) Cost() float64 {
	return PerHour * float64(vm.endTime.Hour()-vm.startTime.Hour())
}

type Spot struct {
	VM
}

func (s Spot) Cost() float64 {
	p := s.VM.Cost()
	return p * 0.8
}

type Costable interface {
	Cost() float64
}

func TotalCost(vms []Costable) float64 {
	total := 0.0
	for _, vm := range vms {
		total += vm.Cost()
	}
	return total
}

func main() {
	loc, err := NewLocation(32.5253837, 34.9427434)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(loc)
	fmt.Println("------------------")

	loc.Move(0, 0)
	fmt.Printf("%#v\n", loc)
	fmt.Println("------------------")

	car, err := NewCar("g0ph3r", 32.5253837, 34.9427434)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	car.Move(32.0641339, 34.8742343)
	fmt.Printf("%#v\n", car)
	fmt.Println("------------------")

	items := []Mover{
		&Location{32.3477669, 34.9160405},
		&Car{
			ID: "g0ph3r",
			Location: Location{
				Lat: 32.5253837,
				Lng: 34.9427434,
			},
		},
	}
	moveAll(items, 32.0641339, 34.8742343)
	for _, item := range items {
		fmt.Printf("%#v\n", item)

	}
	fmt.Println("------------------")

	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.0, 2.0))
	fmt.Println(Add("G", "o"))
	fmt.Println("------------------")

	vms := []Costable{
		&VM{
			startTime: time.Date(2022, time.April, 12, 17, 30, 0, 0, time.UTC),
			endTime:   time.Date(2022, time.April, 12, 19, 54, 0, 0, time.UTC),
		},
		Spot{
			VM: VM{
				startTime: time.Date(2022, time.April, 13, 9, 46, 0, 0, time.UTC),
				endTime:   time.Date(2022, time.April, 15, 12, 7, 0, 0, time.UTC),
			},
		},
	}
	fmt.Println(TotalCost(vms))

}
