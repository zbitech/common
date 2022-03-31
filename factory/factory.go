package factory

import (
	"context"
	"fmt"
	"time"

	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/common/pkg/logger"
	"github.com/zbitech/common/pkg/rctx"
	"github.com/zbitech/common/pkg/utils"
	"github.com/zbitech/common/pkg/vars"
)

func SetRepositoryFactory(ctx context.Context, f interfaces.RepositoryFactoryIF, create_db, load_db bool) {
	ctx = rctx.BuildContext(ctx, rctx.Context(rctx.Component, "SetRepositoryFactory"), rctx.Context(rctx.StartTime, time.Now()))
	defer logger.LogComponentTime(ctx)

	logger.Infof(ctx, "Initializing Repository Factory")

	err := f.Init(ctx, create_db, load_db)
	if err != nil {
		panic(err)
	}

	vars.RepositoryFactory = f
}

func SetManagerFactory(ctx context.Context, f interfaces.ResourceManagerFactoryIF) {
	ctx = rctx.BuildContext(ctx, rctx.Context(rctx.Component, "SetManagerFactory"), rctx.Context(rctx.StartTime, time.Now()))
	defer logger.LogComponentTime(ctx)

	logger.Infof(ctx, "Initializing Manager Factory")

	err := f.Init(ctx)
	if err != nil {
		panic(err)
	}

	vars.ManagerFactory = f
}

func SetKubernetesFactory(ctx context.Context, f interfaces.KlientFactoryIF) {
	ctx = rctx.BuildContext(ctx, rctx.Context(rctx.Component, "SetKubernetesFactory"), rctx.Context(rctx.StartTime, time.Now()))
	defer logger.LogComponentTime(ctx)

	logger.Infof(ctx, "Initializing Kubernetes Factory")

	err := f.Init(ctx)
	if err != nil {
		panic(err)
	}

	vars.KlientFactory = f
}

func SetAuthorizerFactory(ctx context.Context, f interfaces.AuthorizationFactoryIF) {
	ctx = rctx.BuildContext(ctx, rctx.Context(rctx.Component, "SetAuthorizerFactory"), rctx.Context(rctx.StartTime, time.Now()))
	defer logger.LogComponentTime(ctx)

	logger.Infof(ctx, "Initializing Authorizer Factory")
	err := f.Init(ctx)
	if err != nil {
		panic(err)
	}

	vars.AuthorizationFactory = f
}

func InitConfig(ctx context.Context) {

	configPath := fmt.Sprintf("%s/config.yaml", vars.ASSET_PATH_DIRECTORY)
	logger.Infof(ctx, "Initializing application config from %s", configPath)

	//	password_keys := []string{"database.mongodb.url"} //TODO - set parameters
	if err := utils.ReadConfig(configPath, nil, &vars.AppConfig); err != nil {
		panic(err)
	}
	logger.Infof(ctx, "AppConfig - %s", utils.MarshalObject(vars.AppConfig))

	resourcePath := fmt.Sprintf("%s/project.yaml", vars.ASSET_PATH_DIRECTORY)
	logger.Infof(ctx, "Initializing resource config from %s", resourcePath)

	if err := utils.ReadConfig(resourcePath, nil, &vars.ResourceConfig); err != nil {
		panic(err)
	}
	logger.Infof(ctx, "ResourceConfig - %s", utils.MarshalObject(vars.ResourceConfig))
}
