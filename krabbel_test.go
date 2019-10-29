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

func Test_parseStartURL(t *testing.T) {
	type args struct {
		aURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{" 1", args{``}, ``},
		{" 2", args{`http://127.0.0.1`}, `http://127.0.0.1`},
		{" 3", args{`http://127.0.0.1:8080`}, `http://127.0.0.1:8080`},
		{" 4", args{`http://127.0.0.1:8080/dir/`}, `http://127.0.0.1:8080`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStartURL(tt.args.aURL); got != tt.want {
				t.Errorf("parseStartURL() = %v,\nwant %v", got, tt.want)
			}
		})
	}
} // Test_parseStartURL()

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
	bu1 := getStartURL("http://dev.mwat.de/")
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
			// if !reflect.DeepEqual(gotRList, tt.wantRList) {
			// 	t.Errorf("pageLinks() = %v, want %v", gotRList, tt.wantRList)
			// }
			if 0 == len(gotRList) {
				t.Errorf("pageLinks() = %v, want (!nil)", gotRList)
			}
		})
	}
} // Test_pageLinks()
