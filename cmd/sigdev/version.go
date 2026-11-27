package main

import "fmt"

func showVersion() {
	fmt.Printf(
		"app: %s\nversion: %s\ncommit: %s\ndate: %s\n",
		appname, version, commit, date,
	)
}
