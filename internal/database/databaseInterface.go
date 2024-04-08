package database

import "avitotestgo2024/internal/entitys"

type DatabaseInterface interface {
	GetListBannersByListId(id []int) entitys.Banners
	GetBannerByTagIdAndFutureId(tag_id int, future_id int) *entitys.Banner
	GetListBannerByTagAndFutureIdWithOffsetAndLimit(tag_id, future_id, offset, limit int) entitys.Banners
	UpdateBannerById(bannerWithId *entitys.Banner) error
	DeleteBannerById(id int) error
	DeleteBannerByFutureId(future_id int) error
	CreateBanner(bannerWithoutId *entitys.Banner) (banner_id int, err error)
}
