package data

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strings"
	"time"
	"user/internal/biz"
)

// User 是用户数据表结构体
type User struct {
	ID          int64      `gorm:"primarykey"`
	Mobile      string     `gorm:"index:idx_mobile;unique;type:varchar(11) comment '手机号码，用户唯一标识';not null"`
	Password    string     `gorm:"type:varchar(100);not null "` // 用户密码的保存需要注意是否加密
	NickName    string     `gorm:"type:varchar(25) comment '用户昵称'"`
	Birthday    *time.Time `gorm:"type:datetime comment '出生日期'"`
	Gender      string     `gorm:"column:gender;default:male;type:varchar(16) comment 'female:女,male:男'"`
	Role        int        `gorm:"column:role;default:1;type:int comment '1:普通用户，2:管理员'"`
	CreatedAt   time.Time  `gorm:"column:add_time"`
	UpdatedAt   time.Time  `gorm:"column:update_time"`
	DeletedAt   gorm.DeletedAt
	IsDeletedAt bool
}
type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepo) ListUser(ctx context.Context, pageNum, pageSize int64) (users []*biz.User, total int, err error) {
	var results []*User
	offset, limit := (pageNum-1)*pageSize, pageSize
	result := ur.data.db.
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&results)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	total = int(result.RowsAffected)
	for _, user := range results {
		users = append(users, toBizUser(user))
	}
	return users, total, nil
}

func (ur *userRepo) UserByMobile(ctx context.Context, mobile string) (user *biz.User, err error) {
	var u User
	result := ur.data.db.Where(&User{Mobile: mobile}).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	user = toBizUser(&u)
	return user, nil
}

func (ur *userRepo) UserByID(ctx context.Context, id int64) (user *biz.User, err error) {
	var u User
	result := ur.data.db.Where(&User{ID: id}).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	user = toBizUser(&u)
	return user, nil
}

func (ur *userRepo) UpdateUser(ctx context.Context, user *biz.User) (isSuccess bool, err error) {
	var u User
	result := ur.data.db.Where(&User{ID: user.ID}).First(&u)
	if result.RowsAffected == 0 {
		return false, status.Errorf(codes.NotFound, "用户不存在")
	}
	u.NickName = user.NickName
	u.Gender = user.Gender
	u.Birthday = user.Birthday

	err = ur.data.db.Save(&u).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ur *userRepo) CheckPassword(ctx context.Context, psd, encryptedPassword string) (check bool, err error) {
	options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
	passwordInfo := strings.Split(encryptedPassword, "$")
	check = password.Verify(psd, passwordInfo[2], passwordInfo[3], options)
	return check, nil
}

func (ur *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	var user User
	result := ur.data.db.Model(&User{}).Where(&User{Mobile: u.Mobile}).First(&user)
	if result.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Mobile = u.Mobile
	user.Password = encrypt(u.Password)
	user.NickName = u.NickName
	res := ur.data.db.Create(&user)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, "创建用户失败:%v", res.Error)
	}
	return &biz.User{
		ID:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     user.Role,
	}, nil
}

func encrypt(str string) string {
	options := &password.Options{
		SaltLen:      16,
		Iterations:   10000,
		KeyLen:       32,
		HashFunction: sha512.New}
	salt, encodedStr := password.Encode(str, options)
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedStr)
}

func toBizUser(user *User) *biz.User {
	return &biz.User{
		ID:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     user.Role,
		Birthday: user.Birthday,
	}
}
