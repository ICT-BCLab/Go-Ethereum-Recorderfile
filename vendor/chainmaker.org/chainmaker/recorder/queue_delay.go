package recorder

import "time"

// When record InTime: OutTime = "1890-01-02 23:09:48.000"
// When record OutTime: InTime = "1890-01-02 23:09:48.000"
type QueueDelay struct {
	Tx_timestamp time.Time `gorm:"type:char(64);index"`
	TxID_qude    string    `gorm:"type:char(64);index"`
	InTime       time.Time `gorm:"index"`
	OutTime      time.Time `gorm:"index"`
}

func init() {
	RegisterModel(&QueueDelay{})
}
