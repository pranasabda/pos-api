package models

// Product representasi tabel produk dengan relasi ke Category
type Product struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID uint   `json:"category_id"`

	// Category field ini digunakan untuk Challenge Join/Preload
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}
