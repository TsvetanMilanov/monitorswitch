package services

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/TsvetanMilanov/monitorswitch/models"
)

// MonitorsService type
type MonitorsService struct {
}

// GetAllMonitors returns information for all connected monitors
func (service *MonitorsService) GetAllMonitors() []*models.Monitor {
	var monitorsOutput = executeXrandrCommand()
	return parseMonitorOutput(monitorsOutput)
}

// SwitchMonitorOn switches monitor on and puts it to the left or the right of the referenceMonitor
func (service *MonitorsService) SwitchMonitorOn(monitor *models.Monitor, referenceMonitor *models.Monitor, side string) {
	executeXrandrCommand("--output", monitor.Name, "--auto", side, referenceMonitor.Name)
}

// SwitchMonitorOff switches the monitor off.
func (service *MonitorsService) SwitchMonitorOff(monitor *models.Monitor) {
	executeXrandrCommand("--output", monitor.Name, "--off")
}

// GetMonitor gets the primary monitor or the first not - primary monitor
func (service *MonitorsService) GetMonitor(monitors []*models.Monitor, shouldGetPrimaryMonitor bool) *models.Monitor {
	for _, monitor := range monitors {
		if shouldGetPrimaryMonitor {
			if monitor.IsPrimary {
				return monitor
			}
		} else if !monitor.IsPrimary {
			return monitor
		}
	}

	return monitors[len(monitors)-1]
}

func executeXrandrCommand(arg ...string) string {
	var output, err = exec.Command("xrandr", arg...).Output()

	if err != nil {
		panic(err)
	}

	return string(output[:len(output)])
}

func parseMonitorOutput(output string) []*models.Monitor {
	var regExp = regexp.MustCompile(`(.*) connected (primary)?`)
	var allLines = strings.Split(output, "\n")
	var result = []*models.Monitor{}

	for _, line := range allLines {
		var matches = regExp.FindStringSubmatch(line)

		if len(matches) != 0 {
			var monitor = new(models.Monitor)
			monitor.Name = matches[1]
			if len(matches) >= 3 && matches[2] == "primary" {
				monitor.IsPrimary = true
			}

			result = append(result, monitor)
		}
	}

	return result
}
