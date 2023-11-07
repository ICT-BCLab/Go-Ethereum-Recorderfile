package recorder

import "time"

type Txconfirmdelay_start struct {
	Tx_timestamp          time.Time `gorm:"type:char(64);index"`
	Txid_confirmdelay_sta string
	Txsendtime            time.Time
}

type Txconfirmdelay_end struct {
	Tx_timestamp          time.Time `gorm:"type:char(64);index"`
	Txid_confirmdelay_end string
	Txconfirmtime         time.Time
}

func init() {
	RegisterModel(&Txconfirmdelay_start{})
	RegisterModel(&Txconfirmdelay_end{})
}
