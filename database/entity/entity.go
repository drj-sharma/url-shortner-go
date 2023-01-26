package entity

type Url struct {
	ID           uint   `gorm:"primaryKey"`
	OriginalUrl  string `gorm:"not null"`
	ShortUrlCode string `gorm:"unique;not null"`
}
