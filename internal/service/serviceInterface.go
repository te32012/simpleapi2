package service

import "avitotestgo2024/internal/entitys"

type ServiceInterface interface {
	CovertErrorToBytes(err entitys.Error) []byte
	GetUserBanner(tag_id, feature_id int, use_last_revission bool) (ans []byte, err error)
	GetAllBanners(tag_id, feature_id, limit, offset int) (ans []byte, err error)
	CreateBanner(banner []byte) (ans []byte, err error)
	UpdateBanner(id int, banner []byte) (err error)
	Delete(id int) (err error)
}
