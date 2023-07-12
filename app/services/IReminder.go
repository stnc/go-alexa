package services

import (
	"avia/app/domain/entity"
)

// ReminderAppInterface interface
type ReminderAppInterface interface {
	Save(*entity.Reminder) (*entity.Reminder, map[string]string)
	GetByID(uint64) (*entity.Reminder, error)
	GetAll() ([]entity.Reminder, error)
	GetAllPagination(int, int) ([]entity.Reminder, error)
	Update(*entity.Reminder) (*entity.Reminder, map[string]string)
	Count(*int64)
	Delete(uint64) error
}
type ReminderApp struct {
	request ReminderAppInterface
}

var _ ReminderAppInterface = &ReminderApp{}

func (f *ReminderApp) Save(Reminder *entity.Reminder) (*entity.Reminder, map[string]string) {
	return f.request.Save(Reminder)
}

func (f *ReminderApp) GetByID(ReminderID uint64) (*entity.Reminder, error) {
	return f.request.GetByID(ReminderID)
}

func (f *ReminderApp) GetAll() ([]entity.Reminder, error) {
	return f.request.GetAll()
}

func (f *ReminderApp) GetAllPagination(RemindersPerPage int, offset int) ([]entity.Reminder, error) {
	return f.request.GetAllPagination(RemindersPerPage, offset)
}

func (f *ReminderApp) Update(Reminder *entity.Reminder) (*entity.Reminder, map[string]string) {
	return f.request.Update(Reminder)
}

func (f *ReminderApp) Count(ReminderTotalCount *int64) {
	f.request.Count(ReminderTotalCount)
}

func (f *ReminderApp) Delete(ReminderID uint64) error {
	return f.request.Delete(ReminderID)
}
