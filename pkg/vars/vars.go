package vars

import (
	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/common/pkg/model/config"
	"github.com/zbitech/common/pkg/utils"
	"strconv"
)

var (
	DATABASE_FACTORY         = utils.GetEnv("DATABASE_FACTORY", "mongo")
	DATABASE_URL             = utils.GetEnv("DATABASE_URL", "mongodb://root:wzZxK2YCCn@db.zbitech.local:27017")
	DATABASE_NAME            = utils.GetEnv("DATABASE_NAME", "zbiRepo")
	ASSET_PATH_DIRECTORY     = utils.GetEnv("ASSET_PATH_DIRECTORY", "tests/files/etc/zbi")
	KUBECONFIG               = utils.GetEnv("KUBECONFIG", "cfg/kubeconfig")
	USE_KUBERNETES_CONFIG, _ = strconv.ParseBool(utils.GetEnv("USEKUBERNETESCONFIG", "false"))
	K8S_INCLUSTER            = true
	EXPIRATION_HOURS, _      = strconv.Atoi(utils.GetEnv("EXPIRATION_HOURS", "8760"))
	ADMIN_USER               = utils.GetEnv("ZBI_ADMIN_USER", "admin")
	ADMIN_EMAIL              = utils.GetEnv("ZBI_ADMIN_EMAIL", "admin@alphegasolutions.com")
	ADMIN_PASSWORD           = utils.GetEnv("ZBI_ADMIN_PASSWORD", "password")

	HOURS_IN_YEAR = 8760

	KlientFactory     interfaces.KlientFactoryIF
	RepositoryFactory interfaces.RepositoryFactoryIF
	ManagerFactory    interfaces.ResourceManagerFactoryIF
	IAMFactory        interfaces.IAMFactoryIF
	AuthorizerFactory interfaces.AccessAuthorizerFactoryIF

	ResourceConfig config.ResourceConfig
	AppConfig      config.AppConfig
)
