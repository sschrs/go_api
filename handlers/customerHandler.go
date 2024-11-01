package handlers

import (
	"data_app/database"
	"data_app/models"
	"data_app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type CustomerHandler struct{}

func (h CustomerHandler) Get(c *fiber.Ctx) error {
	customerID := c.Query("customer_id")
	dateStart := c.Query("date_start")
	dateEnd := c.Query("date_end")

	var customers []models.Customer

	if err := database.DB.Where("customer_id = ? AND date > ? AND date < ?", customerID, dateStart, dateEnd).Find(&customers).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "record not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Internal Server Error: %v", err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"length": len(customers), "data": customers})
}

// func (h CustomerHandler) GetTransactionSum(c *fiber.Ctx) error {
// 	customerID := c.Query("customer_id")
// 	dateStart := c.Query("date_start")
// 	dateEnd := c.Query("date_end")

// 	var total float32

// 	if err := database.DB.Model(models.Customer{}).Select("SUM(transaction_sum)").Where("customer_id = ? AND date > ? AND date < ?", customerID, dateStart, dateEnd).Scan(&total).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "record not found"})
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Internal Server Error: %v", err.Error())})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"total_transaction": total})
// }

func (h CustomerHandler) GetTransactionSum(c *fiber.Ctx) error {
	customerID := c.Query("customer_id")
	dateStart := c.Query("date_start")
	dateEnd := c.Query("date_end")

	cacheKey := fmt.Sprintf("%s:%s:%s", customerID, dateStart, dateEnd)

	cachedTotal, err := database.RDB.Get(database.RCtx, cacheKey).Result()
	if err == nil {
		var total float32
		if err := json.Unmarshal([]byte(cachedTotal), &total); err == nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"total_transaction": total})
		}
	}

	var total float32
	if err := database.DB.Model(models.Customer{}).
		Select("SUM(transaction_sum)").
		Where("customer_id = ? AND date > ? AND date < ?", customerID, dateStart, dateEnd).
		Scan(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "record not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Internal Server Error: %v", err.Error())})
	}

	totalJson, err := json.Marshal(total)
	if err == nil {
		database.RDB.Set(database.RCtx, cacheKey, totalJson, time.Hour)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"total_transaction": total})
}

func (h CustomerHandler) Create(c *fiber.Ctx) error {
	var customer models.Customer

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	if isValidated, validationError := utils.Validate(customer); !isValidated {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": validationError})
	}

	if err := database.DB.Create(&customer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Internal Server Error: %v", err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "record inserted successfully", "data": customer})
}

var segments = []string{"Gold", "Silver", "Bronze", "Platinum"}

func (h CustomerHandler) CreateRandom(c *fiber.Ctx) error {
	countParam := c.Params("count")

	count, err := strconv.Atoi(countParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "count parameter must be an integer"})
	}

	rand.Seed(uint64(time.Now().Unix()))

	batchSize := 1000
	customers := make([]models.Customer, 0, batchSize)

	for i := 0; i < count; i++ {
		customer := models.Customer{
			CustomerID:     fmt.Sprintf("CUST%05d", rand.Intn(100000)),
			Date:           utils.GenerateRandomDate(),
			Segment:        segments[rand.Intn(len(segments))],
			TransactionSum: rand.Float32() * 10000,
			CreatedDate:    time.Now(),
		}
		customers = append(customers, customer)

		if len(customers) == batchSize || i == count-1 {
			if result := database.DB.Create(&customers); result.Error != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
			}
			log.Printf("%d record inserted into database", i)
			customers = customers[:0]
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("%d records inserted into database", count)})
}
