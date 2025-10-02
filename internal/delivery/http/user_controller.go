package http

import (
	"finenance-app/internal/entity"
	"finenance-app/internal/model"
	"finenance-app/internal/usecase"
	"time"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type UserController struct {
	Log         *zap.Logger
	UserUsecase usecase.UserUsecaseImplementation
}

type UserControllerImplementation interface {
	Register(c fiber.Ctx) error
	Login(c fiber.Ctx) error
	GetProfile(c fiber.Ctx) error
	Logout(c fiber.Ctx) error
}

func NewUserController(log *zap.Logger, us usecase.UserUsecaseImplementation) UserControllerImplementation {
	return &UserController{
		Log:         log,
		UserUsecase: us,
	}
}

func (uc *UserController) Register(c fiber.Ctx) error {
	request := new(model.UsersRegisterRequest)

	if err := c.Bind().Body(request); err != nil {
		uc.Log.Error("Failed to parse request body : %+v", zap.Error(err))
		return fiber.ErrBadRequest
	}

	response, err := uc.UserUsecase.Create(request)
	if err != nil {
		uc.Log.Error("failed to register", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.UsersResponseRegister]{
		Code:    fiber.StatusCreated,
		Message: "success register user",
		Data:    &response,
	})
}

func (uc *UserController) Login(c fiber.Ctx) error {
	request := new(model.UserLoginRequest)
	if err := c.Bind().Body(request); err != nil {
		uc.Log.Error("Failed to parse request body : %+v", zap.Error(err))
		return fiber.ErrBadRequest
	}

	response, err := uc.UserUsecase.Login(request)
	if err != nil {
		uc.Log.Error("failed to login", zap.Error(err))
		return fiber.ErrNotFound
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    response.RefreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	finalResponse := model.UserLoginResponse{
		AccessToken: response.AccessToken,
	}

	return c.Status(fiber.StatusOK).JSON(model.WebResponse[*model.UserLoginResponse]{
		Code:    fiber.StatusOK,
		Message: "success Login",
		Data:    &finalResponse,
	})
}

func (uc *UserController) GetProfile(c fiber.Ctx) error {
	user := c.Locals("user").(*entity.CustomClaims)

	userProfile, err := uc.UserUsecase.GetProfile(user.UserID)
	if err != nil {
		uc.Log.Error("failed to get user profile", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success hello user",
		"data":    userProfile,
	})
}

func (uc *UserController) Logout(c fiber.Ctx) error {

	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return fiber.ErrUnauthorized
	}

	if err := uc.UserUsecase.Logout(refreshToken); err != nil {
		uc.Log.Error("failed to logout", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	// Clear cookie
	c.ClearCookie("refresh_token")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success logout",
		"data":    nil,
	})
}
