package tcmodels

type IJSON interface {
	ToJSON() ([]byte, error)
	FromJSON() error
}

// ProductLineInfo
type ProductLine struct {
	ID        uint   `gorm:";primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(25);unique;not null"`
	URLName   string `gorm:"type:varchar(25);unique;not null"`
	SetCount  uint   `gorm:"not null"`
	CardCount uint   `gorm:"not null"`
}

// SetInfo
type SetInfo struct {
	// Model
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"type:varchar(70);unique;not null;"`
	URLName       string `gorm:"type:varchar(75);unique;not null"`
	CardCount     uint   `gorm:"not null"`
	ProductLineID uint
	ProductLine   ProductLine `gorm:"foreignKey:ProductLineID"` // Defines a "Belongs To" relationship with ProductLine model
}

// CardInfo
type CardInfo struct {
	CardID        uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"type:varchar(100);not null"`
	URLName       string `gorm:"type:varchar(110);not null"`
	AttributesID  uint
	SetID         uint
	ProductLineID uint
	SetInfo       SetInfo     `gorm:"foreignKey:SetID"`         // Defines a "Belongs To" relationship with SetInfo model
	ProductLine   ProductLine `gorm:"foreignKey:ProductLineID"` // Defines a "Belongs To" relationship with ProductLine model
}

// YugiohAttributes
type YugiohAttributes struct {
	AttributesID uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"type:varchar(100)"`
	Number       string `gorm:"type:varchar(15)"`
	Attack       string `gorm:"type:varchar(5)"`
	Attribute    string `gorm:"type:varchar(15)"` // (e.g. Earth, Fire, Water, etc..)
	CardType     string `gorm:"type:varchar(40)"`
	CardTypeB    string `gorm:"type:varchar(40)"` // monster, spell, or trap (e.g. Effect Monster, Ritual Monster)
	Defense      string `gorm:"type:varchar(5)"`
	LinkRating   string `gorm:"type:varchar(1)"`
	LinkArrows   string `gorm:"type:varchar(70)"` // Contains one or more comma separated string values
	Level        string `gorm:"type:varchar(4)"`
	MonsterType  string `gorm:"type:varchar(50)"`
	Rarity       string `gorm:"type:varchar(25);"`
	Description  string `gorm:"type:text"`
	CardID       uint
	CardInfo     CardInfo `gorm:"foreignKey:CardID"` // Defines a "Belongs To" relationship with CardInfo model
}
