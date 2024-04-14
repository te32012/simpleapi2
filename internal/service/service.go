package service

import (
	"avitotestgo2024/internal/auth"
	"avitotestgo2024/internal/database"
	"avitotestgo2024/internal/entitys"
	"errors"
	"log/slog"
	"strconv"

	"github.com/mailru/easyjson"
)

type Service struct {
	Cash CashInterface
	// Доступ к основым данным в PostgreSQL
	MainBase database.DatabaseInterface
	Logger   *slog.Logger
	Auth     auth.AuthInterface
}

func NewService(base database.DatabaseInterface, logger *slog.Logger, auth auth.AuthInterface) *Service {
	return &Service{MainBase: base, Logger: logger, Auth: auth, Cash: NewChashUsedMap(base)}
}

func (s *Service) CovertErrorToBytes(err entitys.Error) []byte {
	data, _ := easyjson.Marshal(&err)
	return data
}

func (s *Service) GetUserBanner(tag_id, feature_id int, use_last_revission bool, token string) (ans []byte, err error) {
	s.Logger.Info("Получение банеров по tag_id = " + strconv.Itoa(tag_id) + " feature_id = " + strconv.Itoa(feature_id) + " use_last_revission = " + strconv.FormatBool(use_last_revission) + " token = " + token)
	if !use_last_revission {
		ok, ans := s.Cash.GetShortByFutureIdAndTagId(tag_id, feature_id)
		if ok && len(ans) > 0 {
			s.Logger.Info("Банер получен из кеша!!!")
			return ans, nil
		}
	}
	banner, errr := s.MainBase.GetBannerByTagIdAndFutureId(tag_id, feature_id)
	if errr != nil {
		err = errr
		ans = make([]byte, 0)
		return ans, err
	}
	if banner != nil {
		s.Cash.Used(banner.Id)
		if banner.Is_active {
			ans, err = easyjson.Marshal(banner.Content)
		} else if ok, _ := s.Auth.HasPermission(token, 2); ok {
			ans, err = easyjson.Marshal(banner.Content)
		} else {
			err = errors.New("404")
			ans = make([]byte, 0)
			return
		}
	} else {
		err = errors.New("404")
		ans = make([]byte, 0)
		return
	}
	s.Logger.Info("статус получения банера ok")
	return
}

func (s *Service) GetAllBanners(tag_id, feature_id, limit, offset int) (ans []byte, err error) {
	s.Logger.Info("Получение банеров по tag_id = " + strconv.Itoa(tag_id) + " feature_id = " + strconv.Itoa(feature_id) + " limit = " + strconv.Itoa(limit) + " offset = " + strconv.Itoa(offset))
	banners, err := s.MainBase.GetListBannerByTagAndFutureIdWithOffsetAndLimit(tag_id, feature_id, offset, limit)
	if err != nil {
		return nil, err
	}
	if len(banners) > 0 {
		ans, err = easyjson.Marshal(banners)
	} else {
		err = errors.New("404")
	}
	s.Logger.Info("Статус получения банеров ok")
	return
}

func (s *Service) CreateBanner(banner []byte) (ans []byte, err error) {
	var bannerJSON *entitys.Banner = new(entitys.Banner)
	err = easyjson.Unmarshal(banner, bannerJSON)
	if err != nil {
		s.Logger.Error("Ошибка конвертации в структуру банера при создании " + err.Error())
		return
	}
	s.Logger.Info("Создаем баннер ")
	s.Logger.Info("баннер", slog.Any("баннер", bannerJSON))
	id, err := s.MainBase.CreateBanner(bannerJSON)
	if err != nil {
		s.Logger.Error("Ошибка работы с базой при создании банера " + err.Error())
		return
	}
	ans, err = easyjson.Marshal(&entitys.Ans201{Id: id})
	s.Logger.Info("Был создан баннер с id = " + strconv.Itoa(id))
	s.Cash.Used(id)
	return
}

func (s *Service) UpdateBanner(id int, banner []byte) (err error) {
	s.Logger.Info("Обновляем банер с id = " + strconv.Itoa(id))
	var bannerJSON *entitys.Banner = new(entitys.Banner)
	s.Cash.Used(id)
	err = easyjson.Unmarshal(banner, bannerJSON)
	if err != nil {
		s.Logger.Error("Ошибка конвертации в структуру банера с id = " + strconv.Itoa(id))
		return
	}
	bannerJSON.Id = id
	err = s.MainBase.UpdateBannerById(bannerJSON)
	s.Logger.Info("Обновили банер с id = " + strconv.Itoa(id))
	return
}

func (s *Service) Delete(id int) (err error) {
	s.Logger.Info("удаляем банер с id = " + strconv.Itoa(id))
	err = s.MainBase.DeleteBannerById(id)
	s.Logger.Info("удаление банера с id = " + strconv.Itoa(id))
	return
}

func (s *Service) GetUserBannerThreeVersion(id int, token string) (ans []byte, err error) {
	s.Logger.Info("Получение 3х версий банеров по id = " + strconv.Itoa(id) + " token = " + token)
	banners, err := s.MainBase.GetThreeVersionBannerById(id)
	s.Cash.Used(id)
	if err != nil {
		return
	}
	if ok, _ := s.Auth.HasPermission(token, 2); ok {
		data, _ := easyjson.Marshal(banners)
		ans = data
		return
	}
	var tmp entitys.Banners
	for i := 0; i < len(banners); i++ {
		if banners[i].Is_active {
			tmp = append(tmp, banners[i])
		}
	}
	if len(tmp) == 0 {
		err = errors.New("404")
	}
	data, _ := easyjson.Marshal(tmp)
	ans = data
	return
}
func (s *Service) DeleteByFuture(future_id int) {
	s.MainBase.(*database.Database).Thread <- future_id
}
