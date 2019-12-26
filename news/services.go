package News

import "github.com/getme04/original/entity"

// NewsService specifies News menu News News_Services
type NewsService interface {
	News() ([]entity.News, error)
	NewsById(id int) (entity.News, error)
	UpdateNews(news entity.News) error
	DeleteNews(id int) error
	StoreNews(news entity.News) error
}
