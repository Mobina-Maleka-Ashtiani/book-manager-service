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
