package exercism

func TriType(side1, side2, side3 int) string {
	if !isTri(side1, side2, side3) {
		return "Invalid Triangle"
	} else if side1 == side2 && side2 == side3 {
		return "Equilateral Triangle"
	} else if side1 == side2 || side2 == side3 || side3 == side1 {
		return "Isosceles Triangle"
	} else {
		return "Scalene Triangle"
	}
}

// checks if a triangle is valid
func isTri(side1, side2, side3 int) bool {
	if side1 <= 0 || side2 <= 0 || side3 <= 0 {
		return false
	}
	if side1+side2 > side3 && side2+side3 > side1 && side3+side1 > side2 {
		return true
	}
	return false
}
