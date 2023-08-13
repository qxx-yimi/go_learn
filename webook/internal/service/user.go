package service

import (
	"context"
	"errors"
	"go_learn/webook/internal/domain"
	"go_learn/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail    = repository.ErrUserDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("邮箱或密码错误")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Signup(c context.Context, u domain.User) error {
	//bcrypt是一个号称最安全的加密算法
	//优点:
	//不需要你自己去生成盐值
	//不需要额外存储盐值
	//可以通过控制cost来控制加密性能
	//同样的文本，加密后的结果不同
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(c, u)
}

func (svc *UserService) Login(c context.Context, email, password string) (domain.User, error) {
	//先找用户
	u, err := svc.repo.FindByEmail(c, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	//比较密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, err
}

func (svc *UserService) Edit(c context.Context, u domain.User) error {
	return svc.repo.Update(c, u)
}

func (svc *UserService) Profile(c context.Context, id int64) (domain.User, error) {
	return svc.repo.FindById(c, id)
}
