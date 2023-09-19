package models

import (
	"time"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	CommunityID   int64  `json:"id" db:"community_id"`
	CommunityName string `json:"name" db:"community_name"`
	Introduction  string `json:"introduction,omitempty" db:"introduction"`
}

type ApiCommunityDetail struct {
	CreatedAt     time.Time
	CommunityName string `json:"name" db:"community_name"`
	Introduction  string `json:"introduction,omitempty" db:"introduction"`
}
