/*
*Traffic Control System*

Problem Statement:
Design a Smart Traffic Management System where multiple traffic-related components automatically update their behavior when traffic conditions change.

Requirements:
TrafficControlCenter: Maintains live traffic data for different roads.

The traffic data includes:
type TrafficData struct {
    RoadID    string   // road/section identifier
    Density   int      // traffic density % (0–100)
    AvgSpeed  float64  // average vehicle speed (km/h)
    Accidents int      // reported accidents
    Weather   string   // "Clear", "Rainy", "Foggy"
}

The components which are dependednt on this traffic control system are:

1. TrafficLight
State: "RED", "YELLOW", "GREEN"
Conditions:
- If Density > 80 OR Accidents > 0 → switch to RED (block traffic).
- If Density > 50 AND AvgSpeed < 20 → switch to YELLOW (slow traffic).
- Else → GREEN (flow normal).

2. DigitalRoadSign
State: message displayed on the signboard
Conditions:
- If Accidents > 0 → message: "Accident Ahead, Use Diversion".
- Else if Density > 70 → message: "Heavy Traffic Ahead, Expect Delays".
- Else if Weather == "Foggy" → message: "Drive Carefully, Low Visibility".
- Else → message: "Traffic Normal".

3. ConnectedCar
State: route (current or recalculated)
Conditions:
- If Density > 70 OR Accidents > 0 → recalculate route (show "Rerouting...").
- If Weather == "Rainy" OR Weather == "Foggy" → reduce speed suggestion ("Drive Slow").
- Else → "Continue on current route".

Components can be added or removed at runtime.

The system must be extensible so that adding new components (PolicePatrolUnit, EmergencyServices) requires no changes to the TrafficControlCenter.
*/

package main

import "log"

type TrafficControlCenter interface {
	Subscribe()
	Unsubscribe()
	NotifyAll()
	SetTrafficDensitiy(density int)
	SetAvarageSpeed(avgSpeed float64)
	ReportAccident(accidents int)
	SetWeather(weather string)
}

type TrafficData struct {
	Components []TrafficComponents
	RoadID     string  // road/section identifier
	Density    int     // traffic density % (0–100)
	AvgSpeed   float64 // average vehicle speed (km/h)
	Accidents  int     // reported accidents
	Weather    string  // "Clear", "Rainy", "Foggy"
}

func (d *TrafficData) Subscribe(trafficComponent TrafficComponents) {
	d.Components = append(d.Components, trafficComponent)
	log.Printf("[%s-%s] Added to the system", trafficComponent.GetName(), trafficComponent.GetID())
}

func (d *TrafficData) Unsubscribe(trafficComponent TrafficComponents) {
	componentID := trafficComponent.GetID()
	for i, component := range d.Components {
		if component.GetID() == componentID {
			d.Components = append(d.Components[:i], d.Components[i+1:]...)
			break
		}
	}
	log.Printf("[%s-%s] Removed from the system", trafficComponent.GetName(), trafficComponent.GetID())
}

func (d *TrafficData) NotifyAll() {
	log.Println("[ADMIN] Notifying to all the components")
	for _, component := range d.Components {
		component.Update(d)
	}
}

func (d *TrafficData) SetTrafficDensitiy(density int) {
	log.Println("[ADMIN] Setting traffic density to", density)
	d.Density = density
	d.NotifyAll()
}

func (d *TrafficData) SetAvarageSpeed(avgSpeed float64) {
	log.Println("[ADMIN] Setting avarage speed to", avgSpeed)
	d.AvgSpeed = avgSpeed
	d.NotifyAll()
}

func (d *TrafficData) ReportAccident(accidents int) {
	log.Printf("[ADMIN] %d accident is reported", accidents)
	d.Accidents = accidents
	d.NotifyAll()
}

func (d *TrafficData) SetWeather(weather string) {
	log.Println("[ADMIN] Setting weather to", weather)
	d.Weather = weather
	d.NotifyAll()
}

type TrafficComponents interface {
	Update(trafficData *TrafficData)
	GetName() string
	GetID() string
}

/*
*******************
Traffic Components and their concrete implementations
*******************
*/
type TrafficLight struct {
	ComponentID string
	Name        string
}

