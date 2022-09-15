package main;
type TimesTwoFactoryᐸᐳ struct {};
type Intᐸᐳ struct {};

//----generated by mono approach
//We still need this function, because we need an implementation of the basic
//operation.
func (this Intᐸᐳ) Plusᐸᐳ(that Intᐸᐳ) Intᐸᐳ { return that };

//(compiled from fgg)
func (this Intᐸᐳ) Plus___Int___Int__() Top { return this };

//This will be replaced by TimesTwo_dict_pass
func (this TimesTwoFactoryᐸᐳ) TimesTwoᐸIntᐸᐳᐳ(x Intᐸᐳ) Intᐸᐳ { return x.Plusᐸᐳ(x) };

//(compiled from fgg)
func (this TimesTwoFactoryᐸᐳ) TimesTwo__β1_Num_β1___β1_β1() Top { return this };

//use it as an empty interface.
type Top interface {};
//----end

//The type of the function to be put into the dictionary.
//Because the parameters have to be fit into many types, we use Top,
//which is actually interface{} that can be put everywhere.
type PlusFuncType func(this Num, that Num) Num;

type Num interface {
	getNumDict() NumDict;
};

//The dictionary.
type NumDict struct {
	PlusFunc PlusFuncType
	//We add fields here when there's more functions to be added into the dictionary.
	//Each func may have different signatures.
};

//The wrapper is necessary in Rust because of low level code of stack variables.
//In Go, it is also necessary, because we have to match the dict's function types.
//PlusFuncType is all about Top.
//The type of Plusᐸᐳ is not compatible with PlusFuncType in Go.
func NumIntPlusWrapper(this Num, that Num) Num {
	_this := this.(Intᐸᐳ);
	_that := that.(Intᐸᐳ);
	ret := _this.Plusᐸᐳ(_that);
	return ret//You may wonder why it is not ret.(Top),
	//ret is of type Intᐸᐳ. It will be converted to Top automatically.
}

//Go doesn't allow const complex types. So we have to initialize the fields like this.
func (this Intᐸᐳ) getNumDict() NumDict {
	return NumDict{
		PlusFunc: NumIntPlusWrapper,
	}
}

//This is the TimesTwo function implemented in dictionary passing.
func (this TimesTwoFactoryᐸᐳ) TimesTwo_dict_pass(x Num) Num {
	Plus := x.getNumDict().PlusFunc;
	return Plus(x, x)
};

func main() {
	_ = TimesTwoFactoryᐸᐳ{}.TimesTwoᐸIntᐸᐳᐳ(Intᐸᐳ{}); //mono
	//_NumDictofInt := getNumDictOfInt();
	//_ = TimesTwoFactoryᐸᐳ{}.TimesTwo_dict_pass(Intᐸᐳ{}, _NumDictofInt).(Intᐸᐳ) //dict pass
	_ = TimesTwoFactoryᐸᐳ{}.TimesTwo_dict_pass(Intᐸᐳ{}).(Intᐸᐳ) //dict pass
}