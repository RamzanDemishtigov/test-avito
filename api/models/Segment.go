package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Segment struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Slug      string    `gorm:"size:255;not null;unique" json:"slug"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *Segment) Prepare() {
	s.ID = 0
	s.Slug = html.EscapeString(strings.TrimSpace(s.Slug))
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}

func (s *Segment) Validate() error {

	if s.Slug == "" {
		return errors.New("required slug")
	}

	return nil
}

func (s *Segment) SaveSegment(db *gorm.DB) (*Segment, error) {
	var err error  = db.Debug().Model(&Segment{}).Create(&s).Error
	if err != nil {
		return &Segment{}, err
	}
	// if s.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", s.AuthorID).Take(&s.Author).Error
	// 	if err != nil {
	// 		return &Segment{}, err
	// 	}
	// }
	return s, nil
}

func (s *Segment) FindAllSegments(db *gorm.DB) (*[]Segment, error) {
	var err error
	segments := []Segment{}
	err = db.Debug().Model(&Segment{}).Limit(100).Find(&segments).Error
	if err != nil {
		return &[]Segment{}, err
	}
	return &segments, nil
}

func (s *Segment) FindSegmentByID(db *gorm.DB, sid uint64) (*Segment, error) {
	var err error = db.Debug().Model(&Segment{}).Where("id = ?", sid).Take(&s).Error
	if err != nil {
		return &Segment{}, err
	}
	return s, nil
}

func (s *Segment) UpdateSegment(db *gorm.DB, sid uint64) (*Segment, error) {

	var err error
	db = db.Debug().Model(&Segment{}).Where("id = ?", sid).Take(&Segment{}).UpdateColumns(
		map[string]interface{}{
			"slug":      s.Slug,
			"updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Segment{}).Where("id = ?", sid).Take(&s).Error
	if err != nil {
		return &Segment{}, err
	}
	return s, nil
}

func (s *Segment) DeleteSegment(db *gorm.DB, sslug uint64) (int64, error) {

	db = db.Debug().Model(&Segment{}).Where("slug = ?", sslug).Take(&Segment{}).Delete(&Segment{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Segment not found")
		}
		return 0, db.Error
	}

	db = db.Debug().Model(&UsersSegments{}).Where("segment_slug = ?", sslug).Take(&UsersSegments{}).Delete(&UsersSegments{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Segment not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
