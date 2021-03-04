package tcmodels

import "encoding/json"

type IJSON interface {
	ToJSON() ([]byte, error)
	FromJSON() error
}

// ProductLine represents the type of information available for a given product line.
// It maps to a record in a database table.
type ProductLine struct {
	ID        uint   `gorm:";primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(25);unique;not null"`
	URLName   string `gorm:"type:varchar(25);unique;not null"`
	SetCount  uint   `gorm:"not null"`
	CardCount uint   `gorm:"not null"`
}

// ToJSON implements the corresponding IJSON interface method for ProductLine model
func (pl *ProductLine) ToJSON() ([]byte, error) {
	return json.Marshal(pl)
}

// FromJSON implements the corresponding IJSON interface method for ProductLine model
func (pl *ProductLine) FromJSON(data []byte) error {
	return json.Unmarshal(data, pl)
}

// SetInfo represents the type of information available for a given card set.
// It maps to a record in a database table.
type SetInfo struct {
	// Model
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"type:varchar(70);unique;not null;"`
	URLName       string `gorm:"type:varchar(75);unique;not null"`
	CardCount     uint   `gorm:"not null"`
	ProductLineID uint
	ProductLine   ProductLine `gorm:"foreignKey:ProductLineID"` // Defines a "Belongs To" relationship with ProductLine model
}

// ToJSON implements the corresponding IJSON interface method for SetInfo model
func (si *SetInfo) ToJSON() ([]byte, error) {
	return json.Marshal(si)
}

// FromJSON implements the corresponding IJSON interface method for SetInfo model
func (si *SetInfo) FromJSON(data []byte) error {
	return json.Unmarshal(data, si)
}

// CardInfo represents the type of information available for a given card.
// It maps to a record in a database table.
type CardInfo struct {
	ID            uint    `gorm:"primaryKey;unique;autoIncrement"`
	Number        string  `gorm:"type:varchar(15);primaryKey"`
	Attack        string  `gorm:"type:varchar(5)"`
	Attribute     string  `gorm:"type:varchar(15)"` // (e.g. Earth, Fire, Water, etc..)
	CardType      string  `gorm:"type:varchar(40)"`
	CardTypeB     string  `gorm:"type:varchar(40);"` // Specific type of monster, spell, or trap (e.g. Effect Monster, Ritual Monster)
	Defense       string  `gorm:"type:varchar(5)"`
	LinkRating    string  `gorm:"type:varchar(1)"`
	LinkArrows    string  `gorm:"type:varchar(70)"` // Contains one or more comma separated string values
	Name          string  `gorm:"type:varchar(100);primaryKey"`
	Level         string  `gorm:"type:varchar(4)"`
	MonsterType   string  `gorm:"type:varchar(50)"`
	Rarity        string  `gorm:"type:varchar(25);primaryKey"`
	MarketPrice   float32 `gorm:"type:float(2) unsigned;not null"`
	Description   string  `gorm:"type:text"`
	SetID         uint    `gorm:"primaryKey"`
	ProductLineID uint
	SetInfo       SetInfo     `gorm:"foreignKey:SetID"` // Defines a "Belongs To" relationship with SetInfo model
	ProductLine   ProductLine `gorm:"foreignKey:ProductLineID"`
}

// ToJSON implements the corresponding IJSON interface method for CardInfo model
func (ci *CardInfo) ToJSON() ([]byte, error) {
	return json.Marshal(ci)
}

// FromJSON implements the corresponding IJSON interface method for CardInfo model
func (ci *CardInfo) FromJSON(data []byte) error {
	return json.Unmarshal(data, ci)
}
