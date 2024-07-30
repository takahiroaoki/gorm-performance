package main

import (
	"fmt"
	"log"

	"github.com/takahiroaoki/gorm-performance/app/cmd"
)

func main() {
	bundleCmd := cmd.NewBundleCmd()
	if err := bundleCmd.Execute(); err != nil {
		log.Fatalf(fmt.Sprintf("Failed to execute the command. Error: %v", err))
	}
}
