package presentation

type Response struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   string                 `json:"error,omitempty"`
}

type ClaimTokenUser struct {
	Authorized   bool    `json:"authorized"`
	UserID       string  `json:"user_id"`
	UserCode     string  `json:"user_code"`
	Exp          float32 `json:"exp"`
}
