package response_service

type ResponseServiceMapper struct {
	RoleResponseMapper         RoleResponseMapper
	RefreshTokenResponseMapper RefreshTokenResponseMapper
	UserResponseMapper         UserResponseMapper
}

func NewResponseServiceMapper() *ResponseServiceMapper {
	return &ResponseServiceMapper{
		UserResponseMapper:         NewUserResponseMapper(),
		RefreshTokenResponseMapper: NewRefreshTokenResponseMapper(),
		RoleResponseMapper:         NewRoleResponseMapper(),
	}
}
