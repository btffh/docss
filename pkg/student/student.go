package student

import (
	"fmt"
	"time"
)

// Student represents functionality students for both school and university.
//
// As an example, this struct has unexported, private fields that are not shown in documentation.
type Student struct {
	name    string
	surname string
	city    string

	Birthday Birthday
}

// Birthday contains day, month and year student was born.
type Birthday struct {
	day   int
	month time.Month
	year  int
}

// NewStudent returns new Student instance. Note: use it to create new SchoolStudent or UniversityStudent.
func NewStudent(
	name string,
	surname string,
	city string,
	day int,
	month time.Month,
	year int,
) *Student {
	return &Student{
		name:     name,
		surname:  surname,
		city:     city,
		Birthday: Birthday{day, month, year},
	}
}

func (s *Student) fullName() string {
	return fmt.Sprintf("%v %v", s.name, s.surname)
}

func (s *Student) birthdayString() string {
	return fmt.Sprintf("%v/%v/%v", s.Birthday.day, s.Birthday.month.String(), s.Birthday.year)
}

// String implements Stringer interface. Usage:
//
// 	fmt.Printf(Student{...})
func (s *Student) String() string {
	return fmt.Sprintf("%v (%v), %v", s.fullName(), s.birthdayString(), s.city)
}

// SchoolStudent represents shool's student.
type SchoolStudent struct {
	Student
	Grade int // Grade is known to be age of studying in school from 1 to 11.
}

// String implements Stringer interface. Usage:
//
// 	fmt.Printf(SchoolStudent{...})
func (ss *SchoolStudent) String() string {
	return fmt.Sprintf("%v studies at %v grade", ss.Student.String(), ss.Grade)
}

// UniversityStudent represents university's student.
type UniversityStudent struct {
	Student
	Course int // Course is known to be age of studying in school from 1 to 11.
}

// String implements Stringer interface. Usage:
//
// 	fmt.Printf(UniversityStudent{...})
func (us *UniversityStudent) String() string {
	return fmt.Sprintf("%v studies at %v course", us.Student.String(), us.Course)
}
