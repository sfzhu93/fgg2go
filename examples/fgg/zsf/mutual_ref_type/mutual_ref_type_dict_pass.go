package main;
type Intᐸᐳ struct {};
func (this Intᐸᐳ) Firstᐸᐳ(that Intᐸᐳ) Intᐸᐳ { return this };
func (this Intᐸᐳ) First___Int___Int__() Top { return this };
type FooImplᐸᐳ struct { inner Intᐸᐳ };
type BarImplᐸᐳ struct { inner Intᐸᐳ };
func (this FooImplᐸᐳ) Innerᐸᐳ() Intᐸᐳ { return this.inner };
func (this FooImplᐸᐳ) Inner___Int__() Top { return this };
func (this FooImplᐸᐳ) FooFuncᐸᐳ(that BarImplᐸᐳ) FooImplᐸᐳ { return FooImplᐸᐳ{this.Innerᐸᐳ().Firstᐸᐳ(that.Innerᐸᐳ())} };

func (this FooImplᐸᐳ) FooFunc(that Bar) FooImplᐸᐳ {
	innerFunc := that.getBarDict().InnerFunc
	return FooImplᐸᐳ{this.Innerᐸᐳ().Firstᐸᐳ(innerFunc(that))}
};

func (this FooImplᐸᐳ) FooFunc___BarImpl___FooImpl__() Top { return this };
func (this BarImplᐸᐳ) Innerᐸᐳ() Intᐸᐳ { return this.inner };
func (this BarImplᐸᐳ) Inner___Int__() Top { return this };
func (this BarImplᐸᐳ) BarFuncᐸᐳ(that FooImplᐸᐳ) BarImplᐸᐳ { return BarImplᐸᐳ{this.Innerᐸᐳ().Firstᐸᐳ(that.Innerᐸᐳ())} };

func (this BarImplᐸᐳ) BarFunc(that Foo) BarImplᐸᐳ {
	innerFunc := that.getFooDict().InnerFunc
	return BarImplᐸᐳ{this.Innerᐸᐳ().Firstᐸᐳ(innerFunc(that))}
}

func (this BarImplᐸᐳ) BarFunc___FooImpl___BarImpl__() Top { return this };


type Foo interface {
	//FooFunc(that Bar) Foo
	getFooDict() FooDict;
}

type FooDictFooFuncFuncType func(this Foo, that Bar) Foo;
type FooDictInnerFuncType func(this Foo) Intᐸᐳ;

type FooDict struct {
	FooFuncFunc FooDictFooFuncFuncType;
	InnerFunc FooDictInnerFuncType;
}

type Bar interface {
	//BarFunc(that Foo) Bar
	getBarDict() BarDict;
}

func (this FooImplᐸᐳ) getFooDict() FooDict {
	return FooDict{FooFuncFunc: FooFooImplFooFuncWrapper, InnerFunc: FooFooImplInnerWrapper}
}

type BarDictBarFuncFuncType func(this Bar, that Foo) Bar;
type BarDictInnerFuncType func(this Bar) Intᐸᐳ;

type BarDict struct {
	BarFuncFunc BarDictBarFuncFuncType;
	InnerFunc BarDictInnerFuncType
}

func (this BarImplᐸᐳ) getBarDict() BarDict {
	return BarDict{BarFuncFunc: BarBarImplBarFuncWrapper, InnerFunc: BarBarImplInnerWrapper}
}

//nomenclature: interface name+implementation struct name+method name+"Wrapper"
func FooFooImplFooFuncWrapper(this Foo, that Bar) Foo {
	_this := this.(FooImplᐸᐳ);
	ret := _this.FooFunc(that)
	return ret
}

func BarBarImplBarFuncWrapper(this Bar, that Foo) Bar {
	_this := this.(BarImplᐸᐳ);
	return _this.BarFunc(that);
}

func FooFooImplInnerWrapper(this Foo) Intᐸᐳ {
	return this.(FooImplᐸᐳ).Innerᐸᐳ()
}

func BarBarImplInnerWrapper(this Bar) Intᐸᐳ {
	return this.(BarImplᐸᐳ).Innerᐸᐳ()
}

type Top interface {};
func main() {
	_ = FooImplᐸᐳ{Intᐸᐳ{}}.FooFunc(BarImplᐸᐳ{Intᐸᐳ{}}.BarFunc(FooImplᐸᐳ{Intᐸᐳ{}}))
}
