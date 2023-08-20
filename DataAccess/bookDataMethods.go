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

func (gdb *GormDB) UpdateBook(book Book, updateName string, updateCategory string) error {
	book.Name = updateName
	book.Category = updateCategory
	if err := gdb.db.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func (gdb *GormDB) DeleteBook(book Book) error {
	if err := gdb.db.Delete(&book).Error; err != nil {
		return err
	}
	return nil
}
