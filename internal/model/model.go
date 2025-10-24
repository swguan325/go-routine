package model

type Card struct {
    Number string `json:"number"`
    Brand  string `json:"brand"`
}

type Dashboard struct {
    Username string  `json:"username"`
    Balance  float64 `json:"balance"`
    Cards    []Card  `json:"cards"`
}

type LoginRequest struct {
    UserID   string `json:"userId"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token     string    `json:"token"`
    Dashboard Dashboard `json:"dashboard"`
}
