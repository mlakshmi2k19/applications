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

func getTotalRevenue(ctx *gin.Context) {
	start_date, err := time.Parse("2006-01-02", ctx.Query("start_date"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format"})
		return

	}
	end_date, err := time.Parse("2006-01-02", ctx.Query("end_date"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format"})
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
	start_date, err := time.Parse("2006-01-02", ctx.Query("start_date"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format"})
		return

	}
	end_date, err := time.Parse("2006-01-02", ctx.Query("end_date"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format"})
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
