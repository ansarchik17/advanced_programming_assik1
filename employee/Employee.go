package employee

import (
	"fmt"
)

type FullTime struct {
	MonthlySalary float64
	BonusRate     float64
}

type PartTime struct {
	HourlyRate  float64
	HoursWorked int
}

type Contractor struct {
	ProjectRate       float64
	ProjectsCompleted int
}

type Intern struct {
	DailyRate  float64
	DaysWorked int
}

type Employee struct {
	Intern     []Intern
	FullTime   []FullTime
	PartTime   []PartTime
	Contractor []Contractor
}

type SalaryCalculator interface {
	CalculateSalary() float64
	CalculateBonus() float64
}

// FullTime
func (fullTime FullTime) CalculateSalary() float64 {
	return fullTime.MonthlySalary + fullTime.CalculateBonus()
}

func (fullTime FullTime) CalculateBonus() float64 {
	if fullTime.MonthlySalary > 3000 {
		return fullTime.MonthlySalary * fullTime.BonusRate
	}
	return 0
}

// PartTime
func (partTime PartTime) CalculateSalary() float64 {
	return partTime.HourlyRate*float64(partTime.HoursWorked) + partTime.CalculateBonus()
}

func (partTime PartTime) CalculateBonus() float64 {
	if partTime.HoursWorked > 30 {
		return float64(partTime.HoursWorked) * partTime.HourlyRate * 0.2
	}
	return 0
}

// Contractor
func (contractor Contractor) CalculateSalary() float64 {
	return contractor.ProjectRate*float64(contractor.ProjectsCompleted) + contractor.CalculateBonus()
}

func (contractor Contractor) CalculateBonus() float64 {
	if contractor.ProjectsCompleted > 7 {
		return contractor.ProjectRate * float64(contractor.ProjectsCompleted) * 0.1
	}
	return 0
}

// Intern
func (intern Intern) CalculateSalary() float64 {
	return intern.DailyRate*float64(intern.DaysWorked) + intern.CalculateBonus()
}

func (intern Intern) CalculateBonus() float64 {
	if intern.DaysWorked > 20 {
		return intern.DailyRate * float64(intern.DaysWorked) * 0.05
	}
	return 0
}

func RunEmployeeMenu(employee *Employee) {
	for {
		fmt.Println("\nEmployee Menu")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Show Employees")
		fmt.Println("3. Exit")
		fmt.Print("Enter your option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var employeeType string
			fmt.Print("Type of employee to add (Intern, FullTime, PartTime, Contractor): ")
			fmt.Scanln(&employeeType)

			switch employeeType {
			case "Intern":
				var dailyRate float64
				var daysWorked int
				fmt.Print("Enter daily rate: ")
				fmt.Scanln(&dailyRate)
				fmt.Print("Enter days worked: ")
				fmt.Scanln(&daysWorked)

				employee.Intern = append(employee.Intern, Intern{
					DailyRate:  dailyRate,
					DaysWorked: daysWorked,
				})

			case "FullTime":
				var monthlySalary, bonusRate float64
				fmt.Print("Enter monthly salary: ")
				fmt.Scanln(&monthlySalary)
				fmt.Print("Enter bonus rate: ")
				fmt.Scanln(&bonusRate)

				employee.FullTime = append(employee.FullTime, FullTime{
					MonthlySalary: monthlySalary,
					BonusRate:     bonusRate,
				})

			case "PartTime":
				var hourlyRate float64
				var hoursWorked int
				fmt.Print("Enter hourly rate: ")
				fmt.Scanln(&hourlyRate)
				fmt.Print("Enter hours worked: ")
				fmt.Scanln(&hoursWorked)

				employee.PartTime = append(employee.PartTime, PartTime{
					HourlyRate:  hourlyRate,
					HoursWorked: hoursWorked,
				})

			case "Contractor":
				var projectRate float64
				var projectsCompleted int
				fmt.Print("Enter project rate: ")
				fmt.Scanln(&projectRate)
				fmt.Print("Enter projects completed: ")
				fmt.Scanln(&projectsCompleted)

				employee.Contractor = append(employee.Contractor, Contractor{
					ProjectRate:       projectRate,
					ProjectsCompleted: projectsCompleted,
				})

			default:
				fmt.Println("We do not have this employee type.")
			}

		case 2:
			fmt.Println("\nEmployees")
			fmt.Println("Interns:")
			for _, intern := range employee.Intern {
				fmt.Printf("DailyRate: %.2f, DaysWorked: %d, Salary: %.2f\n", intern.DailyRate, intern.DaysWorked, intern.CalculateSalary())
			}

			fmt.Println("FullTime:")
			for _, fullTime := range employee.FullTime {
				fmt.Printf("MonthlySalary: %.2f, BonusRate: %.2f, Salary: %.2f\n", fullTime.MonthlySalary, fullTime.BonusRate, fullTime.CalculateSalary())
			}

			fmt.Println("PartTime:")
			for _, partTime := range employee.PartTime {
				fmt.Printf("HourlyRate: %.2f, HoursWorked: %d, Salary: %.2f\n", partTime.HourlyRate, partTime.HoursWorked, partTime.CalculateSalary())
			}

			fmt.Println("Contractors:")
			for _, contractor := range employee.Contractor {
				fmt.Printf("ProjectRate: %.2f, ProjectsCompleted: %d, Salary: %.2f\n", contractor.ProjectRate, contractor.ProjectsCompleted, contractor.CalculateSalary())
			}

		case 3:
			fmt.Println("Exiting menu")
			return

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
