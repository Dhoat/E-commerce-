package db


import(
	"fmt"
	"log"
	"ecommerce-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var  DB *gorm.DB
var err error

func Init() {
	    dsn := "root:password@tcp(localhost:3306)/ecommerce?charset=utf8&parseTime=True&loc=Local"
		DB, err = gorm.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
		}
	
		// Log the successful connection
		fmt.Println("Connected to MySQL database!")
	
		// Automigrate to create tables for the models
		DB.AutoMigrate(&models.User{}, &models.Product{},&models.Cart{},&models.Order{},&models.URL{})
}
func Close() {
    if err := DB.Close(); err != nil {
        log.Fatalf("Error closing the database: %v", err)
    }
}

