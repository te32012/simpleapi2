package database

import (
	"avitotestgo2024/internal/entitys"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mailru/easyjson"
)

// поправить функцию получения по тегу и айди
// добавить логи к up del ск
// реформатировать код

// Pool is an interface representing a pgx connection pool.

type Database struct {
	Pool   *pgxpool.Pool
	Logger *slog.Logger
	Thread chan int
}

func NewDatabase(pool *pgxpool.Pool, logger *slog.Logger) *Database {
	return &Database{Pool: pool, Logger: logger, Thread: make(chan int, 10000)}
}

func (db *Database) GetListBannersByListId(id []int) (baners entitys.Banners, err error) {
	tx, err := db.Pool.Begin(context.Background())
	if err != nil {
		db.Logger.Error("не смогли подключиться к базе " + err.Error())
		return
	}
	var count int
	defer func() {
		_ = tx.Rollback(context.Background())
	}()
	for i := 0; i < len(id); i++ {
		db.Logger.Info("пытаемся достать из базы  " + strconv.Itoa(id[i]) + " банер")
		var banner entitys.Banner
		banner.Id = id[i]
		q := `with first_query as (select id_banner, content, updated_at, is_active from version_banner where id_banner = $1 order by updated_at desc limit 1),
		second_query as (select id_banner, created_at from banner where id_banner = $1), 
		third_query as (select id_banner, id_future from features_banner where id_banner = $1)
		select distinct first_query.content, first_query.updated_at, first_query.is_active, second_query.created_at, third_query.id_future from third_query join first_query on third_query.id_banner=first_query.id_banner join second_query on third_query.id_banner=second_query.id_banner;;
		`
		rows, err := tx.Query(context.Background(), q, id[i])
		if err != nil {
			count += 1
			db.Logger.Error("ошибка получения баннера " + err.Error())
			continue
		}
		if rows.Next() {
			var tmp []byte
			_ = rows.Scan(&tmp, &banner.Updatet_at, &banner.Is_active, &banner.Created_at, &banner.Feature_ids)
			_ = easyjson.Unmarshal(tmp, &banner.Content)
		} else {
			count += 1
			rows.Close()
			db.Logger.Error("ошибка получения баннера ")
			continue
		}
		rows.Close()
		q = "select id_tag from tags_banner where id_banner = $1;"
		rows, err = tx.Query(context.Background(), q, id[i])
		if err != nil {
			count += 1
			db.Logger.Error("ошибка получения баннера " + err.Error())
			continue
		}
		var id_tags []int
		for rows.Next() {
			var tmp int
			_ = rows.Scan(&tmp)
			id_tags = append(id_tags, tmp)
		}
		banner.Tag_ids = id_tags
		baners = append(baners, banner)
		db.Logger.Info("добавили в ответ баннер " + strconv.Itoa(id[i]))
		rows.Close()
	}
	if count != 0 {
		err = errors.New("в базе не смогли найти " + strconv.Itoa(count) + " баннеров")
	}
	db.Logger.Info("достали из базы  " + strconv.Itoa(len(id)-count) + " банеров")
	db.Logger.Info(fmt.Sprintf("%+v", baners))
	return
}

func (db *Database) GetBannerByTagIdAndFutureId(tag_id int, future_id int) (baner *entitys.Banner, err error) {
	db.Logger.Info("ищем банер по тегу и фиче в базе tag = " + strconv.Itoa(tag_id) + " future = " + strconv.Itoa(future_id))
	baner = new(entitys.Banner)
	baner.Feature_ids = future_id
	q := `
	select features_banner.id_banner from features_banner join (select tags_banner.id_banner from tags_banner where id_tag = $1) as first_query on first_query.id_banner = features_banner.id_banner where features_banner.id_future = $2;;
	`
	tx, err := db.Pool.Begin(context.Background())
	defer func() {
		_ = tx.Rollback(context.Background())
	}()
	if err != nil {
		db.Logger.Error("ошибка получения баннера " + err.Error())
		return
	}
	rows, err := tx.Query(context.Background(), q, tag_id, future_id)
	if err != nil {
		db.Logger.Error("ошибка получения баннера " + err.Error())
		return
	}
	if rows.Next() {
		_ = rows.Scan(&baner.Id)
	} else {
		err = errors.New("404")
		db.Logger.Error("ошибка получения баннера " + err.Error())
		return
	}
	db.Logger.Info("нашли айди")
	rows.Close()
	q = `select id_tag from tags_banner where id_banner = $1;
   		`
	rows, err = tx.Query(context.Background(), q, baner.Id)
	if err != nil {
		db.Logger.Error("ошибка получения баннера " + err.Error())
		return
	}
	var tags_id []int
	for rows.Next() {
		var tmp int
		_ = rows.Scan(&tmp)
		tags_id = append(tags_id, tmp)
	}
	db.Logger.Info("загрузили теги")
	rows.Close()
	baner.Tag_ids = tags_id
	q = `select created_at from banner where id_banner = $1;
	`
	rows, err = tx.Query(context.Background(), q, baner.Id)
	if err != nil {
		return
	}
	if rows.Next() {
		_ = rows.Scan(&baner.Created_at)
	} else {
		err = errors.New("404")
		db.Logger.Error("ошибка получения баннера " + err.Error())
		rows.Close()
		return
	}
	db.Logger.Info("загрузили дату создания")
	rows.Close()
	q = `select content, updated_at, is_active from version_banner where id_banner = $1 order by updated_at desc limit 1;
	`
	rows, err = tx.Query(context.Background(), q, baner.Id)
	if err != nil {
		db.Logger.Error("ошибка получения баннера " + err.Error())
		return
	}
	if rows.Next() {
		var tmp []byte
		_ = rows.Scan(&tmp, &baner.Updatet_at, &baner.Is_active)
		_ = easyjson.Unmarshal(tmp, &baner.Content)
	} else {
		err = errors.New("404")
		db.Logger.Error("ошибка получения баннера " + err.Error())
		return
	}
	db.Logger.Info("загрузили основные данные")
	db.Logger.Info("баннер " + baner.Content.Text)
	rows.Close()
	return
}

