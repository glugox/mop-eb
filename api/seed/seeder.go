package seed

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/glugox/mop/api/models"
	"github.com/jaswdr/faker"
	"github.com/jinzhu/gorm"
	"log"
)

var maxPurchases = 10

// Seed TODO: For mor complex db, create migration versioning
func Seed(db *gorm.DB) {

	faker := faker.New()

	// Store created products in slice, so we can use it to seed random purchases for each user when created.
	var createdProductIds []uint32

	// Drop tables
	err := db.Debug().DropTableIfExists(&models.Purchase{}, &models.User{}, &models.Product{}, &models.Widget{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	// Create tables
	err = db.Debug().AutoMigrate(&models.User{}, &models.Product{}, &models.Purchase{}, &models.Widget{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// Register Foreign keys
	err = db.Debug().Model(&models.Purchase{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Purchase{}).AddForeignKey("product_id", "products(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	// Seed Products
	for i:=0; i < 10; i++ {

		dummyP := models.Product{
			Name: faker.Car().Model(),
			Price: faker.IntBetween(10_000, 1_000_000),
		}

		err = db.Debug().Model(&models.Product{}).Create(&dummyP).Error
		if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
			log.Fatalf("cannot seed products table: %v", err)
		}

		// If faker chooses same product name, we will not create any product
		if dummyP.ID != 0 {
			// created!
			createdProductIds = append(createdProductIds, dummyP.ID)
		}


		fmt.Printf("%+v", dummyP)
	}

	// Seed Widgets
	rpWidget := models.Widget{
		Name: "Recent Purchases",
		Ordering: 1,
		Type: int(models.TypeTable),
		Alias: "recent_purchases",
		IsPrivate: true,
	}
	err = db.Debug().Model(&models.Product{}).Create(&rpWidget).Error
	if err != nil {
		log.Fatalf("cannot seed widgets table: %v", err)
	}

	// Seed Widgets
	pbdWidget := models.Widget{
		Name: "Purchases By Day",
		Ordering: 2,
		Type: int(models.TypeTable),
		Alias: "purchases_by_day",
		IsPrivate: false,
	}
	err = db.Debug().Model(&models.Product{}).Create(&pbdWidget).Error
	if err != nil {
		log.Fatalf("cannot seed widgets table: %v", err)
	}

	// Seed Widgets
	lvtWidget := models.Widget{
		Name: "Purchases Last Year vs This Year",
		Ordering: 1,
		Type: int(models.TypeChart),
		Alias: "purchases_last_vs_this_year",
		IsPrivate: true,
	}
	err = db.Debug().Model(&models.Product{}).Create(&lvtWidget).Error
	if err != nil {
		log.Fatalf("cannot seed widgets table: %v", err)
	}


	// Seed Users
	for i:=0; i < 10; i++ {
		dummuU := models.User{
			Email: faker.Internet().Email(),
			FirstName: faker.Person().FirstName(),
			LastName: faker.Person().LastName(),
			Password: faker.Internet().Password(),
		}
		err = db.Debug().Model(&models.Product{}).Create(&dummuU).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		for j := 0; j < maxPurchases ; j++ {

			// User can purchase multiple products with same id, if rand happens
			// to chose same productId over loop
			productIdIndex := rand.Intn(len(createdProductIds))
			productId := createdProductIds[productIdIndex]

			// Create user purchase record
			up := models.Purchase{
				UserID:    dummuU.ID,
				ProductID: productId,
			}
			err = db.Debug().Model(&models.Purchase{}).Create(&up).Error
			if err != nil {
				log.Fatalf("cannot seed user purchase table. Row index: %v, createdProductIds: %v , %v", productIdIndex, createdProductIds, err)
			}
		}

		// Seed user widgets
		// Seed Widgets
		lvtWidget := models.Widget{
			Name: "Purchases Last Year vs This Year",
			Ordering: 1,
			UserID: dummuU.ID,
			Type: int(models.TypeChart),
			Alias: "purchases_last_vs_this_year",
		}
		err = db.Debug().Model(&models.Product{}).Create(&lvtWidget).Error
		if err != nil {
			log.Fatalf("cannot seed widgets table: %v", err)
		}
	}


}