package service

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"vega-server/api"
	"vega-server/internal/model"
	"vega-server/internal/repository"
	"vega-server/pkg/jwt"

	"errors"
)

type UserService struct {
	*Service
	userRepository *repository.UserRepository
	jwtService     *jwt.JWTService
}

func NewUserService(service *Service, userRepository *repository.UserRepository, jwtService *jwt.JWTService) *UserService {
	return &UserService{
		Service:        service,
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (userService *UserService) SignUp(req *api.SignUpRequest) (*api.SignUpResponseData, error) {
	user, err := userService.userRepository.QueryUserByEmail(req.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to sign up")
		}
	} else {
		return nil, errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to sign up")
	}

	user = &model.User{
		Email:          req.Email,
		HashedPassword: string(hashedPassword),
		Avatar:         "https://example.com/default-avatar.png",
		Name:           "Vega",
	}
	err = userService.userRepository.Create(user)
	if err != nil {
		return nil, errors.New("failed to sign up")
	}

	token, err := userService.jwtService.GenerateJWT(user.ID)
	if err != nil {
		return nil, errors.New("failed to sign up")
	}

	resp := &api.SignUpResponseData{
		User: api.User{
			Avatar: user.Avatar,
			Name:   user.Name,
		},
		Token: token,
	}
	return resp, nil
}

func (userService *UserService) SignIn(req *api.SignInRequest) (*api.SignInResponseData, error) {
	user, err := userService.userRepository.QueryUserByEmail(req.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to sign in")
		} else {
			return nil, errors.New("invalid email")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := userService.jwtService.GenerateJWT(user.ID)
	if err != nil {
		return nil, errors.New("failed to sign in")
	}

	resp := &api.SignInResponseData{
		User: api.User{
			Avatar: user.Avatar,
			Name:   user.Name,
		},
		Token: token,
	}
	return resp, nil
}

func (userService *UserService) GetUser(id uint) (*api.GetUserResponseData, error) {
	user, err := userService.userRepository.QueryUserByID(id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to get user")
		} else {
			return nil, errors.New("invalid id")
		}
	}

	resp := &api.GetUserResponseData{
		User: api.User{
			Avatar: user.Avatar,
			Name:   user.Name,
		},
	}
	return resp, nil
}

func (userService *UserService) GetAdvancedUser(id uint) (*api.GetAdvancedUserResponseData, error) {
	user, err := userService.userRepository.QueryUserByID(id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to get user")
		} else {
			return nil, errors.New("invalid id")
		}
	}

	reviews := make([]uint, 0, len(user.Reviews))
	for _, review := range user.Reviews {
		reviews = append(reviews, review.ID)
	}

	resp := &api.GetAdvancedUserResponseData{
		User: api.AdvancedUser{
			Avatar:    user.Avatar,
			Name:      user.Name,
			Gender:    user.Gender,
			BirthDate: user.BirthDate,
			Location:  user.Location,
			Bio:       user.Bio,
		},
		Reviews: reviews,
	}
	return resp, nil
}
