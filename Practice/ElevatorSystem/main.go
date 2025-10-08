/*
Question - https://www.geeksforgeeks.org/system-design/elevator-system-low-level-design-lld/
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type building struct {
	numberOfFloors int
	floors         []floor
}

func (d *building) initialise() {
	fmt.Println("Intializing building with", d.numberOfFloors, "floors")

	esc := &elevatorSystemControl{numberOfElevator: 3}
	esc.initialise()

	var floors []floor
	for i := 0; i < d.numberOfFloors; i++ {
		f := floor{
			floorNumber: i + 1,
			panel: outsideControlPanel{
				system: esc,
			},
		}
		floors = append(floors, f)
	}
	d.floors = floors

}

func (d *building) getFloorPanel(floorNumber int) outsideControlPanel {
	if floorNumber > d.numberOfFloors {
		fmt.Println("Number of floor is more than the actual floor in the building")
	}
	return d.floors[floorNumber-1].panel
}

type floor struct {
	floorNumber int
	panel       outsideControlPanel
}

type outsideControlPanel struct {
	display externalDisplay
	system  *elevatorSystemControl
}

func (d *outsideControlPanel) goDown(currentFloor int) insideControlPanel {
	fmt.Println("Outside Control Panel of floor:", currentFloor, " is used to go up")
	return d.system.callElevator(currentFloor, "down")
}

func (d *outsideControlPanel) goUp(currentFloor int) insideControlPanel {
	fmt.Println("Outside Control Panel of floor:", currentFloor, " is used to go down")
	return d.system.callElevator(currentFloor, "up")

}

type externalDisplay struct {
	currentFloor      int
	movementDirection string
}

type elevatorSystemControl struct {
	numberOfElevator int
	elevators        []*elevator
	mu               sync.Mutex
}

func (d *elevatorSystemControl) initialise() {
	fmt.Println("Intializing elevator control system with", d.numberOfElevator, "elevators")
	var elevators []*elevator
	for i := 0; i < d.numberOfElevator; i++ {
		e := elevator{
			elevatorNumber: i + 1,
			capacityInKG:   15,
			cpacityInPax:   8,
			panel: insideControlPanel{
				elevatorNumber: i + 1,
				system:         d,
			},
			door:         door{},
			status:       "IDLE",
			currentFloor: 15,
		}
		elevators = append(elevators, &e)
	}
	d.elevators = elevators
}

func (d *elevatorSystemControl) callElevator(currentFloor int, direction string) insideControlPanel {

	d.mu.Lock()
	defer d.mu.Unlock()

	// finding which elevator is idle
	floorsToMove := 0
	for _, e := range d.elevators {
		if e.status == "IDLE" {

			ef := e.getCurrentFloor()
			floorsToMove = ef - currentFloor
			if ef < currentFloor {
				floorsToMove = currentFloor - ef
			}
			e.moveToFloor(floorsToMove, direction)
			e.door.openDoor(e.elevatorNumber)
			return e.getInsideControlPanel()
		}
	}
	return insideControlPanel{}

}

func (d *elevatorSystemControl) goToFloor(elevatorNumber, destinationFloor int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	fmt.Println("elevator control system is moving elevator", elevatorNumber, "to floor", destinationFloor)

	e := d.elevators[elevatorNumber-1]
	getCurrentFloor := e.getCurrentFloor()
	direction := "up"
	floorsToMove := destinationFloor - getCurrentFloor
	if getCurrentFloor > destinationFloor {
		direction = "down"
		floorsToMove = getCurrentFloor - destinationFloor
	}

	e.moveToFloor(floorsToMove, direction)
	e.door.openDoor(elevatorNumber)
	e.status = "IDLE"
}

type elevator struct {
	elevatorNumber int
	capacityInKG   int
	cpacityInPax   int
	door           door
	status         string
	currentFloor   int
	panel          insideControlPanel
}

func (d *elevator) moveToFloor(floorsToMove int, directionToMove string) {
	d.status = "MOVING"
	for i := 0; i < floorsToMove; i++ {
		time.Sleep(300 * time.Millisecond) // simulate movement per floor
		if directionToMove == "down" {
			d.currentFloor--
		} else {
			d.currentFloor++
		}
		fmt.Printf("Elevator %d moving %s → floor %d\n", d.elevatorNumber, directionToMove, d.currentFloor)
	}

	// fmt.Println("Moving elevator:", d.elevatorNumber, directionToMove, "by", floorsToMove, "floor.")

}

func (d *elevator) getCurrentFloor() int {
	return d.currentFloor
}

func (d *elevator) getInsideControlPanel() insideControlPanel {
	return d.panel
}

type insideControlPanel struct {
	elevatorNumber int
	display        internalDisplay
	system         *elevatorSystemControl
}

func (d *insideControlPanel) goToFloor(destinationFloor int) {
	// go to a particular floor
	d.system.goToFloor(d.elevatorNumber, destinationFloor)

}

func (d *insideControlPanel) closeDoor() {
	d.system.elevators[d.elevatorNumber-1].door.closeDoor(d.elevatorNumber)
}

type internalDisplay struct {
	currentFloor      int
	movementDirection string
	capacityInKG      int
	capacityInPax     int
}

type door struct{}

func (d *door) openDoor(elevatorNumber int) {
	fmt.Println("Opening door of elevator", elevatorNumber)

}

func (d *door) closeDoor(elevatorNumber int) {
	fmt.Println("Closing door of elevator", elevatorNumber)
}

func main() {
	building := &building{numberOfFloors: 15}
	building.initialise()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		currentFloor := 2
		panel := building.getFloorPanel(currentFloor)
		insidePanel := panel.goUp(currentFloor)
		destinationFloor := 10
		insidePanel.closeDoor()
		insidePanel.goToFloor(destinationFloor)
	}()

	go func() {
		defer wg.Done()
		currentFloor := 15
		panel := building.getFloorPanel(currentFloor)
		insidePanel := panel.goDown(currentFloor)

		destinationFloor := 1
		insidePanel.closeDoor()
		insidePanel.goToFloor(destinationFloor)
	}()

	wg.Wait()
	fmt.Println("All elevators have completed movement.")

}
