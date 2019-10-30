/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
               All rights reserved
           EMail : <support@mwat.de>
*/

package main

//lint:file-ignore ST1017 - I prefer Yoda conditions

import (
	"log"
	"os"
	"time"

	"github.com/mwat56/krabbel"
)

func main() {
	if 2 > len(os.Args) {
		log.Fatal("call:\n\n\t", os.Args[0], " <pageURL>\n\n")
	}
	URL := os.Args[1]
	startTime := time.Now()
	checked := krabbel.Crawl(URL)
	elapsed := time.Since(startTime)
	log.Printf("checked %d pages in %s", checked, elapsed)
} // main()

/* EoF */
