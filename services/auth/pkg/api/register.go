package api

type RegisterRequestV1 struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type RegisterResponseV1 struct {
	UUID string `json:"uuid"`
}
