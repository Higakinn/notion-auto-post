package notion

import (
	"os"
	"fmt"
	"net/http"

	// "github.com/pkg/errors"
	"github.com/higakinn/notion-auto-post/lib"
	"github.com/higakinn/notion-auto-post/domain/model"
	"github.com/higakinn/notion-auto-post/domain/repository"
)

type HTTPAPI interface {
	Do(req *http.Request) (*http.Response, error)
}

type ItemRepository struct {
	notionClient lib.NotionClient
}

func NewItemRepository() repository.IItemRepository {
	accessToken := os.Getenv("NOTION_ACCESS_TOKEN")
	nClient := lib.NewNotionClient(accessToken)
	return &ItemRepository{
		notionClient: *nClient,
	}
}

func (r *ItemRepository) FindAll() ([]model.Item, error) {
	dbId := os.Getenv("NOTION_DATABASE_ID")
	
	dbs, _ := r.notionClient.QueryDatabases(dbId)
	fmt.Println(dbs)
	items := []model.Item{}
	for _,v := range dbs.Results {
		item := model.Item{
			ID: v.ID,
			Title: v.Properties.Name.Title[0].Text.Content,
			Url: v.Properties.URL.URL,
			CreateDate: "test",
		}
		items = append(items, item)
	}

	return items, nil
}