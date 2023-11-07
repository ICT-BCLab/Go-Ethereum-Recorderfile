package recorder

import "time"

type Blockvalidefficy_total struct {
	Tx_timestamp         time.Time `gorm:"type:char(64);index"`
	Block_valideff_total uint64
	Total                int64
}

type Blockvalidefficy_txs struct {
	Tx_timestamp             time.Time `gorm:"type:char(64);index"`
	Blockheight_valideff_txs uint64
	Txs                      uint32
}

func init() {
	RegisterModel(&Blockvalidefficy_total{})
	RegisterModel(&Blockvalidefficy_txs{})
}
