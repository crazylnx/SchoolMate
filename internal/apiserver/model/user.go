package model

import "github.com/google/uuid"

// User 用户信息
type User struct {
	uid  uuid.UUID
	name string
}
