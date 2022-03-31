package utils

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/zbitech/common/pkg/model/config"
)

func Test_ReadConfig(t *testing.T) {
	app_config := config.AppConfig{}
	app_config_path := "tests/files/config.yaml"
	if err := ReadConfig(app_config_path, nil, &app_config); err != nil {
		t.Fatalf("Unable read config at %s", app_config_path)
	}

	name := strings.Split(filepath.Base(app_config_path), ".")[0]
	t.Logf("App Config (%s) -> %s", name, MarshalIndentObject(app_config))

	rsc_config_path := "tests/files/project.yaml"
	rsc_config := config.ResourceConfig{}
	if err := ReadConfig(rsc_config_path, nil, &rsc_config); err != nil {
		t.Fatalf("Unable to read config at %s", rsc_config_path)
	}

	t.Logf("Resource config - %s", MarshalIndentObject(rsc_config))
}
