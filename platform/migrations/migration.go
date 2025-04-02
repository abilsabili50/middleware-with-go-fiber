package migrations

import (
	"log"

	"github.com/abilsabili50/middleware-with-go-fiber/app/model"
	"gorm.io/gorm"
)

func AutoMigrate(conn *gorm.DB) {
	// drop table before migrations (ONLY FOR DEV STAGE)
	if err := conn.Migrator().DropTable(model.Models...); err != nil {
		log.Fatalln("error occured while drop table before migrations:", err.Error())
	}

	// auto migrate all entities
	if err := conn.AutoMigrate(model.Models...); err != nil {
		log.Fatalln("error occured while performing migrations:", err.Error())
	}

	log.Println("Database reset and migrated successfully")
}
