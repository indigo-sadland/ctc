package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/projectdiscovery/gologger"
	"io/ioutil"
	"os"
	"strings"
)

var files *string
var output *string

// WriteToFile is for writing array of strings to a file.
func writeToFile(fileName string, data []string) {

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		gologger.Error().Msgf(err.Error())
	}

	wrt := bufio.NewWriter(file)

	for _, line := range data {
		_, err2 := wrt.WriteString(line + "\n")
		if err2 != nil {
			gologger.Error().Msgf(err2.Error())
		}

	}

	wrt.Flush()
	file.Close()

}

// comUniq gets list and deletes repeated strings.
// So, only one of the repeated strings is left.
func comUniq(combinedSlice []string) []string {

	keys := make(map[string]bool)
	uniqueList := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. Else we jump on another element.
	for _, elem := range combinedSlice {
		if _, value := keys[elem]; !value {
			keys[elem] = true
			if elem == "" { // remove empty string
				continue
			}
			uniqueList = append(uniqueList, elem)
		}
	}

	if *output == "" {
		for _, l := range uniqueList {
			fmt.Printf(l + "\n")
		}
	} else {
		writeToFile(*output, uniqueList)
	}

	return uniqueList

}

// readFiles takes slice of file path, reads line and combine them in one slice.
func readFiles(paths []string) {

	var lines []string

	for _, p := range paths {
		body, err := ioutil.ReadFile(p)
		if err != nil {
			gologger.Error().Msgf(err.Error())
			return
		}

		str := strings.Split(string(body), "\n")
		for _, s := range str {
			lines = append(lines, s)
		}
	}

	comUniq(lines)

}

func main() {

	var slice []string

	files = flag.String("fs", "", "")
	output = flag.String("o", "", "")
	flag.Usage = func() {
		fmt.Printf("Usage:\n\t" +
			"-fs, <FILE1>,<FILE2>,<FILEn>   Define files separated by comma.\n\t" +
			"-o,  <OUTPUT_FILE>             Save results to file. Prints into terminal if not specified.",
		)
	}
	flag.Parse()

	if *files == "" {
		flag.Usage()
		return
	}

	// Split flag's value and combine it in slice.
	for _, s := range strings.Split(*files, ",") {

		s = strings.TrimSpace(s)

		if s == "" {
			continue
		}

		slice = append(slice, s)
	}

	readFiles(slice)
}
