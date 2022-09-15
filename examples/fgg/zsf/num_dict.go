package main;
type Intᐸᐳ struct {};
func (this Intᐸᐳ) Plusᐸᐳ(that Intᐸᐳ) Intᐸᐳ { return that };
func (this Intᐸᐳ) Plus___Int___Int__() Top { return this };
func (this TimesTwoFactoryᐸᐳ) TimesTwoᐸIntᐸᐳᐳ(x Intᐸᐳ) Intᐸᐳ { return x.Plusᐸᐳ(x) };
func (this TimesTwoFactoryᐸᐳ) TimesTwo__β1_Num_β1___β1_β1() Top { return this };
type TimesTwoFactoryᐸᐳ struct {};
type Top interface {};
func main() { _ = TimesTwoFactoryᐸᐳ{}.TimesTwoᐸIntᐸᐳᐳ(Intᐸᐳ{}) }
