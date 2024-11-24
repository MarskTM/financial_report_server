package do

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	Owner        string `json:"owner"` // Người đại diện

	SIC string `json:"sic"` // Mã SIC
	ICB string `json:"icb"` // Mã ngành ICB

	Category            string    `json:"category"`              // Loại hình doanh nghiệp
	Major               string    `json:"major"`                 // Tên ngành hoạt động
	MajorCode           string    `json:"major_code"`            // Mã ngành
	TaxCode             string    `json:"tax_code"`              // Mã số thuế
	DateOfEstablishment time.Time `json:"date_of_establishment"` // ngày thành lập
	ChapterCapital      int64     `json:"chapter_capital"`       // VĐL
	Employees           int32     `json:"employees"`             // Số lượng nhân viên
	Branches            int32     `json:"branches"`              // Số lượng chi nhánh

	Address  string `json:"address"`  // Địa chỉ
	Activity string `json:"activity"` // Tình trạng hoạt động

	ListingDate   time.Time `json:"listing_date"`   // Ngày niêm yết
	ListingFloor  string    `json:"listing_floor"`  // Nơi được niêm yết
	IPOPrice      int64     `json:"ipo_price"`      // Giá chào sàn
	ListendVolume int64     `json:"listend_volume"` // Khối lượng niêm yết
	MarketCap     int64     `json:"market_cap"`     // Thị giá vốn
	SLCP          int64     `json:"slcp"`           // SLCP lưu hành

	// CompanyManagements []CompanyManagement `json:"company_managers"`
	// CompanyTidings     []CompanyTiding     `json:"company_tidings"`
	// Shareholders       []Shareholder       `json:"shareholders"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// ------------------------------------------------------------------------------------------------
// Danh sách công ti con & công ty liên kết
type SubCompany struct {
	ID           int    `json:"id"`
	CompanyID    int32  `json:"company_id"`
	SubCompanyID int32  `json:"sub_company_id"`
	Type         string `json:"type"`  // Cty con | Cty liên kết
	Ratio        int    `json:"ratio"` // TL nắm giữ

	// SubCompanies []Company `json:"sub_companies"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// ------------------------------------------------------------------------------------------------
// "Thông tin ban lãnh đạo"
type CompanyManagement struct {
	ID        int64 `json:"id"`
	CompanyID int64 `json:"company_id"`

	Name string `json:"name"`
	Role string `json:"role"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// ------------------------------------------------------------------------------------------------
// Thông tin các bài báo liên quan đến doanh nghiệp
type CompanyTiding struct {
	ID        int64 `json:"id"`
	CompanyID int64 `json:"company_id"`
	TidingID  int64 `json:"tiding_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// ------------------------------------------------------------------------------------------------
// Thông tin cổ đông
type Shareholder struct {
	ID        int32  `json:"id"`
	CompanyID int32  `json:"company_id"`
	Owner     string `json:"owner"` // Tên cổ đông
	Type      string `json:"type"`  // Cá nhân/Tổ chức

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
