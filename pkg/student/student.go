package student

// Student represents functionality students for both school and university.
type Student struct {
}

type SchoolStudent struct {
	Student
	Grade int // Grade is known to be age of studying in school from 1 to 11.
}

type UniversityStudent struct {
	Student
	Course int // Course is known to be age of studying in school from 1 to 11.
}
