package conference

import (
	"docss/internal/student"
)

type Conference struct {
	SchoolStudents     []*student.SchoolStudent
	UniversityStudents []*student.UniversityStudent
}

func NewConference(
	schoolStudents []*student.SchoolStudent,
	universityStudents []*student.UniversityStudent,
) *Conference {
	return &Conference{
		SchoolStudents:     schoolStudents,
		UniversityStudents: universityStudents,
	}
}

func (c *Conference) CountSchoolStudents() int {
	return len(c.SchoolStudents)
}

func (c *Conference) CountUniversityStudents() int {
	return len(c.UniversityStudents)
}

func (c *Conference) CountAllStudents() int {
	return len(c.SchoolStudents) + len(c.UniversityStudents)
}
