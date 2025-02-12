package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Total     float64        `gorm:"not null"`
	Items     []InvoiceItem  `gorm:"foreignKey:InvoiceID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
}

type InvoiceItem struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	InvoiceID uuid.UUID  `gorm:"type:uuid;not null;index"`
	ProductID uuid.UUID  `gorm:"type:uuid;not null;index"`
	Availabe  int        `gorm:"not null"`
	Price     float64    `gorm:"not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}


func (invoice *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	invoice.ID = uuid.New()

	return
}

func (item *InvoiceItem) BeforeCreate(tx *gorm.DB) (err error) {
	if item.ID == uuid.Nil {
		item.ID = uuid.New()
	}
	return
}