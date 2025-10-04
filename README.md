# System Design LLD Real-World Examples

This repository contains practical implementations of common design patterns used in low-level system design (LLD), demonstrated through real-world examples in Go. Each example illustrates how design patterns can be applied to solve everyday programming challenges.

## Design Patterns Included

### 1. Decorator Pattern

**File:** [decoratorDesignPatternExample.go](./decoratorDesignPatternExample.go)

**Real-World Example:** Coffee Shop Billing System

This example demonstrates how to implement a flexible billing system for a coffee shop where:
- Base coffee costs $5
- Customers can add various add-ons (milk, sugar, whipped cream)
- Each add-on has its own price
- Add-ons can be combined in any order and quantity
- New add-ons can be easily introduced without modifying existing code

**When to Use:**
- When you need to add responsibilities to objects dynamically
- When extending functionality through subclassing is impractical
- When you need a flexible alternative to subclassing for extending functionality

### 2. Observer Pattern

**File:** [observerDesignPatternExample.go](./observerDesignPatternExample.go)

**Real-World Example:** Smart Traffic Management System

This example shows how to implement a traffic control system where:
- A central traffic control center maintains live traffic data
- Multiple components (traffic lights, digital road signs, connected cars) observe traffic conditions
- Each component reacts differently to changes in traffic data
- Components can be added or removed at runtime
- The system is extensible for adding new types of observers

**When to Use:**
- When changes to one object require changing others, and you don't know how many objects need to change
- When an object should notify other objects without making assumptions about those objects
- When you need a one-to-many dependency between objects

### 3. Strategy Pattern

**File:** [strategyDesignPatternExample.go](./strategyDesignPatternExample.go)

**Real-World Example:** Food Delivery App

This example demonstrates how to implement a food delivery application where:
- Different delivery strategies can be selected (time, mode)
- Various payment methods are supported
- Strategies can be changed at runtime
- New strategies can be added without modifying existing code

**When to Use:**
- When you want to define a family of algorithms and make them interchangeable
- When you need different variants of an algorithm
- When an algorithm uses data that clients shouldn't know about
- When a class has multiple behaviors that appear as multiple conditional statements

### 4. Factory Pattern

**File:** [factoryDesignPatternExample.go](./factoryDesignPatternExample.go)

**Real-World Example:** Vehicle Creation System

This example demonstrates how to implement a vehicle creation system where:
- Different types of vehicles (Car, Bike, Truck) can be created through a common interface
- Each vehicle has specific behaviors (Drive, MaxSpeed, FuelType)
- Clients can request a vehicle by specifying only its type name
- New vehicle types can be added without modifying client code

**When to Use:**
- When a class can't anticipate the type of objects it must create
- When you want to encapsulate object creation logic in a separate class
- When you need to decouple client code from concrete product classes
- When you want to provide a way to extend the product family easily

### 5. Abstract Factory Pattern

**File:** [abstractFactoryDesignPatternExample.go](./abstractFactoryDesignPatternExample.go)

**Real-World Example:** Cross-Platform UI Toolkit

This example demonstrates how to implement a cross-platform UI toolkit where:
- Different UI components (Button, Checkbox) can be created for different platforms (Desktop, Mobile)
- Each component maintains platform-specific look and behavior
- A single entry point provides access to all components for a given platform
- New platforms can be added without modifying existing client code

**When to Use:**
- When your system needs to be independent from how its products are created
- When you need to work with multiple families of related products
- When you want to provide a library of products without exposing implementation details
- When you need to enforce constraints on product combinations

## How to Run

Each example can be run independently using the Go command:

```bash
go run decoratorDesignPatternExample.go
go run observerDesignPatternExample.go
go run strategyDesignPatternExample.go
go run factoryDesignPatternExample.go
go run abstractFactoryDesignPatternExample.go
```

## Benefits of Design Patterns

- **Reusability:** Proven solutions to common problems
- **Maintainability:** Well-structured code that's easier to understand and modify
- **Scalability:** Easier to extend and adapt as requirements change
- **Communication:** Common vocabulary for discussing design solutions

## Contributing

Feel free to contribute additional design pattern examples or improvements to existing ones. Please follow the established format for consistency.

## License

[MIT License](LICENSE)