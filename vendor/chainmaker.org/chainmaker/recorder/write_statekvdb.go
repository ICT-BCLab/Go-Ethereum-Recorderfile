package recorder

import "time"

type WriteStateDB struct {
	Height       string    `gorm:"type:char(64)"` //Tx_t
	Tx_timestamp time.Time `gorm:"index"`
	WriteSpeed   float64
	WriteTime    float64
	WriteLen     int
}

func init() {
	RegisterModel(&WriteStateDB{})
}
