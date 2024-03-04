package recorder

import "time"

type ContractTime struct {
	TxHash       string
	ContractAddr string
	StartTime    time.Time     `gorm:"type:char(64);index"`
	EndTime      time.Time     `gorm:"type:char(64);index"`
	ExecTime     time.Duration `gorm:"type:char(64);index"`
}

func init() {
	RegisterModel(&ContractTime{})
}
