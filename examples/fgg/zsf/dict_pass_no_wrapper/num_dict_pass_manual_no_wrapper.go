package main;
type TimesTwoFactoryᐸᐳ struct {};
type Intᐸᐳ struct {};


func (this Intᐸᐳ) Plusᐸᐳ(that Intᐸᐳ) Intᐸᐳ { return that };
func (this Intᐸᐳ) Plus___Int___Int__() Top { return this };
func (this TimesTwoFactoryᐸᐳ) TimesTwoᐸIntᐸᐳᐳ(x Intᐸᐳ) Intᐸᐳ { return x.Plusᐸᐳ(x) };
func (this TimesTwoFactoryᐸᐳ) TimesTwo__β1_Num_β1___β1_β1() Top { return this };


type Top interface {};

type PlusFuncType func(this Top, that Top) Top;
type NumDict struct {
	dict map[string]PlusFuncType
};

//The wrapper is necessary in Rust. But in Go, it seems only a way to fit into a dictionary.
func NumIntPlusWrapper(this Top, that Top) Top {
	_this := this.(Intᐸᐳ);
	_that := that.(Intᐸᐳ);
	ret := _this.Plusᐸᐳ(_that);
	return ret//You may wonder why it is not ret.(Top),
	// because in standard go 𝑒 has to be an interface type in 𝑒.(𝜏)
}

//Go doesn't allow const arrays or maps. So we have to have this to somehow represent a dictionary
func getNumDictOfInt() map[string]PlusFuncType {
	return map[string]PlusFuncType{
		"Plus": (*Intᐸᐳ).Plusᐸᐳ,
	}
}

func (this TimesTwoFactoryᐸᐳ) TimesTwo_dict_pass(x Top, dict1 NumDict) Top {
	Plusᐸᐳ := dict1.dict["Plus"];
	return Plusᐸᐳ(x, x)
};

func main() {
	_ = TimesTwoFactoryᐸᐳ{}.TimesTwoᐸIntᐸᐳᐳ(Intᐸᐳ{});
	_NumDictofInt := NumDict{
		dict:getNumDictOfInt(),
	};
	_ = TimesTwoFactoryᐸᐳ{}.TimesTwo_dict_pass(Intᐸᐳ{}, _NumDictofInt).(Intᐸᐳ)

}