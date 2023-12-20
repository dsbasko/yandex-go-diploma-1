package api

type JWTPayloadV1 struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type JWTValidationRequestV1 struct {
	Token string `json:"token"`
}

type JWTValidationResponseV1 struct {
	IsValid bool          `json:"is_valid"`
	Payload *JWTPayloadV1 `json:"payload"`
}

var AMQPQueueJWTValidation = "auth.validation.queue"
var AMQPKeyJWTValidation = "auth.validation.key"
