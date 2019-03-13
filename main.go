package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var redactedFlags redactFlags

func init() {

	flag.Var(&redactedFlags, "kind", "Kubernetes kind to ignore")

	flag.Parse()
}
func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Split(splitOnDashes)

	for reader.Scan() {
		manifest := reader.Text()
		obj := manifestStub{}
		yaml.Unmarshal([]byte(manifest), &obj)

		if contains(redactedFlags, obj.Kind) {
			continue
		}

		fmt.Fprint(os.Stdout, manifest)
	}
	fmt.Fprintln(os.Stdout)
}

// contains tells whether a contains x.
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
