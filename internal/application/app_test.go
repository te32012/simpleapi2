package application_test

import (
	"avitotestgo2024/internal/entitys"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/mailru/easyjson"
)

func TestUseApplicationOne(t *testing.T) {
	client := http.Client{}
	var ban entitys.Banner
	ban.Feature_ids = 10
	ban.Content = entitys.Content{Title: "abc", Text: "example", Url: "URI"}
	ban.Is_active = true
	ban.Tag_ids = append(ban.Tag_ids, 23)
	ban.Tag_ids = append(ban.Tag_ids, 24)

	reqBody, err := easyjson.Marshal(ban)
	req, err := http.NewRequest("POST", "http://application:2024/banner", bytes.NewBuffer(reqBody))
	req.Header.Set("token", "admin")
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		t.Fatalf("Error making request: %v", err)
	}
	data, err := io.ReadAll(resp.Body)
	var id entitys.Ans201
	err = easyjson.Unmarshal(data, &id)
	if err != nil {
		t.Fatalf("Error unmarshal: %v", err)
	}
	req, err = http.NewRequest("GET", "http://application:2024/user_banner", nil)
	req.Header.Set("token", "user")
	q := req.URL.Query()
	q.Set("tag_id", "23")
	q.Set("feature_id", "10")
	q.Set("use_last_revision", "true")
	req.URL.RawQuery = q.Encode()
	resp, err = client.Do(req)
	// Проверяем статус код ответа
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	var cont entitys.Content
	data, _ = io.ReadAll(resp.Body)
	fmt.Println(len(data))
	err = easyjson.Unmarshal(data, &cont)
	if err != nil {
		t.Fatalf("Error get banner: %s", err.Error())

	}
	if cont.Text != ban.Content.Text || cont.Title != ban.Content.Title || cont.Url != ban.Content.Url {
		t.Fatalf("Error get banner: %s, %s, %s || %s, %s, %s", cont.Text, cont.Title, cont.Url, ban.Content.Text, ban.Content.Title, ban.Content.Url)
	}
	ban.Is_active = false
	ban.Tag_ids = append(ban.Tag_ids, 25)
	reqBody, err = easyjson.Marshal(&ban)
	req, err = http.NewRequest("PATCH", "http://application:2024/banner/"+strconv.Itoa(id.Id), bytes.NewBuffer(reqBody))
	req.Header.Set("token", "admin")
	resp, err = client.Do(req)
	if resp.StatusCode != 200 {
		t.Fatalf("Error get banner: %v", resp.StatusCode)
	}
	req, err = http.NewRequest("GET", "http://application:2024/banner", nil)
	req.Header.Set("token", "admin")
	q = req.URL.Query()
	q.Set("tag_id", "25")
	req.URL.RawQuery = q.Encode()
	resp, err = client.Do(req)
	if resp.StatusCode != 200 {
		t.Fatalf("Error status answer banner: %v", err)
	}
	var banners entitys.Banners
	data, _ = io.ReadAll(resp.Body)
	easyjson.Unmarshal(data, &banners)
	if len(banners) != 1 {
		t.Fatalf("Error status answer banner: %d", len(banners))
	}
	if banners[0].Feature_ids != 10 {
		t.Fatalf("Error status answer banner: %d", banners[0].Feature_ids)
	}
}
