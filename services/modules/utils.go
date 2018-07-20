package modules

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

func moduleRegistered(a string) bool {
	return stringInSlice(a, registered)
}

func moduleConnected(a string, list []string) bool {
	return stringInSlice(a, connected)
}
