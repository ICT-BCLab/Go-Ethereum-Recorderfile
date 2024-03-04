package recorder

import "time"

type TransmissionLatency struct {
	Tx_timestamp time.Time `gorm:"type:char(64);index"`
	TxID_tranla  string    `gorm:"index"`
	PeerID       string    `gorm:"index"`
	T1           int       `gorm:"index"`
	T2           int
	T3           int
	T4           int
}

func init() {
	RegisterModel(&TransmissionLatency{})
}
