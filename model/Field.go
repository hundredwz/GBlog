package model

type Field struct {
	Cid        int     `gorm:"PRIMARY_KEY;NOT NULL"`
	Name       string  `gorm:"PRIMARY_KEY;NOT NULL"`
	Type       string  `gorm:"SIZE:8;DEFAULT:'str'"`
	StrValue   string  `gorm:"TYPE:text"`
	IntValue   string  `gorm:"DEFAULT:0"`
	FloatValue float64 `gorm:"DEFAULT:0"`
}

func (f *Field) TableName() string {
	return "gblog_fields"
}
