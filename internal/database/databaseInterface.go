package database

import "avitotestgo2024/internal/entitys"

type DatabaseInterface interface {
	GetListBannersByListId(id []int) (baners entitys.Banners, err error)
	GetBannerByTagIdAndFutureId(tag_id int, future_id int) (baner *entitys.Banner, err error)
	GetListBannerByTagAndFutureIdWithOffsetAndLimit(tag_id, future_id, offset, limit int) (baners entitys.Banners, err error)
	UpdateBannerById(bannerWithId *entitys.Banner) error
	DeleteBannerById(id int) error
	DeleteBannerByFutureId(future_id int)
	CreateBanner(bannerWithoutId *entitys.Banner) (banner_id int, err error)
	GetThreeVersionBannerById(id int) (baners entitys.Banners, err error)
}
