package main;
//import "fmt";
type World struct{};

//func (this World) newgoroutine(y World) World {
//return this
//};

func (this World) doit(y World) World {
	go this.newgoroutine(y);
	t := make(chan World );
	t <- y;
    x := <- t;
	close(t);
	return x//our syntax forces an ending of return or select
};

func (this World) select_parse(y <-chan World) World {
	select {
	case y <- this:
		return this//method body always terminates with select
	case x := <- y:
		select {}
	default:
		return World{}
	}
};

func (this World) test_chan_ambiguity_parsing(y chan <-chan World) World {
	return World{}
};

func main() {
    _ = World{}.doit(World{})
    //fmt.Printf("%#v", World{}.hello())
}
