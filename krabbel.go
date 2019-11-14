/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
               All rights reserved
           EMail : <support@mwat.de>
*/

package krabbel

//lint:file-ignore ST1017 - I prefer Yoda conditions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	// Arbitrary list of file extensions to exclude from reading/parsing.
	binExts = []string{
		".amr", ".avi", ".azw3", ".bak", ".bibtex", ".bin", ".bz2",
		".cfg", ".com", ".conf", ".css", ".csv",
		".db", ".deb", ".doc", ".docx", ".dia", ".epub", ".exe",
		".flv", ".gif", ".gz", ".ics", ".iso",
		".jar", ".jpeg", ".jpg", ".json",
		".log", ".mobi", ".mp3", ".mp4", ".mpeg",
		".odf", ".odg", ".odp", ".ods", ".odt", ".otf", ".oxt",
		".pas", ".pdf", ".pl", ".png", ".ppd", ".ppt", ".pptx",
		".rip", ".rpm", ".sh", ".spk", ".sql", ".sxg", ".sxw",
		".ttf", ".txt", ".vbox", ".vmdk", ".vcs",
		".wav", ".xls", ".xpi", ".xsl", ".zip",
	}

	// RegEx to match complete link tags.
	hrefRE = regexp.MustCompile(`(?si)(<a[^>]+href=")([^"#]+)[^"]*"[^>]*>`)
	//                                1              2

	// RegEx to check whether an URL starts with a scheme.
	schemeRE = regexp.MustCompile(`^\w+:`)

	// RegEx to get the base of an URL.
	startRE = regexp.MustCompile(`^(\w+://[^/]+)`)
)

// `goProcessURL()` reads `aURL` and sends the links therein to `aList`.
//
//	`aBaseURL` The start of all the local URLs.
//	`aURL` The page URL to process.
//	`aList` The channel to receive the list of links.
func goProcessURL(aBaseURL, aURL string, aList chan<- []string) {
	if page, err := readPage(aURL); nil == err {
		if links := pageLinks(aBaseURL, page); nil != links {
			aList <- links
		}
	}
} // goProcessURL()

// `pageLinks()` extracts the A/HREF links in `aPage`
// returning them in a list.
//
//	`aBaseURL` The start of all the local URLs.
//	`aPage` The web-page to handle.
func pageLinks(aBaseURL string, aPage []byte) (rList []string) {
	linkMatches := hrefRE.FindAllSubmatch(aPage, -1)
	if nil == linkMatches {
		return
	}

	for cnt, l := 0, len(linkMatches); cnt < l; cnt++ {
		link := string(linkMatches[cnt][2])
		if 0 < len(link) {
			// check for certain binary file extensions
			for _, ext := range binExts {
				if strings.HasSuffix(link, ext) {
					link = ""
					break
				}
			}
		}
		if 0 == len(link) {
			continue
		}
		if strings.HasPrefix(link, aBaseURL) {
			rList = append(rList, link)
		} else if strings.HasPrefix(link, `/`) {
			rList = append(rList, aBaseURL+link)
		} else if !schemeRE.MatchString(link) { // skip external links
			rList = append(rList, aBaseURL+`/`+link)
		}
	}

	return
} // pageLinks()

const (
	tenSex = 10 * time.Second
)

// `readPage()` requests a single page identified by `aURL`
// returning its contents.
func readPage(aURL string) ([]byte, error) {
	req, err := http.NewRequest(`GET`, aURL, nil)
	if nil != err {
		return nil, err
	}
	req.Header.Set(`Referer`, `https://github.com/mwat56/krabbel`)

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   tenSex,
				KeepAlive: tenSex,
			}).DialContext,
			ExpectContinueTimeout: tenSex,
			ResponseHeaderTimeout: tenSex,
			TLSHandshakeTimeout:   tenSex,
		},
		Timeout: 10 * time.Minute, // prepare for looong response bodies
	}

	fmt.Println(`Reading`, aURL)
	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	if http.StatusOK == resp.StatusCode {
		return ioutil.ReadAll(resp.Body)
	}

	return nil, errors.New(http.StatusText(resp.StatusCode))
} // readPage()

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

const (
	// Half a second to sleep in `Crawl()`.
	halfSecond = 500 * time.Millisecond
)

// Crawl reads web-page links starting with `aStartURL`
// returning the number pages checked.
func Crawl(aStartURL string) int {
	var (
		checked  = make(map[string]bool)
		empty    int
		linkList = make(chan []string, 63)
	)
	linkList <- []string{aStartURL}
	baseURL := startRE.FindString(aStartURL)

	for {
		select {
		case list, more := <-linkList:
			if !more { // channel closed
				return len(checked)
			}
			empty = 0
			for _, link := range list {
				if done, ok := checked[link]; ok && done {
					continue
				}
				checked[link] = true
				go goProcessURL(baseURL, link, linkList)
			}

		default:
			empty++
			if 3 < empty {
				return len(checked) // nothing more to do
			}
			time.Sleep(halfSecond)
		}
	}
} // Crawl()

/* EoF */
