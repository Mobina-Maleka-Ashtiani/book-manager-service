package DataAccess

func (gdb *GormDB) GetAllBooks() ([]Book, error) {
	var books []Book
	if err := gdb.db.Preload("Author").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (gdb *GormDB) GetBookByID(id int) (*Book, error) {
	var book Book
	if err := gdb.db.Preload("Author").First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil

}
