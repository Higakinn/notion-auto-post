package batch

import (
	"fmt"
	"os"

	"github.com/higakinn/notion-auto-post/infrastructure/notion"
	"github.com/higakinn/notion-auto-post/usecase"
	"github.com/higakinn/notion-auto-post/lib"
	"github.com/pkg/errors"
)

type IRunner interface {
	Execute()
}
type Runner struct {}

func NewRunner() IRunner {
	return &Runner{}
}

func (r *Runner) Execute() {
	fmt.Println("batch start")
	itemRepo := notion.NewItemRepository()
	qClient := lib.NewQiitaClient()

	itemUC := usecase.NewItemUseCase(itemRepo)
	pages,_ := itemUC.GetAll()

	qBody := ""
	for _,p := range pages {
		fmt.Println(p)
		qBody += "## " + p.Title + "\n\n" + p.Url + "\n\n"
	}

	tags := []lib.QiitaTag{ {"Qiita", [] string{}} }
	if err := qClient.UpdateItem(os.Getenv("QIITA_ARTICLE_ID"), "読みたい記事", qBody, tags); err != nil {
			errors.WithStack(err)
	}

	fmt.Println("batch end")

}