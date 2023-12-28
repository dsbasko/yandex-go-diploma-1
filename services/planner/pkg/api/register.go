package api

type CreateTaskRequestV1 struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateTaskResponseV1 struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
