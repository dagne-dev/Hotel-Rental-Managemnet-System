package news

import "github.com/getme04/Hotel-Rental-Managemnet-System/entity"

// CategoryService specifies food menu category News_Services
type NewsRepository interface {
	News() ([]entity.News, error)
	NewsById(id int) (entity.News, error)
	UpdateNews(news entity.News) error
	DeleteNews(id int) error
	StoreNews(news entity.News) error
}