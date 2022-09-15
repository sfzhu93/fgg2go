//This program tries to mock Fig. 1 and Fig. 2 in paper walder88
//https://people.csail.mit.edu/dnj/teaching/6898/papers/wadler88.pdf
//This has been compiled by fgg. So it is less "syntactic sweet" than
//the code on the paper...

package main;

type Any interface {};

//------start defining a Num interface with an Int declaration------
type Num interface {
Plus (that Num) Num
};

type Int struct {
};

func (this Int) Plus(that Num) Num {
	return that//A fake plus operation! Peano numbers are too wordy, so just make it simple...
};
//------end defining a Num interface with an Int declaration------


//define a method that adds a number to itself, so the number is doubled.
//We call it "TimesTwo"...
func (this TimesTwoFactory) TimesTwo(x Num) Num {
	return x.Plus(x)
};

//Because FGG methods must have a receiver (𝑒 is necessary in 𝑒.𝑚(𝜏)(𝑒) ),
//We define a dummy struct as the receiver of TimesTwo function.
type TimesTwoFactory[type ] struct {

};


func main() {
	_ = TimesTwoFactory[]{}.TimesTwo[Int[]](Int[]{})
}