func (db *Database) GetListBannerByTagAndFutureIdWithOffsetAndLimit(tag_id, future_id, offset, limit int) (baners entitys.Banners, err error) {
	// возможно быстрее через джоин
	q := `select distinct tags_banner.id_banner from tags_banner join (select id_banner from features_banner where id_future = $1 or $1<=0 ) as first_query on first_query.id_banner = tags_banner.id_banner where tags_banner.id_tag = $2 or $2 <= 0 order by tags_banner.id_banner desc LIMIT all offset $3;
	`
	// `select id_banner from features_banner join tmp as (select id_banner from tags_banner where tags_banner.id_tag = $2) using features_banner.id_banner=tmp.id_banner where where id_future = $1;`
	tx, err := db.Pool.Begin(context.Background())
	if err != nil {
		db.Logger.Error("ошибка получения всех баннеров по шаблону " + err.Error())
		return
	}
	defer func() {
		_ = tx.Rollback(context.Background())
	}()
	rows, err := tx.Query(context.Background(), q, future_id, tag_id, offset)
	if err != nil {
		db.Logger.Error("ошибка получения всех баннеров по шаблону " + err.Error())
		return
	}
	var ids []int
	for (len(ids) < limit-offset || limit == 0) && rows.Next() {
		var tmp int
		_ = rows.Scan(&tmp)
		ids = append(ids, tmp)
	}
	rows.Close()
	baners, err = db.GetListBannersByListId(ids)
	return
}

func (db *Database) UpdateBannerById(bannerWithId *entitys.Banner) (err error) {
	q := "select id_banner from banner where id_banner=$1;"
	tx, err := db.Pool.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		return
	}
	rows, err := tx.Query(context.Background(), q, bannerWithId.Id)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return
	}
	if rows.Next() {
		var tmp int
		_ = rows.Scan(&tmp)
		rows.Close()
	} else {
		rows.Close()
		err = errors.New("404")
		_ = tx.Rollback(context.Background())
		return
	}
	q = "insert into version_banner(id_banner, content, updated_at, is_active) values ($1, $2, $3, $4) returning id_version;"
	data, err := easyjson.Marshal(bannerWithId.Content)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return
	}
	var t time.Time = time.Now()
	rows, err = tx.Query(context.Background(), q, bannerWithId.Id, data, t, bannerWithId.Is_active)
	// вставляем новую версию!!
	if err != nil {
		_ = tx.Rollback(context.Background())
		return
	}
	var version int
	if rows.Next() {
		_ = rows.Scan(&version)
	}
	rows.Close()
	q = "insert into features_banner(id_version, id_future, id_banner) values ($1, $2, $3);"
	rows, err = tx.Query(context.Background(), q, version, bannerWithId.Feature_ids, bannerWithId.Id)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return
	}
	rows.Close()
	q = "insert into tags_banner(id_version, id_tag, id_banner) values ($1, $2, $3);"
	for i := 0; i < len(bannerWithId.Tag_ids); i++ {
		rows, err = tx.Query(context.Background(), q, version, bannerWithId.Tag_ids[i], bannerWithId.Id)
		if err != nil {
			_ = tx.Rollback(context.Background())
			return
		}
		rows.Close()
	}
	_ = tx.Commit(context.Background())
	return
}

func (db *Database) DeleteBannerById(id int) (err error) {
	tx, err := db.Pool.Begin(context.Background())
	if err != nil {
		return
	}
	err = del(id, tx)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return
	} else {
		_ = tx.Commit(context.Background())
		return
	}
}

