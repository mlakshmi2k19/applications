package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	m "github.com/mlakshmi2k19/applications/go-projects/cmd/data-lumen/model"
	util "github.com/mlakshmi2k19/applications/go-projects/internal"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	var err error
	DB, err = util.ConnectDB()
	if err != nil {
		slog.Error("while connecting to DB", "err", err)
	}
	router := gin.Default()
	router.GET("/revenue", getTotalRevenue)
	router.GET("/revenue_by_product", getTotalRevenueByProduct)
	router.GET("/revenue_by_category", getTotalRevenueByCategory)
	router.GET("/revenue_by_region", getTotalRevenueByRegion)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			slog.Error("starting server", "err", err.Error())
		}
	}()

	<-quit

	slog.Info("Shutting down the server.")
	if err := server.Shutdown(context.Background()); err != nil {
		slog.Error("shutting down the server", "err", err.Error())
	}
	slog.Info("server stopped.")
}

func getQueryParams(ctx *gin.Context) (time.Time, time.Time, error) {
	start_date, err := time.Parse("2006-01-02", ctx.Query("start_date"))
	if err != nil {
		return time.Time{}, time.Time{}, err

	}
	end_date, err := time.Parse("2006-01-02", ctx.Query("end_date"))
	if err != nil {
		return time.Time{}, time.Time{}, err

	}
	return start_date, end_date, nil
}

func getTotalRevenue(ctx *gin.Context) {
	start_date, end_date, err := getQueryParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var orders []m.Order
	if err = DB.Preload("Items.Product").
		Where("date_of_sale BETWEEN ? AND ?", start_date, end_date).
		Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	var revenue float64
	for _, order := range orders {
		for _, item := range order.Items {
			revenue += float64(item.QuantitySold) * float64(item.Product.UnitPrice)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Start Date":    start_date,
		"End Date":      end_date,
		"Total Revenue": revenue,
	})
}

func getTotalRevenueByProduct(ctx *gin.Context) {
	start_date, end_date, err := getQueryParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var orders []m.Order
	if err = DB.Preload("Items.Product").
		Where("date_of_sale BETWEEN ? AND ?", start_date, end_date).
		Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	var revenueByProd = make(map[string]float64)
	for _, order := range orders {
		for _, item := range order.Items {
			revenueByProd[item.ProductID] += float64(item.QuantitySold) * float64(item.Product.UnitPrice)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Start Date":         start_date,
		"End Date":           end_date,
		"Revenue By Product": revenueByProd,
	})
}

func getTotalRevenueByCategory(ctx *gin.Context) {
	start_date, end_date, err := getQueryParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var orders []m.Order
	if err = DB.Preload("Items.Product").
		Where("date_of_sale BETWEEN ? AND ?", start_date, end_date).
		Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	var revenueByCategory = make(map[string]float64)
	for _, order := range orders {
		for _, item := range order.Items {
			revenueByCategory[item.Product.Category] += float64(item.QuantitySold) * float64(item.Product.UnitPrice)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Start Date":          start_date,
		"End Date":            end_date,
		"Revenue By Category": revenueByCategory,
	})
}

func getTotalRevenueByRegion(ctx *gin.Context) {
	start_date, end_date, err := getQueryParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var orders []m.Order
	if err = DB.Preload("Items.Product").
		Where("date_of_sale BETWEEN ? AND ?", start_date, end_date).
		Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	var revenueByRegion = make(map[string]float64)
	for _, order := range orders {
		for _, item := range order.Items {
			revenueByRegion[order.RegionName] += float64(item.QuantitySold) * float64(item.Product.UnitPrice)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Start Date":        start_date,
		"End Date":          end_date,
		"Revenue By Region": revenueByRegion,
	})
}
