package service

type CashInterface interface {
	GetShortByFutureIdAndTagId(tag_id int, future_id int) (bool, []byte)
	Used(id int)
	UpdateCash()
}
