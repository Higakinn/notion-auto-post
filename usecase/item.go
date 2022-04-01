package usecase

import (
	
	"github.com/higakinn/notion-auto-post/domain/repository"
	"github.com/higakinn/notion-auto-post/domain/model"
)
type ItemUseCase interface {
	GetAll() ([]model.Item, error)
}

type itemUseCase struct {
	itemRepository repository.ItemRepository
}

func NewItemUseCase(ir repository.ItemRepository) ItemUseCase {
	return &itemUseCase {
		itemRepository: ir,
	}
}

func (iu *itemUseCase) GetAll() (items []model.Item, err error){
	items, err = iu.itemRepository.FindAll()
	if err != nil {
    return nil, err
  }
  return items, nil
}