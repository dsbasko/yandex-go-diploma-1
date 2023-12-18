package api

type RegisterRequestV1 struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponseV1 struct {
	UUID string `json:"uuid"`
}
