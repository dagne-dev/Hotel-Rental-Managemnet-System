
package News_Repository


import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/getme04/Hotel-Rental-Managemnet-System/entity"
)

// RoomRepositoryImpl implements the rooms.RoomRepository interface
type NewsRepositoryImpl struct {
	Con *sql.DB
}

// NewRoomRepositoryImpl will create an object of PsqlRoomRepository
func NewNewsRepositoryImpl(Conn *sql.DB) *NewsRepositoryImpl {
	return &NewsRepositoryImpl{Con: Conn}
}

// News returns all rooms from the database
func (rri *NewsRepositoryImpl) News() ([]entity.News, error) {

	rows, err := rri.Con.Query("SELECT * FROM News;")
	if err != nil {

		return nil, errors.New("Could not query the database")
	}
	fmt.Printf("data fetched exactlly")
	defer rows.Close()

	newsList := []entity.News{}

	for rows.Next() {
		new := entity.News{}
		err = rows.Scan(&new.Id, &new.Header,&new.Description, &new.Image)
		if err != nil {
			return nil, err
		}
		newsList = append(newsList, new)
	}

	return newsList, nil
}