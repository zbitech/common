package vars

import (
	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/common/pkg/model/config"
	"github.com/zbitech/common/pkg/utils"
)

var (
	MONGODB_URL          = utils.GetEnv("MONGODB_URL", "mongodb://root:wzZxK2YCCn@db.zbitech.local:27017")
	MONGODB_NAME         = utils.GetEnv("MONGODB_NAME", "zbiRepo")
	ASSET_PATH_DIRECTORY = utils.GetEnv("ASSET_PATH_DIRECTORY", "tests/files/etc/zbi")
	HOURS_IN_YEAR        = 8760

	KlientFactory        interfaces.KlientFactoryIF
	RepositoryFactory    interfaces.RepositoryFactoryIF
	ManagerFactory       interfaces.ResourceManagerFactoryIF
	AuthorizationFactory interfaces.AuthorizationFactoryIF

	ResourceConfig config.ResourceConfig
	AppConfig      config.AppConfig
)
