package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitializeDB() {
	var (
		err error
		dsn string
	)

	if err = godotenv.Load();
		err != nil {
		panic("Error loading .env file")
	}

	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             1 * time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/model/database",
		OutFile:      "gen.go",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		ModelPkgPath: "database",
	})

	g.UseDB(DB)
	g.GenerateAllTable()
	g.Execute()

	log.Println("Database model generation complete!")

	log.Println("Database connection established successfully!")
	if err := DB.AutoMigrate(); err != nil {
		log.Fatalf("Error during database migration: %v", err)
	}

	log.Println("Database migration completed successfully!")
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database connection is not initialized. Call InitializeDB first.")
	}
	return DB
}
