package docss

import "math/rand"

// RandomHours returns pseudo random hours that employee spend at work weekly.
//
// Returned value is set to be 20, 30 or 40 as it is the most common and profitable
// case for company to hire employees.
//
// 	emp := NewEmployee(RandomHours())
//
// Chances:
//
// 	20: 1/8
// 	30: 2/8
// 	40: 5/8
func RandomHours() int {
	switch rand.Intn(8) {
	case 0:
		return 20
	case 1:
		return 30
	default:
		return 40
	}
}
