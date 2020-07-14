// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/redhat-marketplace/redhat-marketplace-operator/pkg/controller"
	"github.com/redhat-marketplace/redhat-marketplace-operator/pkg/managers"
	"github.com/redhat-marketplace/redhat-marketplace-operator/pkg/utils/reconcileutils"
)

// Injectors from wire.go:

func InitializeMarketplaceController() *managers.ControllerMain {
	controllerFlagSet := controller.ProvideControllerFlagSet()
	defaultCommandRunnerProvider := reconcileutils.ProvideDefaultCommandRunnerProvider()
	marketplaceController := controller.ProvideMarketplaceController(defaultCommandRunnerProvider)
	meterbaseController := controller.ProvideMeterbaseController(defaultCommandRunnerProvider)
	meterDefinitionController := controller.ProvideMeterDefinitionController(defaultCommandRunnerProvider)
	razeeDeployController := controller.ProvideRazeeDeployController()
	olmSubscriptionController := controller.ProvideOlmSubscriptionController()
	controllerList := controller.ProvideControllerList(marketplaceController, meterbaseController, meterDefinitionController, razeeDeployController, olmSubscriptionController)
	opsSrcSchemeDefinition := controller.ProvideOpsSrcScheme()
	monitoringSchemeDefinition := controller.ProvideMonitoringScheme()
	olmV1SchemeDefinition := controller.ProvideOLMV1Scheme()
	olmV1Alpha1SchemeDefinition := controller.ProvideOLMV1Alpha1Scheme()
	localSchemes := controller.ProvideLocalSchemes(opsSrcSchemeDefinition, monitoringSchemeDefinition, olmV1SchemeDefinition, olmV1Alpha1SchemeDefinition)
	controllerMain := makeMarketplaceController(controllerFlagSet, controllerList, localSchemes)
	return controllerMain
}
