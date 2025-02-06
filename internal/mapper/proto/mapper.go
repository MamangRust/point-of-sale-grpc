package protomapper

type ProtoMapper struct {
	AuthProtoMapper AuthProtoMapper
	RoleProtoMapper RoleProtoMapper
	UserProtoMapper UserProtoMapper
}

func NewProtoMapper() *ProtoMapper {
	return &ProtoMapper{
		AuthProtoMapper: NewAuthProtoMapper(),
		RoleProtoMapper: NewRoleProtoMapper(),
		UserProtoMapper: NewUserProtoMapper(),
	}
}
