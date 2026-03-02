package seeder

import (
	"context"
	"fmt"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/logger"

	"go.uber.org/zap"
	"golang.org/x/exp/rand"
)

type productSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewProductSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *productSeeder {
	return &productSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *productSeeder) Seed() error {
	merchants, err := r.db.GetMerchants(r.ctx, db.GetMerchantsParams{
		Column1: "",
		Limit:   20,
		Offset:  0,
	})
	if err != nil {
		r.logger.Error("Failed to get merchants:", zap.Any("error", err))
		return err
	}

	categories, err := r.db.GetCategories(r.ctx, db.GetCategoriesParams{
		Column1: "",
		Limit:   20,
		Offset:  0,
	})
	if err != nil {
		r.logger.Error("Failed to get categories:", zap.Any("error", err))
		return err
	}

	if len(merchants) == 0 || len(categories) == 0 {
		r.logger.Error("No merchants or categories found, skipping seeding")
		return nil
	}

	productNames := []string{
		"Smartphone", "Laptop", "Wireless Earbuds", "Gaming Mouse", "Mechanical Keyboard",
		"Smartwatch", "Power Bank", "Bluetooth Speaker", "External Hard Drive", "Monitor",
	}
	brands := []string{"Samsung", "Apple", "Sony", "Logitech", "Razer", "Xiaomi", "HP", "Dell", "Acer", "Asus"}
	images := []string{
		"image1.jpg", "image2.jpg", "image3.jpg", "image4.jpg", "image5.jpg",
		"image6.jpg", "image7.jpg", "image8.jpg", "image9.jpg", "image10.jpg",
	}

	for i := 0; i < 10; i++ {
		merchant := merchants[rand.Intn(len(merchants))]
		category := categories[rand.Intn(len(categories))]

		name := productNames[rand.Intn(len(productNames))]
		price := int32(rand.Intn(5_000_000) + 50_000)
		countInStock := int32(rand.Intn(100) + 1)

		brand := brands[rand.Intn(len(brands))]
		weight := int32(rand.Intn(5000) + 100)
		slug := fmt.Sprintf("%s-%d", name, rand.Intn(1000))
		image := images[rand.Intn(len(images))]
		barcode := fmt.Sprintf("BC-%d", rand.Intn(9_999_999))
		desc := fmt.Sprintf("Description for %s", name)

		_, err := r.db.CreateProduct(r.ctx, db.CreateProductParams{
			MerchantID:   merchant.MerchantID,
			CategoryID:   category.CategoryID,
			Name:         name,
			Description:  &desc,
			Price:        price,
			CountInStock: countInStock,
			Brand:        &brand,
			Weight:       &weight,
			SlugProduct:  &slug,
			ImageProduct: &image,
			Barcode:      &barcode,
		})

		if err != nil {
			r.logger.Error("Failed to create product:", zap.Any("error", err))
			return err
		}

	}

	r.logger.Info("Product seeding completed successfully.")
	return nil
}
