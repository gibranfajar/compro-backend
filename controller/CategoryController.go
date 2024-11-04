package controller

import (
	"compro-backend/config"
	"compro-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CategoryResponse adalah struktur untuk menyimpan data kategori
type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// GetCategories adalah fungsi untuk mengambil semua kategori
func GetCategories(c *gin.Context) {
	// Inisialisasi slice kategori
	var categories []models.Categories

	// Mencari semua kategori dari database
	config.DB.Find(&categories)

	// Slice untuk menyimpan respons kategori
	var categoriesSlice []CategoryResponse

	// Looping untuk mengambil ID dan Name
	for _, category := range categories {
		categoriesSlice = append(categoriesSlice, CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	// Mengembalikan respons
	c.JSON(http.StatusOK, gin.H{"data": categoriesSlice})
}

// CreateCategory adalah fungsi untuk membuat kategori baru
func GetCategory(c *gin.Context) {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Inisialisasi kategori
	var category models.Categories // Ubah dari slice menjadi tipe tunggal

	// Mencari kategori berdasarkan ID
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Mengembalikan respons kategori yang dipilih sebagai objek
	c.JSON(http.StatusOK, gin.H{
		"data": CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		},
	})
}

// CreateCategory adalah fungsi untuk membuat kategori baru
func CreateCategory(c *gin.Context) {
	// inisialisasi kategori
	var category models.Categories

	// Melakukan binding data dari form ke kategori
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validasi form is null
	if category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	// Cek error saat menyimpan ke database
	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Cek error saat menyimpan ke database
	config.DB.Create(&category)

	// Mengembalikan respons
	c.JSON(http.StatusOK, gin.H{
		"message": "Category created successfully",
	})
}

// UpdateCategory adalah fungsi untuk memperbarui kategori
func UpdateCategory(c *gin.Context) {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Inisialisasi kategori
	var category models.Categories

	// Mencari kategori berdasarkan ID
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Melakukan binding data dari form ke kategori
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi jika name kosong
	if category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	// Cek error saat menyimpan ke database
	if err := config.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan respons
	c.JSON(http.StatusOK, gin.H{
		"message": "Category updated successfully",
	})
}

// DeleteCategory adalah fungsi untuk menghapus kategori
func DeleteCategory(c *gin.Context) {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Inisialisasi kategori
	var category models.Categories

	// Mencari kategori berdasarkan ID
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Cek error saat menyimpan ke database
	if err := config.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan respons
	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
	})
}
