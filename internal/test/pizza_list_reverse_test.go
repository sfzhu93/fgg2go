package frontend

import (
	"github.com/sfzhu93/fgg2go/internal/frontend"
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkDictPassPizzaListReverse(b *testing.B) {
	testNames := []string{
		//"pizza_list_reverse.fgg",
		"pizza_list_reverse.fgg",
		//"motiv-ex.fgg",
		//"rwcell.fgg",
		//"chan_close.fgg",
		//"micro-m-exp-9.fgg",
		//"micro-n-40.fgg",
		//"chan_interface.fgg",
	}
	var tests []struct {
		name string
		src  string
	}
	for _, testName := range testNames {
		content, err := ioutil.ReadFile(testName)
		if err != nil {
			log.Fatal(err)
		}
		tests = append(tests, struct {
			name string
			src  string
		}{name: testName, src: string(content)})
	}
	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			src := tt.src
			intrpFgg := frontend.NewFGGInterp(false, src, true)
			intrpFgg.DictPass(false)
		}
	}
}