func (d *TrafficLight) Update(trafficData *TrafficData) {

	density := trafficData.Density
	avgSpeed := trafficData.AvgSpeed
	accidents := trafficData.Accidents

	if density > 80 || accidents > 0 {
		log.Printf("[%s-%s] Updating Traffic Light to RED", d.Name, d.ComponentID)
		return
	}
	if density > 50 && avgSpeed < 20 {
		log.Printf("[%s-%s] Updating Traffic Light to YELLOW", d.Name, d.ComponentID)
		return
	}
	log.Printf("[%s-%s] Updating Traffic Light to GREEN", d.Name, d.ComponentID)
}

func (d *TrafficLight) GetName() string {
	return d.Name
}

func (d *TrafficLight) GetID() string {
	return d.ComponentID
}

type DigitalRoadSign struct {
	ComponentID string
	Name        string
}

func (d *DigitalRoadSign) Update(trafficData *TrafficData) {

	density := trafficData.Density
	weather := trafficData.Weather
	accidents := trafficData.Accidents

	if accidents > 0 {
		log.Printf("[%s-%s] Accident Ahead, Use Diversion", d.Name, d.ComponentID)
		return
	}
	if density > 70 {
		log.Printf("[%s-%s] Heavy Traffic Ahead, Expect Delays", d.Name, d.ComponentID)
		return
	}
	if weather == "Foggy" {
		log.Printf("[%s-%s] Drive Carefully, Low Visibility", d.Name, d.ComponentID)
		return
	}
	log.Printf("[%s-%s] Traffic Normal", d.Name, d.ComponentID)
}

func (d *DigitalRoadSign) GetName() string {
	return d.Name
}

func (d *DigitalRoadSign) GetID() string {
	return d.ComponentID
}

type ConnectedCar struct {
	ComponentID string
	Name        string
}

func (d *ConnectedCar) Update(trafficData *TrafficData) {

	density := trafficData.Density
	weather := trafficData.Weather
	accidents := trafficData.Accidents

	if density > 70 || accidents > 0 {
		log.Printf("[%s-%s] Rerouting...", d.Name, d.ComponentID)
		return
	}
	if weather == "Foggy" || weather == "Rainy" {
		log.Printf("[%s-%s] Drive Slow", d.Name, d.ComponentID)
		return
	}
	log.Printf("[%s-%s] Continue on current route", d.Name, d.ComponentID)
}

func (d *ConnectedCar) GetName() string {
	return d.Name
}

func (d *ConnectedCar) GetID() string {
	return d.ComponentID
}

func main() {
	// Initialize subject
	trafficData := &TrafficData{}

	// Attach observers
	trafficLight := &TrafficLight{ComponentID: "TL1", Name: "traffic_light"}
	digitalRoadSign := &DigitalRoadSign{ComponentID: "DRS1", Name: "digital_road_sign"}
	connectedCar := &ConnectedCar{ComponentID: "CC1", Name: "connected_car"}

	trafficData.Subscribe(trafficLight)
	trafficData.Subscribe(digitalRoadSign)
	trafficData.Subscribe(connectedCar)

	// === Test Case 1: Normal traffic ===
	println("\n--- Test Case 1: Normal Traffic ---")
	trafficData.SetAvarageSpeed(55)
	trafficData.SetTrafficDensitiy(30)
	trafficData.ReportAccident(0)
	trafficData.SetWeather("Clear")

	// === Test Case 2: Heavy traffic ===
	println("\n--- Test Case 2: Heavy Traffic ---")
	trafficData.SetAvarageSpeed(15)
	trafficData.SetTrafficDensitiy(85)

	// === Test Case 3: Accident on Road ===
	println("\n--- Test Case 3: Accident ---")
	trafficData.ReportAccident(2)
	trafficData.SetAvarageSpeed(0)
	trafficData.SetTrafficDensitiy(90)
	trafficData.Unsubscribe(digitalRoadSign)

	// === Test Case 4: Foggy Weather ===
	println("\n--- Test Case 4: Foggy Weather ---")
	trafficData.SetWeather("Foggy")
	trafficData.SetAvarageSpeed(40)
	trafficData.SetTrafficDensitiy(40)

	// === Test Case 5: Unsubscribing ConnectedCar ===
	println("\n--- Test Case 5: Unsubscribing ConnectedCar ---")
	trafficData.Unsubscribe(connectedCar)
	trafficData.ReportAccident(1)
	trafficData.SetWeather("Rainy")
	trafficData.SetTrafficDensitiy(75)
	trafficData.SetAvarageSpeed(20)
}
