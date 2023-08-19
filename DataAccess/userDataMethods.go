package DataAccess

func (gdb *GormDB) UsernameExistence(username string) bool {
	var count int64
	gdb.db.Model(&User{}).Where("username = ?", username).Count(&count)

	return count > 0
}

func (gdb *GormDB) EmailExistence(email string) bool {
	var count int64
	gdb.db.Model(&User{}).Where("email = ?", email).Count(&count)

	return count > 0
}

func (gdb *GormDB) PhoneNumberExistence(phoneNumber string) bool {
	var count int64
	gdb.db.Model(&User{}).Where("phone_number = ?", phoneNumber).Count(&count)

	return count > 0
}

func (gdb *GormDB) AddUserToDatabase(user *User) error {
	return gdb.db.Create(user).Error
}

func (gdb *GormDB) FindUserByUsername(username string) (*User, error) {
	var user User
	err := gdb.db.Where("username = ?", username).Preload("Books").First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (gdb *GormDB) AddBookToUser(user User, book Book) error {
	user.Books = append(user.Books, &book)
	if err := gdb.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
