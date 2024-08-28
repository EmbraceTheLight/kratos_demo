package service

import "github.com/google/wire"

// ProviderSet is service providers.
// var ProviderSet = wire.NewSet(NewGreeterService)
var ProviderSet = wire.NewSet(NewStudentService)
var a int = 1
