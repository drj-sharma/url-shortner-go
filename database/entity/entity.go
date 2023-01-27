package entity

type Url struct {
	ID           uint   `gorm:"primaryKey"`
	OriginalUrl  string `gorm:"not null"`
	ShortUrlCode string `gorm:"index:idx_short_url_code,unique;not null"`
}
