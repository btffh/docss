package main

import (
	"docss/internal/conference"
	"docss/internal/employee"
	"docss/internal/logger"
	"docss/internal/organization"
	"docss/internal/student"
	"math/rand"

	"go.uber.org/zap/zapcore"
)

func main() {
	logger.NewLogger(zapcore.DebugLevel)

	schoolStudents := make([]*student.SchoolStudent, 0)
	for i := 0; i < 11; i++ {
		schoolStudents = append(schoolStudents, &student.SchoolStudent{
			Student: student.Student{},
			Grade:   i + 1,
		})
	}

	universityStudents := make([]*student.UniversityStudent, 0)
	for i := 0; i < 4; i++ {
		universityStudents = append(universityStudents, &student.UniversityStudent{
			Student: student.Student{},
			Course:  i + 1,
		})
	}

	conference := conference.NewConference(schoolStudents, universityStudents)
	logger.Infof("conference has %v school students", conference.CountSchoolStudents())
	logger.Infof("conference has %v university students", conference.CountUniversityStudents())
	logger.Infof("conference has %v students (both university and school)", conference.CountAllStudents())

	employees := make([]*employee.Employee, 0)
	for i := 0; i < 10; i++ {
		employees = append(employees, employee.NewEmployee(randomHours()))
	}

	org := organization.NewOrganization(employees)
	logger.Infof("organization has %v full rate employees", org.CounFullRate())

	logger.Fatal("exiting process")
}

func randomHours() int {
	switch rand.Intn(6) {
	case 0:
		return 20
	case 1:
		return 30
	default:
		return 40
	}
}
