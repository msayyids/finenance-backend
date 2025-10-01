package usecase

import (
	"context"
	"finenance-app/internal/entity"
	"finenance-app/internal/model"
	"finenance-app/internal/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUsecase) Create(userRequest *model.UsersRegisterRequest) (model.UsersResponseRegister, error) {
	tx, err := uc.DB.Beginx()
	if err != nil {
		uc.Log.Error("failed to create transaction", zap.Error(err))
		return model.UsersResponseRegister{}, err
	}
	defer tx.Rollback()

	//validasi
	if err := uc.Validator.Struct(userRequest); err != nil {
		uc.Log.Warn("invalid user request", zap.Error(err))
		return model.UsersResponseRegister{}, fiber.ErrBadRequest
	}

	password, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		uc.Log.Error("failed to generate password", zap.Error(err))
		return model.UsersResponseRegister{}, err
	}

	//find role

	user := &entity.Users{
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Password:  string(password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	//create/insert new users
	if err := uc.UserRepo.AddUser(tx, user); err != nil {
		uc.Log.Warn("failed to create user", zap.Error(err))
		return model.UsersResponseRegister{}, fiber.ErrInternalServerError
	}
	//commit
	if err := tx.Commit(); err != nil {
		uc.Log.Warn("failed to commit transaction", zap.Error(err))
		return model.UsersResponseRegister{}, err

	}

	response := model.ToUserResponse(userRequest.Name, user.Id)
	return response, nil

}

func (uc *UserUsecase) Login(userRequest *model.UserLoginRequest) (model.UserLoginResponse, error) {
	tx, err := uc.DB.Beginx()
	if err != nil {
		uc.Log.Error("failed to create transaction", zap.Error(err))
		return model.UserLoginResponse{}, err
	}
	defer tx.Rollback()

	if err := uc.Validator.Struct(userRequest); err != nil {
		uc.Log.Warn("invalid user request", zap.Error(err))
		return model.UserLoginResponse{}, err
	}

	user, err := uc.UserRepo.FindUserByEmail(tx, userRequest.Email)
	if err != nil {
		uc.Log.Warn("failed to find user by email", zap.Error(err))
		return model.UserLoginResponse{}, fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err != nil {
		uc.Log.Warn("invalid user request", zap.Error(err))
		return model.UserLoginResponse{}, fiber.ErrInternalServerError
	}

	stringUserId := strconv.Itoa(user.Id)

	//create token jwt (implement reddis or not?)
	expAccessToken := 15 * time.Minute
	accessToken, err := utils.GenerateToken(stringUserId, "user", expAccessToken)
	if err != nil {
		uc.Log.Warn("failed to generate token", zap.Error(err))
	}

	exprefreshToken := 7 * 24 * time.Hour
	refreshToken, err := utils.GenerateToken(stringUserId, "user", exprefreshToken)
	if err != nil {
		uc.Log.Warn("failed to generate token", zap.Error(err))
	}

	uc.ReddisClient.Set(context.Background(), "user_id_"+stringUserId, refreshToken, exprefreshToken)
	response := model.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := tx.Commit(); err != nil {
		uc.Log.Warn("failed to commit transaction", zap.Error(err))
		return model.UserLoginResponse{}, err

	}

	return response, nil
}

func (uc *UserUsecase) Logout(refreshToken string) error {

	tx, err := uc.DB.Beginx()
	if err != nil {
		uc.Log.Error("failed to create transaction", zap.Error(err))
		return err
	}
	defer tx.Rollback()

	// Parse refresh token → ambil user_id dari claim
	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		uc.Log.Error("failed to validate token", zap.Error(err))
		return err
	}

	userId := claims.UserID

	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		uc.Log.Error("failed to convert user id to int", zap.Error(err))
		return err
	}

	_, err = uc.UserRepo.FindUserById(tx, intUserId)
	if err != nil {
		uc.Log.Warn("user not found", zap.Error(err))
		return err
	}

	// Hapus refresh token dari Redis
	if err := uc.ReddisClient.Del(context.Background(), "user_id_"+userId).Err(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		uc.Log.Warn("failed to commit transaction", zap.Error(err))
		return err

	}

	return nil
}

func (uc *UserUsecase) GetProfile(userId string) (model.UserGetProfileResponse, error) {

	tx, err := uc.DB.Beginx()
	if err != nil {
		uc.Log.Error("failed to create transaction", zap.Error(err))
	}
	defer tx.Rollback()

	Intid, err := strconv.Atoi(userId)
	if err != nil {
		uc.Log.Error("failed to convert user id to int", zap.Error(err))
	}

	profile, err := uc.UserRepo.FindUserById(tx, Intid)
	if err != nil {
		uc.Log.Warn("failed to find user by id", zap.Error(err))
		return model.UserGetProfileResponse{}, err
	}

	if err := tx.Commit(); err != nil {
		uc.Log.Warn("failed to commit transaction", zap.Error(err))
		return model.UserGetProfileResponse{}, err

	}

	getProfile := model.UserGetProfileResponse{
		Id:    profile.Id,
		Name:  profile.Name,
		Email: profile.Email,
	}

	return getProfile, nil
}

//func (uc *UserUsecase) Logout(userId string) error {
//
//}
