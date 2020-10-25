package main

import (
	"docss"
	"docss/internal/logger"
	"docss/pkg/conference"
	"docss/pkg/employee"
	"docss/pkg/organization"
	"docss/pkg/student"
	"time"

	"go.uber.org/zap/zapcore"
)

func main() {
	logger.NewLogger(zapcore.DebugLevel)

	schoolStudents := make([]*student.SchoolStudent, 0)
	for i := 0; i < 11; i++ {
		st := &student.SchoolStudent{
			Student: *student.NewStudent("William", "Turner Jr.", "Caribbean Sea", i+1, time.June, 1681),
			Grade:   i + 1,
		}
		schoolStudents = append(schoolStudents, st)
		logger.Infof("New school student %s", st)
	}

	universityStudents := make([]*student.UniversityStudent, 0)
	for i := 0; i < 4; i++ {
		st := &student.UniversityStudent{
			Student: *student.NewStudent("Davy", "Jones", "Caribbean Sea", i+1, time.July, 1654),
			Course:  i + 1,
		}
		universityStudents = append(universityStudents, st)
		logger.Infof("New university student %s", st)
	}

	conference := conference.NewConference(schoolStudents, universityStudents)
	logger.Infof("conference has %v school students", conference.CountSchoolStudents())
	logger.Infof("conference has %v university students", conference.CountUniversityStudents())
	logger.Infof("conference has %v students (both university and school)", conference.CountAllStudents())

	employees := make([]*employee.Employee, 0)
	for i := 0; i < 10; i++ {
		employees = append(employees, employee.NewEmployee(docss.RandomHours()))
	}

	org := organization.NewOrganization(employees)
	logger.Infof("organization has %v full rate employees", org.CounFullRate())

	logger.Fatal("exiting process")
}
