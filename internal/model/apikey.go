package models

import "github.com/google/uuid"

type ApiKey struct {
	Id                 int
	UserId             int
	KeyName            string
	Key                string
	AllowedAttemts     int
	AllowedIntervalSec int
	IsDeleted          bool
}

func NewKey(userId int, keyName string) *ApiKey {
	return &ApiKey{
		UserId:             userId,
		KeyName:            keyName,
		Key:                uuid.NewString(),
		AllowedAttemts:     3,
		AllowedIntervalSec: 600,
		IsDeleted:          false}
}

func NewKeyWithCustomAllowed(
	userId int,
	keyName string,
	allowedAttempts int,
	allowedIntervalSec int) *ApiKey {

	return &ApiKey{
		UserId:             userId,
		KeyName:            keyName,
		Key:                uuid.NewString(),
		AllowedAttemts:     allowedAttempts,
		AllowedIntervalSec: allowedIntervalSec,
		IsDeleted:          false}
}
