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

func SetRepositoryFactory(ctx context.Context, f interfaces.RepositoryFactoryIF) {
	ctx = rctx.BuildContext(ctx, rctx.Context(rctx.Component, "SetRepositoryFactory"), rctx.Context(rctx.StartTime, time.Now()))
	defer logger.LogComponentTime(ctx)

	logger.Infof(ctx, "Initializing Repository Factory")

	err := f.Init(ctx)
	if err != nil {
		panic(err)
	}

	logger.Debugf(ctx, "Project Repo - %v", f.GetProjectRepository())
	logger.Debugf(ctx, "Admin Repo - %v", f.GetAdminRepository())

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

	logger.Debugf(ctx, "Ingress Resource Manager - %v", f.GetIngressResourceManager(ctx))
	logger.Debugf(ctx, "Project Resource Manager - %v", f.GetProjectDataManager(ctx))

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

	logger.Debugf(ctx, "ZBI Client - %v", f.GetZBIClient())
	//	logger.Debugf(ctx, "Kubernetes Client - %v", f.GetKubernesClient())

	vars.KlientFactory = f
}

func SetIAMFactory(ctx context.Context, f interfaces.IAMFactoryIF) {
	ctx = rctx.BuildContext(ctx, rctx.Context(rctx.Component, "SetIAMFactory"), rctx.Context(rctx.StartTime, time.Now()))
	defer logger.LogComponentTime(ctx)

	logger.Infof(ctx, "Initializing IAM Factory")
	err := f.Init(ctx)
	if err != nil {
		panic(err)
	}

	logger.Debugf(ctx, "IAM - %v", f.GetIAMService())
	logger.Debugf(ctx, "JWT - %v", f.GetJwtServer())
	vars.IAMFactory = f
}

func SetAccessAuthorizerFactory(ctx context.Context, f interfaces.AccessAuthorizerFactoryIF) {
	ctx = rctx.BuildContext(ctx, rctx.Context(rctx.Component, "SetAccessAuthorizerFactory"), rctx.Context(rctx.StartTime, time.Now()))
	defer logger.LogComponentTime(ctx)

	logger.Infof(ctx, "Initializing Access Authorizer Factory")
	err := f.Init(ctx)
	if err != nil {
		panic(err)
	}

	logger.Debugf(ctx, "Authorizer Service - %v", f.GetAccessAuthorizer())
	vars.AuthorizerFactory = f
}

func InitConfig(ctx context.Context) {

	configPath := fmt.Sprintf("%s/config.yaml", vars.ASSET_PATH_DIRECTORY)
	logger.Infof(ctx, "Initializing application config from %s", configPath)

	if err := utils.ReadConfig(configPath, nil, &vars.AppConfig); err != nil {
		panic(err)
	}

	vars.AppConfig.Repository.Database.Factory = vars.DATABASE_FACTORY
	vars.AppConfig.Repository.Database.Url = vars.DATABASE_URL
	vars.AppConfig.Kubernetes.InCluster = vars.K8S_INCLUSTER
	vars.AppConfig.Kubernetes.KubeConfig = vars.KUBECONFIG

	logger.Infof(ctx, "AppConfig - %s", utils.MarshalObject(vars.AppConfig))
}

func InitProjectResourceConfig(ctx context.Context) {
	resourcePath := fmt.Sprintf("%s/project.yaml", vars.ASSET_PATH_DIRECTORY)
	logger.Infof(ctx, "Initializing resource config from %s", resourcePath)

	if err := utils.ReadConfig(resourcePath, nil, &vars.ResourceConfig); err != nil {
		panic(err)
	}
	logger.Infof(ctx, "ResourceConfig - %s", utils.MarshalObject(vars.ResourceConfig))
}
