// Instructions:
// go build wordCount.go
// ./wordCount filename substring

// Output: 1 line containing number of substrings

package main
import (
	"fmt"
	"runtime"
	"sync/atomic"
	s "strings"
	"bufio"
	"os"
	"log"
)

func reader(id int, counter *uint64, line chan string, done chan string, txt string) {
	for {
		lintxt, exists := <- line
		if exists {
			if s.Contains(lintxt, txt) {
		    	atomic.AddUint64(counter, 1)
		    }
		} else {
			done <- "Done"	
			return
		}
	}
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: ./wordCount filename substring")
		return
	}

	var counter uint64 = 0
	ncpu := runtime.NumCPU()
	done := make(chan string)
	line := make(chan string)
	
	file, err := os.Open(os.Args[1])
	if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	txt := os.Args[2]

    scanner := bufio.NewScanner(file)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	for i := 0; i < ncpu; i++ {
		go reader(i, &counter, line, done, txt)
    }
    for scanner.Scan() {
    	line <- scanner.Text()
    }
    close(line)

    for i := 0; i < ncpu; i++ {
    	<- done
    }
    result := atomic.LoadUint64(&counter)
    fmt.Println(result)

}