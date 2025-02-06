package recordmapper

type RecordMapper struct {
	UserRecordMapper         UserRecordMapping
	RoleRecordMapper         RoleRecordMapping
	UserRoleRecordMapper     UserRoleRecordMapping
	RefreshTokenRecordMapper RefreshTokenRecordMapping
}

func NewRecordMapper() *RecordMapper {
	return &RecordMapper{
		UserRecordMapper:         NewUserRecordMapper(),
		RoleRecordMapper:         NewRoleRecordMapper(),
		UserRoleRecordMapper:     NewUserRoleRecordMapper(),
		RefreshTokenRecordMapper: NewRefreshTokenRecordMapper(),
	}
}
