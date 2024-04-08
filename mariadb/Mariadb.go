package mariadb

import (
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type Hotel struct {
	ID        int64  `json:"id"`                  // Hotel ID
	Name      string `json:"name"`                // Hotel Name
	Address   string `json:"address"`             // Hotel Address
	Price     int    `json:"price"`               // Hotel Price
	Score     int    `json:"score"`               // Hotel Score
	Brand     string `json:"brand"`               // Hotel Brand
	City      string `json:"city"`                // City
	StarName  string `json:"star_name,omitempty"` // Hotel Star Rating, from 1 star to 5 stars, or 1 diamond to 5 diamonds
	Business  string `json:"business,omitempty"`  // Business District
	Latitude  string `json:"latitude"`            // Latitude
	Longitude string `json:"longitude"`           // Longitude
	Pic       string `json:"pic,omitempty"`       // Hotel Picture
}

func (Hotel) TableName() string {
	return "tb_hotel"
}

const MySQLDSN = "louis:louis268@tcp(192.168.1.72:3306)/assignment?charset=utf8mb4&parseTime=True"

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       MySQLDSN,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         256,
		DisableWithReturning:      true,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
	}), &gorm.Config{})

	return db, err
}
