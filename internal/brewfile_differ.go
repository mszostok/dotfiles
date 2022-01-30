package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type (
	Brewfile map[string]Entries
	Entries  map[string]struct{}
)

const (
	reset  = "\033[m"
	yellow = "\033[33m"
)

func main() {
	source := flag.String("source", "./Brewfile", "Define Brewfile source path")
	flag.Parse()

	sourceFile, err := os.Open(*source)
	fatalOnErr(err)
	sourceBrew := parseBrewfile(sourceFile)

	cmd := exec.Command("brew", "bundle", "dump", "--all", "--file=-")
	out, err := cmd.Output()
	fatalOnErr(err)
	targetBrew := parseBrewfile(bytes.NewReader(out))

	missingByKind := getMissingEntriesFromTarget(sourceBrew, targetBrew)

	printDiff(os.Stdout, "source", missingByKind)
}

func printDiff(w io.Writer, entry string, missingByKind map[string][]string) {
	fmt.Fprint(w, yellow)
	fmt.Fprintf(w, "Missing in %s\n", entry)
	fmt.Fprint(w, reset)

	for kind, apps := range missingByKind {
		for _, app := range apps {
			fmt.Fprintf(w, "%s %s\n", kind, app)
		}
		fmt.Fprintln(w)
	}
}

func getMissingEntriesFromTarget(source, target Brewfile) map[string][]string {
	missingByKind := map[string][]string{}
	for kind, apps := range target {
		k, found := source[kind]
		if !found {
			continue
		}
		for app := range apps {
			if _, found := k[app]; found {
				continue
			}

			missingByKind[kind] = append(missingByKind[kind], app)
		}
	}

	return missingByKind
}

func parseBrewfile(in io.Reader) Brewfile {
	out := Brewfile{}

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		fields := strings.SplitN(scanner.Text(), " ", 2)
		if len(fields) < 2 {
			continue
		}

		kind := fields[0]
		app := fields[1]

		if out[kind] == nil {
			out[kind] = Entries{}
		}
		out[kind][app] = struct{}{}
	}

	return out
}

func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
