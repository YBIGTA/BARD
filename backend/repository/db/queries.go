package repository

func GetOne[T any](query T) (T, error) {
	var model T

	err := DB.Where(query).First(&model).Error

	return model, err
}
