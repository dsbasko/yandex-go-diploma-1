package api

type JWTValidationRequestV1 struct {
	JWT string `json:"jwt"`
}

type JWTValidationResponseV1 struct {
	IsValid bool `json:"is_valid"`
}
