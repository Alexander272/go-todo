package service

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/Alexander272/go-todo/internal/repository"
	"github.com/Alexander272/go-todo/pkg/auth"
	"github.com/Alexander272/go-todo/pkg/hash"
)

type SessionService struct {
	repoUsers       repository.Users
	repoSes         repository.Session
	tokenManager    auth.TokenManager
	hasher          hash.PasswordHasher
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	domain          string
}

func NewSessionService(repoUsers repository.Users, repoSes repository.Session, tokenManager auth.TokenManager, hasher hash.PasswordHasher,
	accessTokenTTL time.Duration, refreshTokenTTL time.Duration, domain string) *SessionService {
	return &SessionService{
		repoUsers:       repoUsers,
		repoSes:         repoSes,
		tokenManager:    tokenManager,
		hasher:          hasher,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
		domain:          domain,
	}
}

func (s *SessionService) SignIn(ctx context.Context, input SignInInput, ua, ip string) (*http.Cookie, *domain.Token, error) {
	user, err := s.repoUsers.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, nil, errors.New("invalid credentials")
	}
	if ok := s.hasher.CheckPasswordHash(input.Password, user.Password); !ok {
		return nil, nil, errors.New("invalid credentials")
	}

	iat, accessToken, err := s.tokenManager.NewJWT(user.Id, user.Email, user.Role, s.accessTokenTTL)
	if err != nil {
		return nil, nil, err
	}
	refreshToken, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		return nil, nil, err
	}

	data := repository.SessionData{
		UserId:   user.Id,
		UserName: user.Name,
		Email:    user.Email,
		Role:     user.Role,
		Ua:       ua,
		Ip:       ip,
		Exp:      s.refreshTokenTTL,
	}

	if err := s.repoSes.CreateSession(ctx, refreshToken, data); err != nil {
		return nil, nil, err
	}

	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    refreshToken,
		MaxAge:   int(s.refreshTokenTTL.Seconds()),
		Path:     "/",
		Domain:   s.domain,
		Secure:   false,
		HttpOnly: true,
	}

	token := &domain.Token{
		Token: domain.TokenData{
			AccessToken: accessToken,
			Exp:         iat.Add(s.accessTokenTTL).Unix(),
		},
		UserId: user.Id,
		Role:   user.Role,
		Name:   user.Name,
	}

	return cookie, token, nil
}

func (s *SessionService) SingOut(ctx context.Context, token string) (*http.Cookie, error) {
	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    "",
		MaxAge:   1,
		Path:     "/",
		Domain:   s.domain,
		Secure:   false,
		HttpOnly: true,
	}

	err := s.repoSes.RemoveSession(ctx, token)
	if err != nil {
		return cookie, err
	}

	return cookie, nil
}

func (s *SessionService) Refresh(ctx context.Context, token, ua, ip string) (*domain.Token, *http.Cookie, error) {
	data, err := s.repoSes.GetDelSession(ctx, token)
	if err != nil {
		return nil, nil, err
	}
	if ua != data.Ua || ip != data.Ip {
		return nil, nil, errors.New("invalid data")
	}

	iat, accessToken, err := s.tokenManager.NewJWT(data.UserId, data.Email, data.Role, s.accessTokenTTL)
	if err != nil {
		return nil, nil, err
	}
	refreshToken, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		return nil, nil, err
	}

	newData := repository.SessionData{
		UserId:   data.UserId,
		UserName: data.UserName,
		Email:    data.Email,
		Role:     data.Role,
		Ua:       ua,
		Ip:       ip,
		Exp:      s.refreshTokenTTL,
	}

	if err := s.repoSes.CreateSession(ctx, refreshToken, newData); err != nil {
		return nil, nil, err
	}

	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    refreshToken,
		MaxAge:   int(s.refreshTokenTTL.Seconds()),
		Path:     "/",
		Domain:   s.domain,
		Secure:   false,
		HttpOnly: true,
	}

	newToken := &domain.Token{
		Token: domain.TokenData{
			AccessToken: accessToken,
			Exp:         iat.Add(s.accessTokenTTL).Unix(),
		},
		UserId: data.UserId,
		Role:   data.Role,
		Name:   data.UserName,
	}

	return newToken, cookie, nil
}

func (s *SessionService) TokenParse(token string) (userId string, role string, err error) {
	claims, err := s.tokenManager.Parse(token)
	if err != nil {
		return "", "", err
	}
	return claims["userId"].(string), claims["role"].(string), err
}
