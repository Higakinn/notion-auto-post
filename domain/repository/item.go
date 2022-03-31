package repository

import (
	"github.com/higakinn/notion-auto-post/domain/model"
) 
type IItemRepository interface {
  FindAll() ([]model.Item, error)
}