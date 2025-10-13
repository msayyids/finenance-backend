package usecase

import (
	"context"
	"errors"
	"finenance-app/internal/entity"
	"finenance-app/internal/model"
	"finenance-app/internal/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUsecase) Create(userRequest *model.UsersRegisterRequest) (model.UsersResponseRegister, error) {

	tx, err := uc.DB.Beginx()
	if err != nil {
		uc.Log.Error("failed to create transaction", zap.Error(err))
		return model.UsersResponseRegister{}, err
	}
	defer utils.CommitOrRollback(tx, &err)

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

	defaultCategory := []entity.Categories{
		{
			User_Id:    user.Id,
			Name:       "Salary",
			Type:       "income",
			Created_at: time.Now(),
		},
		{
			User_Id:    user.Id,
			Name:       "food and beverage",
			Type:       "expense",
			Created_at: time.Now(),
		},
		{
			User_Id:    user.Id,
			Name:       "transport",
			Type:       "expense",
			Created_at: time.Now(),
		},
		{
			User_Id:    user.Id,
			Name:       "entertainment",
			Type:       "expense",
			Created_at: time.Now(),
		},
	}

	if err := uc.CategoryRepo.AddDefaultCategory(tx, defaultCategory); err != nil {
		uc.Log.Warn("failed to add default category", zap.Error(err))
		return model.UsersResponseRegister{}, fiber.ErrInternalServerError
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
	defer utils.CommitOrRollback(tx, &err)

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
	expAccessToken := 60 * time.Minute
	accessToken, err := utils.GenerateToken(stringUserId, expAccessToken)
	if err != nil {
		uc.Log.Warn("failed to generate token", zap.Error(err))
		return model.UserLoginResponse{}, fiber.ErrInternalServerError
	}

	exprefreshToken := 7 * 24 * time.Hour
	refreshToken, err := utils.GenerateToken(stringUserId, exprefreshToken)
	if err != nil {
		uc.Log.Warn("failed to generate token", zap.Error(err))
		return model.UserLoginResponse{}, fiber.ErrInternalServerError
	}

	err = uc.ReddisClient.Set(context.Background(), "user_id_"+stringUserId, refreshToken, exprefreshToken).Err()
	if err != nil {
		uc.Log.Warn("failed to refresh token", zap.Error(err))
		return model.UserLoginResponse{}, fiber.ErrInternalServerError
	}
	response := model.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (uc *UserUsecase) RefreshToken(refreshToken string, user_id int) (string, error) {
	tx, err := uc.DB.Beginx()
	if err != nil {
		uc.Log.Error("failed to create transaction", zap.Error(err))
		return "", err
	}
	defer utils.CommitOrRollback(tx, &err)

	// cek user
	user, err := uc.UserRepo.FindUserById(tx, user_id)
	if err != nil {
		uc.Log.Warn("failed to find user by id", zap.Error(err))
		return "", err
	}

	// cek refresh token di Redis
	key := "user_id_" + strconv.Itoa(user.Id)
	storedToken, err := uc.ReddisClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		uc.Log.Warn("refresh token not found in redis")
		return "", errors.New("refresh token not found")
	} else if err != nil {
		uc.Log.Error("failed to get refresh token from redis", zap.Error(err))
		return "", err
	}

	// validasi token
	if storedToken != refreshToken {
		uc.Log.Warn("refresh token mismatch")
		return "", errors.New("invalid refresh token")
	}
	stringUserId := strconv.Itoa(user.Id)

	// generate access token baru
	newAccessToken, err := utils.GenerateToken(stringUserId, 15*time.Minute)
	if err != nil {
		uc.Log.Error("failed to generate new access token", zap.Error(err))
		return "", err
	}

	return newAccessToken, nil
}

func (uc *UserUsecase) Logout(refreshToken string) error {

	tx, err := uc.DB.Beginx()
	if err != nil {
		uc.Log.Error("failed to create transaction", zap.Error(err))
		return err
	}
	defer utils.CommitOrRollback(tx, &err)

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

	return nil
}

func (uc *UserUsecase) GetProfile(userId string) (model.UserGetProfileResponse, error) {

	tx, err := uc.DB.Beginx()
	if err != nil {
		uc.Log.Error("failed to create transaction", zap.Error(err))
	}
	defer utils.CommitOrRollback(tx, &err)

	Intid, err := strconv.Atoi(userId)
	if err != nil {
		uc.Log.Error("failed to convert user id to int", zap.Error(err))
	}

	profile, err := uc.UserRepo.FindUserById(tx, Intid)
	if err != nil {
		uc.Log.Warn("failed to find user by id", zap.Error(err))
		return model.UserGetProfileResponse{}, err
	}

	getProfile := model.UserGetProfileResponse{
		Id:    profile.Id,
		Name:  profile.Name,
		Email: profile.Email,
	}

	return getProfile, nil
}
