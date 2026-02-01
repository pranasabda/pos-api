package database

import (
	"KASIR-API/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB menginisialisasi koneksi database
func ConnectDB() {
	// Membaca konfigurasi dari file .env via Viper
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Peringatan: File .env tidak ditemukan")
	}
	viper.ReadInConfig()

	// Konfigurasi Logger GORM agar Query SQL muncul di terminal
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable default_query_exec_mode=exec",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Alternatif lain: matikan prepared statement cache di tingkat GORM
		PrepareStmt: false,
		Logger:      newLogger,
	})
	if err != nil {
		log.Fatal("Koneksi Database Gagal: ", err)
	}

	// Connection Pooling
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Sinkronisasi tabel otomatis
	db.AutoMigrate(&models.Category{}, &models.Product{})

	DB = db
	log.Println("Database terkoneksi dan migrasi selesai")
}
