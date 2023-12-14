package entities

import "time"

type Transfers struct {
	Transfer_id     uint
	User_id_sdr     uint
	User_id_rcv     uint
	Transfer_amount uint
	Transfer_time   time.Time
}
