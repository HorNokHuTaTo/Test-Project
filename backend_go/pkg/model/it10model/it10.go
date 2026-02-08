package it10model

type SubmitRequest struct {
	Name    string `json:"name"`
	Answers []int  `json:"answers"`
}

type ExamResult struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
	Total int    `json:"total"`
}

type Question struct {
	QuestionText  string   `json:"question_text"`
	Choices       []string `json:"choices"`
	CorrectChoice int      `json:"correct_choice"`
}
