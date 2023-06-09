package internal

// Project Ricotta: Bechamel API
//
//	This file contains JWT (JSON Web Token) utility functions for the Bechamel API to use internally.

import (
	"errors"
	"fmt"
	"project-ricotta/bechamel-api/config"
	"project-ricotta/bechamel-api/model"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessJWT(userName string) (string, error) {
	return GenerateAccessJWTWithTTL(userName, config.RuntimeConfig.AccessJWTTTL())
}

func GenerateAccessJWTWithTTL(userName string, ttl int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": userName,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
	})
	return token.SignedString(config.RuntimeConfig.AccessJWTSigningKey())
}

func GenerateRefreshJWT(userName string) (string, error) {
	return GenerateAccessJWTWithTTL(userName, config.RuntimeConfig.RefreshJWTTTL())
}

func GenerateRefreshJWTWithTTL(userName string, ttl int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": userName,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
	})
	return token.SignedString(config.RuntimeConfig.RefreshJWTSigningKey())
}

// Validates a supplied access JWT token and returns the userName for the user
// represented by the token if the token is valid and not expired.

func VerifyAccessJWT(jwtTokenString string) (string, error) {
	if jwtTokenString == "" {
		return "", errors.New("JWT token to verify must be non-empty")
	}

	token, err := jwt.Parse(jwtTokenString, func(token *jwt.Token) (interface{}, error) {
		return config.RuntimeConfig.AccessJWTSigningKey(), nil
	})
	if token == nil || err != nil {
		// TODO: propogate parsing error from JWT up in a more useful form?
		return "", fmt.Errorf("could not parse supplied JWT: %w", err)
	}

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if err = claims.Valid(); err == nil {
				return fmt.Sprint(claims["userName"]), nil
			} else {
				return "", errors.New("supplied JWT expired or not yet valid")
			}
		} else {
			return "", errors.New("error retrieving user name from JWT token")
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", errors.New("malformed or invalid JWT token supplied")
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			return "", errors.New("supplied JWT token is expired")
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			return "", errors.New("supplied JWT does not contain valid time")
		} else {
			return "", errors.New("error validating supplied JWT token")
		}
	} else {
		return "", fmt.Errorf("could not parse supplied JWT: %w", err)
	}
}

// Validates a supplied access JWT token and returns the userName for the user
// represented by the token if the token is valid and not expired.

func VerifyRefreshJWT(jwtTokenString string) (string, error) {
	if jwtTokenString == "" {
		return "", errors.New("JWT token to verify must be non-empty")
	}

	token, err := jwt.Parse(jwtTokenString, func(token *jwt.Token) (interface{}, error) {
		return config.RuntimeConfig.RefreshJWTSigningKey(), nil
	})
	if token == nil || err != nil {
		// TODO: propogate parsing error from JWT up in a more useful form?
		return "", fmt.Errorf("could not parse supplied JWT: %w", err)
	}

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if err = claims.Valid(); err == nil {
				return fmt.Sprint(claims["userName"]), nil
			} else {
				return "", errors.New("supplied JWT expired or not yet valid")
			}
		} else {
			return "", errors.New("error retrieving user name from JWT token")
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", errors.New("malformed or invalid JWT token supplied")
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			return "", errors.New("supplied JWT token is expired")
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			return "", errors.New("supplied JWT does not contain valid time")
		} else {
			return "", errors.New("error validating supplied JWT token")
		}
	} else {
		return "", fmt.Errorf("could not parse supplied JWT: %w", err)
	}
}

func GetUserFromAuthHeader(authHeader string) (model.LasagnaLoveUser, error) {
	if authHeader == "" {
		return model.LasagnaLoveUser{}, errors.New("authorization header with JWT token required, not provided")
	}

	tokenString := authHeader[len("Bearer "):]
	userName, err := VerifyAccessJWT(tokenString)
	if err != nil {
		return model.LasagnaLoveUser{}, errors.New("provided authorization token could not be authorized")
	}

	userProfile, err := GetUserByUserName(userName)
	if err != nil {
		return model.LasagnaLoveUser{}, errors.New("profile not found for user with supplied userName")
	}

	return userProfile, nil
}
