package entitys

import "time"

type Banner struct {
	Id int `json:"banner_id"` 
	Tag_ids []int `json:"tag_ids"` 
	Feature_ids []int `json:"feature_id"` 
	Content Content `json:"content"` 
	Is_active bool `json:"is_active"` 
	Created_at time.Time `json:"created_at"` 
	Updatet_at time.Time `json:"updated_at"` 
}

type Content struct {
	Title string `json:"title"`
	Text string `json:"text"`
	Url string `json:"url"`
}