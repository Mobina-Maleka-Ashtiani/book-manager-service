package DataAccess

func (gdb *GormDB) UsernameExistence(username string) bool {
	var count int64
	gdb.db.Model(&User{}).Where(&User{Username: username}).Count(&count)

	return count > 0
}

func (gdb *GormDB) EmailExistence(email string) bool {
	var count int64
	gdb.db.Model(&User{}).Where(&User{Email: email}).Count(&count)

	return count > 0
}

func (gdb *GormDB) PhoneNumberExistence(phoneNumber string) bool {
	var count int64
	gdb.db.Model(&User{}).Where(&User{PhoneNumber: phoneNumber}).Count(&count)

	return count > 0
}

func (gdb *GormDB) AddUserToDatabase(user *User) error {
	return gdb.db.Create(user).Error
}

func (gdb *GormDB) FindUserByUsername(username string) (*User, error) {
	var user User
	err := gdb.db.Where(&User{Username: username}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
