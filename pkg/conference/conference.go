package conference

import (
	"docss/pkg/student"
)

// Conference represents a conference that has students as a partisipants.
type Conference struct {
	SchoolStudents     []*student.SchoolStudent
	UniversityStudents []*student.UniversityStudent
}

// NewConference creates new Conference instance.
func NewConference(
	schoolStudents []*student.SchoolStudent,
	universityStudents []*student.UniversityStudent,
) *Conference {
	return &Conference{
		SchoolStudents:     schoolStudents,
		UniversityStudents: universityStudents,
	}
}

// CountSchoolStudents returns amount of conference's students who studies at school and participate at conference.
func (c *Conference) CountSchoolStudents() int {
	return len(c.SchoolStudents)
}

// CountUniversityStudents returns amount of conference's students who studies at university and participate at conference.
func (c *Conference) CountUniversityStudents() int {
	return len(c.UniversityStudents)
}

// CountAllStudents returns amount of students who participates at conference.
func (c *Conference) CountAllStudents() int {
	return len(c.SchoolStudents) + len(c.UniversityStudents)
}
