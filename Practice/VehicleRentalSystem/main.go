/*
 * System Requirements:

	1) System will support the renting of different automobiles like cars, trucks, SUVs, vans, and motorcycles.

	2) Each vehicle should be added with a unique barcode and other details, including a parking stall number which helps to locate the vehicle.

	3) System should be able to retrieve information like which member took a particular vehicle or what vehicles have been rented out by a specific member.

	4) System should collect a late-fee for vehicles returned after the due date.

	5) Members should be able to search the vehicle inventory and reserve any available vehicle.

	6) The system should be able to send notifications whenever the reservation is approaching the pick-up date, as well as when the vehicle is nearing the due date or has not been returned within the due date.

	7) The system will be able to read barcodes from vehicles.

	8) Members should be able to cancel their reservations.

	9) The system should maintain a vehicle log to track all events related to the vehicles.

	10) Members can add rental insurance to their reservation.

	11) Members can rent additional equipment, like navigation, child seat, ski rack, etc.

	12) Members can add additional services to their reservation, such as roadside assistance, additional driver, wifi, etc.
 *
 * */

package main

import (
	"errors"
	"log"
)

type member struct {
	id           int
	vehicle      vehicle
	notification notification
}

func addMember(id int) *member {
	log.Println("Adding member with id:", id)
	return &member{
		id: id,
	}
}

func (d *member) reserveVehicle(vehicle vehicle) {
	log.Println("Reserving a vehicle for member with id:", d.id)
	d.vehicle = vehicle
	d.vehicle.reserveVehicle(d.id)
	d.sendNotification("reserving vehicle" + vehicle.getVehicleNumber() + "for you")
}

func (d *member) cancelReservation() {
	log.Println("Cancelling vehicle reservation for member with id:", d.id)
	d.vehicle.cancelReservation()
	d.sendNotification("cancelling vehicle" + d.vehicle.getVehicleNumber() + "for you")
}

func (d *member) getReservedVehicle() (vehicle, error) {
	if d.vehicle == nil {
		log.Println("Member with id:", d.id, "has no vehicle on rent")
		return nil, errors.New("no rented vehicle for this memeber")
	}
	return d.vehicle, nil
}

func (d *member) sendNotification(message string) {
	notification := notification{
		id:      1,
		message: message,
	}
	notification.sendNotification(d.id)
	d.notification = notification
}

type memberManager struct {
	members []*member
}

func (d *memberManager) init() {
	log.Println("initializing member manager")
	var members []*member
	members = append(members, addMember(1))
	members = append(members, addMember(2))
	members = append(members, addMember(3))
	d.members = members
}

func (d *memberManager) getMember(id int) (*member, error) {
	for _, m := range d.members {
		if m.id == id {
			return m, nil
		}
	}
	log.Println("[SYS] Member with id:", id, "not found")
	return &member{}, errors.New("member not found")
}

type vehicle interface {
	reserveVehicle(int)
	getAvailablity() bool
	getVehicleNumber() string
	getDetails(int)
	getReservation() *reservation
	cancelReservation()
}

type vehicleBase struct {
	parkingStallNumber int
	barCodeString      string
	reservation        reservation
	isAvailable        bool
}

func (d *vehicleBase) reserveVehicle(memberId int) {
	log.Println("Reserving vehicle:", d.barCodeString, "for member with id:", memberId)
	d.reservation = addReservation(d.barCodeString, memberId)
	d.isAvailable = false
}

func (d *vehicleBase) getAvailablity() bool {
	return d.isAvailable
}

func (d *vehicleBase) getVehicleNumber() string {
	return d.barCodeString
}

func (d *vehicleBase) getDetails(memberId int) {
	log.Println("member with Id:", memberId, "has rented vehicle:", d.barCodeString)
}

func (d *vehicleBase) getReservation() *reservation {
	return &d.reservation
}

func (d *vehicleBase) cancelReservation() {
	d.isAvailable = true
}

type car struct {
	vehicleBase
}

