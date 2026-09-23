package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

var employees []Employee

func addEmployee(id int, name string, salary float64) {
	employee := Employee{
		ID:     id,
		Name:   name,
		Salary: salary,
	}

	employees = append(employees, employee)
	fmt.Println("Employee added successfully")
}

func displayEmployees() {
	fmt.Println("\nEmployee List:")

	for _, employee := range employees {
		fmt.Printf(
			"ID: %d | Name: %s | Salary: %.2f\n",
			employee.ID,
			employee.Name,
			employee.Salary,
		)
	}
}

func searchEmployee(id int) {
	for _, employee := range employees {
		if employee.ID == id {
			fmt.Printf(
				"Employee Found: %d | %s | %.2f\n",
				employee.ID,
				employee.Name,
				employee.Salary,
			)
			return
		}
	}

	fmt.Println("Employee not found")
}

func deleteEmployee(id int) {
	for i, employee := range employees {
		if employee.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted successfully")
			return
		}
	}

	fmt.Println("Employee not found")
}

func updateEmployee(id int, newSalary float64) {
	for i := range employees {
		if employees[i].ID == id {
			employees[i].Salary = newSalary
			fmt.Println("Employee updated successfully")
			return
		}
	}

	fmt.Println("Employee not found")
}

func main() {
	addEmployee(101, "Sanskriti", 45000)
	addEmployee(102, "Rahul", 50000)

	displayEmployees()

	searchEmployee(101)

	deleteEmployee(102)

	displayEmployees()
}
