package springcloud

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudApp -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1 -rewrite=true
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudAppAssociation -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/bindings/bind1 -rewrite=true
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudAPIPortal -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/apiPortals/apiPortal1 -rewrite=true
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudAPIPortalCustomDomain -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/apiPortals/apiPortal1/domains/domain1 -rewrite=true
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudBuildServiceBuilder -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/buildServices/buildService1/builders/builder1 -rewrite=true
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudBuildPackBinding -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/buildServices/buildService1/builders/builder1/buildPackBindings/buildPackBinding1 -rewrite=true
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudDeployment -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/deployments/deploy1 -rewrite=true
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudCertificate -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/certificates/cert1 -rewrite=true
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudCustomDomain -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/domains/domain.com -rewrite=true
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudConfigurationService -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/configurationServices/configurationService1 -rewrite=true
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudGateway -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1 -rewrite=true
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudGatewayCustomDomain -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/Spring/service1/gateways/gateway1/domains/domain1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudGatewayRouteConfig -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/Spring/service1/gateways/gateway1/routeConfigs/routeConfig1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudService -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/Spring/spring1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudServiceRegistry -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/Spring/spring1/serviceRegistries/serviceRegistry1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SpringCloudStorage -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/Spring/service1/storages/storage1
