package config

import (
	"path/filepath"
	"testing"
	"url_pinger/pkg/config"

	"github.com/stretchr/testify/require"
)


func TestLoadConfig(t *testing.T) {
	path := filepath.Join("test.env")
	cfg, err := config.LoadConfig(path)

	require.NoError(t, err)
    require.NotEmpty(t, cfg.DBConn)

}