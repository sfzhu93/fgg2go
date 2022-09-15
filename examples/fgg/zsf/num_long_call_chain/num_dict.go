package main;
type Num interface {
	//Plus(that Num, dict PlusDict) Num
}

type PlusDict struct {

}

type Intᐸᐳ struct {};

func (this Intᐸᐳ) Plusᐸᐳ(that Intᐸᐳ, dict PlusDict) Intᐸᐳ { return that };

func PlusWrapper(this Num, that Num, dict PlusDict) Num {
	return this.(Intᐸᐳ).Plusᐸᐳ(that.(Intᐸᐳ), dict)
}
func (this Intᐸᐳ) Plus___Int___Int__() Top { return this };

type TimesTwoDict struct {
	PlusFunc func (this Num, that Num, dict PlusDict) Num
	PlusFuncDict PlusDict
}

func (this TimesTwoFactoryᐸᐳ) TimesTwo(x Num, dict TimesTwoDict) Num {
	plusFunc := dict.PlusFunc
	return plusFunc(x, x, dict.PlusFuncDict)
};

func TimesTwoWrapper(this TimesTwoFactoryᐸᐳ, that Num, dict TimesTwoDict) Num {
	return this.TimesTwo(that, dict)
}

func (this TimesTwoFactoryᐸᐳ) TimesTwo__β1_Num_β1___β1_β1() Top { return this };
type TimesTwoFactoryᐸᐳ struct {};

type TimesThreeDict struct {
	PlusFunc func (this Num, that Num, dict PlusDict) Num
	TimesTwoFunc func (this TimesTwoFactoryᐸᐳ, that Num, dict TimesTwoDict) Num
	PlusFuncDict PlusDict
	TimesTwoFuncDict TimesTwoDict
}

func (this TimesThreeFactoryᐸᐳ) TimesThree(x Num, dict TimesThreeDict) Num {
	plusFunc := dict.PlusFunc;
	timesTwo := dict.TimesTwoFunc;
	tmp := timesTwo(TimesTwoFactoryᐸᐳ{}, x, dict.TimesTwoFuncDict)
	return plusFunc(tmp, tmp, dict.PlusFuncDict)
};

func TimesThreeWrapper() {

}

func (this TimesThreeFactoryᐸᐳ) TimesThree__β1_Num_β1___β1_β1() Top { return this };
type TimesThreeFactoryᐸᐳ struct {};
type Top interface {};


func main() {
	dict := TimesThreeDict{
		PlusFunc:         PlusWrapper,//we can use an anonymous function here in std. go
		TimesTwoFunc:     TimesTwoWrapper,
		PlusFuncDict:     PlusDict{},
		TimesTwoFuncDict: TimesTwoDict{
			PlusFunc:     PlusWrapper,
			PlusFuncDict: PlusDict{},
		},
	}
	_ = TimesThreeFactoryᐸᐳ{}.TimesThree(Intᐸᐳ{}, dict)
}
