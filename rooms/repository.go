package rooms

import "github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"

// RoomRepository specifies room related database operations
type RoomRepository interface {
	Rooms() ([]entity.Room, error)
	Room(id int) (entity.Room, error)
	UpdateRoom(category entity.Room) error
	DeleteRoom(id int) error
	StoreRoom(category entity.Room) error
}