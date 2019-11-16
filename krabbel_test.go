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
		{" 1", args{"http://www.mwat.de/", true}, w1, false},
		{" 2", args{"http://mmm.mwat.de/bla", true}, w1, true},
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

func Test_pageLinks(t *testing.T) {
	bu1 := startRE.FindString("http://dev.mwat.de/")
	p1, _ := readPage(bu1, true)
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
