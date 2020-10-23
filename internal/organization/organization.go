package organization

import (
	"docss/internal/employee"
)

// An Organization type represents an organization that holds employees.
type Organization struct {
	Employees []*employee.Employee
}

// NewOrganization creates new Organization instance with it's employees passed.
func NewOrganization(employees []*employee.Employee) *Organization {
	return &Organization{Employees: employees}
}

// CounFullRate returns amount of employees that works for a full rate on company.
func (o *Organization) CounFullRate() int {
	count := 0
	for _, e := range o.Employees {
		if e.IsFullRate() {
			count++
		}
	}

	return count
}
