package main

import (
	"os/exec"
	"regexp"
	"strings"
)

// MonitorsService type
type MonitorsService struct {
}

// GetAllMonitors returns information for all connected monitors
func (service *MonitorsService) GetAllMonitors() []*Monitor {
	var monitorsOutput = executeXrandrCommand()
	return parseMonitorOutput(monitorsOutput)
}

// SwitchMonitorOn switches monitor on and puts it to the left or the right of the referenceMonitor
func (service *MonitorsService) SwitchMonitorOn(monitor *Monitor, referenceMonitor *Monitor, side string) {
	executeXrandrCommand("--output", monitor.name, "--auto", side, referenceMonitor.name)
}

// SwitchMonitorOff switches the monitor off.
func (service *MonitorsService) SwitchMonitorOff(monitor *Monitor) {
	executeXrandrCommand("--output", monitor.name, "--off")
}

// GetMonitor gets the primary monitor or the first not - primary monitor
func (service *MonitorsService) GetMonitor(monitors []*Monitor, shouldGetPrimaryMonitor bool) *Monitor {
	// TODO: refactor this to work with monitor name
	for _, monitor := range monitors {
		if shouldGetPrimaryMonitor {
			if monitor.isPrimary {
				return monitor
			}
		} else if !monitor.isPrimary {
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

func parseMonitorOutput(output string) []*Monitor {
	var regExp = regexp.MustCompile(`(.*) connected (primary)?`)
	var allLines = strings.Split(output, "\n")
	var result = []*Monitor{}

	for _, line := range allLines {
		var matches = regExp.FindStringSubmatch(line)

		if len(matches) != 0 {
			var monitor = new(Monitor)
			monitor.name = matches[1]
			if len(matches) >= 3 && matches[2] == "primary" {
				monitor.isPrimary = true
			}

			result = append(result, monitor)
		}
	}

	return result
}
