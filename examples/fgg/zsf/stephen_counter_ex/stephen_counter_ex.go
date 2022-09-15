package main;

type Barᐸᐳ interface { BarMemberFunc___Int__() Top };
type BarImplᐸᐳ struct {};
func (this BarImplᐸᐳ) BarMemberFunc___Int__() Top { return this };

type FooᐸBarᐸᐳᐳ interface { idᐸᐳ(in Barᐸᐳ) Barᐸᐳ; id___Bar___Bar__() Top };

type BazᐸBarᐸᐳᐳ struct {};
func (this BazᐸBarᐸᐳᐳ) idᐸᐳ(in Barᐸᐳ) Barᐸᐳ { return in };
func (this BazᐸBarᐸᐳᐳ) id___Bar___Bar__() Top { return this };
type Dummyᐸᐳ struct {};

type BazᐸBarᐸᐳᐳ__id__Dict struct {
	//no further function calls
}

func Baz_Bar_id_Wrapper(this FooᐸBarᐸᐳᐳ, in Barᐸᐳ, dict BazᐸBarᐸᐳᐳ__id__Dict) Barᐸᐳ {
	_this := this.(BazᐸBarᐸᐳᐳ);
	return _this.idᐸᐳ(in)
}

type Dummyᐸᐳ__f__Dict struct {
	idFunc func(FooᐸBarᐸᐳᐳ, Barᐸᐳ, BazᐸBarᐸᐳᐳ__id__Dict) Barᐸᐳ;
	idDict BazᐸBarᐸᐳᐳ__id__Dict;
}

func (this Dummyᐸᐳ) f(in FooᐸBarᐸᐳᐳ, dict Dummyᐸᐳ__f__Dict) Barᐸᐳ {
	return dict.idFunc(in, BarImplᐸᐳ{}, dict.idDict)
};

func (this Dummyᐸᐳ) f___Foo_Bar____Bar__() Top { return this };
type Top interface {};
func main() {
	f_dict := Dummyᐸᐳ__f__Dict{
		idFunc: Baz_Bar_id_Wrapper,
		idDict: BazᐸBarᐸᐳᐳ__id__Dict{},
	}
	_ = Dummyᐸᐳ{}.f(BazᐸBarᐸᐳᐳ{}, f_dict)
}

