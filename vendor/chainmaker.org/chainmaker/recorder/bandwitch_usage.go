package recorder

import "time"

type BandwitchUsage struct {
	Tx_timestamp      time.Time `gorm:"type:char(64);index"`
	MessageArriveTime time.Time `gorm:"index"`
	MessageType       string
	MessageSize       int
}

func init() {
	RegisterModel(&BandwitchUsage{})
}
