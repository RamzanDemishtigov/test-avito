package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type UsersSegments struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Segment_Slug string `gorm:"size:255;not null" json:"segment_slug"`
	User_Id		uint64	`gorm:"size:255;not null" json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
type UsersSegmentsDTO struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Segment_Slug []string `gorm:"size:255;not null" json:"segment_slug"`
	User_Id		uint64	`gorm:"size:255;not null" json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *UsersSegments) Prepare() {
	p.ID = 0
	p.Segment_Slug = html.EscapeString(strings.TrimSpace(p.Segment_Slug))
	p.User_Id = 0
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *UsersSegments) Validate() error {

	if p.Segment_Slug == "" {
		return errors.New("required segment_slug")
	}
	if p.User_Id == 0 {
		return errors.New("required user_id")
	}

	return nil
}

func (u *UsersSegmentsDTO) AddUserToSegments(db *gorm.DB)  (*[]UsersSegments,error) {
	CreatedUsersSegments := []UsersSegments{}
    for _, slug := range u.Segment_Slug {
        userSegment := UsersSegments{
            Segment_Slug: slug,
            User_Id:      u.User_Id,
        }
        err := db.Debug().Model(&UsersSegments{}).Create(&userSegment).Error
		CreatedUsersSegments = append(CreatedUsersSegments,userSegment)
        if err != nil {
            return &[]UsersSegments{},err
        }
    }
    return &CreatedUsersSegments,nil
}


func (p *UsersSegments) FindAllUsersSegments(db *gorm.DB) (*[]UsersSegments, error) {
	var err error
	AllUsersSegments := []UsersSegments{}
	err = db.Debug().Model(&UsersSegments{}).Limit(100).Find(&AllUsersSegments).Error
	if err != nil {
		return &[]UsersSegments{}, err
	}
	return &AllUsersSegments, nil
}

func (u *UsersSegments) FindUsersSegmentsByID(db *gorm.DB, uid uint64) (*[]UsersSegments, error) {
	var err error
	userssegments := []UsersSegments{}
	err = db.Debug().Model(&UsersSegments{}).Where("user_id = ?", uid).Limit(100).Find(&userssegments).Error
	if err != nil {
		return &[]UsersSegments{}, err
	}
	return &userssegments, nil
}