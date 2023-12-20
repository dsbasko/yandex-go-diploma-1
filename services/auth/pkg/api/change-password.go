package api

type ChangePasswordRequestV1 struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordResponseV1 struct {
	UUID string `json:"uuid"`
}
