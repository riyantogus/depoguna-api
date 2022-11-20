package seeders

import (
	"depoguna-api/models"
	"depoguna-api/utils"

	"gorm.io/gorm"
)

// CreateUser is seeder to create user
func CreateUser(db *gorm.DB) error {
	encrypt := utils.NewEncryptUtil()
	var users = []models.User{
		{Name: "Sallie Walker", Email: "nchamplin@hulas.me", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Velma D. Brodsky", Email: "velmadbrodsky@teleworm.us", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Ronald L. Neil", Email: "ronaldLneil@teleworm.us", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "James E. McNair", Email: "jamesemcnair@dayrep.com", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Clarence J. Baker", Email: "clarencejbaker@teleworm.us", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Frederick R. Myler", Email: "frederickrmyler@jourrapide.com", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "James H. Campbell", Email: "jameshcampbell@armyspy.com", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Matthew S. Torres", Email: "matthewstorres@dayrep.com", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Donald L. Cain", Email: "donaldlcain@dayrep.com", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Daniel J. Helton", Email: "danieljhelton@jourrapide.com", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Billy T. Moore", Email: "billytmoore@rhyta.com", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Gerald L. Ball", Email: "geraldlball@jourrapide.com", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Walter C. Cervantes", Email: "walterccervantes@armyspy.com", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Harry T. Niven", Email: "harrytniven@rhyta.com", Password: encrypt.HashAndSalt([]byte("12345678"))},
		{Name: "Horace C. Miles", Email: "horacecmiles@rhyta.com", Password: encrypt.HashAndSalt([]byte("12345678"))},
	}
	return db.Create(&users).Error
}
