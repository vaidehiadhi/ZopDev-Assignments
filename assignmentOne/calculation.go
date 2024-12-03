package assignmentOne

const Pi = 3.14

func PerimeterCircle(radius int) float64 {
	var p = 2 * Pi * (float64(radius))
	return p
}

func PerimeterSquare(a int) int {
	return 4 * a
}

func PerimeterRectangle(l, w int) float64 {
	var r = (float64(l) + float64(w)) * 2
	return r

}
