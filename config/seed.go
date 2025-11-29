package config

import (
	"log"
	"ticketing-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin() {
	var count int64
	DB.Model(&models.User{}).Where("role = ?", "admin").Count(&count)

	if count > 0 {
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	admin := models.User{
		Name:     "Administrator",
		Email:    "admin@ticketing.com",
		Password: string(hash),
		Role:     "admin",
	}

	if err := DB.Create(&admin).Error; err != nil {
		log.Println("Failed seeding admin:", err)
	} else {
		log.Println("Admin account seeded (email: admin@ticketing.com / pass: admin123)")
	}
}
