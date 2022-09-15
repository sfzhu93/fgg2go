package main

type Any interface {};
type Function interface {
        //Apply(x Any) Any
};
type incr struct { n int };
func (this incr) Apply(x int) int {
        return x + this.n
};
type pos struct {};  // We already have IsNeg, though
func (this pos) Apply(x int) bool {
        return x > 0
};

type compose struct {//(type a Any(), b Any(), c Any())
        f Function;//(a, b);
        g Function;//(b, c)
};
func (this compose) Apply(x Any, dict composeDict) Any {
        return dict.this_g_Apply(this.g, dict.this_f_Apply(this.f, x))
};

type Function_Int_Int interface {
Apply(x int) int
}

type Function_Int_Bool interface {
Apply(x int) bool
}

type composeDict struct {
    this_g_Apply func (this Function, x Any) Any;
    this_f_Apply func (this Function, x Any) Any;
    //...
}

func main() {
        /*var f Function(int, bool) =
                                compose(int, int, bool){incr{-5}, pos{}}
        var b bool = f.Apply(3)*/
        dict := composeDict {
                this_g_Apply: func (this Function, x Any) Any {
                        return this.(Function_Int_Bool).Apply(x.(int))
                },
                this_f_Apply: func (this Function, x Any) Any {
                        return this.(Function_Int_Int).Apply(x.(int))
                },
        }
        print(compose{incr{-5} , pos{}}.Apply(3, dict).(bool))
}
