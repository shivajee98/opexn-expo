package main

import (
    "fmt"
    "log"

    "github.com/shivajee98/opexn-exhibitors/internal/config"
    "github.com/shivajee98/opexn-exhibitors/internal/db"
    "github.com/shivajee98/opexn-exhibitors/internal/model"
)

func main() {
    fmt.Println("Starting DB Drop...")

    cfg := config.LoadEnv()
    fmt.Println("Using database URL:", cfg.DatabaseURL)

    dbConn, err := db.Connect(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Drop tables in order accounting for foreign key dependencies to avoid errors
    // Drop child tables first, then parent tables
    err = dbConn.Migrator().DropTable(
        &model.ProductImage{},
        &model.Product{},
        &model.UserType{},
        &model.RevenueInfo{},
        &model.FundingInfo{},
        &model.EventIntent{},
        &model.SPOC{},
        &model.Director{},
        &model.Startup{},
        &model.Address{},
    )
    if err != nil {
        log.Fatalf("Dropping tables failed: %v", err)
    }

    fmt.Println("Tables dropped successfully.")
}
