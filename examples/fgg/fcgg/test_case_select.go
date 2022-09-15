package main

func (this World) select_parse() World {
	select {
	case y <- this:
		select {
		case this <- this:
			return World{}
		}
	case x := <-y:
		select {
		default:
			return World{}
		}
	default:
		return World{}
	}
}

func main() {
	_ = make(<-chan World)

}
