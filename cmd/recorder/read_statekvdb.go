package recorder

import "time"

type ReadStateDB struct {
	Tx_timestamp time.Time `gorm:"type:char(64);index"`
	ReadSpeed    float64
	ReadTime     float64
	ReadLen      int
}

func init() {
	RegisterModel(&ReadStateDB{})
}
