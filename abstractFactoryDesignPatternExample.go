/*
Design a Cross-Platform UI Toolkit:

# Problem Statement:
Your team is building a cross-platform design system that supports multiple environments — such as Desktop and Mobile.
The system needs to ensure that all UI components for a given platform share the same look and behavior, while keeping the code easily extensible for future platforms (like Web or TV).

You are asked to design a component creation system that can generate platform-specific UI components such as Buttons and Checkboxes, while allowing the rest of the codebase to remain independent of the platform type.

# Detailed Requirements:

- The system must support two platforms initially:
	- Desktop
	- Mobile

- Each platform will have its own look and behavior for two UI components:
	- Button
	- Checkbox

- Each component should expose the following methods:
	- Render() — displays the component visually (for example, logs “Rendering desktop button” or “Rendering mobile checkbox”)
	- Click() — performs the click action appropriate for that platform (for example, logs “Desktop button clicked”)

- There must be a single entry point (like UIProvider) that the client can use to:
- Request a Button and a Checkbox for a given platform type ("desktop" or "mobile").
- The client should not directly create or know about any DesktopButton, MobileCheckbox, etc.
- The system should be easy to extend to new platforms (e.g., adding WebButton and WebCheckbox) without modifying existing logic that uses Buttons and Checkboxes.

Example client usage:
ui := GetUIComponents("mobile")
button := ui.CreateButton()
checkbox := ui.CreateCheckbox()

button.Render()
button.Click()
checkbox.Render()
checkbox.Click()


Example console output for "desktop" platform:
Rendering a desktop button
Desktop button clicked
Rendering a desktop checkbox
Desktop checkbox checked
*/

package main

import "log"

type UIComponents interface {
	Render()
	Click()
}

type Button struct {
	ProviderName string
}

func (d *Button) Render() {
	log.Println(d.ProviderName, "button is rendered")
}

func (d *Button) Click() {
	log.Println(d.ProviderName, "button is clicked")
}

type Checkbox struct {
	ProviderName string
}

func (d *Checkbox) Render() {
	log.Println(d.ProviderName, "checkbox is rendered")
}

func (d *Checkbox) Click() {
	log.Println(d.ProviderName, "checkbox is clicked")
}

type UIProviders interface {
	CreateButton() UIComponents
	CreateCheckbox() UIComponents
}

type Mobile struct{}

func (d *Mobile) CreateButton() UIComponents {
	return &Button{
		ProviderName: "Mobile",
	}
}

func (d *Mobile) CreateCheckbox() UIComponents {
	return &Checkbox{
		ProviderName: "Mobile",
	}
}

type Desktop struct{}

func (d *Desktop) CreateButton() UIComponents {
	return &Button{
		ProviderName: "Desktop",
	}
}

func (d *Desktop) CreateCheckbox() UIComponents {
	return &Checkbox{
		ProviderName: "Desktop",
	}
}

func GetUIComponents(device string) UIProviders {
	switch device {
	case "mobile":
		return &Mobile{}
	case "desktop":
		return &Desktop{}
	default:
		return nil
	}
}

func main() {
	ui := GetUIComponents("desktop")
	button := ui.CreateButton()
	checkbox := ui.CreateCheckbox()

	button.Render()
	button.Click()
	checkbox.Render()
	checkbox.Click()
}
