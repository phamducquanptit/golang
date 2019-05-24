package models

import (
	"sync"
	"time"
)

type BaseModel struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}

type UserInfo struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TodoInfo struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Status    uint64 `json:"status"`
	UserID    uint64 `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TodoList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*TodoInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
