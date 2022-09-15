package dictpass

import (
	"github.com/sfzhu93/fgg2go/internal/fg"
	"github.com/sfzhu93/fgg2go/internal/fgg"
)

func AuxBracket(s fgg.Name) fg.Name {
	return s + "ǂ"
}

//fgg.TParam <: fgg.Type
//fgg.TNamed <: fgg.Type
func AuxRootType(ty fgg.Type) fg.TypeBase {
	switch ty_ := ty.(type) {
	case fgg.TNamed:
		return fg.Type(AuxBracket(ty_.TName))
	/*case fgg.TParam:
		return fg.Type("Any")
	case fgg.ChannelType:
		return fg.ChannelType{
			ElementType: AuxRootType(ty_.ElementType),
			ChType:      ty_.ChType,
		}*/
	default:
		panic("Only TNamed is used when channel is not implemented")
	}
}

func AuxTypeDict(ty fg.TypeBase) fg.Type {
	switch ty_ := ty.(type) {
	case fg.ChannelType:
		panic("no dict for channel types")
	case fg.Type:
		return fg.Type(AuxBracket(ty_.String() + "Dict")) //AuxRootType(ty)
	default:
		panic("not implemented type.")
	}

}
