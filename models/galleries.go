package models

import (
	"github.com/jinzhu/gorm"
)

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
	ById(id uint) (*Gallery, error)
	Create(gallery *Gallery) error
	Update(gallery *Gallery) error
}

var _ GalleryDB = &GalleryGorm{}

func (gg *GalleryGorm) Create(gallery *Gallery) error {
	return gg.db.Create(gallery).Error
}

func (gg *GalleryGorm) Update(gallery *Gallery) error {
	return gg.db.Save(gallery).Error
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

func (gg *GalleryGorm) ById(id uint) (*Gallery, error) {
	var gallery Gallery
	db := gg.db.Where("id = ?", id)
	err := first(db, &gallery)
	return &gallery, err
}

func (gv *galleryValidator) Create(gallery *Gallery) error {
	if err := runGalleryValFuncs(gallery,
		gv.userIDRequired,
		gv.titleRequired); err != nil {
		return err
	}
	return gv.GalleryDB.Create(gallery)
}

func (gv *galleryValidator) Update(gallery *Gallery) error {
	if err := runGalleryValFuncs(gallery,
		gv.userIDRequired,
		gv.titleRequired); err != nil {
		return err
	}
	return gv.GalleryDB.Update(gallery)
}

type galleryValFunc func(*Gallery) error

func runGalleryValFuncs(gallery *Gallery, fns ...galleryValFunc) error {
	for _, fn := range fns {
		if err := fn(gallery); err != nil {
			return err
		}
	}
	return nil
}

func (gv *galleryValidator) userIDRequired(g *Gallery) error {
	if g.UserID <= 0 {
		return ErrUserIDRequired
	}
	return nil
}

func (gv *galleryValidator) titleRequired(g *Gallery) error {
	if g.Title == "" {
		return ErrTitleRequired
	}
	return nil
}
