package main;
import "fmt";
type Any interface {

};

type Intᐸᐳ struct {};
type Eventᐸᐳ interface { ProcessᐸIntᐸᐳᐳ(y Intᐸᐳ) Intᐸᐳ; Process__β1_Any____β1_Int__() Top };
type UIEventᐸᐳ struct {};
func (x UIEventᐸᐳ) ProcessᐸIntᐸᐳᐳ(y Intᐸᐳ) Intᐸᐳ { return Intᐸᐳ{} };
func (x UIEventᐸᐳ) Process(y Any) Intᐸᐳ { return Intᐸᐳ{} };
func (x UIEventᐸᐳ) Process__β1_Any____β1_Int__() Top { return x };
type Dispatcherᐸᐳ struct {};
func (x Dispatcherᐸᐳ) Dispatchᐸᐳ(y Eventᐸᐳ) Intᐸᐳ { return y.ProcessᐸIntᐸᐳᐳ(Intᐸᐳ{}) };

func (x Dispatcherᐸᐳ) Dispatch(y Eventᐸᐳ, dict DispatchDict) Intᐸᐳ {
	return dict.ProcessFunc(y, Intᐸᐳ{})
};

func (x Dispatcherᐸᐳ) Dispatch___Event___Int__() Top { return x };
type Top interface {};
type ProcessDict struct {

}
type DispatchDict struct {
	ProcessFunc func (this Eventᐸᐳ, y Any) Intᐸᐳ;
	ProcessFuncDict ProcessDict;
}

func ProcessFuncWrapper(this Eventᐸᐳ, y Any) Intᐸᐳ {
	return this.(UIEventᐸᐳ).Process(y)
}

func main() {
	dict := DispatchDict{
		ProcessFunc:     ProcessFuncWrapper,
		ProcessFuncDict: ProcessDict{},
	}
	fmt.Printf("%#v", Dispatcherᐸᐳ{}.Dispatch(UIEventᐸᐳ{}, dict))
}
