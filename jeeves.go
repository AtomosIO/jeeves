// Jeeves does not clean up the environment after executing. This is because it is
// assumed that Jeeves is executed in an ephemeral environment such as Docker or
// any similiar environment.
package main

import (
	//	"errors"
	"flag"
	"fmt"
	//	"github.com/atomosio/common"
	"github.com/atomosio/oxygen-fuse"
	//	"github.com/atomosio/oxygen-go"
	"os"
	"os/exec"
)

const (
	//	defaultTitaniumEndpoint = common.TitaniumEndpoint
	//	defaultOxgenEndpoint    = common.OxygenEndpoint
	//	defaultTitaniumEndpoint = "http://10.240.0.2/"
	//	defaultOxgenEndpoint    = "http://10.240.0.3/"
	defaultTitaniumEndpoint = "http://localhost:9002/"
	defaultOxgenEndpoint    = "http://localhost:9000/"
)

var (
	titaniumEndpoint, oxygenEndpoint, tokenString, mountPoint string
)

func init() {
	flag.StringVar(&titaniumEndpoint, "titanium", defaultTitaniumEndpoint, "Titanium Endpoint")
	flag.StringVar(&oxygenEndpoint, "oxygen", defaultOxgenEndpoint, "Oxygen Endpoint")
	flag.StringVar(&mountPoint, "mount", "/atomos", "Mount Point")
	flag.StringVar(&tokenString, "token", "", "Token")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  jeeves -token=<token> [-titanium=<endpoint>] [-oxygen=<endpoint>] [-mount=<path>]\n")
}

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

	// Run jeeves
	output, err := exec.Command("echo", "test").CombinedOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(output))

	// Unmount fuse
	oxygenfuse.Unmount(mountPoint)
}

func prepEnv() error {
	// Get task info

	// Mount at ./atomos/
	err := os.Mkdir(mountPoint, 0750)
	if err != nil {
		return err
	}

	go oxygenfuse.MountAndServeOxygen(mountPoint, oxygenEndpoint, tokenString)

	// TODO Wait until oxygenfuse has mounted and is ready

	return nil
}
