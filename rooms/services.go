package rooms

import "github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"

// RoomServices specifies room menu services
type RoomServices interface {
	Rooms() ([]entity.Room, error)
	Room(id int) (entity.Room, error)
	UpdateRoom(category entity.Room) error
	DeleteRoom(id int) error
	StoreRoom(category entity.Room) error
}