func addCar(carNumber string, parkingStallNumber int) vehicle {
	log.Println("Adding car with number:", carNumber)
	return &car{
		vehicleBase{
			parkingStallNumber: parkingStallNumber,
			barCodeString:      carNumber,
			reservation:        reservation{},
			isAvailable:        true,
		},
	}
}

type motocycle struct {
	vehicleBase
}
type suv struct {
	vehicleBase
}
type van struct {
	vehicleBase
}
type truck struct {
	vehicleBase
}

type vehicleInventory struct {
	vehicles []vehicle
}

func (d *vehicleInventory) init() {
	log.Println("Initializing vehicle inventory")
	var vehicles []vehicle
	vehicles = append(vehicles, addCar("BR_14_8900", 1))
	vehicles = append(vehicles, addCar("TN_34_1234", 1))
	d.vehicles = vehicles
}

func (d *vehicleInventory) getAvailableVehicle() (vehicle, error) {
	for _, vi := range d.vehicles {
		if vi.getAvailablity() {
			log.Println("Vehicle with number:", vi.getVehicleNumber(), "is available")
			return vi, nil
		}
	}
	log.Println("No vehicle available")
	return nil, errors.New("no vehicle available")
}

type equipment interface{}

type equipmentBase struct {
}

type navigation struct {
	equipmentBase
}

type childSeat struct {
	equipmentBase
}

type skiRack struct {
	equipmentBase
}

type service interface{}
type serviceBase struct{}

type roadSideAssistance struct {
	serviceBase
}

type additionalDriver struct {
	serviceBase
}

type wifi struct {
	serviceBase
}

type reservation struct {
	id        int
	insurance insurance
	service   service
	equipment equipment
}

func (d *reservation) addInsurance() {
	d.insurance = addInsurance(d.id)
	log.Println("insurance with id:", d.insurance.id, "added for reservationId: ", d.id)
}

func (d *reservation) addService() {

}

func (d *reservation) addEquipment() {

}

type insurance struct {
	id int
}

func addInsurance(reservationId int) insurance {
	return insurance{
		id: reservationId + 1,
	}
}

func addReservation(vehicleNumber string, memberId int) reservation {
	log.Println("Reservation created for vehicle:", vehicleNumber, "for member with Id:", memberId)
	return reservation{
		id: memberId,
	}
}

type notification struct {
	id      int
	message string
}

func (d *notification) sendNotification(memeberId int) {
	log.Println("Notification with message:", d.message, "id:", d.id, "sent to member with id:", memeberId)
}

type rentalSystem struct {
	vehicleInventory *vehicleInventory
	memberManager    *memberManager
}

func (d *rentalSystem) init() {
	log.Println("Initializing rental system")
	vehicleInventory := &vehicleInventory{}
	vehicleInventory.init()

	memberManager := &memberManager{}
	memberManager.init()

	d.vehicleInventory = vehicleInventory
	d.memberManager = memberManager

}

func main() {
	rentalSystem := &rentalSystem{}
	rentalSystem.init()

	member1, err := rentalSystem.memberManager.getMember(1)
	if err == nil {
		vehicle, err := rentalSystem.vehicleInventory.getAvailableVehicle()
		if err == nil {
			member1.reserveVehicle(vehicle)
		}
	}

	member2, err := rentalSystem.memberManager.getMember(2)
	if err == nil {
		vehicle, err := rentalSystem.vehicleInventory.getAvailableVehicle()
		if err == nil {
			member2.reserveVehicle(vehicle)
		}
	}

	member2.cancelReservation()

	member3, err := rentalSystem.memberManager.getMember(3)
	if err == nil {
		vehicle, err := rentalSystem.vehicleInventory.getAvailableVehicle()
		if err == nil {
			member3.reserveVehicle(vehicle)
		}
	}

	member4, err := rentalSystem.memberManager.getMember(4)
	if err == nil {
		vehicle, err := rentalSystem.vehicleInventory.getAvailableVehicle()
		if err == nil {
			member4.reserveVehicle(vehicle)
		}
	}

	vehicle1, err := member1.getReservedVehicle()
	if err == nil {
		vehicle1.getDetails(member1.id)
		vehicle1.getReservation().addInsurance()
	}
}
