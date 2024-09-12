// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: realworld/v1/realworld.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RealWorld_Login_FullMethodName             = "/realworld.v1.RealWorld/Login"
	RealWorld_Register_FullMethodName          = "/realworld.v1.RealWorld/Register"
	RealWorld_GetCurrentUser_FullMethodName    = "/realworld.v1.RealWorld/GetCurrentUser"
	RealWorld_UpdateUser_FullMethodName        = "/realworld.v1.RealWorld/UpdateUser"
	RealWorld_GetProfile_FullMethodName        = "/realworld.v1.RealWorld/GetProfile"
	RealWorld_FollowUser_FullMethodName        = "/realworld.v1.RealWorld/FollowUser"
	RealWorld_UnFollowUser_FullMethodName      = "/realworld.v1.RealWorld/UnFollowUser"
	RealWorld_ListArticles_FullMethodName      = "/realworld.v1.RealWorld/ListArticles"
	RealWorld_FeedArticles_FullMethodName      = "/realworld.v1.RealWorld/FeedArticles"
	RealWorld_GetArticle_FullMethodName        = "/realworld.v1.RealWorld/GetArticle"
	RealWorld_CreateArticle_FullMethodName     = "/realworld.v1.RealWorld/CreateArticle"
	RealWorld_UpdateArticle_FullMethodName     = "/realworld.v1.RealWorld/UpdateArticle"
	RealWorld_DeleteArticle_FullMethodName     = "/realworld.v1.RealWorld/DeleteArticle"
	RealWorld_AddComment_FullMethodName        = "/realworld.v1.RealWorld/AddComment"
	RealWorld_GetComments_FullMethodName       = "/realworld.v1.RealWorld/GetComments"
	RealWorld_DeleteComments_FullMethodName    = "/realworld.v1.RealWorld/DeleteComments"
	RealWorld_FavoriteArticle_FullMethodName   = "/realworld.v1.RealWorld/FavoriteArticle"
	RealWorld_UnFavoriteArticle_FullMethodName = "/realworld.v1.RealWorld/UnFavoriteArticle"
	RealWorld_GetTags_FullMethodName           = "/realworld.v1.RealWorld/GetTags"
)

