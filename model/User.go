package model

import (
	"encoding/base64"
	"errors"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` //是否管理员
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password) //对密码进行加密，或者对数据库使用钩子函数
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200

}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users) //todo 优化RowsAffected
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS //200
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error //分页查询
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error fetching users: %v", err)
		return []User{} // 返回空的切片而不是 nil
	}
	return users
}

// EditUser 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// BeforeCreate 或者使用钩子函数进行加密
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
