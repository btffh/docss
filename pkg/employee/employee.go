package employee

import "docss/pkg/logger"

// An Employee represents a single employee that work for a company
// on a full or partial rate.
type Employee struct {
	Hours int // Hours employee works per week.
}

// NewEmployee creates new Employee instance with amount of hours passed.
func NewEmployee(hours int) *Employee {
	return &Employee{
		Hours: hours,
	}
}

// IsFullRate indicates if employee works on a full rate or not.
func (e *Employee) IsFullRate() bool {
	logger.Debugf("employee's hours: %v", e.Hours)
	return e.Hours == 40
}
