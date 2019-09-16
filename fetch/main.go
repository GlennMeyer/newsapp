package main

import ()

func main() {
	createTables()
	fetchSources()
	sleepFetch()
	for {
		select{}
	}
}
