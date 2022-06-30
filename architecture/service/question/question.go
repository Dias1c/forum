package question

import "forum/architecture/models"

type QuestionService struct {
	repo models.IQuestionService
}

func NewQuestionService(repo models.IQuestionService) *QuestionService {
	return &QuestionService{repo}
}
