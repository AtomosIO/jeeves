// Jeeves does not clean up the environment after executing. This is because it is
// assumed that Jeeves is executed in an ephemeral environment such as Docker or
// any similiar environment.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/atomosio/oxygen-fuse"
	"github.com/atomosio/titanium-go"
)

const (
	//	defaultTitaniumEndpoint = common.TitaniumEndpoint
	//	defaultOxgenEndpoint    = common.OxygenEndpoint
	//	defaultTitaniumEndpoint = "http://10.240.0.2"
	//	defaultOxgenEndpoint    = "http://10.240.0.3"
	defaultTitaniumEndpoint = "http://localhost:9002"
	defaultOxgenEndpoint    = "http://localhost:9000"
	defaultMountPoint       = "/atomos"
)

var (
	titaniumEndpoint, oxygenEndpoint, tokenString, mountPoint string
)

func init() {
	flag.StringVar(&titaniumEndpoint, "titanium", defaultTitaniumEndpoint, "Titanium Endpoint")
	flag.StringVar(&oxygenEndpoint, "oxygen", defaultOxgenEndpoint, "Oxygen Endpoint")
	flag.StringVar(&mountPoint, "mount", defaultMountPoint, "Mount Point")
	flag.StringVar(&tokenString, "token", "", "Token")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  jeeves -token=<token> [-titanium=<endpoint>] [-oxygen=<endpoint>] [-mount=<path>]\n")
}

var titaniumClient *titanium.HttpClient

func main() {
	flag.Parse()

	// Make sure we specified a token
	if tokenString == "" {
		usage()
		os.Exit(1)
	}

	// Prep environment
	if err := prepEnv(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	titaniumClient = titanium.NewHttpClient(titaniumEndpoint, tokenString)

	// Get task executable and arguments
	instance, err := titaniumClient.GetTokenInstance()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create command for execution
	cmd := exec.Command(instance.Executable, instance.Arguments...)

	// Pass stdout/stderr through
	cmd.Dir = instance.Directory
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmount fuse
	oxygenfuse.Unmount(mountPoint)
}

func prepEnv() error {
	// Get task info

	// Mount at /atomos/
	err := os.Mkdir(mountPoint, 0750)
	if err != nil {
		return err
	}

	readyChan := make(chan bool)
	go oxygenfuse.MountAndServeOxygen(mountPoint, oxygenEndpoint, tokenString, readyChan)
	// Wait until oxygenfuse has mounted and is ready
	for {
		select {
		case <-readyChan:
			return nil
		default:
		}
	}

	return nil
}
