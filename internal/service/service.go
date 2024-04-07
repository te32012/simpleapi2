package service

import (
	"avitotestgo2024/internal/database"
	"avitotestgo2024/internal/entitys"
	"encoding/json"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	CashOne  *redis.Client
	CashTwo  *redis.Client
	MainBase database.DatabaseInterface
	Logger   *slog.Logger
}

func (s *Service) CovertErrorToBytes(err entitys.Error) []byte {
	json.Marshal()
}
func (s *Service) GetUserBanner(tag_id, feature_id int, use_last_revission bool) (ans []byte, err error) {

}
func (s *Service) GetAllBanners(tag_id, feature_id, limit, offset int) (ans []byte, err error) {

}
func (s *Service) CreateBanner(banner []byte) (ans []byte, err error) {

}
func (s *Service) UpdateBanner(id int, banner []byte) (err error) {

}
func (s *Service) Delete(id int) (err error) {

}
