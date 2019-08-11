package models

import "github.com/jinzhu/gorm"

// image container resource that visitors view
type Gallery struct {
	gorm.Model
	UserID uint   `gorm:"not_null;index"`
	Title  string `gorm:"not_null"`
}

type GalleryService interface {
	GalleryDB
}

type GalleryDB interface {
	Create(gallery *Gallery) error
}

var _ GalleryDB = &GalleryGorm{}

func (gg *GalleryGorm) Create(gallery *Gallery) error {
	return gg.db.Create(gallery).Error
}

type GalleryGorm struct {
	db *gorm.DB
}

type galleryValidator struct {
	GalleryDB
}

type galleryService struct {
	GalleryDB
}

func NewGalleryService(db *gorm.DB) GalleryService {
	return &galleryService{
		GalleryDB: &galleryValidator{
			&GalleryGorm{db},
		},
	}
}
