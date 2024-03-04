package recorder

import "time"

type TransactionThroughput struct {
	Tx_timestamp      time.Time `gorm:"type:char(64);index"`
	TxID_tranthr      string
	OccurTime_tranthr time.Time `gorm:"index"`
	Source_tranthr    int
}

func init() {
	RegisterModel(&TransactionThroughput{})
}
