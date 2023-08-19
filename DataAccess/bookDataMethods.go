package DataAccess

func (gdb *GormDB) GetAllBooks() ([]Book, error) {
	var books []Book
	if err := gdb.db.Preload("Authors").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
