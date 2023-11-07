package recorder

import "time"

type SampleConflict struct {
	Tx_timestamp  time.Time `gorm:"type:char(64);index"`
	CoFL          int
	OccurTime_con time.Time
	Blokheit      int
}

func init() {
	RegisterModel(&SampleConflict{})
}
