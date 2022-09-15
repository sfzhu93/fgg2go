package fgg

import (
	"fmt"

	"github.com/sfzhu93/fgg2go/internal/base"
	"github.com/sfzhu93/fgg2go/internal/fg"
)

type fg2fgg struct {
	fgProg  fg.FGProgram
	fggProg FGGProgram
}

// FromFG converts a FG program prog into a FGG program
// with empty type parameters
func FromFG(prog fg.FGProgram) (FGGProgram, error) {
	c := &fg2fgg{fgProg: prog}
	if err := c.convert(); err != nil {
		return FGGProgram{}, err
	}
	return c.fggProg, nil
}

func (c *fg2fgg) convert() error {
	// convert toplevel declarations
	for _, decl := range c.fgProg.GetDecls() {
		switch decl := decl.(type) {
		case fg.STypeLit:
			sTypeLit, err := c.convertSTypeLit(decl)
			if err != nil {
				return err
			}
			c.fggProg.decls = append(c.fggProg.decls, sTypeLit)

		case fg.ITypeLit:
			iTypeLit, err := c.convertITypeLit(decl)
			if err != nil {
				return err
			}
			c.fggProg.decls = append(c.fggProg.decls, iTypeLit)

		case fg.MethDecl:
			mDecl, err := c.convertMDecl(decl)
			if err != nil {
				return err
			}
			c.fggProg.decls = append(c.fggProg.decls, mDecl)

		default:
			return fmt.Errorf("unknown declaration type: %T", decl)
		}
	}

	expr, err := c.convertExpr(c.fgProg.GetMain())
	if err != nil {
		return err
	}
	c.fggProg.e_main = expr

	return nil
}

// convertType converts a plain type to a parameterised type
func (c *fg2fgg) convertType(t fg.TypeBase) (Name, BigPsi) {
	return Name(t.String()), BigPsi{TFormals: nil} // 0 formal parameters
}

func (c *fg2fgg) convertSTypeLit(s fg.STypeLit) (STypeLit, error) {
	typeName, typeFormals := c.convertType(s.GetType())

	var fieldDecls []FieldDecl
	for _, f := range s.GetFieldDecls() {
		fd, err := c.convertFieldDecl(f)
		if err != nil {
			return STypeLit{}, err
		}
		fieldDecls = append(fieldDecls, fd)
	}

	return STypeLit{T_name: typeName, Psi: typeFormals, fDecls: fieldDecls}, nil
}

func (c *fg2fgg) convertITypeLit(i fg.ITypeLit) (ITypeLit, error) {
	typeName, typeFormals := c.convertType(i.GetType())

	var specs []Spec
	for _, s := range i.GetSpecs() {
		sig := s.(fg.Sig)
		var paramDecls []ParamDecl
		for _, p := range sig.GetParamDecls() {
			pd, err := c.convertParamDecl(p)
			if err != nil {
				return ITypeLit{}, nil
			}
			paramDecls = append(paramDecls, pd)
		}
		retTypeName, _ := c.convertType(sig.GetReturn())

		specs = append(specs, Sig{
			meth:   Name(sig.GetMethod()),
			Psi:    BigPsi{TFormals: nil},
			pDecls: paramDecls,
			u_ret:  TNamed{TName: retTypeName},
		})
	}

	return ITypeLit{t_I: typeName, Psi: typeFormals, specs: specs}, nil
}

func (c *fg2fgg) convertFieldDecl(fd fg.FieldDecl) (FieldDecl, error) {
	typeName, _ := c.convertType(fd.GetType())
	return FieldDecl{field: fd.GetName(), u: TNamed{TName: typeName}}, nil
}

func (c *fg2fgg) convertParamDecl(pd fg.ParamDecl) (ParamDecl, error) {
	typeName, _ := c.convertType(pd.GetType())
	return ParamDecl{Name: pd.GetName(), U: TNamed{TName: typeName}}, nil
}

func (c *fg2fgg) convertMDecl(md fg.MethDecl) (MethDecl, error) {
	recvTypeName, recvTypeFormals := c.convertType(md.GetReceiver().GetType())

	var paramDecls []ParamDecl
	for _, p := range md.GetParamDecls() {
		pd, err := c.convertParamDecl(p)
		if err != nil {
			return MethDecl{}, err
		}
		paramDecls = append(paramDecls, pd)
	}

	retTypeName, _ := c.convertType(md.GetReturn())
	methImpl, err := c.convertExpr(md.GetBody())
	if err != nil {
		return MethDecl{}, err
	}

	return MethDecl{
		XRecv:    md.GetReceiver().GetName(),
		TRecv:    recvTypeName,
		Psi_recv: recvTypeFormals,
		name:     Name(md.GetName()),
		Psi_meth: BigPsi{}, // empty parameter
		pDecls:   paramDecls,
		U_ret:    TNamed{TName: retTypeName},
		E_body:   methImpl,
	}, nil
}

func (c *fg2fgg) convertExpr(expr base.Expr) (FGGExpr, error) {
	switch expr := expr.(type) {
	case fg.Variable:
		return Variable{name: expr.String()}, nil

	case fg.StructLit:
		sLitExpr, err := c.convertStructLit(expr)
		if err != nil {
			return nil, err
		}
		return sLitExpr, nil

	case fg.Call:
		callExpr, err := c.convertCall(expr)
		if err != nil {
			return nil, err
		}
		return callExpr, nil

	case fg.Select:
		selExpr, err := c.convertExpr(expr.GetExpr())
		if err != nil {
			return nil, err
		}
		return Select{E_S: selExpr, field: Name(expr.GetField())}, nil

	case fg.Assert:
		assertExpr, err := c.convertExpr(expr.GetExpr())
		if err != nil {
			return nil, err
		}
		assType, _ := c.convertType(expr.GetType())
		return Assert{E_I: assertExpr, U_cast: TNamed{TName: assType}}, nil
	}

	return nil, fmt.Errorf("unknown expression type: %T", expr)
}

func (c *fg2fgg) convertStructLit(sLit fg.StructLit) (StructLit, error) {
	structType, _ := c.convertType(sLit.GetType())

	var es []FGGExpr
	for _, expr := range sLit.GetElems() {
		fieldExpr, err := c.convertExpr(expr)
		if err != nil {
			return StructLit{}, err
		}
		es = append(es, fieldExpr)
	}

	return StructLit{U_S: TNamed{TName: structType}, Elems: es}, nil
}

func (c *fg2fgg) convertCall(call fg.Call) (Call, error) {
	e, err := c.convertExpr(call.GetReceiver())
	if err != nil {
		return Call{}, err
	}

	var args []FGGExpr
	for _, arg := range call.GetArgs() {
		argExpr, err := c.convertExpr(arg)
		if err != nil {
			return Call{}, err
		}
		args = append(args, argExpr)
	}

	return Call{E_recv: e, Meth: Name(call.GetMethod()), Args: args}, nil
}
