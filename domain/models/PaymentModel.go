package models

import "time"

type PaymentModel struct {
	ID          uint64     `json:"id" gorm:"primarykey"`
	Ok          bool       `json:"ok" gorm:"ok"`
	Msg         string     `json:"msg" gorm:"msg"`
	TraceCode   string     `json:"trace_code" gorm:"trace_code"`
	Refnum      string     `json:"refnum" gorm:"refnum"`
	Paytype     uint       `json:"paytype" gorm:"paytype"`
	Info        string     `json:"info" gorm:"info"`
	CheckNum    string     `json:"check_num" gorm:"check_num"`
	CheckBank   string     `json:"check_bank" gorm:"check_bank"`
	CheckStatus uint       `json:"check_status" gorm:"check_status"`
	Amount      float64    `json:"amount" gorm:"amount"`
	Discount    float64    `json:"discount" gorm:"discount"`
	Income      uint       `json:"income" gorm:"income"`
	PaidFor     uint       `json:"paid_for" gorm:"paid_for"`
	PaidTo      string     `json:"paid_to" gorm:"paid_to"`
	UserID      uint64     `json:"user_id" gorm:"user_id"`
	User        *UserModel `json:"user" gorm:"foreignkey:UserID"`
	StaffID     uint64     `json:"staff_id" gorm:"staff_id"`
	Staff       *UserModel `json:"staff" gorm:"foreignkey:StaffID"`
	Created     *time.Time `json:"created" gorm:"created"`
	CheckDate   *time.Time `json:"check_date" gorm:"check_date"`
}

func (PaymentModel) TableName() string {
	return "payment"
}
