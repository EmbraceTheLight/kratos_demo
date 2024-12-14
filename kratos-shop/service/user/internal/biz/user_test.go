package biz_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"user/internal/biz"
	"user/internal/mocks/mrepo"
)

var _ = Describe("UserUsecase", func() {
	var userUsecase *biz.UserUsecase
	var mUserRepo *mrepo.MockUserRepo

	BeforeEach(func() {
		mUserRepo = mrepo.NewMockUserRepo(ctl)
		userUsecase = biz.NewUserUsecase(mUserRepo, nil)
	})
	It("Create", func() {
		info := &biz.User{
			ID:       1,
			Mobile:   "13800138088",
			Password: "admin123456",
			NickName: "admina3d45",
			Role:     1,
			Birthday: 693629981,
		}
		mUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(info, nil)
		l, err := userUsecase.Create(ctx, info)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l.ID).To(Equal(int64(1)))
		Ω(l.Mobile).To(Equal("13800138088"))
	})
})
