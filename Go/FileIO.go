/* Author : Syed Mujtaba
   Seek Knowledge from the Cradle to the Grave
   This program is a part of my Cookbook at : https://github.com/mujtaba1747/Cookbook
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func writeString(fname string, str string) error {
	fp, err := os.OpenFile(fname, os.O_WRONLY, 0666)
	defer fp.Close() // Why do we close files : So that the resources allocated to them can be used elsewhere
	if err != nil {
		return err
	}
	_, err = io.WriteString(fp, str)
	if err != nil {
		return err
	}
	return nil
}

func readLineByLine(fname string) (lines []string, err error) {
	fp, err := os.Open(fname)
	defer fp.Close() 
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
func readLineByLine2(fname string) (lines []string, err error) {
	fp, err := os.Open(fname)
	defer fp.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(fp)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			// Some other error
			return nil, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}
func main() {
	// fmt.Println(os.Stdout) // os.Stdout is a *os.File pointing to /dev/stdout

	writeString("/dev/stdout", "This program behaves like command tee if path is stdin\n") // Will Work only on Linux, use os.Stdin for safety // same as io.WriteString(os.Stdout, message) 

	lines, _ := readLineByLine("/dev/stdin") // Does work Similar to tee // You can have ANY file instead of stdin as well, in case of stdin use Ctrl+D on Linux to stop reading
	for _, str := range lines {
		fmt.Print(str)
	}
	fmt.Println()
}
