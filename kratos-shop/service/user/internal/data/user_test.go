package data_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"user/internal/biz"
	"user/internal/data"
)

var _ = ginkgo.Describe("User", func() {
	var ur biz.UserRepo
	var user *biz.User
	ginkgo.BeforeEach(func() {
		// DB is defined in the data_suite_test
		ur = data.NewUserRepo(DB, nil)
		user = &biz.User{
			ID:       2,
			Mobile:   "13800138088",
			Password: "admin123456",
			NickName: "admina3d45",
			Role:     1,
			Birthday: 693629981,
		}
	})
	ginkgo.It("CreateUser", func() {
		u, err := ur.CreateUser(ctx, user)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(u.Mobile).Should(gomega.Equal("13800138088"))
	})
})
