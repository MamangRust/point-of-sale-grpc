package api

import (
	"net/http"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	response_api "pointofsale/internal/mapper/response/api"
	"pointofsale/internal/pb"
	"pointofsale/pkg/logger"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userHandleApi struct {
	client  pb.UserServiceClient
	logger  logger.LoggerInterface
	mapping response_api.UserResponseMapper
}

func NewHandlerUser(router *echo.Echo, client pb.UserServiceClient, logger logger.LoggerInterface, mapping response_api.UserResponseMapper) *userHandleApi {
	userHandler := &userHandleApi{
		client:  client,
		logger:  logger,
		mapping: mapping,
	}
	routerUser := router.Group("/api/user")

	routerUser.GET("", userHandler.FindAllUser)
	routerUser.GET("/:id", userHandler.FindById)
	routerUser.GET("/active", userHandler.FindByActive)
	routerUser.GET("/trashed", userHandler.FindByTrashed)

	routerUser.POST("/create", userHandler.Create)
	routerUser.POST("/update/:id", userHandler.Update)

	routerUser.POST("/trashed/:id", userHandler.TrashedUser)
	routerUser.POST("/restore/:id", userHandler.RestoreUser)
	routerUser.DELETE("/permanent/:id", userHandler.DeleteUserPermanent)

	routerUser.POST("/restore/all", userHandler.RestoreAllUser)
	routerUser.POST("/permanent/all", userHandler.DeleteAllUserPermanent)

	return userHandler
}

// @Security Bearer
// @Summary Find all users
// @Tags User
// @Description Retrieve a list of all users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationUser "List of users"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve user data"
// @Router /api/user [get]
func (h *userHandleApi) FindAllUser(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllUserRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAll(ctx, req)

	if err != nil {
		h.logger.Error("Failed to fetch users", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "server_error",
			Message: "We couldn't retrieve the users. Please try again later.",
			Code:    http.StatusInternalServerError,
		})
	}

	so := h.mapping.ToApiResponsePaginationUser(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find user by ID
// @Tags User
// @Description Retrieve a user by ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ApiResponseUser "User data"
// @Failure 400 {object} response.ErrorResponse "Invalid user ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve user data"
// @Router /api/user/{id} [get]
func (h *userHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid user ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid user ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdUserRequest{
		Id: int32(id),
	}

	user, err := h.client.FindById(ctx, req)

	if err != nil {
		h.logger.Error("Failed to fetch user details", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "server_error",
			Message: "We couldn't retrieve the user details. Please try again later.",
			Code:    http.StatusInternalServerError,
		})
	}

	so := h.mapping.ToApiResponseUser(user)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active users
// @Tags User
// @Description Retrieve a list of active users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationUserDeleteAt "List of active users"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve user data"
// @Router /api/user/active [get]
func (h *userHandleApi) FindByActive(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllUserRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByActive(ctx, req)

	if err != nil {
		h.logger.Error("Failed to fetch active users", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "server_error",
			Message: "We couldn't retrieve the active users list. Please try again later.",
			Code:    http.StatusInternalServerError,
		})
	}

	so := h.mapping.ToApiResponsePaginationUserDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// FindByTrashed retrieves a list of trashed user records.
// @Summary Retrieve trashed users
// @Tags User
// @Description Retrieve a list of trashed user records
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationUserDeleteAt "List of trashed user data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve user data"
// @Router /api/user/trashed [get]
func (h *userHandleApi) FindByTrashed(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllUserRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByTrashed(ctx, req)

	if err != nil {
		h.logger.Error("Failed to fetch archived users", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "server_error",
			Message: "We couldn't retrieve the archived users list. Please try again later.",
			Code:    http.StatusInternalServerError,
		})
	}

	so := h.mapping.ToApiResponsePaginationUserDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// Create handles the creation of a new user.
// @Summary Create a new user
// @Tags User
// @Description Create a new user with the provided details
// @Accept json
// @Produce json
// @Param request body requests.CreateUserRequest true "Create user request"
// @Success 200 {object} response.ApiResponseUser "Successfully created user"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create user"
// @Router /api/user/create [post]
func (h *userHandleApi) Create(c echo.Context) error {
	var body requests.CreateUserRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Invalid request format", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Invalid request format. Please check your input.",
			Code:    http.StatusBadRequest,
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation failed", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "validation_error",
			Message: "Please provide valid transaction information.",
			Code:    http.StatusBadRequest,
		})
	}

	ctx := c.Request().Context()

	req := &pb.CreateUserRequest{
		Firstname:       body.FirstName,
		Lastname:        body.LastName,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	}

	res, err := h.client.Create(ctx, req)

	if err != nil {
		h.logger.Error("transaction creation failed", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "creation_failed",
			Message: "We couldn't create the transaction. Please try again.",
			Code:    http.StatusInternalServerError,
		})
	}

	so := h.mapping.ToApiResponseUser(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// Update handles the update of an existing user record.
// @Summary Update an existing user
// @Tags User
// @Description Update an existing user record with the provided details
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param UpdateUserRequest body requests.UpdateUserRequest true "Update user request"
// @Success 200 {object} response.ApiResponseUser "Successfully updated user"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update user"
// @Router /api/user/update/{id} [post]
func (h *userHandleApi) Update(c echo.Context) error {
	id := c.Param("id")

	idStr, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Invalid id parameter", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid id parameter",
			Code:    http.StatusBadRequest,
		})
	}

	var body requests.UpdateUserRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Invalid request format", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Invalid request format. Please check your input.",
			Code:    http.StatusBadRequest,
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation failed", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "validation_error",
			Message: "Please provide valid transaction information.",
			Code:    http.StatusBadRequest,
		})
	}

	ctx := c.Request().Context()

	req := &pb.UpdateUserRequest{
		Id:              int32(idStr),
		Firstname:       body.FirstName,
		Lastname:        body.LastName,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	}

	res, err := h.client.Update(ctx, req)

	if err != nil {
		h.logger.Error("transaction update failed", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "update_failed",
			Message: "We couldn't update the transaction information. Please try again.",
			Code:    http.StatusInternalServerError,
		})
	}

	so := h.mapping.ToApiResponseUser(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// TrashedUser retrieves a trashed user record by its ID.
// @Summary Retrieve a trashed user
// @Tags User
// @Description Retrieve a trashed user record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ApiResponseUserDeleteAt "Successfully retrieved trashed user"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve trashed user"
// @Router /api/user/trashed/{id} [get]
func (h *userHandleApi) TrashedUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid user ID format", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "invalid_input",
			Message: "Please provide a valid user ID.",
			Code:    http.StatusBadRequest,
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdUserRequest{
		Id: int32(id),
	}

	user, err := h.client.TrashedUser(ctx, req)

	if err != nil {
		h.logger.Error("Failed to archive user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "archive_failed",
			Message: "We couldn't archive the user. Please try again.",
			Code:    http.StatusInternalServerError,
		})
	}

	so := h.mapping.ToApiResponseUserDeleteAt(user)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// RestoreUser restores a user record from the trash by its ID.
