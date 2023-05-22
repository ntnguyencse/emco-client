package main

import (
	"fmt"
)

const (
	config = "/home/ubuntu/emco-client/.emco.yaml"
)

func main() {
	fmt.Println("Beginning of Cmd")
	fmt.Println("Execute the cmd")
	cfgFile = config
	args := []string{"apply", "-f", "/home/ubuntu/emco-client/test/register-cluster.yml"}
	rootCmd.SetArgs(args)
	rootCmd.DebugFlags()
	Execute()
}
