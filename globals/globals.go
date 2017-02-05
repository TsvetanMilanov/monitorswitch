package globals

import (
	"github.com/TsvetanMilanov/monitorswitch/injector"
)

var injectorInstance *injector.Injector

// SetupInjector registers the injector instance.
func SetupInjector(bootstrap func() *injector.Injector) {
	injectorInstance = bootstrap()
}

// GetInjector returns the global injector instance.
func GetInjector() *injector.Injector {
	return injectorInstance
}
