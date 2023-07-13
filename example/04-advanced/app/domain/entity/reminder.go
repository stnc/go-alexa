package entity

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"html"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
)

// Reminder strcut
type Reminder struct {
	ID             uint64     `gorm:"primary_key;auto_increment" json:"id"`
	PersonName     string     `gorm:"type:varchar(255) ;not null;"  json:"name" `
	NumberOfPeople string     `gorm:"type:text ;"  json:"people"`
	RemindDate     string     `gorm:"type:text ;" json:"time"`
	RemindTime     string     `gorm:"type:text ;" json:"date"`
	Phone          string     `gorm:"type:varchar(255) ;null;" json:"phone"`
	Email          string     `gorm:"type:varchar(255) ;null;" json:"email"`
	CreatedAt      time.Time  ` json:"created_at"`
	UpdatedAt      time.Time  ` json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

// BeforeSave init
func (f *Reminder) BeforeSave() {
	f.PersonName = html.EscapeString(strings.TrimSpace(f.PersonName))
	f.NumberOfPeople = html.EscapeString(strings.TrimSpace(f.NumberOfPeople))

}

/*
func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
 }
*/

// Prepare init
func (f *Reminder) Prepare() {
	f.PersonName = html.EscapeString(strings.TrimSpace(f.PersonName))
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}

// Validate fluent validation
func (f *Reminder) Validate() map[string]string {
	var (
		validate *validator.Validate
		uni      *ut.UniversalTranslator
	)
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate = validator.New()
	tr_translations.RegisterDefaultTranslations(validate, trans)

	errorLog := make(map[string]string)

	err := validate.Struct(f)
	fmt.Println(err)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs)
		for _, e := range errs {
			// can translate each error one at a time.
			lng := strings.Replace(e.Translate(trans), e.Field(), "This Here", 1)
			errorLog[e.Field()+"_error"] = e.Translate(trans)
			// errorLog[e.Field()] = e.Translate(trans)
			errorLog[e.Field()] = lng
			errorLog[e.Field()+"_valid"] = "is-invalid"
		}
	}
	return errorLog
}
