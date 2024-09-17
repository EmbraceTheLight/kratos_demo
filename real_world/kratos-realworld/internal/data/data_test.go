package data

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/stretchr/testify/require"
	"kratos-realworld/internal/conf"
	"testing"
)

func TestNewDB(t *testing.T) {
	flagConf := "D:\\Go\\WorkSpace\\src\\Go_Project\\demo\\kratos\\real_world\\kratos-realworld\\configs\\config.yaml"
	c := config.New(
		config.WithSource(
			file.NewSource(flagConf),
		),
	)
	err := c.Load()
	require.NoError(t, err)
	require.NotNil(t, c)

	var bs = &conf.Bootstrap{}
	err = c.Scan(bs)
	require.NoError(t, err)
	require.NotNil(t, bs)
	NewMySQL(bs.Data)
}
