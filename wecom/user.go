package wecom

type User struct {
	UserID     string `json:"userid"`
	Name       string `json:"name"`
	Mobile     string `json:"mobile,omitempty"`
	Email      string `json:"email,omitempty"`
	Department []int  `json:"department,omitempty"`
	Position   string `json:"position,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Enable     int    `json:"enable,omitempty"`
}

type userGetResponse struct {
	APIError
	User
}
