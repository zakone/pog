// go run listpack.go hash
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Pack struct {
	Dir        string
	ImportPath string
	Name       string
	Doc        string
	Target     string
	Goroot     bool
	Standard   bool
	Root       string
	GoFiles    []string
	Imports    []string
	Deps       []string
}

func main() {
	pack := os.Args[1]
	out, _ := exec.Command("go", "list", "-f", "{{join .Deps \" \"}}", pack).Output()

	var b bytes.Buffer
	b.Write(out[:len(out)-1])

	res := strings.Split(b.String(), " ")
	for _, s := range res {
		output, err := exec.Command("go", "list", "-json", s).Output()
		if err != nil {
			fmt.Printf("package json command failed %s", err)
		}

		var p Pack
		if err := json.Unmarshal(output, &p); err != nil {
			log.Fatalf("Json unmarshaling failed: %s", err)
		}
		fmt.Printf("#Dir\t %s\n#ImportPath:\t %s\n#Name:\t %s\n", p.Dir, p.ImportPath, p.Name)

	}
}

// #Dir	 /usr/local/opt/go/libexec/src/errors
// #ImportPath:	 errors
// #Name:	 errors
// #Dir	 /usr/local/opt/go/libexec/src/internal/race
// #ImportPath:	 internal/race
// #Name:	 race
// #Dir	 /usr/local/opt/go/libexec/src/io
// #ImportPath:	 io
// #Name:	 io
// #Dir	 /usr/local/opt/go/libexec/src/runtime
// #ImportPath:	 runtime
// #Name:	 runtime
// #Dir	 /usr/local/opt/go/libexec/src/runtime/internal/atomic
// #ImportPath:	 runtime/internal/atomic
// #Name:	 atomic
// #Dir	 /usr/local/opt/go/libexec/src/runtime/internal/sys
// #ImportPath:	 runtime/internal/sys
// #Name:	 sys
// #Dir	 /usr/local/opt/go/libexec/src/sync
// #ImportPath:	 sync
// #Name:	 sync
// #Dir	 /usr/local/opt/go/libexec/src/sync/atomic
// #ImportPath:	 sync/atomic
// #Name:	 atomic
// #Dir	 /usr/local/opt/go/libexec/src/unsafe
// #ImportPath:	 unsafe
// #Name:	 unsafe
