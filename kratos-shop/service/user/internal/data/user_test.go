package data_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
	"user/internal/biz"
	"user/internal/data"
	"user/internal/testdata"
)

var _ = Describe("User", func() {
	var ur biz.UserRepo
	var user *biz.User
	BeforeEach(func() {
		// DB is defined in the data_suite_test
		ur = data.NewUserRepo(DB, nil)
		user = testdata.NewUser()
	})
	// Test cases
	It("CreateUser", func() {
		u, err := ur.CreateUser(ctx, user)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(u.Mobile).Should(Equal("13509876789")) //手机号应为组装数据时的，即user.Mobile
	})

	It("ListUser", func() {
		user, total, err := ur.ListUser(ctx, 1, 10)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(user).ShouldNot(BeEmpty())
		Ω(total).Should(Equal(1)) //上面只创建了一条数据，所以total应该为1
		Ω(len(user)).Should(Equal(1))
		Ω(user[0].Mobile).Should(Equal("13509876789"))
	})

	It("UpdateUser", func() {
		birthDay := time.Unix(int64(693646426), 0)
		user.NickName = "new_name"
		user.Birthday = &birthDay
		user.Gender = "female"
		ok, err := ur.UpdateUser(ctx, user)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(ok).Should(BeTrue())
	})

	It("CheckPassword", func() {
		p1 := "admin"
		encryptedPassword := "$pbkdf2-sha512$5p7doUNIS9I5mvhA$b18171ff58b04c02ed70ea4f39bda036029c107294bce83301a02fb53a1bcae0"
		ok, err := ur.CheckPassword(ctx, p1, encryptedPassword)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(ok).Should(BeTrue())

		encryptedPassword = "$should $be $false$"
		ok, err = ur.CheckPassword(ctx, p1, encryptedPassword)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(ok).Should(BeFalse())
	})
})
