package response

type UserProfile struct {
	UID               string `json:"uid"`
	Name              string `json:"name"`
	Avatar            string `json:"avatar"`
	Domain            string `json:"domain"`
	Background        string `json:"background"`
	Intro             string `json:"intro"`
	HarassmentSetting string `json:"harassmentSetting"`
}

type PageQuestionsItem struct {
	ID        uint   `json:"id"`
	CreatedAt Time   `json:"createdAt"`
	Content   string `json:"content"`
	Answer    string `json:"answer"`
}

type PageQuestions struct {
	Total     int64                `json:"total"`
	Cursor    string               `json:"cursor"`
	Questions []*PageQuestionsItem `json:"questions"`
}

type PageQuestion struct {
	ID                uint     `json:"id"`
	IsOwner           bool     `json:"isOwner"`
	CreatedAt         Time     `json:"createdAt"`
	AnsweredAt        Time     `json:"answeredAt"`
	Content           string   `json:"content"`
	Answer            string   `json:"answer"`
	QuestionImageURLs []string `json:"questionImageURLs"`
	AnswerImageURLs   []string `json:"answerImageURLs"`
	HasReplyEmail     bool     `json:"hasReplyEmail"`
	IsPrivate         bool     `json:"isPrivate"`
}
