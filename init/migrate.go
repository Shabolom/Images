package migrate

import (
	"Service_Photo/config"
	"Service_Photo/iternal/domain"
	log "github.com/sirupsen/logrus"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
)

func Migrate() {
	db := config.DB
	id, _ := uuid.NewV4()

	// создаем объект миграции данная строка всегда статична (всегда такая)
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// id всех миграций кторые были проведены
			ID: id.String(),
			// переписываем так при создании таблицы изменяется только структура которую мы передаем
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(&domain.Photos{}).Error
				if err != nil {
					return err
				}
				return nil
			},
			// это метод отмены миграции ни разу не использовал
			Rollback: func(tx *gorm.DB) error {
				err := tx.DropTable("photos").Error
				if err != nil {
					return err
				}
				return nil
			},
		},
	})

	err := m.Migrate()
	if err != nil {
		log.WithField("component", "migration").Panic(err)
	}

	if err == nil {
		log.WithField("component", "migration").Info("Migration did run successfully")
	} else {
		log.WithField("component", "migration").Infof("Could not migrate: %v", err)
	}
}
