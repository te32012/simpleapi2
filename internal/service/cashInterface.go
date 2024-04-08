package service

type CashInterface interface {
	GetShortByFutureIdAndTagId(tag_id int, future_id int) []byte
	Used(id int)
	UpdateCash()
}
