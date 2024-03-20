package models

import "time"

type Call struct {
	Id        int
	ApiKeyId  int
	Recepient string
	Code      int
	Timestamp time.Time
}
