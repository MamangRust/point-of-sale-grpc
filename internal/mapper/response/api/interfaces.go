package response_api

import (
	"pointofsale/internal/domain/response"
	"pointofsale/internal/pb"
)

type AuthResponseMapper interface {
	ToResponseLogin(res *pb.ApiResponseLogin) *response.ApiResponseLogin
	ToResponseRegister(res *pb.ApiResponseRegister) *response.ApiResponseRegister
	ToResponseRefreshToken(res *pb.ApiResponseRefreshToken) *response.ApiResponseRefreshToken
	ToResponseGetMe(res *pb.ApiResponseGetMe) *response.ApiResponseGetMe
}

type RoleResponseMapper interface {
	ToApiResponseRoleAll(pbResponse *pb.ApiResponseRoleAll) *response.ApiResponseRoleAll
	ToApiResponseRoleDelete(pbResponse *pb.ApiResponseRoleDelete) *response.ApiResponseRoleDelete
	ToApiResponseRole(pbResponse *pb.ApiResponseRole) *response.ApiResponseRole
	ToApiResponsesRole(pbResponse *pb.ApiResponsesRole) *response.ApiResponsesRole
	ToApiResponsePaginationRole(pbResponse *pb.ApiResponsePaginationRole) *response.ApiResponsePaginationRole
	ToApiResponsePaginationRoleDeleteAt(pbResponse *pb.ApiResponsePaginationRoleDeleteAt) *response.ApiResponsePaginationRoleDeleteAt
}

type UserResponseMapper interface {
	ToApiResponseUser(pbResponse *pb.ApiResponseUser) *response.ApiResponseUser
	ToApiResponsesUser(pbResponse *pb.ApiResponsesUser) *response.ApiResponsesUser
	ToApiResponseUserDelete(pbResponse *pb.ApiResponseUserDelete) *response.ApiResponseUserDelete
	ToApiResponseUserAll(pbResponse *pb.ApiResponseUserAll) *response.ApiResponseUserAll
	ToApiResponsePaginationUserDeleteAt(pbResponse *pb.ApiResponsePaginationUserDeleteAt) *response.ApiResponsePaginationUserDeleteAt
	ToApiResponsePaginationUser(pbResponse *pb.ApiResponsePaginationUser) *response.ApiResponsePaginationUser
}
