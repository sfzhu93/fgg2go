package main;

type Any interface {};
type Function interface {Apply(in Any) Any};
type Int struct {};

type Pipeline interface {
	Map(f Function) Pipeline
};

type Stream struct {
	in <-chan Any
};

func (this Stream) Map(f Function) Pipeline {
	out := make(chan Any);
	go this._Map(f, out);
	return Stream{out}
};

func (this Stream) _Map(f Function,
	out chan<- Any) Any {
	val := <-this.in;
	out <- f.Apply(val);
	return this._Map(f, out)
};

type StreamGen struct {};
type Pair struct {fst Any; snd Any};

func (this StreamGen) Gen(f Function,
	init Any) Pipeline {
	out := make(chan Any);
	go this._Gen(f, init, out);
	return Stream{out}
};

func (this StreamGen) _Gen(f Function,
	state Any, out chan<- Any) Any {
	next := f.Apply(state).(Pair);
	out <- next.fst;
	return this._Gen(f, next.snd, out)
};

type Add struct {y Int};
func (this Add) Apply(in Any) Any {
	return Pair{in, in}
};

type Top struct{};
func (this Top) _main() Top {
	f := Add{Int{}};
	pipeline := StreamGen{}.Gen(f, Int{}).Map(f).(Stream);
	x1 := <-pipeline.in;
	x2 := <-pipeline.in;
	return Top{}
};

func main() {
	_ = Top{}._main()
}