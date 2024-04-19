# Builder Pattern Demonstration in Go

## Overview
This repository demonstrates the use of the Builder design pattern implemented in Go. The project illustrates how to manage complex object construction with multiple optional configurations and showcases robust error handling within the builder processes. The primary focus is on building customizable pizzas, offering both standard and gourmet options.

## Pattern Description
The Builder Pattern separates the construction of a complex object from its representation, allowing the same construction process to create different representations. This is particularly useful when an object requires multiple parts to be constructed across several steps or if its construction requires an extensive list of parameters. In this project, `PizzaBuilder` and `GourmetPizzaBuilder` are used to construct `Pizza` objects with various properties such as crust type, size, and toppings.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating the use of the Builder pattern to construct pizzas.
- **pkg/**
    - **builder/**: Includes the interfaces and concrete implementations for the pizza builders.
    - **product/**: Contains the `Pizza` class definition.
    - **config/**: (Optional) Contains predefined configurations or recipes for building standard or gourmet pizzas.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/arthurfp/Go_Builder_Pattern.git
cd Go_Builder_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp