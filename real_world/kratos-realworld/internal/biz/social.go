package biz

import "github.com/go-kratos/kratos/v2/log"

// SocialUsecase is a Comment usecase.
type SocialUsecase struct {
	cmtRepo CommentRepo
	atcRepo ArticleRepo
	tagRepo TagRepo
	log     *log.Helper
}

// NewSocialUsecase new a Social usecase.
func NewSocialUsecase(
	cmtRepo CommentRepo,
	atcRepo ArticleRepo,
	tagRepo TagRepo,
	logger log.Logger) *SocialUsecase {

	return &SocialUsecase{
		cmtRepo: cmtRepo,
		atcRepo: atcRepo,
		tagRepo: tagRepo,
		log:     log.NewHelper(logger),
	}
}
