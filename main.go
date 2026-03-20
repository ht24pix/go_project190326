package main

import (
	"go-project-190326/internal/db"
	"go-project-190326/internal/server"
	"go-project-190326/internal/warehouse"
	"log"
)

func main() {
	// 1. Kết nối DB
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Không thể kết nối DB: %v", err)
	}
	defer database.Close()

	// 2. Khởi tạo Repository cho Warehouse
	warehouseRepo := &warehouse.Repository{DB: database}

	// Thử nghiệm lấy dữ liệu (để kiểm tra logic)
	contracts, _ := warehouseRepo.GetAllContracts()
	log.Printf("Số lượng hợp đồng hiện có: %d", len(contracts))

	// 3. Chạy Server
	server.StartServer()
}
