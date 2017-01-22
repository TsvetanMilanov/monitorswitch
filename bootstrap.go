package main

func bootstrap() *Injector {
	var result = new(Injector)

	result.monitorsService = new(MonitorsService)

	return result
}
