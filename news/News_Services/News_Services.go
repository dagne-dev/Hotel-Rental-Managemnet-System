package News_Services
import (
	"github.com/getme04/Hotel-Rental-Managemnet-System/news"
	"github.com/getme04/Hotel-Rental-Managemnet-System/entity"
)

// RoomServiceImpl implements rooms.RoomService interface
type NewsServiceImpl struct {
	newsRepo news.NewsService
}

// NewNewsServiceImpl will create new NewsService object
func NewNewsServiceImpl(NewRepo news.NewsService) *NewsServiceImpl {
	return &NewsServiceImpl{newsRepo: NewRepo}
}

// News returns list of all rooms
func (rs *NewsServiceImpl) News() ([]entity.News, error) {

	news, err := rs.newsRepo.News()

	if err != nil {
		return nil, err
	}

	return news, nil
}


