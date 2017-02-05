package injector

import (
	"github.com/TsvetanMilanov/monitorswitch/services"
)

// Injector type
type Injector struct {
	MonitorsService *services.MonitorsService
}
