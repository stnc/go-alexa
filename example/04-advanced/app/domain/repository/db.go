package dbRepository

import (
	"avia/app/domain/entity"
	services "avia/app/services"
	"fmt"
	"github.com/hypnoglow/gormzap"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"os"

	_ "github.com/lib/pq" // here
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
)

var DB *gorm.DB

// Repositories strcut
type Repositories struct {
	Reminder services.ReminderAppInterface
	DB       *gorm.DB
}

// DbConnect initial
func DbConnect() *gorm.DB {
	dbdriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	gormAdvancedLogger := os.Getenv("GORM_ZAP_LOGGER")
	debug := os.Getenv("MODE")
	//	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword) //bu postresql

	//DBURL := "root:sel123C#@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local" //mysql
	var DBURL string

	if dbdriver == "mysql" {
		DBURL = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"
	} else if dbdriver == "postgres" {
		DBURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", dbHost, dbPort, dbUser, dbPassword, dbName) //Build connection string
	}

	// dsn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s sslmode=disable",
	// HOST, PORT, username, password, database)

	db, err := gorm.Open(dbdriver, DBURL)
	db.Set("gorm:table_options", "charset=utf8")

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	if debug == "DEBUG" && gormAdvancedLogger == "ENABLE" {
		db.LogMode(true)
		log := zap.NewExample()
		db.SetLogger(gormzap.New(log, gormzap.WithLevel(zap.DebugLevel)))
	} else if debug == "DEBUG" || debug == "TEST" && gormAdvancedLogger == "ENABLE" {
		db.LogMode(true)
	} else if debug == "RELEASE" {
		fmt.Println(debug)
		db.LogMode(false)
	}
	DB = db

	db.SingularTable(true)

	return db
}

// RepositoriesInit initial
func RepositoriesInit(db *gorm.DB) (*Repositories, error) {

	return &Repositories{
		Reminder: ReminderRepositoryInit(db),
		DB:       db,
	}, nil
}

func (s *Repositories) Automigrate() error {
	return s.DB.AutoMigrate(&entity.Reminder{}).Error
}
