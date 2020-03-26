package viewmodel

//UserEditResponse ...
type UserEditResponse struct {
	*BaseResponse
	Name  string
	Pass  string
	Email string
	Admin bool
}

//NewUserEditResponse ...
func NewUserEditResponse() *UserEditResponse {
	return &UserEditResponse{&BaseResponse{false, ""}, "", "", "", false}
}
