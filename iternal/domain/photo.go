package domain

type Photos struct {
	Base
	Name      string `gorm:"column:имя" gorm:"type:text"`
	Path      string `gorm:"column:путь_до_файла" gorm:"type:text"`
	Size      int64  `gorm:"column:размер_файла" gorm:"type:int"`
	PathToMin string `gorm:"column:короткий_путь_до_файла" gorm:"type:text"`
}
