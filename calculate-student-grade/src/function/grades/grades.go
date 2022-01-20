package grades

import "math"

const lectureWeight float32 = 0.2
const exerciseWeight float32 = 0.3
const workshopWeight float32 = 0.5
const gradeToPass = 3

func CalculateStudentGrades(lectureGrade float32, exerciseGrade float32, workshopGrade float32) (finalGrade float32, hasPassed bool) {
	finalGrade = lectureWeight*lectureGrade +
		exerciseWeight*exerciseGrade +
		workshopWeight*workshopGrade
	hasPassed = finalGrade > gradeToPass &&
		lectureGrade > gradeToPass &&
		exerciseGrade > gradeToPass

	return finalGrade, hasPassed
}

func roundToNearestHalf(numberToRound float64) float64 {
	return math.Round(numberToRound*2)/2
}
