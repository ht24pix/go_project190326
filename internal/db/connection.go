package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

// Connect thiết lập kết nối tới MySQL
func ConnectDB() (*sql.DB, error) {
	// Định dạng: user:password@tcp(127.0.0.1:3306)/dbname?parseTime=true
	// parseTime=true cực kỳ quan trọng để Go hiểu kiểu DATE/DATETIME của MySQL
	dsn := "root:tri1912@tcp(127.0.0.1:3306)/warehouse_lease?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Cấu hình Connection Pool (Tư duy Architect)
	db.SetMaxOpenConns(25)           // Tối đa 25 kết nối mở cùng lúc
	db.SetMaxIdleConns(25)           // Giữ 25 kết nối chờ sẵn
	db.SetConnMaxLifetime(time.Hour) // Reset kết nối sau mỗi 1 tiếng

	// Kiểm tra xem database có thực sự phản hồi không
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Đã kết nối Database thành công!")
	return db, nil
}
