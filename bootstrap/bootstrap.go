package bootstrap

import (
	"github.com/TsvetanMilanov/monitorswitch/injector"
	"github.com/TsvetanMilanov/monitorswitch/services"
)

// RunBootstrap rejisters all dependencies in the injector.
func RunBootstrap() *injector.Injector {
	var result = new(injector.Injector)

	result.MonitorsService = new(services.MonitorsService)

	return result
}
