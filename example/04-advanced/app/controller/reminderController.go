package controller

import (
	"avia/app/domain/entity"
	repository "avia/app/domain/repository"
	"avia/app/services"
	"github.com/flosch/pongo2/v5"
	"log"
	"net/http"
)

// ReminderControl constructor
type ReminderControl struct {
	ReminderControlApp services.ReminderAppInterface
}

// InitReminderControl ReminderControl controller constructor
func InitReminderControl(KiApp services.ReminderAppInterface) *ReminderControl {
	return &ReminderControl{
		ReminderControlApp: KiApp,
	}
}

// Index list
func (access *ReminderControl) Index(w http.ResponseWriter, req *http.Request) {
	//var reminder entity.Reminder
	//reminder.PersonName = "selman 4"
	//reminder.RemindDate = "2023-07-12"
	//reminder.RemindTime = "04:00"
	//reminder.NumberOfPeople = "4"
	//reminder.Email = "dsds@d.com"
	//reminder.Phone = "5354543543"
	//SaveData(reminder)
	var total int64
	access.ReminderControlApp.Count(&total)
	list, _ := access.ReminderControlApp.GetAll()
	//
	//var Reminder entity.Reminder
	//Reminder.PersonName = "selman"
	//Reminder.RemindDate = "2023-07-12"
	//Reminder.RemindTime = "04:00"
	//Reminder.NumberOfPeople = "4"
	//Reminder.Email = "dsds@d.com"
	//Reminder.Phone = "5354543543"
	//
	//access.ReminderControlApp.Save(&Reminder)

	tpl, err := pongo2.FromFile("app/view/index.html")
	if err != nil {
		log.Fatal(err)
	}
	ctx := pongo2.Context{"title": "list", "total": total, "list": list}
	err2 := tpl.ExecuteWriter(ctx, w)
	if err2 != nil {
		log.Fatal(err2)
	}

}

func SaveData(Reminder entity.Reminder) {
	db := repository.DbConnect()
	services, _ := repository.RepositoriesInit(db)
	reminder := InitReminderControl(services.Reminder)
	reminder.ReminderControlApp.Save(&Reminder)
}
