package warehouse

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

// GetAllContracts lấy toàn bộ danh sách hợp đồng từ DB
func (r *Repository) GetAllContracts() ([]Contract, error) {
	query := "SELECT id, contract_number, customer_id, warehouse_id, area_m2, status FROM contract"
	
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []Contract
	for rows.Next() {
		var c Contract
		// Mapping từng cột trong SQL vào từng trường trong Struct
		err := rows.Scan(&c.ID, &c.ContractNumber, &c.CustomerID, &c.WarehouseID, &c.AreaM2, &c.Status)
		if err != nil {
			return nil, err
		}
		contracts = append(contracts, c)
	}

	return contracts, nil
}