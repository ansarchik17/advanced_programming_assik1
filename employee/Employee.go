package employee

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

type SalaryCalculator interface {
	CalculateSalary() float64
	CalculateBonus() float64
}

// Full time
func (fullTime FullTime) CalculateSalary() float64 {
	return fullTime.MonthlySalary + fullTime.CalculateBonus()
}

func (fullTime FullTime) CalculateBonus() float64 {
	if fullTime.MonthlySalary > 3000 {
		return fullTime.MonthlySalary * fullTime.BonusRate
	}
	return 0
}

// Part time
func (partTime PartTime) CalculateSalary() float64 {
	return partTime.HourlyRate*float64(partTime.HoursWorked) + partTime.CalculateBonus()
}

func (partTime PartTime) CalculateBonus() float64 {
	if partTime.HoursWorked > 30 {
		return float64(partTime.HoursWorked) * float64(partTime.HourlyRate) * 0.2 // 20% bonus
	}
	return 0
}

// Contractor
func (contractor Contractor) CalculateSalary() float64 {
	return contractor.ProjectRate*float64(contractor.ProjectsCompleted) + contractor.CalculateBonus()
}

func (contractor Contractor) CalculateBonus() float64 {
	if contractor.ProjectsCompleted > 7 {
		return contractor.ProjectRate * float64(contractor.ProjectsCompleted) * 0.1 //10% bonus here
	}
	return 0
}

// Intern
func (intern Intern) CalculateSalary() float64 {
	return intern.DailyRate*float64(intern.DaysWorked) + intern.CalculateBonus()
}

func (intern Intern) CalculateBonus() float64 {
	if intern.DaysWorked > 20 {
		return intern.DailyRate * float64(intern.DaysWorked) * 0.05 // 5% bonus
	}
	return 0
}
