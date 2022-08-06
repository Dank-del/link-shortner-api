package types

type ShortenedLink struct {
	LinkId       string `gorm:"primaryKey"`
	ProvidedLink string `json:"provided_link"`
	ShortenedAt  int64  `json:"autoCreateTime"`
}
