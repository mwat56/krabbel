/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
               All rights reserved
           EMail : <support@mwat.de>
*/

package main

//lint:file-ignore ST1017 - I prefer Yoda conditions

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mwat56/krabbel"
)

// `getArguments()` returns the values of commandline arguments.
func getArguments() (rURL string, rCGI bool) {
	flag.StringVar(&rURL, "url", rURL, "the URL to start crawling")
	rCGI = true
	flag.BoolVar(&rCGI, "cgi", rCGI, "<bool> use CGI arguments")

	flag.Usage = showHelp
	flag.Parse()

	return
} // getArguments()

// `main()` runs the application.
func main() {
	URL, useCGI := getArguments()
	if 0 == len(URL) {
		showHelp()
		os.Exit(1)
	}

	startTime := time.Now()
	checked := krabbel.Crawl(URL, useCGI)
	elapsed := time.Since(startTime)
	log.Printf("checked %d pages in %s", checked, elapsed)
} // main()

// `showHelp()` lists the commandline options to `Stderr`.
func showHelp() {
	fmt.Fprintf(os.Stderr, "\n  Usage: %s [OPTIONS]\n\n", os.Args[0])
	flag.PrintDefaults()
} // showHelp()

/* EoF */
