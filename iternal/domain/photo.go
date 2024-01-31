package domain

type Photos struct {
	Base
	Name      string `gorm:"type:text"`
	Path      string `gorm:"type:text"`
	Size      int64  `gorm:"type:int"`
	PathToMin string `gorm:"type:text"`
}
