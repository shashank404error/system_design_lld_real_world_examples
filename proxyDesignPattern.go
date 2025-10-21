package main

import "fmt"

type DBOps interface {
	Create(string) int
	Read(int) string
	Update(int, string) bool
	Delete(int) bool
}

type Connector struct {
	name []string
}

func (d *Connector) Create(name string) int {
	d.name = append(d.name, name)
	id := len(d.name) - 1
	fmt.Printf("Creating user with name: %s, id: %d\n", name, id)
	return id
}

func (d *Connector) Read(id int) string {
	if id > len(d.name)-1 {
		fmt.Printf("id: %d donot exits\n", id)
		return ""
	}
	name := d.name[id]
	fmt.Printf("user found. Name: %s, id: %d\n", name, id)
	return name
}

func (d *Connector) Update(id int, name string) bool {
	if id > len(d.name)-1 {
		fmt.Printf("id: %d donot exits\n", id)
		return false
	}
	fmt.Printf("updating user with id: %d\n", id)
	d.name[id] = name
	return true
}

func (d *Connector) Delete(id int) bool {
	if id > len(d.name)-1 {
		fmt.Printf("id: %d donot exits\n", id)
		return false
	}
	if id == len(d.name) {
		d.name = d.name[:id]
		return true
	}
	d.name = append(d.name[:id], d.name[id+1:len(d.name)]...)
	fmt.Printf("deleting user with id: %d\n", id)
	return true
}

type ProxyConnector struct {
	role      string // admin,user
	connector Connector
}

func (d *ProxyConnector) Create(name string) int {
	if isRoleAllowed(d.role) {
		return d.connector.Create(name)
	}
	fmt.Printf("access denied\n")
	return -1
}

func (d *ProxyConnector) Read(id int) string {
	if isRoleAllowed(d.role) {
		return d.connector.Read(id)
	}
	fmt.Printf("access denied\n")
	return ""
}

func (d *ProxyConnector) Update(id int, name string) bool {
	if isRoleAllowed(d.role) {
		return d.connector.Update(id, name)
	}
	fmt.Printf("access denied\n")
	return false
}

func (d *ProxyConnector) Delete(id int) bool {
	if isRoleAllowed(d.role) {
		return d.connector.Delete(id)
	}
	fmt.Printf("access denied\n")
	return false
}

func isRoleAllowed(role string) bool {
	if role != "admin" {
		return false
	}
	return true
}

func main() {
	connector := &ProxyConnector{role: "user"} // run with role: "admin" for it to work
	connector.Create("Shashank")
	connector.Create("Prakash")
	connector.Create("Sharma")
	connector.Read(1)
	connector.Update(0, "Random")
	connector.Delete(2)
	connector.Read(2)
}
