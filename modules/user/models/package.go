package usermodels

type Package struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Thumb       string `json:"thumb" gorm:"thumb"`
	Title       string `json:"title" gorm:"title"`
	Code        string `json:"code" gorm:"code"`
	Duration    int    `json:"duration" gorm:"duration"`
	Description string `json:"description" gorm:"description"`
}

func (Package) TableName() string {
	return "packages"
}

type PackagePrice struct {
	ID          int      `json:"id" gorm:"primaryKey"`
	PackageID   int      `json:"-,omitempty" gorm:"package"`
	Package     *Package `json:"package,omitempty" gorm:"foreignKey:PackageID"`
	Price       float64  `json:"price" gorm:"price"`
	Description string   `json:"description" gorm:"description"`
}

func (PackagePrice) TableName() string {
	return "package_prices"
}
