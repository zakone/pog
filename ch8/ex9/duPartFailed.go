package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "sync"
    "time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

//!+
func main() {
    // ...determine roots...

    //!-
    flag.Parse()

    // Determine the initial directories.
    roots := flag.Args()
    if len(roots) == 0 {
        roots = []string{"."}
    }
    fileSizes := make([]chan int64, len(roots))
    for i := range fileSizes {
        fileSizes[i] = make(chan int64)
    }
    //!+
    // Traverse each root of the file tree in parallel.

    var n sync.WaitGroup
    for i, root := range roots {
        n.Add(1)
        go walkDir(root, &n, fileSizes[i])
    }
    go func() {
        n.Wait()
        for _, filesize := range fileSizes {
            close(filesize)
        }
    }()
    //!-

    // Print the results periodically.
    var tick <-chan time.Time
    if *vFlag {
        tick = time.Tick(500 * time.Millisecond)
    }

    nfiles := make([]int64, len(roots))
    nbytes := make([]int64, len(roots))

loop:
    for {
        for i := 0; i < len(roots); i++ {
            select {
            case size, ok := <-fileSizes[i]:
                if !ok {
                    break loop // fileSizes was closed
                }
                nfiles[i]++
                nbytes[i] += size
            case <-tick:
                printDiskUsage(nfiles, nbytes, roots)
            }
        }

    }

    printDiskUsage(nfiles, nbytes, roots) // final totals

}

//!-

func printDiskUsage(nfiles, nbytes []int64, roots []string) {
    for i := 0; i < len(nfiles); i++ {
        fmt.Printf("root dir: %s\n", roots[i])
        fmt.Printf("%d files  %.1f GB\n", nfiles[i], float64(nbytes[i])/1e9)
    }
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(dir string, n *sync.WaitGroup, filesize chan<- int64) {
    defer n.Done()
    for _, entry := range dirents(dir) {
        if entry.IsDir() {
            n.Add(1)
            subdir := filepath.Join(dir, entry.Name())
            go walkDir(subdir, n, filesize)
        } else {
            filesize <- entry.Size()
        }
    }
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
    sema <- struct{}{}        // acquire token
    defer func() { <-sema }() // release token
    // ...
    //!-sema

    entries, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "du: %v\n", err)
        return nil
    }
    return entries
}
