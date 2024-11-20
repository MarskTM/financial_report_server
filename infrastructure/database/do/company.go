package do

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	Owner        string `json:"owner"`

	SIC string `json:"sic"`
	ICB string `json:"icb"`

	Category            string    `json:"category"`
	Major               string    `json:"major"`
	MajorCode           string    `json:"major_code"`
	TaxCode             string    `json:"tax_code"`
	DateOfEstablishment time.Time `json:"date_of_establishment"`
	ChapterCapital      int64     `json:"chapter_capital"`
	Employees           int32     `json:"employees"`
	Branches            int32     `json:"branches"`

	Address  string `json:"address"`
	Activity string `json:"activity"`

	ListingDate   time.Time `json:"listing_date"`
	ListingFloor  string    `json:"listing_floor"`
	IPOPrice      int64     `json:"ipo_price"`
	ListendVolume int64     `json:"listend_volume"`
	MarketCap     int64     `json:"market_cap"`
	SLCP          int64     `json:"slcp"`

	CompanyManagements []CompanyManagements `json:"company_managers"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type SubCompanies struct {
	ID           int    `json:"id"`
	CompanyID    int32  `json:"company_id"`
	SubCompanyID int32  `json:"sub_company_id"`
	Type         string `json:"type"`
	Ratio        int    `json:"ratio"`

	SubCompanies []SubCompanies `json:"sub_companies"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// "Thông tin ban lãnh đạo"
type CompanyManagements struct {
	ID        int64 `json:"id"`
	CompanyID int64 `json:"company_id"`

	Name string `json:"name"`
	Role string `json:"role"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