func (db *Database) DeleteBannerByFutureId(future_id int) {
	q := "select id_banner from features_banner where id_future=$1;"
	tx, err := db.Pool.Begin(context.Background())
	if err != nil {
		return
	}
	rows, err := tx.Query(context.Background(), q, future_id)
	if err != err {
		return
	}
	var ids []int
	for rows.Next() {
		var tmp int
		_ = rows.Scan(&tmp)
		ids = append(ids, tmp)
	}
	rows.Close()
	for i := 0; i < len(ids); i++ {
		err = del(ids[i], tx)
		if err != nil {
			_ = tx.Rollback(context.Background())
			return
		}
	}
	_ = tx.Commit(context.Background())
}

func (db *Database) CreateBanner(bannerWithoutId *entitys.Banner) (banner_id int, err error) {
	q := "insert into banner(created_at) values ($1) returning id_banner;"
	tx, err := db.Pool.Begin(context.Background())
	if err != nil {
		return
	}
	t := time.Now()
	rows, err := tx.Query(context.Background(), q, t)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return
	}
	if rows.Next() {
		_ = rows.Scan(&bannerWithoutId.Id)
	}
	banner_id = bannerWithoutId.Id
	rows.Close()
	q = "insert into version_banner(id_banner, content, updated_at, is_active) values ($1, $2, $3, $4) returning id_version;"
	data, err := easyjson.Marshal(bannerWithoutId.Content)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return
	}
	rows, err = tx.Query(context.Background(), q, bannerWithoutId.Id, data, t, bannerWithoutId.Is_active)
	// вставляем новую версию!!
	if err != nil {
		_ = tx.Rollback(context.Background())
		return
	}
	var version int
	if rows.Next() {
		_ = rows.Scan(&version)
	}
	rows.Close()
	q = "insert into features_banner(id_version, id_future, id_banner) values ($1, $2, $3);"
	rows, err = tx.Query(context.Background(), q, version, bannerWithoutId.Feature_ids, bannerWithoutId.Id)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return
	}
	rows.Close()
	q = "insert into tags_banner(id_version, id_tag, id_banner) values ($1, $2, $3);"
	for i := 0; i < len(bannerWithoutId.Tag_ids); i++ {
		rows, err = tx.Query(context.Background(), q, version, bannerWithoutId.Tag_ids[i], bannerWithoutId.Id)
		if err != nil {
			_ = tx.Rollback(context.Background())
			return
		}
		rows.Close()
	}
	_ = tx.Commit(context.Background())
	return
}

func del(id int, tx pgx.Tx) (err error) {
	q := "delete from features_banner where id_banner=$1;"
	rows, err := tx.Query(context.Background(), q, id)
	if err != nil {
		return
	}
	rows.Close()
	q = "delete from tags_banner where id_banner=$1;"
	rows, err = tx.Query(context.Background(), q, id)
	if err != nil {
		return
	}
	rows.Close()
	q = "delete from version_banner where id_banner=$1;"
	rows, err = tx.Query(context.Background(), q, id)
	if err != nil {
		return
	}
	rows.Close()
	q = "delete from banner where id_banner=$1;"
	rows, err = tx.Query(context.Background(), q, id)
	if err != nil {
		return
	}
	rows.Close()
	return
}

func (db *Database) GetThreeVersionBannerById(id int) (baners entitys.Banners, err error) {
	tx, err := db.Pool.Begin(context.Background())
	if err != nil {
		db.Logger.Error("ошибка получения трех баннеров по шаблону " + err.Error())
		return
	}
	defer func() {
		_ = tx.Rollback(context.Background())
	}()
	var created time.Time
	q := "select created_at from banner where id_banner=$1;"
	rows, err := tx.Query(context.Background(), q, id)
	if err != nil {
		return
	}
	if rows.Next() {
		_ = rows.Scan(&created)
	} else {
		err = errors.New("404")
		rows.Close()
		return
	}
	rows.Close()
	q = "select id_version, id_banner, content, updated_at, is_active from version_banner where id_banner=$1 order by updated_at DESC limit 3;"
	rows, err = tx.Query(context.Background(), q, id)
	if err != nil {
		return
	}
	var versions []int
	for rows.Next() {
		var banner entitys.Banner
		banner.Created_at = created
		var tmp int
		var tmp1 []byte
		_ = rows.Scan(&tmp, &banner.Id, &tmp1, &banner.Updatet_at, &banner.Is_active)
		_ = easyjson.Unmarshal(tmp1, &banner.Content)
		versions = append(versions, tmp)
		baners = append(baners, banner)
	}
	rows.Close()
	for i := 0; i < len(versions); i++ {
		q = "select id_tag from tags_banner where id_version=$1;"
		rows, err = tx.Query(context.Background(), q, versions[i])
		var tags []int
		for rows.Next() {
			var tmp int
			_ = rows.Scan(&tmp)
			tags = append(tags, tmp)
		}
		baners[i].Tag_ids = tags
		rows.Close()
		q = "select id_future from features_banner where id_version=$1;"
		rows, err = tx.Query(context.Background(), q, versions[i])
		var future int
		for rows.Next() {
			_ = rows.Scan(&future)
		}
		baners[i].Feature_ids = future
		rows.Close()
	}
	db.Logger.Info(fmt.Sprintf("%+v", baners))
	return
}
