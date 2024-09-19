package test

import (
	"testing"
	"os"
	"fmt"
	"bufio"
	"reflect"
	"github.com/FilterX/logic"
	"github.com/FilterX/pkg"
)

func init_SearchExTest()*logic.SearchEx{
	file, err := os.Open("../testdata/BadWord.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	list := make([]string, 0)
    for scanner.Scan() {
		text:=pkg.RemoveBOM(scanner.Text())
		list = append(list, text)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
	search := logic.NewSearchEx()
	search.SetKeyWords(list)
	return search
}

func TestSearchEx_GetStringFindFirst(t *testing.T) {
	search:=init_SearchExTest()
	
	type args struct {
		text string
	}
	tests := []struct {
		name string
		w    *logic.SearchEx
		args args
		want string
	}{
		{
			name: "1",
			w:    search,
			args: args{
				text: "15768靠北柯賜海柯庆施警靖国静坐纠察员",
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.GetStringFindFirst(tt.args.text); got != tt.want {
				t.Errorf("SearchEx.FindFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSearchEx_GetStringFindAll(t *testing.T) {
	search:=init_SearchExTest()

	type args struct {
		text string
	}
	tests := []struct {
		name string
		w    *logic.SearchEx
		args args
		want []string
	}{
		{
			name: "1",
			w:    search,
			args: args{
				text: "15768靠北柯賜海柯庆施警靖国静坐纠察员",
			},
			want: []string{"1","15768","靠","靠北","柯賜海","柯庆施","靖国","静坐","纠察员"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.GetStringFindAll(tt.args.text);!reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchEx.GetStringFindAll() = %v, want %v", got,tt.want)
			}
		})
	}
}