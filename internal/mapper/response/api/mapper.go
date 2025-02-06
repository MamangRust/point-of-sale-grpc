package response_api

type ResponseApiMapper struct {
	AuthResponseMapper AuthResponseMapper
	RoleResponseMapper RoleResponseMapper
	UserResponseMapper UserResponseMapper
}

func NewResponseApiMapper() *ResponseApiMapper {
	return &ResponseApiMapper{
		AuthResponseMapper: NewAuthResponseMapper(),
		UserResponseMapper: NewUserResponseMapper(),
		RoleResponseMapper: NewRoleResponseMapper(),
	}
}
