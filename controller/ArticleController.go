package controller

import (
	"compro-backend/config"
	"compro-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct articlesController
type ArticleResponse struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	CategoryName string `json:"category_name"`
}

// GetArticles adalah fungsi untuk mengambil semua artikel
func GetArticles(c *gin.Context) {
	var articles []models.Articles

	// Mencari semua artikel
	if err := config.DB.Preload("Category").Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// slice untuk menyimpan response articles
	var articlesSlice []ArticleResponse

	// Looping untuk mengambil ID dan Title
	for _, article := range articles {
		articlesSlice = append(articlesSlice, ArticleResponse{
			ID:           article.ID,
			Title:        article.Title,
			Description:  article.Description,
			CategoryName: article.Category.Name,
		})
	}

	// Mengembalikan respons dengan data articles yang dipilih
	c.JSON(http.StatusOK, gin.H{"data": articlesSlice})
}

// GetArticle adalah fungsi untuk mengambil satu artikel
func GetArticle(c *gin.Context) {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Inisialisasi artikel
	var article models.Articles

	// Mencari artikel berdasarkan ID
	if err := config.DB.Preload("Category").First(&article, id).First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Mengembalikan respons dengan data artikel yang dipilih
	c.JSON(http.StatusOK, gin.H{
		"data": ArticleResponse{
			ID:           article.ID,
			Title:        article.Title,
			Description:  article.Description,
			CategoryName: article.Category.Name,
		},
	})
}

// CreateArticle adalah fungsi untuk membuat artikel baru
func CreateArticle(c *gin.Context) {
	// Inisialisasi artikel
	var article models.Articles

	// Melakukan binding data dari form ke artikel
	if err := c.ShouldBind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validasi form is null
	if article.Title == "" || article.Description == "" || article.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field can not be null"})
		return
	}

	// Cek error saat menyimpan ke database
	if err := config.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Cek error saat menyimpan ke database
	config.DB.Create(&article)

	// Mengembalikan respons
	c.JSON(http.StatusOK, gin.H{
		"message": "Article created successfully",
	})
}

// UpdateArticle adalah fungsi untuk mengubah artikel
func UpdateArticle(c *gin.Context) {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Inisialisasi artikel
	var article models.Articles

	// Mencari artikel berdasarkan ID
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Melakukan binding data dari form ke artikel
	if err := c.ShouldBind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek error saat menyimpan ke database
	if err := config.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan respons
	c.JSON(http.StatusOK, gin.H{
		"message": "Article updated successfully",
	})
}

// DeleteArticle adalah fungsi untuk menghapus artikel
func DeleteArticle(c *gin.Context) {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Inisialisasi artikel
	var article models.Articles

	// Mencari artikel berdasarkan ID
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Cek error saat menyimpan ke database
	if err := config.DB.Delete(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan respons
	c.JSON(http.StatusOK, gin.H{
		"message": "Article deleted successfully",
	})
}
