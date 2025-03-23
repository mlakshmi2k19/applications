package main

import (
	"errors"
	"log/slog"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	m "github.com/mlakshmi2k19/applications/go-projects/cmd/data-lumen/model"
	util "github.com/mlakshmi2k19/applications/go-projects/internal"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var tables = []interface{}{&m.Customer{}, &m.Product{}, &m.Region{}, &m.Item{}, &m.Order{}}

func doAction(ctx *cli.Context) error {
	file := ctx.String("file-path")
	if filepath.Ext(file) != ".csv" {
		return errors.New("Supported file formats: csv")
	}
	data, err := util.ReadFile(file)
	if err != nil {
		return err
	}
	db, err := util.ConnectDB()
	if err != nil {
		return err
	}
	if ctx.Bool("create-tables") {
		if err = createTablesIfNotExists(db, tables); err != nil {
			return err
		}
		slog.Info("Tables created.", "names",
			strings.Join([]string{
				"customers",
				"products",
				"regions",
				"items",
				"orders",
			}, ","))
	} else {
		slog.Info("Tables already exists.")
	}
	return convertAndUpsert(db, data)
}

func convertAndUpsert(db *gorm.DB, data [][]string) error {
	for ind, row := range data[1:] {
		dateOfSale, _ := time.Parse("2006-01-02", row[6])
		quantitySold, _ := strconv.Atoi(row[7])
		unitPrice, _ := strconv.ParseFloat(row[8], 64)
		discount, _ := strconv.ParseFloat(row[9], 64)
		shippingCost, _ := strconv.ParseFloat(row[10], 64)

		d := m.Data{
			Cust: m.Customer{
				ID:      row[2],
				Name:    row[12],
				Email:   row[13],
				Address: row[14],
			},
			Prod: m.Product{
				ID:        row[1],
				Name:      row[3],
				Category:  row[4],
				UnitPrice: unitPrice,
			},
			Ord: m.Order{
				ID:            row[0],
				CustomerID:    row[2],
				DateOfSale:    dateOfSale,
				RegionName:    row[5],
				PaymentMethod: row[11],
				ShippingCost:  shippingCost,
			},
			Itm: m.Item{
				OrderID:      row[0],
				ProductID:    row[1],
				QuantitySold: quantitySold,
				Discount:     discount,
			},
			Reg: m.Region{Name: row[5]},
		}
		if err := upsert(db, d); err != nil {
			return err
		}
		slog.Info("Record inserted/updated.", "no", ind)
	}
	return nil
}

func upsert(db *gorm.DB, d m.Data) error {
	return db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "email", "address"}),
		}).Create(&d.Cust).Error; err != nil {
			return err
		}

		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "category", "unit_price"}),
		}).Create(&d.Prod).Error; err != nil {
			return err
		}

		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoNothing: true,
		}).Create(&d.Reg).Error; err != nil {
			return err
		}

		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}, {Name: "customer_id"}},
			DoUpdates: clause.AssignmentColumns(
				[]string{
					"region_name",
					"date_of_sale",
					"payment_method",
					"shipping_cost"}),
		}).Create(&d.Ord).Error; err != nil {
			return err
		}

		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "order_id"}, {Name: "product_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"quantity_sold", "discount"}),
		}).Create(&d.Itm).Error; err != nil {
			return err
		}

		return nil
	})
}

func createTablesIfNotExists(db *gorm.DB, tables []interface{}) error {
	for _, t := range tables {
		if err := db.AutoMigrate(t); err != nil {
			return err
		}
	}
	return db.Exec(`ALTER TABLE orders ADD CONSTRAINT unique_order_customer UNIQUE (id, customer_id)`).Error
}
