package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type FkValidator struct {
	db *gorm.DB
}

func NewFkValidator(db *gorm.DB) FkValidator {
	return FkValidator{
		db: db,
	}
}

// fk validator
// please send destionation table name as foreign key
func (v FkValidator) Handler() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().Uint()
		table := fl.Param()
		var count int64
		v.db.Table(table).Where("id=?", value).Count(&count)
		return count > 0
	}
}

func (v FkValidator) Setup() error {
	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := validator.RegisterValidation("fkGorm", v.Handler()); err != nil {
			return err
		}
	}
	return nil
}
