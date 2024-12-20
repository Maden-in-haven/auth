package handler

import (
	"auth/internal/gen"
	to"auth/internal/token"
	"context"

	// "github.com/Maden-in-haven/crmlib/pkg/myjwt"
	"log"

	"github.com/Maden-in-haven/crmlib/pkg/user"
)

func (s *AuthService) APIAuthLoginPost(ctx context.Context, req *gen.APIAuthLoginPostReq) (gen.APIAuthLoginPostRes, error) {
	// Логируем начало запроса
	log.Printf("Начало запроса на авторизацию для пользователя: %s", req.Username.Value)

	// Проверка корректности логина и пароля
	user, err := user.AuthenticateUser(req.Username.Value, req.Password.Value)
	if err != nil {
		log.Printf("Ошибка авторизации пользователя %s: %v", req.Username.Value, err)
		return &gen.APIAuthLoginPostUnauthorized{
			Message: gen.OptString{
				Value: "Неверный указан пользователь или пароль",
				Set:   true,
			},
		}, nil
	}
	log.Printf("Пользователь %s успешно аутентифицирован, генерируем токен", req.Username.Value)

	// Генерация JWT токена
	token, err := to.GenerateJWT(user.ID)
	if err != nil {
		log.Printf("Ошибка генерации JWT токена для пользователя %s: %v", user.ID, err)
		// Возвращаем 500 Internal Server Error в случае ошибки при генерации JWT
		return &gen.APIAuthLoginPostInternalServerError{
			Message: gen.OptString{
				Value: "Ошибка при генерации JWT токена",
				Set:   true,
			},
		}, nil
	}
	log.Printf("JWT токен успешно сгенерирован для пользователя %s", user.ID)

	// Генерация Refresh токена
	refreshToken, err := to.GenerateRefreshToken(user.ID)
	if err != nil {
		log.Printf("Ошибка генерации Refresh токена для пользователя %s: %v", user.ID, err)
		// Возвращаем 500 Internal Server Error в случае ошибки при генерации Refresh токена
		return &gen.APIAuthLoginPostInternalServerError{
			Message: gen.OptString{
				Value: "Ошибка при генерации рефреш токена",
				Set:   true,
			},
		}, nil
	}
	log.Printf("Refresh токен успешно сгенерирован для пользователя %s", user.ID)

	// Логируем успешную авторизацию
	log.Printf("Авторизация пользователя %s успешно завершена", req.Username.Value)

	// Возвращаем ответ с токеном и рефреш токеном
	return &gen.APIAuthLoginPostOK{
		AccessToken:  gen.OptString{Value: token, Set: true},
		RefreshToken: gen.OptString{Value: refreshToken, Set: true},
	}, nil
}
