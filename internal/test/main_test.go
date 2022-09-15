package frontend

import (
	"github.com/sfzhu93/fgg2go/internal/frontend"
	"io/ioutil"
	"log"
	"testing"
)

func TestFGGInterp(t *testing.T) {
	testNames := []string{
		"pizza_list_reverse.fgg",
		"micro_benchmark.fgg",
		"motiv-ex.fgg",
		"rwcell.fgg",
		"chan_close.fgg",
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
	for _, tt := range tests {
		src := tt.src
		intrpFgg := frontend.NewFGGInterp(false, src, true)
		t.Run(tt.name, func(t *testing.T) {
			intrpFgg.DictPass(false)
		})
	}
	for _, tt := range tests {
		src := tt.src
		intrpFgg := frontend.NewFGGInterp(false, src, true)
		t.Run(tt.name, func(t *testing.T) {
			intrpFgg.Dyn()
		})
	}
}

func BenchmarkDictPassMExp(b *testing.B) {
	testNames := []string{
		//"pizza_list_reverse.fgg",
		//"micro_benchmark.fgg",
		//"motiv-ex.fgg",
		//"rwcell.fgg",
		//"chan_close.fgg",
		"micro-m-exp-9.fgg",
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

func BenchmarkDictPassN(b *testing.B) {
	testNames := []string{
		//"pizza_list_reverse.fgg",
		//"micro_benchmark.fgg",
		//"motiv-ex.fgg",
		//"rwcell.fgg",
		//"chan_close.fgg",
		//"micro-m-exp-9.fgg",
		"micro-n-40.fgg",
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

func RunIntrpFggWithDictPass(intrpFgg *frontend.FGGInterp) {
	intrpFgg.DictPass(false)
}

func RunIntrpFggWithDyn(intrpFgg *frontend.FGGInterp) {
	intrpFgg.Dyn()
}

func TestPanic(t *testing.T) {
	testNames := []string{
		"translateMethod_first_panic.fgg",
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
	for _, tt := range tests {
		src := tt.src
		intrpFgg := frontend.NewFGGInterp(false, src, true)
		shouldPanic(t, RunIntrpFggWithDictPass, intrpFgg)
	}
	for _, tt := range tests {
		src := tt.src
		intrpFgg := frontend.NewFGGInterp(false, src, true)
		shouldPanic(t, RunIntrpFggWithDyn, intrpFgg)
	}
}

func shouldPanic(t *testing.T, f func(*frontend.FGGInterp), interp *frontend.FGGInterp) {
	defer func() { recover() }()
	f(interp)
	t.Errorf("should have panicked")
}
