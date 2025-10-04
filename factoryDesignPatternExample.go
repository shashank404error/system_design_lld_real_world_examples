/*
Design a Vehicle Creation System

# Problem Statement
You are asked to design a Vehicle Creation System for a transportation company.
The company needs a central module that can create different types of vehicles — such as Car, Bike, and Truck — based on user input.

The system should allow clients (other parts of the application) to request a vehicle object by specifying only its type name (like "car" or "bike") — without having to know how that vehicle object is created internally.

# Detailed Requirements:

- The system must support three types of vehicles initially:

1. Car
2. Bike
3. Truck

- Each vehicle should expose the following behaviors through a common interface:
	- Drive() → prints or logs how that vehicle moves (e.g., "Driving a car", "Riding a bike")
	- MaxSpeed() → returns the vehicle’s maximum speed in km/h (integer)
	- FuelType() → returns the type of fuel it uses (e.g., "Petrol", "Diesel", "Electric")

- The client code should:
	- Ask the system to create a vehicle of a given type ("car", "bike", or "truck").

- The design should allow easy addition of new vehicle types in the future (e.g., Bus, ElectricScooter) without modifying existing client code.
*/

package main

import "log"

type Vehicle interface {
	Drive()
	MaxSpeed() float64
	FuelType() string
}

type Car struct{}

func (d *Car) Drive() {
	log.Println("Driving a car")
}

func (d *Car) MaxSpeed() float64 {
	return 180.00
}

func (d *Car) FuelType() string {
	return "Petrol"
}

type Bike struct{}

func (d *Bike) Drive() {
	log.Println("Driving a bike")
}

func (d *Bike) MaxSpeed() float64 {
	return 150.00
}

func (d *Bike) FuelType() string {
	return "Electric"
}

type Truck struct{}

func (d *Truck) Drive() {
	log.Println("Driving a truck")
}

func (d *Truck) MaxSpeed() float64 {
	return 120.00
}

func (d *Truck) FuelType() string {
	return "Diesel"
}

type vehicleCreatorFunction func() Vehicle
type VehicleFactory struct {
	vehicleRegistry map[string]vehicleCreatorFunction
}

func (d *VehicleFactory) CreateVehicleFactory() {
	d.vehicleRegistry = map[string]vehicleCreatorFunction{
		"car":   func() Vehicle { return &Car{} },
		"bike":  func() Vehicle { return &Bike{} },
		"truck": func() Vehicle { return &Truck{} },
	}
}

// Vehicle factory must be created before calling this method
func (d *VehicleFactory) CreateVehicle(vehicleString string) Vehicle {
	if vehicle, ok := d.vehicleRegistry[vehicleString]; ok {
		return vehicle()
	}
	return nil
}

func main() {
	vf := &VehicleFactory{}
	vf.CreateVehicleFactory()

	vehicle := vf.CreateVehicle("bike")
	vehicle.Drive()
	log.Println(vehicle.MaxSpeed())
	log.Println(vehicle.FuelType())

	vehicle = vf.CreateVehicle("car")
	vehicle.Drive()
	log.Println(vehicle.MaxSpeed())
	log.Println(vehicle.FuelType())

	vehicle = vf.CreateVehicle("truck")
	vehicle.Drive()
	log.Println(vehicle.MaxSpeed())
	log.Println(vehicle.FuelType())

}
