package service

import (
	"github.com/baajarmeh/fiber-estate/pkg/model"
	"github.com/baajarmeh/fiber-estate/pkg/repository"
)

type Room interface {
	Create(room *model.Room) (int, error)
	Delete(id int) error
	GetAll(sortField string) ([]*model.Room, error)
}

type Booking interface {
	Create(booking *model.Booking) (int, error)
	Delete(id int) error
	GetByRoomId(roomId int) ([]*model.Booking, error)
}

type Service struct {
	Room
	Booking
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Room:    NewRoomService(repos.Room),
		Booking: NewBookingService(repos.Booking, repos.Room),
	}
}
