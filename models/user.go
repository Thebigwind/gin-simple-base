package models

import (
	tool "gin-simple-base/pkg/tools/time"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model2
	Name     string `gorm:"name" json:"name"`
	Phone    string `gorm:"phone" json:"phone"`
	NickName string `gorm:"nick_name" json:"nickName"`
	Sex      string `gorm:"sex" json:"sex"`
	Age      int    `gorm:"age" json:"age"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {

	scope.SetColumn("CreatedOn", tool.GetCurrentTime())
	scope.SetColumn("ModifiedOn", tool.GetCurrentTime())
	return nil
}

func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", tool.GetCurrentTime())

	return nil
}

////////////////////////////////////////////////////////////

// Add persists patch to database
func (u *User) AddUser() error {

	// 用选定字段的来创建
	// db.Select("Name", "Age").Create(u)
	// INSERT INTO `users` (`name`,`age`) VALUES ("jinzhu", 18)

	// 创建时排除选定字段
	// db.Omit( "Age").Create(u)
	// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")

	if err := db.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) GetUser() (user User, err error) {

	curErr := db.Where("id = ?", u.ID).First(&user).Error
	if curErr != nil && curErr != gorm.ErrRecordNotFound {
		err = curErr
	}

	return
}

func (u *User) GetUserListByCondition(page, pageSize int) ([]User, int64, error) {
	var (
		userList []User
		count    int64
	)
	//db = db.Debug().Where("name = ?", u.Name).Where("nick_name = ?", u.NickName)
	db = db.Debug()
	if u.Name != "" {
		db = db.Where("name = ?", u.Name)
	}
	if u.NickName != "" {
		db = db.Where("nick_name = ?", u.NickName)
	}
	//gorm 方式
	if u.Age != 0 {
		db = db.Where("age = ?", u.Age)
	}
	if u.Phone != "" {
		db = db.Where("phone = ?", u.Phone)
	}
	if u.Sex != "" {
		db = db.Where("sex = ?", u.Sex)
	}

	db.Model(&User{}).Count(&count)

	db = db.Order("id DESC")
	db = db.Limit(pageSize).Offset((page - 1) * pageSize)
	if err := db.Find(&userList).Error; err != nil {
		return nil, 0, err
	}

	return userList, count, nil
}

func (u *User) UpdateUser() error {

	// 使用 `map` 更新多个属性，只会更新那些被更改了的字段 ; 使用 `struct` 更新多个属性，只会更新那些被修改了的和非空的字段
	if err := db.Model(&User{}).
		//Where("id = ? ", v.ID).
		Where("name = ? AND nick_name = ? AND phone = ?", u.Name, u.NickName, u.Phone).
		Updates(map[string]interface{}{
			"age": gorm.Expr("age + ?", 1),
			"sex": u.Sex,
		}).Error; err != nil {
		return err
	}

	return nil
}

// 更新User结构体中所有的属性
func (u *User) UpdateUserAllAttr() error {

	if err := db.Model(&User{}).Where("id = ?", u.ID).Updates(u).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) ExistUserByName(name string) bool {
	var user User
	db.Select("id").Where("name = ?", name).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func (u *User) ExistUserByID(id int) bool {
	var user User
	db.Select("id").Where("id = ?", id).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func (u *User) GetUserTotal(maps interface{}) (count int) {
	db.Model(&User{}).Where(maps).Count(&count)

	return
}

///
// 更新User结构体中所有的属性
func UpdateUserScript(u User) error {

	if err := db.Model(&User{}).Where("id = ?", u.ID).Updates(u).Error; err != nil {
		return err
	}

	return nil
}

func CleanUserScript(p map[string]interface{}) error {

	return nil
}
