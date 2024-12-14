package biz_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"testing"
)

func TestBiz(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "biz user test")
}

var (
	ctl     *gomock.Controller
	cleaner func()
	ctx     context.Context
)

var _ = ginkgo.BeforeEach(func() {
	ctl = gomock.NewController(ginkgo.GinkgoT())
	cleaner = ctl.Finish
	ctx = context.Background()
})
var _ = ginkgo.AfterEach(func() {
	cleaner()
})
