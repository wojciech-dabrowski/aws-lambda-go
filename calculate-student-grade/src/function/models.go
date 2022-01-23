package main

type CalculateStudentGradeRequest struct {
	LectureGrade  float32 `json:"lectureGrade"`
	ExerciseGrade float32 `json:"exerciseGrade"`
	WorkshopGrade float32 `json:"workshopGrade"`
}

type CalculateStudentGradeResponse struct {
	FinalGrade float32 `json:"finalGrade"`
	HasPassed  bool    `json:"hasPassed"`
}
