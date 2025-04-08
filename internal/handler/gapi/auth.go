package gapi

import (
	"context"
	"net/http"
	"pointofsale/internal/domain/requests"
	protomapper "pointofsale/internal/mapper/proto"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authHandleGrpc struct {
	pb.UnimplementedAuthServiceServer
	authService service.AuthService
	mapping     protomapper.AuthProtoMapper
}

func NewAuthHandleGrpc(auth service.AuthService, mapping protomapper.AuthProtoMapper) *authHandleGrpc {
	return &authHandleGrpc{authService: auth, mapping: mapping}
}

func (s *authHandleGrpc) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.ApiResponseLogin, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	request := &requests.AuthRequest{
		Email:    strings.TrimSpace(req.Email),
		Password: req.Password,
	}

	res, errRes := s.authService.Login(request)

	if errRes != nil {
		if errRes.Code == http.StatusUnauthorized {
			return nil, status.Error(
				codes.Unauthenticated,
				"invalid credentials",
			)
		}
		return nil, status.Errorf(
			codes.Code(errRes.Code),
			"%s: %s",
			errRes.Status,
			errRes.Message,
		)
	}

	if res == nil || res.AccessToken == "" {
		return nil, status.Error(
			codes.Internal,
			"authentication service returned empty token",
		)
	}

	return s.mapping.ToProtoResponseLogin("success", "Login successful", res), nil
}

func (s *authHandleGrpc) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.ApiResponseRefreshToken, error) {
	res, err := s.authService.RefreshToken(req.RefreshToken)

	if err != nil {
		switch {
		case err.Code == http.StatusUnauthorized:
			return nil, status.Error(codes.Unauthenticated, "invalid or expired refresh token")
		case err.Code == http.StatusInternalServerError:
			return nil, status.Error(codes.InvalidArgument, err.Message)
		default:
			return nil, status.Errorf(codes.Code(err.Code), "%s: %s", err.Status, err.Message)
		}
	}

	if res == nil || res.AccessToken == "" || res.RefreshToken == "" {
		return nil, status.Error(codes.Internal, "failed to generate valid tokens")
	}

	return s.mapping.ToProtoResponseRefreshToken("success", "tokens refreshed successfully", res), nil
}

func (s *authHandleGrpc) GetMe(ctx context.Context, req *pb.GetMeRequest) (*pb.ApiResponseGetMe, error) {
	if req.AccessToken == "" {
		return nil, status.Error(codes.InvalidArgument, "access token is required")
	}

	res, err := s.authService.GetMe(req.AccessToken)

	if err != nil {
		switch {
		case err.Code == http.StatusUnauthorized:
			return nil, status.Error(codes.Unauthenticated, "invalid or expired access token")
		case err.Code == http.StatusInternalServerError:
			return nil, status.Error(codes.Internal, err.Message)
		default:
			return nil, status.Errorf(codes.Code(err.Code), "%s: %s", err.Status, err.Message)
		}
	}

	if res == nil {
		return nil, status.Error(codes.Internal, "failed to get user information")
	}

	return s.mapping.ToProtoResponseGetMe("success", "user data retrieved successfully", res), nil
}

func (s *authHandleGrpc) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.ApiResponseRegister, error) {
	request := &requests.CreateUserRequest{
		FirstName:       req.Firstname,
		LastName:        req.Lastname,
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Invalid registration data. Please check your input.",
		})
	}

	res, err := s.authService.Register(request)
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseRegister("success", "Registration successful", res), nil
}
