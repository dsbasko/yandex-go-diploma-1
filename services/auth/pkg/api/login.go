package api

type LoginRequestV1 struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseV1 struct {
	UUID  string `json:"uuid"`
	Token string `json:"token"`
}
