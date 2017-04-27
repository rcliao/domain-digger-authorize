package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()
	filePath := flag.Arg(0)
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	resultFile, _ := os.Create("result.txt")

	for scanner.Scan() {
		result := dig(scanner.Text())
		resultFile.WriteString(result)
	}
}

func dig(domain string) string {
	cmd := exec.Command("dig", domain, "+trace", "+all")
	cmd2 := exec.Command("sed", "-n", "/AUTHORITY SECTION/,/^$/p")
	cmd2.Stdin, _ = cmd.StdoutPipe()
	var out bytes.Buffer
	cmd2.Stdout = &out
	_ = cmd2.Start()
	_ = cmd.Run()
	err := cmd2.Wait()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("output for domain \"%s\":\n%s\n\n", domain, out.String())
}
