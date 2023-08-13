package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

// User 直接对应于数据库表结构
type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	// 创建时间与更新时间，毫秒数
	Ctime int64
	Utime int64
	// edit部分字段,昵称，生日，个人简介
	// 昵称不容许为空，生日和个人简介允许不设置
	Nickname        string `gorm:"size:60,not null"`
	Birthday        string
	PersonalProfile string `gorm:"size:600"`
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	// 默认的昵称为ctime
	u.Nickname = strconv.FormatInt(now, 10)
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			//邮箱冲突
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

func (dao *UserDAO) Update(ctx context.Context, u User) error {
	var user User
	err := dao.db.WithContext(ctx).Where("id=?", u.Id).First(&user).Error
	if err != nil {
		return err
	}
	now := time.Now().UnixMilli()
	user.Utime = now
	user.Nickname = u.Nickname
	user.Birthday = u.Birthday
	user.PersonalProfile = u.PersonalProfile
	dao.db.Save(&user)
	return err
}

func (dao *UserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("id=?", id).First(&u).Error
	return u, err
}
