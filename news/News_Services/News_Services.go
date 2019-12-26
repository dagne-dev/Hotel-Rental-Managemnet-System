package News_Services

import (

	"github.com/getme04/original/News"
	"github.com/getme04/original/entity"
)

// RoomServiceImpl implements rooms.RoomService interface
type NewsServiceImpl struct {
	newsRepo News.NewsService
}

// NewNewsServiceImpl will create new RoomService object
func NewNewsServiceImpl(NewRepo News.NewsService) *NewsServiceImpl {
	return &NewsServiceImpl{newsRepo: NewRepo}
}

// News returns list of all rooms
func (rs *NewsServiceImpl) News() ([]entity.News, error) {

	return news, nil
}

// StoreNews persists new room information
func (rs *NewsServiceImpl) StoreNews(neww entity.News) error {


	return nil
}

// NewById returns a room object with a given id
func (rs *NewsServiceImpl) NewsById(id int) (entity.News, error) {

	return r, nil
}

// UpdateNews updates a cateogory with new data
func (rs *NewsServiceImpl) UpdateNews(neww entity.News) error {



	return nil
}

// DeleteNews delete a room by its id
func (rs *NewsServiceImpl) DeleteNews(id int) error {

	return nil
}


