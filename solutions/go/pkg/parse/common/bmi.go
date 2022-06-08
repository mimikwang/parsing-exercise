package common

// Bmi returns the corresponding bmi label given a numeric body mass index
func Bmi(bmi float64) string {
	switch bmi := bmi; {
	case bmi < 18.5:
		return "underweight"
	case bmi <= 24.9:
		return "normal"
	case bmi <= 29.9:
		return "overweight"
	default:
		return "very overweight"
	}
}
