package fg

/*
import (
	"testing"

	"github.com/rhu1/fgg/internal/parser"  // TODO: internal package not visible from here?
)

func TestDistinctDecl(t *testing.T) {
	prog := MakeFgProgram(
		"type A struct{}", // [0] not ok, clash with [6]
		"type X struct{}",
		"func (x X) m1() X { return x }", // [2] not ok, clash with [4]
		"func (x X) m2() X { return x }", // ok
		"func (z X) m1() X { return z }", // [4] not OK, clash with [2]
		"type B struct{}",
		"type A interface{}", // [6] not ok, clash with [0]
		"A{}",
	)
	expectDistinct := []bool{false, true, false, true, false, true, false}

	var a parser.FGAdaptor
	ast := a.Parse(true, prog)
	if want, got := len(expectDistinct), len(ast.GetDecls()); want != got {
		t.Fatalf("expected %d Decls but got %d", want, got)
	}
	Decls := ast.GetDecls()
	for i := range expectDistinct {
		t.Logf("expected unique: %t %s", expectDistinct[i], Decls[i])
		if want, got := expectDistinct[i], isDistinctDecl(Decls[i], Decls); want != got {
			t.Fatalf("decl[%d] expected unique decl = %t", i, want)
		}
	}
}
*/
