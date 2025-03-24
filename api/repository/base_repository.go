package repository

import (
	"gorm.io/gorm"
	"reflect"
)

type BaseRepository[T any] struct{}

func (r *BaseRepository[T]) CreateWithID(tx *gorm.DB, entity *T) (id int, err error) {
	err = r.Create(tx, entity)
	id = r.getID(entity)
	return
}

func (r *BaseRepository[T]) Create(tx *gorm.DB, entity *T) error {
	return tx.Create(&entity).Error
}

func (r *BaseRepository[T]) getID(entity *T) int {
	idField := reflect.ValueOf(entity).Elem().FieldByName("ID")
	if idField.IsValid() && idField.CanInterface() {
		return idField.Interface().(int)
	}
	return 0
}

func (r *BaseRepository[T]) Update(tx *gorm.DB, entity *T) error {
	return tx.Save(&entity).Error
}
