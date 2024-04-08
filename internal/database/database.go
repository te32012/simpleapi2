package database

import (
	"avitotestgo2024/internal/entitys"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)


type Database struct {
	Pool *pgxpool.Pool
	Logger *slog.Logger
}
func (db *Database) GetListBannersByListId(id []int) entitys.Banners {
	q :=  "select id_banner, id_tag from tags_banner where "
}
func (db *Database) GetBannerByTagIdAndFutureId(tag_id int, future_id int) *entitys.Banner {
	var ans = new(entitys.Banner)
	q := `with first_query as (select id_banner from tags_banner where id_tag = $1),
	 select id_banner from features_banner where id_future = $2 and first_query.id_banner = features_banner.id_banner;
	`
	db.Pool.

	return nil
}	
func (db *Database) GetListBannerByTagAndFutureIdWithOffsetAndLimit(tag_id, future_id, offset, limit int) entitys.Banners {

}
func (db *Database) UpdateBannerById(bannerWithId *entitys.Banner) error {

}
func (db *Database) DeleteBannerById(id int) error {

}
func (db *Database) DeleteBannerByFutureId(future_id int) error {

}
func (db *Database) CreateBanner(bannerWithoutId *entitys.Banner) (banner_id int, err error) {

}	
