package dbRepository

import (
	"avia/app/domain/entity"
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

// ReminderRepo struct
type ReminderRepo struct {
	db *gorm.DB
}

// ReminderRepositoryInit initial
func ReminderRepositoryInit(db *gorm.DB) *ReminderRepo {
	return &ReminderRepo{db}
}

//ReminderRepo implements the repository.ReminderRepository interface
// var _ interfaces.PostAppInterface = &ReminderRepo{}

// Save data
func (r *ReminderRepo) Save(post *entity.Reminder) (*entity.Reminder, map[string]string) {
	dbErr := map[string]string{}
	//The images are uploaded to digital ocean spaces. So we need to prepend the url. This might not be your use case, if you are not uploading image to Digital Ocean.
	var err error
	err = r.db.Debug().Create(&post).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "post title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return post, nil
}

// Update  data
func (r *ReminderRepo) Update(post *entity.Reminder) (*entity.Reminder, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&post).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return post, nil
}

// Count fat
func (r *ReminderRepo) Count(postTotalCount *int64) {
	var post entity.Reminder
	var count int64
	r.db.Debug().Model(post).Count(&count)
	*postTotalCount = count
}

// Delete data
func (r *ReminderRepo) Delete(id uint64) error {
	var post entity.Reminder
	var err error
	err = r.db.Debug().Where("id = ?", id).Delete(&post).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

// GetByID get data
func (r *ReminderRepo) GetByID(id uint64) (*entity.Reminder, error) {
	var post entity.Reminder
	var err error
	err = r.db.Debug().Where("id = ?", id).Take(&post).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return &post, nil
}

// GetAll all data
func (r *ReminderRepo) GetAll() ([]entity.Reminder, error) {
	var posts []entity.Reminder
	var err error
	err = r.db.Debug().Order("created_at desc").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return posts, nil
}

// GetAllPagination pagination all data
func (r *ReminderRepo) GetAllPagination(postsPerPage int, offset int) ([]entity.Reminder, error) {
	var posts []entity.Reminder
	var err error
	err = r.db.Debug().Limit(postsPerPage).Offset(offset).Order("created_at desc").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return posts, nil
}
