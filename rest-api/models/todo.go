package models

import (
	"log"
	"rest-api/pkg/constvar"

	validator "gopkg.in/go-playground/validator.v9"
)

// TodoModel - Desription To Do of User
type TodoModel struct {
	BaseModel
	Status uint64 `gorm:"column:status;type:int(11);default:0;not null" json:"status" validate:"eq=0|eq=1"`
	UserID uint64 `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	Title  string `gorm:"column:title;not null" json:"title" validate:"min=5,max=300"`
}

type Result struct {
	Name 	string
	Title  	string
}

// TableName - mapping name table in database
func (td *TodoModel) TableName() string {
	return "todo"
}

// Create todo
func (td *TodoModel) Create() error {
	return DB.Self.Create(&td).Error
}

func ListTodo(userID, offset, limit int) ([]*TodoModel, uint64, error) {

	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	todos := make([]*TodoModel, 0)
	var count uint64

	if err := DB.Self.Model(&TodoModel{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return todos, count, err
	}

	var result Result

	DB.Self.Debug().Table("user").Select("user.id, user.username, todo.title").Joins("left join todo on todo.user_id = user.id").Where("user.id = ?", userID).Scan(&result)
	log.Println("result: ", result)

	if err := DB.Self.Offset(offset).Limit(limit).Where("user_id = ?", userID).Order("id asc").Find(&todos).Error; err != nil {
		return todos, count, err
	}

	return todos, count, nil
}

// DeleteTodo todo by id
func DeleteTodo(userID, id uint64) error {
	td := &TodoModel{}
	todo := TodoModel{}
	todo.BaseModel.ID = id
	if err := DB.Self.Where("id = ? AND user_id = ?", id, userID).First(&td).Error; err != nil {
		log.Printf("Error when delete todo: %s", err)
		return err
	}
	return DB.Self.Delete(&todo).Error
}

// Update todo
func (td *TodoModel) Update(userID uint64) error {

	if err := DB.Self.Where("id = ? AND user_id = ?", td.ID, userID).Find(&TodoModel{}).Error; err != nil {
		return err
	}
	return DB.Self.Model(&TodoModel{}).Select("title", "status").Where("user_id = ? AND id = ?", userID, td.ID).Updates(map[string]interface{}{"title": td.Title, "status": td.Status}).Error
}

// Validate the fields of TodoModel
func (td *TodoModel) Validate() error {
	validate := validator.New()
	return validate.Struct(td)
}
