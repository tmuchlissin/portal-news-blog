package seed

import (
	"portal-news-blog/internal/core/domain/model"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	bytes, err := bcrypt.GenerateFromPassword([]byte("admin123"), 14)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to hash password")
	}

	admin := model.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: string(bytes),
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: admin.Email}).Error; err != nil {
		log.Fatal().Err(err).Msg("failed to seed admin user")
	}

	log.Info().Msg("admin user seeded successfully")
}
