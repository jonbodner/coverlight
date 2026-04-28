package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	output  = flag.String("o", "", "file for output")
	htmlOut = flag.String("html", "", "generate HTML representation of coverage profile")
	mode    = flag.String("mode", "light", "use either light or dark mode for output (default light)")
)

func main() {
	flag.Parse()
	profile := *htmlOut
	err := htmlOutput(profile, *output, *mode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "coverlight: %v\n", err)
		os.Exit(2)
	}
}
