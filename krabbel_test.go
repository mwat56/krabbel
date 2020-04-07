/*
   Copyright © 2019, 2020 M.Watermann, 10247 Berlin, Germany
               All rights reserved
           EMail : <support@mwat.de>
*/

package krabbel

//lint:file-ignore ST1017 - I prefer Yoda conditions

import (
	"testing"
)

func Test_pageLinks(t *testing.T) {
	bu1 := startRE.FindString("http://dev.mwat.de/")
	p1, _ := readPage(bu1, true)
	bu2 := startRE.FindString("http://192.168.192.234:8181/")
	p2, _ := readPage(bu2, true)
	var w1 []string
	type args struct {
		aBaseURL string
		aPage    []byte
		aUseCGI  bool
	}
	tests := []struct {
		name      string
		args      args
		wantRList []string
	}{
		// TODO: Add test cases.
		{" 1", args{bu1, p1, true}, w1},
		{" 2", args{bu2, p2, true}, w1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRList := pageLinks(tt.args.aBaseURL, tt.args.aPage, tt.args.aUseCGI)
			// if !reflect.DeepEqual(gotRList, tt.wantRList) {
			// 	t.Errorf("pageLinks() = %v, want %v", gotRList, tt.wantRList)
			// }
			if 0 == len(gotRList) {
				t.Errorf("pageLinks() = %v, want (!nil)", gotRList)
			}
		})
	}
} // Test_pageLinks()

func Test_readPage(t *testing.T) {
	var w1 []byte
	type args struct {
		aURL    string
		aUseCGI bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		// {" 1", args{"http://www.mwat.de/", true}, w1, false},
		// {" 2", args{"http://mmm.mwat.de/bla", true}, w1, false},
		{" 3", args{"http://192.168.192.234:8181/", true}, w1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readPage(tt.args.aURL, true)
			if (err != nil) != tt.wantErr {
				t.Errorf("readPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (0 == len(got)) != tt.wantErr {
				t.Errorf("readPage() = %s, want (!nil)", got)
			}
		})
	}
} // Test_readPage()

func TestCrawl(t *testing.T) {
	type args struct {
		aStartURL string
		aUseCGI   bool
		aQuiet    bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{" 1", args{`http://192.168.192.234:8181/`, true, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Crawl(tt.args.aStartURL, tt.args.aUseCGI, tt.args.aQuiet); got <= 0 {
				t.Errorf("Crawl() = %v, want >0", got)
			}
		})
	}
} // TestCrawl()
