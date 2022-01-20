package main

type CalculateStudentGradeRequest struct {
	LectureGrade  float32 `json:"LectureGrade"`
	ExerciseGrade float32 `json:"ExerciseGrade"`
	WorkshopGrade float32 `json:"WorkshopGrade"`
}

type CalculateStudentGradeResponse struct {
	FinalGrade float32 `json:"FinalGrade"`
	HasPassed  bool    `json:"HasPassed"`
}
