package main

import (
	"calculate-student-grade/grades"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(event CalculateStudentGradeRequest) (CalculateStudentGradeResponse, error) {
	finalGrade, hasPassed := grades.CalculateStudentGrades(event.LectureGrade, event.ExerciseGrade, event.WorkshopGrade)
	return CalculateStudentGradeResponse{
		finalGrade, hasPassed,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