// RealWorldClient is the client API for RealWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The greeting service definition.
type RealWorldClient interface {
	// Sends a greeting
	//
	//	rpc SayHello (HelloRequest) returns (HelloReply) {
	//	  option (google.api.http) = {
	//	    get: "/helloworld/{name}"
	//	  };
	//	}
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetCurrentUser(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*UserReply, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	UnFollowUser(ctx context.Context, in *UnFollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error)
	FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleReply, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error)
	DeleteComments(ctx context.Context, in *DeleteCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error)
	FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	UnFavoriteArticle(ctx context.Context, in *UnFavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	GetTags(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TasLIstReply, error)
}

type realWorldClient struct {
	cc grpc.ClientConnInterface
}

func NewRealWorldClient(cc grpc.ClientConnInterface) RealWorldClient {
	return &realWorldClient{cc}
}

func (c *realWorldClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserReply)
	err := c.cc.Invoke(ctx, RealWorld_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserReply)
	err := c.cc.Invoke(ctx, RealWorld_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetCurrentUser(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*UserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserReply)
	err := c.cc.Invoke(ctx, RealWorld_GetCurrentUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserReply)
	err := c.cc.Invoke(ctx, RealWorld_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, RealWorld_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, RealWorld_FollowUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UnFollowUser(ctx context.Context, in *UnFollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, RealWorld_UnFollowUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MultipleArticlesReply)
	err := c.cc.Invoke(ctx, RealWorld_ListArticles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MultipleArticlesReply)
	err := c.cc.Invoke(ctx, RealWorld_FeedArticles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_GetArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_CreateArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_UpdateArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_DeleteArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_AddComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MultipleCommentsReply)
	err := c.cc.Invoke(ctx, RealWorld_GetComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) DeleteComments(ctx context.Context, in *DeleteCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MultipleCommentsReply)
	err := c.cc.Invoke(ctx, RealWorld_DeleteComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_FavoriteArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UnFavoriteArticle(ctx context.Context, in *UnFavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_UnFavoriteArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetTags(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TasLIstReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TasLIstReply)
	err := c.cc.Invoke(ctx, RealWorld_GetTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealWorldServer is the server API for RealWorld service.
// All implementations must embed UnimplementedRealWorldServer
// for forward compatibility.
//
// The greeting service definition.
type RealWorldServer interface {
	// Sends a greeting
	//
	//	rpc SayHello (HelloRequest) returns (HelloReply) {
	//	  option (google.api.http) = {
	//	    get: "/helloworld/{name}"
	//	  };
	//	}
	Login(context.Context, *LoginRequest) (*UserReply, error)
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	GetCurrentUser(context.Context, *EmptyRequest) (*UserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error)
	FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error)
	UnFollowUser(context.Context, *UnFollowUserRequest) (*ProfileReply, error)
	ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error)
	FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*SingleArticleReply, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*SingleArticleReply, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleReply, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleReply, error)
	AddComment(context.Context, *AddCommentRequest) (*SingleArticleReply, error)
	GetComments(context.Context, *GetCommentsRequest) (*MultipleCommentsReply, error)
	DeleteComments(context.Context, *DeleteCommentsRequest) (*MultipleCommentsReply, error)
	FavoriteArticle(context.Context, *FavoriteArticleRequest) (*SingleArticleReply, error)
	UnFavoriteArticle(context.Context, *UnFavoriteArticleRequest) (*SingleArticleReply, error)
	GetTags(context.Context, *EmptyRequest) (*TasLIstReply, error)
	mustEmbedUnimplementedRealWorldServer()
}

// UnimplementedRealWorldServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRealWorldServer struct{}

func (UnimplementedRealWorldServer) Login(context.Context, *LoginRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRealWorldServer) Register(context.Context, *RegisterRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRealWorldServer) GetCurrentUser(context.Context, *EmptyRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (UnimplementedRealWorldServer) UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedRealWorldServer) GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedRealWorldServer) FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedRealWorldServer) UnFollowUser(context.Context, *UnFollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollowUser not implemented")
}
func (UnimplementedRealWorldServer) ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (UnimplementedRealWorldServer) FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedArticles not implemented")
}
func (UnimplementedRealWorldServer) GetArticle(context.Context, *GetArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedRealWorldServer) CreateArticle(context.Context, *CreateArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedRealWorldServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedRealWorldServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedRealWorldServer) AddComment(context.Context, *AddCommentRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedRealWorldServer) GetComments(context.Context, *GetCommentsRequest) (*MultipleCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedRealWorldServer) DeleteComments(context.Context, *DeleteCommentsRequest) (*MultipleCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComments not implemented")
}
func (UnimplementedRealWorldServer) FavoriteArticle(context.Context, *FavoriteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteArticle not implemented")
}
func (UnimplementedRealWorldServer) UnFavoriteArticle(context.Context, *UnFavoriteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFavoriteArticle not implemented")
}
func (UnimplementedRealWorldServer) GetTags(context.Context, *EmptyRequest) (*TasLIstReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
func (UnimplementedRealWorldServer) mustEmbedUnimplementedRealWorldServer() {}
func (UnimplementedRealWorldServer) testEmbeddedByValue()                   {}

// UnsafeRealWorldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealWorldServer will
// result in compilation errors.
type UnsafeRealWorldServer interface {
	mustEmbedUnimplementedRealWorldServer()
}

func RegisterRealWorldServer(s grpc.ServiceRegistrar, srv RealWorldServer) {
	// If the following call pancis, it indicates UnimplementedRealWorldServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RealWorld_ServiceDesc, srv)
}

func _RealWorld_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_GetCurrentUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetCurrentUser(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_FollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).FollowUser(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UnFollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnFollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UnFollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_UnFollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UnFollowUser(ctx, req.(*UnFollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_ListArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).ListArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_ListArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).ListArticles(ctx, req.(*ListArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_FeedArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).FeedArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_FeedArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).FeedArticles(ctx, req.(*FeedArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_GetArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_CreateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_UpdateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_DeleteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_GetComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_DeleteComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).DeleteComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_DeleteComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).DeleteComments(ctx, req.(*DeleteCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_FavoriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).FavoriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_FavoriteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).FavoriteArticle(ctx, req.(*FavoriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UnFavoriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnFavoriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UnFavoriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_UnFavoriteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UnFavoriteArticle(ctx, req.(*UnFavoriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_GetTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetTags(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RealWorld_ServiceDesc is the grpc.ServiceDesc for RealWorld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RealWorld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "realworld.v1.RealWorld",
	HandlerType: (*RealWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _RealWorld_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _RealWorld_Register_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _RealWorld_GetCurrentUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _RealWorld_UpdateUser_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _RealWorld_GetProfile_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _RealWorld_FollowUser_Handler,
		},
		{
			MethodName: "UnFollowUser",
			Handler:    _RealWorld_UnFollowUser_Handler,
		},
		{
			MethodName: "ListArticles",
			Handler:    _RealWorld_ListArticles_Handler,
		},
		{
			MethodName: "FeedArticles",
			Handler:    _RealWorld_FeedArticles_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _RealWorld_GetArticle_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _RealWorld_CreateArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _RealWorld_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _RealWorld_DeleteArticle_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _RealWorld_AddComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _RealWorld_GetComments_Handler,
		},
		{
			MethodName: "DeleteComments",
			Handler:    _RealWorld_DeleteComments_Handler,
		},
		{
			MethodName: "FavoriteArticle",
			Handler:    _RealWorld_FavoriteArticle_Handler,
		},
		{
			MethodName: "UnFavoriteArticle",
			Handler:    _RealWorld_UnFavoriteArticle_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _RealWorld_GetTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "realworld/v1/realworld.proto",
}
