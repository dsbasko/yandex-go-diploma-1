package api

type AuthRequestV1 struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseV1 struct {
	UUID  string `json:"uuid"`
	Token string `json:"token"`
}
