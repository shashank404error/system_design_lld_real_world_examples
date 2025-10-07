/*
https://workat.tech/machine-coding/practice/design-parking-lot-qm6hwq4wkhp8
*/

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type ParkingLotSystem interface {
	CreateParkingLot()
	ParkVehicle()
	UnparkVehicle()
	Display()
	Exit()
}

type ParkingLot struct {
	ID                string
	FloorCount        int
	SlotCountPerFloor int
	FloorToSlotArray  [][]*ParkingLotSlot
	TicketToSlotMap   map[string]ParkingLotSlot
}

type ParkingLotSlot struct {
	Floor        *int
	Slot         *int
	VehicleType  *string
	Availibility *bool
}

type Ticket struct {
	ID string
}

func (d *ParkingLot) CreateParkingLot(parkingLotID string, floorCount, slotCountPerFloor int) {
	log.Println("Created parking lot with ", floorCount, " floors and ", slotCountPerFloor, " slots per floor")
	d.ID = parkingLotID
	d.FloorCount = floorCount
	d.SlotCountPerFloor = slotCountPerFloor
	d.CreateParkingLotSlot()

	for i, floor := range d.FloorToSlotArray {
		for j, slot := range floor {
			fmt.Printf("FloorToSlotArray[%d][%d]:{%d,%d,%s,%v}\n", i+1, j+1, *slot.Floor, *slot.Slot, *slot.VehicleType, *slot.Availibility)
		}
	}
}

func (d *ParkingLot) ParkVehicle(vehicleType, regNo, color string) *Ticket {
	floor, slot, isSlotAvailable := d.FindFirstAvailableSlot(vehicleType, regNo, color)
	if !isSlotAvailable {
		log.Println("Parking Lot Full")
		return nil
	}

	isBookingDone := d.BookAvailableSlot(floor, slot)
	if isBookingDone {
		ticket := d.CreateTicket(floor, slot)
		log.Println("Vehicle parked at floor:", floor, "slot:", slot, "TicketID:", ticket.ID)
		return ticket
	}

	log.Println("Some error in parking vehicle")
	return nil
}

func (d *ParkingLot) UnparkVehicle(ticketID string) {
	if slot, ok := d.TicketToSlotMap[ticketID]; ok {
		if *d.FloorToSlotArray[*slot.Floor-1][*slot.Slot-1].Availibility {
			log.Println("Invalid Ticket")
			return
		}
		availibility := true
		d.FloorToSlotArray[*slot.Floor-1][*slot.Slot-1].Availibility = &availibility
		log.Println("Unparked the vehicle from floor:", *slot.Floor, "slot:", *slot.Slot)
		return
	}
	log.Println("Invalid Ticket")
}

func (d *ParkingLot) Display(displayType, vehicleType string) {
	var freeSlotsPerFloor []int
	var floorToFreeSlotsArray [][]int
	var floorToOccupiedSlotsArray [][]int

	for _, floor := range d.FloorToSlotArray {
		totalFreeSlots := 0
		var freeSlots []int
		var occupiedSlots []int
		for j, slot := range floor {

			if *slot.VehicleType == vehicleType {
				if *slot.Availibility {
					totalFreeSlots++
					freeSlots = append(freeSlots, j+1)
				} else {
					occupiedSlots = append(occupiedSlots, j+1)
				}
			}

		}
		freeSlotsPerFloor = append(freeSlotsPerFloor, totalFreeSlots)
		floorToFreeSlotsArray = append(floorToFreeSlotsArray, freeSlots)
		floorToOccupiedSlotsArray = append(floorToOccupiedSlotsArray, occupiedSlots)
	}

	if displayType == "free_count" {
		for i, s := range freeSlotsPerFloor {
			log.Println("No. of free slots for ", vehicleType, " on Floor ", i+1, ":", s)
		}
	}

	if displayType == "free_slots" {
		for i, slots := range floorToFreeSlotsArray {
			log.Println("Free slots for", vehicleType, "on Floor ", i+1, ": ", slots)
		}
	}

	if displayType == "occupied_slots" {
		for i, slots := range floorToOccupiedSlotsArray {
			log.Println("Occupied slots for", vehicleType, "on Floor ", i+1, ": ", slots)
		}
	}

}
func (d *ParkingLot) Exit() {
	log.Fatal()
}

func (d *ParkingLot) CreateParkingLotSlot() {
	var floors [][]*ParkingLotSlot
	floorCount := d.FloorCount
	slotCount := d.SlotCountPerFloor

	for i := 0; i < floorCount; i++ {
		var parkingLotSlots []*ParkingLotSlot
		for j := 0; j < slotCount; j++ {
			vehicleType := "CAR"
			if j == 0 {
				vehicleType = "TRUCK"
			}
			if j == 1 || j == 2 {
				vehicleType = "BIKE"
			}
			floor := i + 1
			slot := j + 1
			availibility := true
			parkingLotSlot := ParkingLotSlot{
				Floor:        &floor,
				Slot:         &slot,
				VehicleType:  &vehicleType,
				Availibility: &availibility,
			}
			parkingLotSlots = append(parkingLotSlots, &parkingLotSlot)

		}
		floors = append(floors, parkingLotSlots)
	}
	d.FloorToSlotArray = floors
}

func (d *ParkingLot) FindFirstAvailableSlot(vehilceType, regNo, color string) (int, int, bool) {
	for floorIndex, floor := range d.FloorToSlotArray {
		for slotIndex, slot := range floor {
			if *slot.Availibility && *slot.VehicleType == vehilceType {
				return floorIndex + 1, slotIndex + 1, true
			}
		}
	}
	return -1, -1, false
}

func (d *ParkingLot) BookAvailableSlot(floor, slot int) bool {
	if floor == -1 || slot == -1 {
		return false
	}
	availibility := false
	d.FloorToSlotArray[floor-1][slot-1].Availibility = &availibility
	return true
}

func (d *ParkingLot) CreateTicket(floor, slot int) *Ticket {
	ticketID := d.ID + "_" + strconv.Itoa(floor) + "_" + strconv.Itoa(slot)
	d.TicketToSlotMap[ticketID] = ParkingLotSlot{
		Floor: &floor,
		Slot:  &slot,
	}
	return &Ticket{
		ID: ticketID,
	}
}

// ParkingLotFloor
// Ticket
// Vehicle

func main() {
	commands := GetInput()

	ticketToSlotMap := make(map[string]ParkingLotSlot)
	parkingLot := &ParkingLot{
		TicketToSlotMap: ticketToSlotMap,
	}

	for _, command := range commands {
		commandArr := strings.Split(command, " ")
		switch commandArr[0] {
		case "create_parking_lot":
			floorCount, _ := strconv.Atoi(commandArr[2])
			slotCountPerFloor, _ := strconv.Atoi(commandArr[3])
			parkingLot.CreateParkingLot(commandArr[1], floorCount, slotCountPerFloor)
		case "park_vehicle":
			parkingLot.ParkVehicle(commandArr[1], commandArr[2], commandArr[3])
		case "unpark_vehicle":
			parkingLot.UnparkVehicle(commandArr[1])
		case "display":
			parkingLot.Display(commandArr[1], commandArr[2])
		case "exit":
			parkingLot.Exit()
		}
	}

}
