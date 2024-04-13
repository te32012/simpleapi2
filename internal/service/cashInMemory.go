package service

import (
	"avitotestgo2024/internal/database"
	"sync"

	"github.com/mailru/easyjson"
)

type node struct {
	previous *node
	next     *node
	id       int
}

type Lru struct {
	head    *node
	last    *node
	l       int
	storage map[int]*node
	mutex   sync.Mutex
}

func (l *Lru) add(id int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	nd, ok := l.storage[id]
	if ok {
		if nd != l.last {
			if nd == l.head {
				l.head = l.head.next
				l.head.previous = nil
				nd.next = nil
				nd.previous = l.last
				l.last.next = nd
				l.last = nd
			}
			nd.previous.next = nd.next
			nd.next.previous = nd.previous
			nd.next = nil
			nd.previous = l.last
			l.last.next = nd
			l.last = nd
		}
	} else {
		if l.l < 1100 {
			if l.last == nil && l.head == nil {
				nd := new(node)
				nd.id = id
				l.last = nd
				l.head = nd
				l.storage[id] = nd
				return
			}
			nd := new(node)
			nd.id = id
			nd.previous = l.last
			l.last.next = nd
			l.last = nd
			l.storage[id] = nd
			l.l += 1
		} else {
			nd := l.head
			delete(l.storage, nd.id)
			l.head = l.head.next
			l.head.previous = nil
			nd.id = id
			nd.next = nil
			nd.previous = l.last
			l.last.next = nd
			l.last = nd
			l.storage[id] = nd
		}
	}
}

func (l *Lru) GetId() []int {
	l.mutex.Lock()
	var ans []int
	node := l.head
	for node.next != nil {
		ans = append(ans, node.id)
		node = node.next
	}
	l.mutex.Unlock()
	return ans
}

func newLru() *Lru {
	return &Lru{storage: make(map[int]*node)}
}

// Кешированные ответы
type ChashedBanner struct {
	ShortBanner []byte
}

// Объединенные по одному ID группы ответов для кеша по фичам и по тегам.
type GroupOfCash struct {
	Banners sync.Map
}

type TagAndFutureID struct {
	TagID    int
	FutureID int
}

// Экземлпяр кеша из памяти
type CashInMemory struct {
	// map[TagAndFutureID]*ChashedBanner
	FindIdByFeatureIdAndTagID sync.Map
	// map[int]*ChashedBanner
	FindById sync.Map
}

type CashUsedMap struct {
	// Кеш в памяти в данной ситуации используется только на чтение
	// Каждые 3 минуты актуальный кеш обновляется
	// В кеш идут последние 1500 баннеров, которые обрабатывал сервис
	CashOne      *CashInMemory
	CashTwo      *CashInMemory
	NumberOfCash int
	Lru          *Lru
	MainBase     database.DatabaseInterface
}

func NewChashUsedMap(base database.DatabaseInterface) *CashUsedMap {
	return &CashUsedMap{CashOne: &CashInMemory{}, CashTwo: &CashInMemory{}, NumberOfCash: 1, Lru: newLru(), MainBase: base}
}

func (c *CashUsedMap) GetShortByFutureIdAndTagId(tag_id int, future_id int) (bool, []byte) {
	switch c.NumberOfCash {
	case 1:
		id, ok := c.CashOne.FindIdByFeatureIdAndTagID.Load(TagAndFutureID{TagID: tag_id, FutureID: future_id})
		if ok {
			c.Used(id.(int))
			tmp, _ := c.CashOne.FindIdByFeatureIdAndTagID.Load(id.(int))
			return true, tmp.(*ChashedBanner).ShortBanner
		}
		return false, nil
	case 2:
		id, ok := c.CashTwo.FindIdByFeatureIdAndTagID.Load(TagAndFutureID{TagID: tag_id, FutureID: future_id})
		if ok {
			c.Used(id.(int))
			tmp, _ := c.CashTwo.FindIdByFeatureIdAndTagID.Load(id.(int))
			return true, tmp.(*ChashedBanner).ShortBanner
		}
		return false, nil
	}
	return false, nil
}

func (c *CashUsedMap) Used(id int) {
	c.Lru.add(id)
}

func (c *CashUsedMap) UpdateCash() {
	ids := c.Lru.GetId()

	switch c.NumberOfCash {
	case 1:
		tmp, err := c.MainBase.GetListBannersByListId(ids)
		if err != nil {
			return
		}
		c.CashTwo.FindById = sync.Map{}
		c.CashTwo.FindIdByFeatureIdAndTagID = sync.Map{}
		for i := 0; i < len(tmp); i++ {
			data, _ := easyjson.Marshal(tmp[i].Content)
			c.CashTwo.FindById.Store(tmp[i].Id, data)
			for j := 0; j < len(tmp[i].Tag_ids); j++ {
				tmp1 := TagAndFutureID{TagID: tmp[i].Tag_ids[j], FutureID: tmp[i].Feature_ids}
				c.CashTwo.FindIdByFeatureIdAndTagID.Store(tmp1, tmp[i].Id)
			}
		}
	case 2:
		tmp, err := c.MainBase.GetListBannersByListId(ids)
		if err != nil {
			return
		}
		c.CashOne.FindById = sync.Map{}
		c.CashOne.FindIdByFeatureIdAndTagID = sync.Map{}
		for i := 0; i < len(tmp); i++ {
			data, _ := easyjson.Marshal(tmp[i].Content)
			c.CashOne.FindById.Store(tmp[i].Id, data)
			for j := 0; j < len(tmp[i].Tag_ids); j++ {
				tmp1 := TagAndFutureID{TagID: tmp[i].Tag_ids[j], FutureID: tmp[i].Feature_ids}
				c.CashOne.FindIdByFeatureIdAndTagID.Store(tmp1, tmp[i].Id)
			}
		}
	}
	c.NumberOfCash = 3 - c.NumberOfCash
}
