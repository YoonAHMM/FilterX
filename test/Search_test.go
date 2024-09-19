package test

import (
	"reflect"
	"github.com/FilterX/pkg"
	"testing"
	"os"
	"fmt"
	"bufio"
	"github.com/FilterX/logic"
)

func init_SearchTest()*logic.Search{
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
	search := logic.NewSearch()
	search.SetKeywords(list)
	return search
}
func TestSearch_GetStringFindFirst(t *testing.T) {
	search:=init_SearchTest()
	
	type args struct {
		text string
	}
	tests := []struct {
		name string
		w    *logic.Search
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
				t.Errorf("Search.GetStringFindFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch_GetStringFindAll(t *testing.T) {
	search:=init_SearchTest()

	type args struct {
		text string
	}
	tests := []struct {
		name string
		w    *logic.Search
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
				t.Errorf("Search.GetStringFindAll() = %v, want %v", got,tt.want)
			}
		})
	}
}

func TestSearch_ContainsAny(t *testing.T) {
	search:=init_SearchTest()

	type args struct {
		text string
	}
	tests := []struct {
		name string
		w    *logic.Search
		args args
	}{
		{
			name: "1",
			w:    search,
			args: args{
				text: "wwwwwwwwddddaadadadadadhmmloveljxfhhfhffhh1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.ContainsAny(tt.args.text);!got {
				t.Errorf("Search.ContainsAny() = %v,want:true", got)
			}
		})
	}
}

func TestSearch_Replace(t *testing.T) {
	search:=init_SearchTest()

	type args struct {
		text string
	}
	tests := []struct {
		name string
		w    *logic.Search
		args args
		want string
	}{
		{
			name: "1",
			w:    search,
			args: args{
				text: "11111hmmloveljx11111",
			},
			want: "*****hmmloveljx*****",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Replace(tt.args.text,'*');!reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search.Replace() = %v, want %v", got,tt.want)
			}
		})
	}
}
