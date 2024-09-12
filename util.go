package main

func sInSlice(s string, slice []string) bool {
	for _, t := range slice {
		if s == t {
			return true
		}
	}
	return false
}
