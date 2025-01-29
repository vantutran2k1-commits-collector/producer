package models

import (
	"github.com/google/uuid"
	"time"
)

type CollectionJob struct {
	Id            uuid.UUID `json:"id" gorm:"column:id"`
	CollectedFrom time.Time `json:"collected_from" gorm:"column:collected_from"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
}
