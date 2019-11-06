/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
               All rights reserved
           EMail : <support@mwat.de>
*/

package krabbel

//lint:file-ignore ST1017 - I prefer Yoda conditions

import (
	"testing"
)

func Test_readPage(t *testing.T) {
	var w1 []byte
	type args struct {
		aURL string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{" 1", args{"http://www.mwat.de/"}, w1, false},
		{" 2", args{"http://mmm.mwat.de/bla"}, w1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readPage(tt.args.aURL)
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

func Test_pageLinks(t *testing.T) {
	bu1 := startRE.FindString("http://dev.mwat.de/")
	p1, _ := readPage(bu1)
	var w1 []string
	type args struct {
		aBaseURL string
		aPage    []byte
	}
	tests := []struct {
		name      string
		args      args
		wantRList []string
	}{
		// TODO: Add test cases.
		{" 1", args{bu1, p1}, w1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRList := pageLinks(tt.args.aBaseURL, tt.args.aPage)
			if 0 == len(gotRList) {
				t.Errorf("pageLinks() = %v, want (!nil)", gotRList)
			}
		})
	}
} // Test_pageLinks()
