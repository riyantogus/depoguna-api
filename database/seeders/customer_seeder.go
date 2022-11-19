package seeders

import (
	"depoguna-api/models"

	"gorm.io/gorm"
)

// CreateCustomer is seeder to create customer
func CreateCustomer(db *gorm.DB) error {
	var customers = []models.Customer{
		{UserId: 1, Name: "Sallie Walker", Email: "nchamplin@hulas.me", Gender: "female", DateOfBirth: "1985-05-27", Mobile: "14103493830", Address: "8714 Purdy Orchard"},
		{UserId: 2, Name: "Velma D. Brodsky", Email: "velmadbrodsky@teleworm.us", Gender: "female", DateOfBirth: "1992-10-16", Mobile: "6192244211", Address: "2437 Holden Street"},
		{UserId: 3, Name: "Ronald L. Neil", Email: "ronaldLneil@teleworm.us", Gender: "male", DateOfBirth: "1995-05-21", Mobile: "8303363815", Address: "1278 Morris Street"},
		{UserId: 4, Name: "James E. McNair", Email: "jamesemcnair@dayrep.com", Gender: "male", DateOfBirth: "1999-07-17", Mobile: "9856520521", Address: "2499 Woodland Avenue"},
		{UserId: 5, Name: "Clarence J. Baker", Email: "clarencejbaker@teleworm.us", Gender: "male", DateOfBirth: "1969-04-19", Mobile: "3054919234", Address: "592 Poplar Lane"},
		{UserId: 6, Name: "Frederick R. Myler", Email: "frederickrmyler@jourrapide.com", Gender: "male", DateOfBirth: "1973-12-14", Mobile: "7577266946", Address: "2983 Jefferson Street"},
		{UserId: 7, Name: "James H. Campbell", Email: "jameshcampbell@armyspy.com", Gender: "male", DateOfBirth: "1969-07-07", Mobile: "2396439952", Address: "1803 Wilkinson Court"},
		{UserId: 8, Name: "Matthew S. Torres", Email: "matthewstorres@dayrep.com", Gender: "male", DateOfBirth: "1938-06-13", Mobile: "5108241170", Address: "243 Green Avenue"},
		{UserId: 9, Name: "Donald L. Cain", Email: "donaldlcain@dayrep.com", Gender: "male", DateOfBirth: "1971-08-28", Mobile: "2035237260", Address: "1074 Cook Hill Road"},
		{UserId: 10, Name: "Daniel J. Helton", Email: "danieljhelton@jourrapide.com", Gender: "male", DateOfBirth: "1946-12-03", Mobile: "8182749509", Address: "898 Oakway Lane"},
		{UserId: 11, Name: "Billy T. Moore", Email: "billytmoore@rhyta.com", Gender: "male", DateOfBirth: "1959-01-03", Mobile: "6603332437", Address: "3332 Fairmont Avenue"},
		{UserId: 12, Name: "Gerald L. Ball", Email: "geraldlball@jourrapide.com", Gender: "male", DateOfBirth: "1983-12-30", Mobile: "4803969030", Address: "305 Hillside Street"},
		{UserId: 13, Name: "Walter C. Cervantes", Email: "walterccervantes@armyspy.com", Gender: "male", DateOfBirth: "1957-06-07", Mobile: "2155589369", Address: "500 Pheasant Ridge Road"},
		{UserId: 14, Name: "Harry T. Niven", Email: "harrytniven@rhyta.com", Gender: "male", DateOfBirth: "1957-06-07", Mobile: "8067302282", Address: "2460 Smithfield Avenue"},
		{UserId: 15, Name: "Horace C. Miles", Email: "horacecmiles@rhyta.com", Gender: "male", DateOfBirth: "1962-05-01", Mobile: "9405282735", Address: "970 Alexander Drive"},
	}
	return db.Create(&customers).Error
}
