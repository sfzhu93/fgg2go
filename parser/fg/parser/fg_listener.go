// Generated from parser/FG.g4 by ANTLR 4.7.

package parser // FG

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FGListener is a complete listener for a parse tree produced by FGParser.
type FGListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDecls is called when entering the decls production.
	EnterDecls(c *DeclsContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterMethReturn is called when entering the MethReturn production.
	EnterMethReturn(c *MethReturnContext)

	// EnterSequence is called when entering the Sequence production.
	EnterSequence(c *SequenceContext)

	// EnterSprintf is called when entering the Sprintf production.
	EnterSprintf(c *SprintfContext)

	// EnterCaseSelect is called when entering the CaseSelect production.
	EnterCaseSelect(c *CaseSelectContext)

	// EnterMethDecl is called when entering the methDecl production.
	EnterMethDecl(c *MethDeclContext)

	// EnterStructTypeLit is called when entering the StructTypeLit production.
	EnterStructTypeLit(c *StructTypeLitContext)

	// EnterInterfaceTypeLit is called when entering the InterfaceTypeLit production.
	EnterInterfaceTypeLit(c *InterfaceTypeLitContext)

	// EnterChannelTypeLit is called when entering the channelTypeLit production.
	EnterChannelTypeLit(c *ChannelTypeLitContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterFieldDecls is called when entering the fieldDecls production.
	EnterFieldDecls(c *FieldDeclsContext)

	// EnterFieldDecl is called when entering the fieldDecl production.
	EnterFieldDecl(c *FieldDeclContext)

	// EnterSpecs is called when entering the specs production.
	EnterSpecs(c *SpecsContext)

	// EnterSigSpec is called when entering the SigSpec production.
	EnterSigSpec(c *SigSpecContext)

	// EnterInterfaceSpec is called when entering the InterfaceSpec production.
	EnterInterfaceSpec(c *InterfaceSpecContext)

	// EnterSig is called when entering the sig production.
	EnterSig(c *SigContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterParamDecl is called when entering the paramDecl production.
	EnterParamDecl(c *ParamDeclContext)

	// EnterMethCallPrime is called when entering the MethCallPrime production.
	EnterMethCallPrime(c *MethCallPrimeContext)

	// EnterChClose is called when entering the ChClose production.
	EnterChClose(c *ChCloseContext)

	// EnterChDispatch is called when entering the ChDispatch production.
	EnterChDispatch(c *ChDispatchContext)

	// EnterGoroutine is called when entering the Goroutine production.
	EnterGoroutine(c *GoroutineContext)

	// EnterAssignment is called when entering the Assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterGuardDispatch is called when entering the GuardDispatch production.
	EnterGuardDispatch(c *GuardDispatchContext)

	// EnterGuardReceive is called when entering the GuardReceive production.
	EnterGuardReceive(c *GuardReceiveContext)

	// EnterCall is called when entering the Call production.
	EnterCall(c *CallContext)

	// EnterVariable is called when entering the Variable production.
	EnterVariable(c *VariableContext)

	// EnterAssert is called when entering the Assert production.
	EnterAssert(c *AssertContext)

	// EnterSelect is called when entering the Select production.
	EnterSelect(c *SelectContext)

	// EnterChanRecv is called when entering the ChanRecv production.
	EnterChanRecv(c *ChanRecvContext)

	// EnterStructLit is called when entering the StructLit production.
	EnterStructLit(c *StructLitContext)

	// EnterMakeChan is called when entering the MakeChan production.
	EnterMakeChan(c *MakeChanContext)

	// EnterExprs is called when entering the exprs production.
	EnterExprs(c *ExprsContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDecls is called when exiting the decls production.
	ExitDecls(c *DeclsContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitMethReturn is called when exiting the MethReturn production.
	ExitMethReturn(c *MethReturnContext)

	// ExitSequence is called when exiting the Sequence production.
	ExitSequence(c *SequenceContext)

	// ExitSprintf is called when exiting the Sprintf production.
	ExitSprintf(c *SprintfContext)

	// ExitCaseSelect is called when exiting the CaseSelect production.
	ExitCaseSelect(c *CaseSelectContext)

	// ExitMethDecl is called when exiting the methDecl production.
	ExitMethDecl(c *MethDeclContext)

	// ExitStructTypeLit is called when exiting the StructTypeLit production.
	ExitStructTypeLit(c *StructTypeLitContext)

	// ExitInterfaceTypeLit is called when exiting the InterfaceTypeLit production.
	ExitInterfaceTypeLit(c *InterfaceTypeLitContext)

	// ExitChannelTypeLit is called when exiting the channelTypeLit production.
	ExitChannelTypeLit(c *ChannelTypeLitContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitFieldDecls is called when exiting the fieldDecls production.
	ExitFieldDecls(c *FieldDeclsContext)

	// ExitFieldDecl is called when exiting the fieldDecl production.
	ExitFieldDecl(c *FieldDeclContext)

	// ExitSpecs is called when exiting the specs production.
	ExitSpecs(c *SpecsContext)

	// ExitSigSpec is called when exiting the SigSpec production.
	ExitSigSpec(c *SigSpecContext)

	// ExitInterfaceSpec is called when exiting the InterfaceSpec production.
	ExitInterfaceSpec(c *InterfaceSpecContext)

	// ExitSig is called when exiting the sig production.
	ExitSig(c *SigContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitParamDecl is called when exiting the paramDecl production.
	ExitParamDecl(c *ParamDeclContext)

	// ExitMethCallPrime is called when exiting the MethCallPrime production.
	ExitMethCallPrime(c *MethCallPrimeContext)

	// ExitChClose is called when exiting the ChClose production.
	ExitChClose(c *ChCloseContext)

	// ExitChDispatch is called when exiting the ChDispatch production.
	ExitChDispatch(c *ChDispatchContext)

	// ExitGoroutine is called when exiting the Goroutine production.
	ExitGoroutine(c *GoroutineContext)

	// ExitAssignment is called when exiting the Assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitGuardDispatch is called when exiting the GuardDispatch production.
	ExitGuardDispatch(c *GuardDispatchContext)

	// ExitGuardReceive is called when exiting the GuardReceive production.
	ExitGuardReceive(c *GuardReceiveContext)

	// ExitCall is called when exiting the Call production.
	ExitCall(c *CallContext)

	// ExitVariable is called when exiting the Variable production.
	ExitVariable(c *VariableContext)

	// ExitAssert is called when exiting the Assert production.
	ExitAssert(c *AssertContext)

	// ExitSelect is called when exiting the Select production.
	ExitSelect(c *SelectContext)

	// ExitChanRecv is called when exiting the ChanRecv production.
	ExitChanRecv(c *ChanRecvContext)

	// ExitStructLit is called when exiting the StructLit production.
	ExitStructLit(c *StructLitContext)

	// ExitMakeChan is called when exiting the MakeChan production.
	ExitMakeChan(c *MakeChanContext)

	// ExitExprs is called when exiting the exprs production.
	ExitExprs(c *ExprsContext)
}
