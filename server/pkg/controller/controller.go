package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// Book represents a book object
type Book struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Author  string  `json:"author"`
	Rating  float64 `json:"rating"`
	Preview string  `json:"preview"`
}

// ReadBooksHandler handles the request to fetch books
func ReadBooksHandler(c *fiber.Ctx) error {
	q := c.Query("q")
	if q == "" {
		q = "harry"
	}

	resp, err := http.Get("https://www.googleapis.com/books/v1/volumes?q=" + q)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to fetch books")
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to decode response body")
	}

	// Check if "items" key exists
	items, ok := data["items"].([]interface{})
	if !ok {
		return c.Status(http.StatusNotFound).SendString("Books not found")
	}

	var books []Book
	for _, item := range items {
		book := item.(map[string]interface{})
		volumeInfo, ok := book["volumeInfo"].(map[string]interface{})
		if !ok {
			continue
		}
	
		authors, ok := volumeInfo["authors"].([]interface{})
		var author string
		if ok && len(authors) > 0 {
			author = authors[0].(string)
		}
	
		var rating float64
		if avgRating, ok := volumeInfo["averageRating"].(float64); ok {
			rating = avgRating
		} else {
			rating = 0 // Or any default value you prefer
		}
	
		var previewURL string
		if imageLinks, ok := volumeInfo["imageLinks"].(map[string]interface{}); ok {
			if thumbnail, ok := imageLinks["thumbnail"].(string); ok {
				previewURL = thumbnail
			}
		}
	
		books = append(books, Book{
			ID:      book["id"].(string),
			Title:   volumeInfo["title"].(string),
			Author:  author,
			Rating:  rating,
			Preview: previewURL,
		})
	}
	
	return c.JSON(books)
}
