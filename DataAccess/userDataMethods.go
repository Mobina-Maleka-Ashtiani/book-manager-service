package DataAccess

func (gdb *GormDB) UsernameExistence(username string) bool {
	var count int64
	gdb.db.Model(&User{}).Where(&User{Username: username}).Count(&count)

	return count > 0
}
