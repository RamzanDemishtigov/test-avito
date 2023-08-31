package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"main/api/models"
)

var users = []models.User{
	{
	},
	{
	},
}

var segments = []models.Segment{
	{
		Slug:   "AVITO_1",
	},
	{
		Slug:   "AVITO_2",
	},
}

var userssegments = []models.UsersSegments{
	{
		User_Id: 1,
		Segment_Slug: "AVITO_1",
	},
	{
		User_Id: 2,
		Segment_Slug: "AVITO_2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Segment{}, &models.User{}, &models.UsersSegments{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Segment{}, &models.UsersSegments{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.UsersSegments{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.UsersSegments{}).AddForeignKey("segment_slug", "segments(slug)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		userssegments[i].User_Id = users[i].ID

		err = db.Debug().Model(&models.Segment{}).Create(&segments[i]).Error
		if err != nil {
			log.Fatalf("cannot seed segments table: %v", err)
		}
		userssegments[i].Segment_Slug = segments[i].Slug

		err = db.Debug().Model(&models.UsersSegments{}).Create(&userssegments[i]).Error
		if err != nil {
			log.Fatalf("cannot seed userssegments table: %v", err)
		}
	}
}
