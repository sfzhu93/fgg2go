package dyn

import (
	"github.com/sfzhu93/fgg2go/internal/fg"
	"github.com/sfzhu93/fgg2go/internal/fgg"
)

func AuxBracket(s fgg.Name) fg.Name {
	return s + "ǂ"
}

func AuxRequires(tFormals []fgg.TFormal) []fg.Name {
	ret := []fg.Name{}
	for _, formal := range tFormals {
		if tnamed, ok := formal.U_I.(fgg.TNamed); ok {
			ret = append(ret, fg.Name(tnamed.TName))
		}
	}
	return ret
}

func addRequires(formals []fgg.TFormal, requires map[fg.Name]struct{}) {
	for _, name := range AuxRequires(formals) {
		requires[name] = struct{}{}
	}
}

func getDict(ProgDecls []fgg.Decl) map[fg.Name]struct{} {
	ret := make(map[fg.Name]struct{})
	for _, decl := range ProgDecls {
		switch decl_ := decl.(type) {
		case fgg.MethDecl:
			addRequires(decl_.Psi_recv.TFormals, ret)
			addRequires(decl_.Psi_meth.TFormals, ret)
		case fgg.ITypeLit:
			addRequires(decl_.Psi.TFormals, ret)
			for _, spec := range decl_.GetSpecs() {
				if sig, ok := spec.(fgg.Sig); ok {
					addRequires(sig.Psi.TFormals, ret)
				} else {
					panic("should be a fgg.Sig here.")
				}
			}
		case fgg.STypeLit:
			break
		}
	}
	return ret
}

var dictCache map[fg.Name]struct{} = nil
