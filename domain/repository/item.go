package repository

import (
	"github.com/higakinn/notion-auto-post/domain/model"
) 
type ItemRepository interface {
  FindAll() ([]model.Item, error)
}