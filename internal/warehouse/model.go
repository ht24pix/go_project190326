package warehouse

import (
	"math/big"
	"time"
)

// CompanyInfo ánh xạ với bảng 'companyinfo'
type CompanyInfo struct {
	ID             int64  `json:"id"`
	CompanyName    string `json:"company_name"`
	TaxCode        string `json:"tax_code"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Representative string `json:"representative"`
	Position       string `json:"position"`
}

// Customer ánh xạ với bảng 'customer'
type Customer struct {
	ID             int64  `json:"id"`
	CompanyID      int64  `json:"company_id"`
	CustomerName   string `json:"customer_name"`
	TaxCode        string `json:"tax_code"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Representative string `json:"representative"`
}

// Warehouse ánh xạ với bảng 'warehouse'
type Warehouse struct {
	ID            int64   `json:"id"`
	CompanyID     int64   `json:"company_id"`
	WarehouseName string  `json:"warehouse_name"`
	Location      string  `json:"location"`
	AreaM2        float64 `json:"area_m2"`
	Note          string  `json:"note"`
}

// Contract ánh xạ với bảng 'contract'
type Contract struct {
	ID             int64     `json:"id"`
	CustomerID     int64     `json:"customer_id"`
	WarehouseID    int64     `json:"warehouse_id"`
	ContractNumber string    `json:"contract_number"`
	SignDate       time.Time `json:"sign_date"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	AreaM2         float64   `json:"area_m2"`
	// Sử dụng big.Rat cho các giá trị tiền tệ để đảm bảo độ chính xác
	UnitPrice *big.Rat `json:"unit_price"`
	Deposit   *big.Rat `json:"deposit"`
	Status string `json:"status"`
	Note   string `json:"note"`
}

// UtilityService ánh xạ với bảng 'utility_service'
type UtilityService struct {
	ID          int64    `json:"id"`
	ContractID  int64    `json:"contract_id"`
	ServiceName string   `json:"service_name"`
	UnitPrice   *big.Rat `json:"unit_price"`
	Unit        string   `json:"unit"`
}

// UtilityUsage ánh xạ với bảng 'utility_usage'
type UtilityUsage struct {
	ID          int64     `json:"id"`
	ContractID  int64     `json:"contract_id"`
	ServiceID   int64     `json:"service_id"`
	RecordDate  time.Time `json:"record_date"`
	OldIndex    float64   `json:"old_index"`
	NewIndex    float64   `json:"new_index"`
	UsageAmount float64   `json:"usage_amount"`
	TotalAmount *big.Rat  `json:"total_amount"`
}

// Payment ánh xạ với bảng 'payment'
type Payment struct {
	ID            int64     `json:"id"`
	ContractID    int64     `json:"contract_id"`
	PaymentDate   time.Time `json:"payment_date"`
	Amount        *big.Rat  `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	Content       string    `json:"content"`
}
