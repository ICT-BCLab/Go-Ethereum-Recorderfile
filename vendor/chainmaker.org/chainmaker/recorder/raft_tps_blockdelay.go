package recorder

import "time"

type Txtpsblockdelay_start struct {
	Tx_timestamp          time.Time `gorm:"type:char(64);index"`
	Height_delay_sta      uint64
	Total_tps_start       int64 `gorm:"type:int"`
	Txsendtime_blockdelay time.Time
}

type Txtpsblockdelay_end struct {
	Tx_timestamp             time.Time `gorm:"type:char(64);index"`
	Height_delay_end         uint64
	Txs_tps_end              uint32 `gorm:"type:int"`
	Txconfirmtime_blockdelay time.Time
}

func init() {
	RegisterModel(&Txtpsblockdelay_start{})
	RegisterModel(&Txtpsblockdelay_end{})
}
