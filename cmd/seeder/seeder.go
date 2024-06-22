package seeder

import (
	"bbdk/domain/entity"
	gorm3 "bbdk/domain/repository/transaction/gorm"
	gorm2 "bbdk/domain/repository/user/gorm"
	"bbdk/infrastructure/godotenv"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func Seed() error {
	env := godotenv.NewEnv()
	env.Load()
	db, err := gorm.Open(postgres.Open(env.DATABASE_HOST), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := gorm2.NewUserRepository(db)
	transactionRepo := gorm3.NewTransactionRepository(db)

	for i := 0; i < 10; i++ {
		user := entity.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: faker.Password(),
		}
		err := userRepo.CreateUser(&user)
		if err != nil {
			log.Printf("Failed to create user: %s", err.Error())
			return err
		}
		fmt.Printf("Created User: %v\n", user)

		for j := 0; j < 5; j++ {
			transaction := entity.Transaction{
				UserID: user.ID,
				Amount: rand.Float64() * 100,
				Type:   randomTransactionType(),
			}
			err := transactionRepo.CreateTransaction(&transaction)
			if err != nil {
				log.Printf("Failed to create transaction: %s", err.Error())
				return err
			}
			fmt.Printf("Created Transaction: %v\n", transaction)
		}
	}

	return nil
}

func randomTransactionType() string {
	types := []string{"game_referral", "p2e", "seazen_zero"}
	rand.Seed(time.Now().UnixNano())
	return types[rand.Intn(len(types))]
}
