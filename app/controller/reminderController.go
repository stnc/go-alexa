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

/*
	func (access *ReminderControl) Create(c *gin.Context) {
		stncsession.IsLoggedInRedirect(c)
		locale, menuLanguage := lang.LoadLanguages("Reminder")
		flashMsg := stncsession.GetFlashMessage(c)
		region, _ := access.Region.GetAll()
		roles, _ := access.RoleApp.GetAll()

		//#json formatter #stncjson
		empJSON, err := json.MarshalIndent(region, "", "  ")
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))

		viewData := pongo2.Context{
			"title":       "İçerik Ekleme",
			"regions":     region,
			"roles":       roles,
			"flashMsg":    flashMsg,
			"csrf":        csrf.GetToken(c),
			"locale":      locale,
			"localeMenus": menuLanguage,
		}
		c.HTML(
			http.StatusOK,
			viewPathReminderControl+"create.html",
			viewData,
		)
	}

	func (access *ReminderControl) Store(c *gin.Context) {
		stncsession.IsLoggedInRedirect(c)
		locale, menuLanguage := lang.LoadLanguages("Reminder")
		flashMsg := stncsession.GetFlashMessage(c)
		roles, _ := access.RoleApp.GetAll()
		var ReminderSave = ReminderModel(c, "create", "")
		var ReminderSavePostError = make(map[string]string)
		ReminderSavePostError = ReminderSave.Validate()

		if len(ReminderSavePostError) == 0 {
			saveData, saveErr := access.ReminderControlApp.Save(&ReminderSave)
			if saveErr != nil {
				ReminderSavePostError = saveErr
			}
			lastID := strconv.FormatUint(uint64(saveData.ID), 10)
			stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
			c.Redirect(http.StatusMovedPermanently, "/admin/Reminder/edit/"+lastID)
			return
		} else {
			stncsession.SetFlashMessage("Zorunlu alanları lütfen doldurunuz", "danger", c)
		}
		viewData := pongo2.Context{
			"title":       "content add",
			"csrf":        csrf.GetToken(c),
			"err":         ReminderSavePostError,
			"data":        ReminderSave,
			"flashMsg":    flashMsg,
			"roles":       roles,
			"locale":      locale,
			"localeMenus": menuLanguage,
		}
		c.HTML(
			http.StatusOK,
			viewPathReminderControl+"create.html",
			viewData,
		)

}

	func (access *ReminderControl) Edit(c *gin.Context) {
		stncsession.IsLoggedInRedirect(c)
		locale, menuLanguage := lang.LoadLanguages("Reminder")
		flashMsg := stncsession.GetFlashMessage(c)
		if ReminderID, err := strconv.ParseUint(c.Param("ReminderID"), 10, 64); err == nil {
			if ReminderData, err := access.ReminderControlApp.GetByID(ReminderID); err == nil {
				roles, _ := access.RoleApp.GetAll()
				region, _ := access.Region.GetAll()

				dataReminderForBranchID, _ := access.ReminderControlApp.GetByReminderForBranchID(ReminderData.BranchID)
				branchID := dataReminderForBranchID.BranchID
				regionID := dataReminderForBranchID.RegionID

				viewData := pongo2.Context{
					"title":       "kullanıcı düzenleme",
					"data":        ReminderData,
					"csrf":        csrf.GetToken(c),
					"flashMsg":    flashMsg,
					"regions":     region,
					"branchID":    branchID,
					"regionID":    regionID,
					"roles":       roles,
					"locale":      locale,
					"localeMenus": menuLanguage,
				}
				c.HTML(
					http.StatusOK,
					viewPathReminderControl+"edit.html",
					viewData,
				)
			} else {
				c.AbortWithError(http.StatusNotFound, err)
			}

		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	}

	func (access *ReminderControl) Update(c *gin.Context) {
		stncsession.IsLoggedInRedirect(c)
		flashMsg := stncsession.GetFlashMessage(c)
		locale, menuLanguage := lang.LoadLanguages("Reminder")
		roles, _ := access.RoleApp.GetAll()
		id := c.PostForm("ReminderID")
		id2 := stnccollection.StringtoUint64(id)
		var pass string
		if ReminderData, err := access.ReminderControlApp.GetByID(id2); err == nil {
			pass = ReminderData.Password
		}
		var ReminderControl = ReminderModel(c, "edit", pass)
		var ReminderSavePostError = make(map[string]string)
		ReminderSavePostError = ReminderControl.Validate()
		region, _ := access.Region.GetAll()
		if len(ReminderSavePostError) == 0 {
			_, saveErr := access.ReminderControlApp.Update(&ReminderControl)
			if saveErr != nil {
				ReminderSavePostError = saveErr
			}
			stncsession.SetFlashMessage("Save Succesful ", "success", c)
			c.Redirect(http.StatusMovedPermanently, "/"+viewPathReminderControl+"edit/"+id)
			return
		} else {
			stncsession.SetFlashMessage("required field ", "danger", c)
		}

		viewData := pongo2.Context{
			"title":       "Reminder Edit",
			"err":         ReminderSavePostError,
			"csrf":        csrf.GetToken(c),
			"flashMsg":    flashMsg,
			"regions":     region,
			"data":        ReminderControl,
			"roles":       roles,
			"locale":      locale,
			"localeMenus": menuLanguage,
		}

		c.HTML(
			http.StatusOK,
			viewPathReminderControl+"edit.html",
			viewData,
		)
	}

// Delete data

	func (access *ReminderControl) Delete(c *gin.Context) {
		stncsession.IsLoggedInRedirect(c)
		if postID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {
			access.ReminderControlApp.Delete(postID)
			stncsession.SetFlashMessage("Success Delete", "success", c)
			c.Redirect(http.StatusMovedPermanently, "/admin/Reminder")
			return
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	}

// form post model

	func ReminderModel(c *gin.Context) (Reminder entity.Reminder) {
		uintInt, _ = strconv.ParseUint(c.PostForm("ID"), 10, 64)
		Reminder.ID = stnccollection.StringtoUint64(c.PostForm("ID"))
		Reminder.Remindername = c.PostForm("Remindername")
		Reminder.Email = c.PostForm("Email")
		Reminder.FirstName = c.PostForm("FirstName")
		Reminder.LastName = c.PostForm("LastName")
		Reminder.Phone = c.PostForm("Phone")
		return Reminder
	}
*/
func SaveData(Reminder entity.Reminder) {

	db := repository.DbConnect()
	services, _ := repository.RepositoriesInit(db)
	reminder := InitReminderControl(services.Reminder)
	reminder.ReminderControlApp.Save(&Reminder)
}