// @Summary Restore a trashed user
// @Tags User
// @Description Restore a trashed user record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ApiResponseUserDeleteAt "Successfully restored user"
// @Failure 400 {object} response.ErrorResponse "Invalid user ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore user"
// @Router /api/user/restore/{id} [post]
func (h *userHandleApi) RestoreUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid user ID format", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "invalid_input",
			Message: "Please provide a valid user ID.",
			Code:    http.StatusBadRequest,
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdUserRequest{
		Id: int32(id),
	}

	user, err := h.client.RestoreUser(ctx, req)

	if err != nil {
		h.logger.Error("Failed to restore user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "restore_failed",
			Message: "We couldn't restore the user. Please try again.",
			Code:    http.StatusInternalServerError,
		})
	}

	so := h.mapping.ToApiResponseUserDeleteAt(user)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// DeleteUserPermanent permanently deletes a user record by its ID.
// @Summary Permanently delete a user
// @Tags User
// @Description Permanently delete a user record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ApiResponseUserDelete "Successfully deleted user record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete user:"
// @Router /api/user/delete/{id} [delete]
func (h *userHandleApi) DeleteUserPermanent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid user ID format", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "invalid_input",
			Message: "Please provide a valid user ID.",
			Code:    http.StatusBadRequest,
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdUserRequest{
		Id: int32(id),
	}

	user, err := h.client.DeleteUserPermanent(ctx, req)

	if err != nil {
		h.logger.Error("Failed to delete user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "deletion_failed",
			Message: "We couldn't permanently delete the user. Please try again.",
			Code:    http.StatusInternalServerError,
		})
	}

	so := h.mapping.ToApiResponseUserDelete(user)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// RestoreUser restores a user record from the trash by its ID.
// @Summary Restore a trashed user
// @Tags User
// @Description Restore a trashed user record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ApiResponseUserAll "Successfully restored user all"
// @Failure 400 {object} response.ErrorResponse "Invalid user ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore user"
// @Router /api/user/restore/all [post]
func (h *userHandleApi) RestoreAllUser(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.RestoreAllUser(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Bulk user restoration failed", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "restoration_failed",
			Message: "We couldn't restore all user. Please try again later.",
			Code:    http.StatusInternalServerError,
		})
	}

	so := h.mapping.ToApiResponseUserAll(res)

	h.logger.Debug("Successfully restored all user")

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// DeleteUserPermanent permanently deletes a user record by its ID.
// @Summary Permanently delete a user
// @Tags User
// @Description Permanently delete a user record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ApiResponseUserAll "Successfully deleted user record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete user:"
// @Router /api/user/delete/all [post]
func (h *userHandleApi) DeleteAllUserPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.DeleteAllUserPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Bulk user deletion failed", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "deletion_failed",
			Message: "We couldn't permanently delete all user. Please try again later.",
			Code:    http.StatusInternalServerError,
		})
	}

	so := h.mapping.ToApiResponseUserAll(res)

	h.logger.Debug("Successfully deleted all user permanently")

	return c.JSON(http.StatusOK, so)
}
