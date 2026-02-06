package entities

import (
	"fmt"
	"strings"
	"time"
)

type Product struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	CategoryID    string    `json:"category_id"`
	Weight        string    `json:"weight"`
	Brand         string    `json:"brand"`
	Country       string    `json:"country"`
	ImageURL      string    `json:"image_url"`
	ProductURL    string    `json:"product_url"`
	IsAvailable   bool      `json:"is_available"`
	Barcode       string    `json:"barcode"`
	SKU           string    `json:"sku"`
	Rating        float64   `json:"rating"`
	ReviewsCount  int       `json:"reviews_count"`
	StoreID       string    `json:"store_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProductBuilder struct {
	product *Product
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{
		product: &Product{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func (b *ProductBuilder) SetID(id string) *ProductBuilder {
	b.product.ID = id
	return b
}

func (b *ProductBuilder) SetName(name string) *ProductBuilder {
	b.product.Name = name
	return b
}

func (b *ProductBuilder) SetDescription(description string) *ProductBuilder {
	b.product.Description = description
	return b
}

func (b *ProductBuilder) SetCategoryID(categoryID string) *ProductBuilder {
	b.product.CategoryID = categoryID
	return b
}

func (b *ProductBuilder) SetWeight(weight string) *ProductBuilder {
	b.product.Weight = weight
	return b
}

func (b *ProductBuilder) SetBrand(brand string) *ProductBuilder {
	b.product.Brand = brand
	return b
}

func (b *ProductBuilder) SetCountry(country string) *ProductBuilder {
	b.product.Country = country
	return b
}

func (b *ProductBuilder) SetImageURL(imageURL string) *ProductBuilder {
	b.product.ImageURL = imageURL
	return b
}

func (b *ProductBuilder) SetProductURL(productURL string) *ProductBuilder {
	b.product.ProductURL = productURL
	return b
}

func (b *ProductBuilder) SetIsAvailable(isAvailable bool) *ProductBuilder {
	b.product.IsAvailable = isAvailable
	return b
}

func (b *ProductBuilder) SetBarcode(barcode string) *ProductBuilder {
	b.product.Barcode = barcode
	return b
}

func (b *ProductBuilder) SetSKU(sku string) *ProductBuilder {
	b.product.SKU = sku
	return b
}

func (b *ProductBuilder) SetRating(rating float64) *ProductBuilder {
	b.product.Rating = rating
	return b
}

func (b *ProductBuilder) SetReviewsCount(reviewsCount int) *ProductBuilder {
	b.product.ReviewsCount = reviewsCount
	return b
}

func (b *ProductBuilder) SetStoreID(storeID string) *ProductBuilder {
	b.product.StoreID = storeID
	return b
}

// Build создает Product с валидацией
func (b *ProductBuilder) Build() (*Product, error) {
	if err := b.product.Validate(); err != nil {
		return nil, err
	}
	b.product.UpdatedAt = time.Now()
	return b.product, nil
}

func (p *Product) Validate() error {
	if strings.TrimSpace(p.ID) == "" {
		return fmt.Errorf("ID товара не может быть пустым")
	}
	if strings.TrimSpace(p.Name) == "" {
		return fmt.Errorf("название товара не может быть пустым")
	}
	if strings.TrimSpace(p.CategoryID) == "" {
		return fmt.Errorf("ID категории не может быть пустым")
	}
	if strings.TrimSpace(p.StoreID) == "" {
		return fmt.Errorf("ID магазина не может быть пустым")
	}
	return nil
}

type ProductFactory struct{}

func (f *ProductFactory) CreateFromLentaAPI(data map[string]interface{}) (*Product, error) {
	builder := NewProductBuilder()
	
	builder.SetID(fmt.Sprintf("%v", data["id"]))
	builder.SetName(fmt.Sprintf("%v", data["name"]))
	
	if category, ok := data["category"].(map[string]interface{}); ok {
		builder.SetCategoryID(fmt.Sprintf("%v", category["id"]))
	}
	
	if brand, ok := data["brand"].(map[string]interface{}); ok {
		builder.SetBrand(fmt.Sprintf("%v", brand["name"]))
	}
	
	if image, ok := data["image"].(map[string]interface{}); ok {
		builder.SetImageURL(fmt.Sprintf("%v", image["url"]))
	}
	
	return builder.Build()
}