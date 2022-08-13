package service

import (
	v1 "clock-in/api/User/interface/v1"
	recordv1 "clock-in/api/record/v1"
	userv1 "clock-in/api/user/v1"
	"clock-in/app/User/interface/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// UserService
type UserService struct {
	v1.UnimplementedUserInterfaceServer
	au  *biz.AuthUsecase
	cu  *biz.UserUsecase
	log *log.Helper
}

var _ v1.UserInterfaceServer = (*UserService)(nil)

// NewUserService new a greeter service.
func NewUserService(
	au *biz.AuthUsecase,
	cu *biz.UserUsecase,
	logger log.Logger,
) *UserService {
	return &UserService{
		au:  au,
		cu:  cu,
		log: log.NewHelper(logger),
	}
}

func (svc *UserService) Register(ctx context.Context, in *v1.RegisterRequest) (reply *v1.RegisterReply, err error) {
	user, err := svc.au.Register(ctx, &biz.User{
		Name:     in.Name,
		Password: in.Password,
		Phone:    in.Phone,
	})
	if err != nil {
		switch {
		case userv1.IsUserExist(err):
			err = v1.ErrorUserExistedError(err.Error())
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	} else {
		reply = &v1.RegisterReply{
			Id:    user.Id,
			Name:  user.Name,
			Phone: user.Phone,
		}
	}
	return
}

func (svc *UserService) Login(ctx context.Context, in *v1.LoginRequest) (reply *v1.LoginReply, err error) {
	token, err := svc.au.Login(ctx, in.Username, in.Password)
	if err != nil {
		switch {
		case userv1.IsUserNotFound(err):
			err = v1.ErrorUserNotFound(err.Error())
		case err == biz.ErrIncorrectPassword:
			err = v1.ErrorPasswordIncorrect(err.Error())
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	} else {
		reply = &v1.LoginReply{
			Token: token,
		}
	}
	return
}

func (svc *UserService) UserOnWork(ctx context.Context, in *v1.UserOnWorkRequest) (*v1.UserOnWorkReply, error) {
	err := svc.cu.UserOnWork(ctx)
	if err != nil {
		switch {
		case recordv1.IsRecordExisted(err):
			err = v1.ErrorRecordExisted(err.Error())
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	}
	return &v1.UserOnWorkReply{}, err
}

func (svc *UserService) UserOffWork(ctx context.Context, in *v1.UserOffWorkRequest) (*v1.UserOffWorkReply, error) {
	err := svc.cu.UserOffWork(ctx)
	if err != nil {
		switch {
		case recordv1.IsRecordExisted(err):
			err = v1.ErrorRecordExisted(err.Error())
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	}
	return &v1.UserOffWorkReply{}, err
}
