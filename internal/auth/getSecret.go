package auth

import (
	"just_simple_api/internal/utils"
	"sync"
)

var (
	secret string
	once   sync.Once
)

func LoadSecretKey() {
	secret = utils.LoadENV("SECRET_KEY")
}

func GetSecretKey() string {
	once.Do(LoadSecretKey)
	return secret
}
