// Generated from parser/FGG.g4 by ANTLR 4.7.

package parser // FGG

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFGGListener is a complete listener for a parse tree produced by FGGParser.
type BaseFGGListener struct{}

var _ FGGListener = &BaseFGGListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFGGListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFGGListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFGGListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFGGListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterChannelTypeLit is called when production channelTypeLit is entered.
func (s *BaseFGGListener) EnterChannelTypeLit(ctx *ChannelTypeLitContext) {}

// ExitChannelTypeLit is called when production channelTypeLit is exited.
func (s *BaseFGGListener) ExitChannelTypeLit(ctx *ChannelTypeLitContext) {}

// EnterTypeParam is called when production TypeParam is entered.
func (s *BaseFGGListener) EnterTypeParam(ctx *TypeParamContext) {}

// ExitTypeParam is called when production TypeParam is exited.
func (s *BaseFGGListener) ExitTypeParam(ctx *TypeParamContext) {}

// EnterTypeName is called when production TypeName is entered.
func (s *BaseFGGListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production TypeName is exited.
func (s *BaseFGGListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterChannelType is called when production ChannelType is entered.
func (s *BaseFGGListener) EnterChannelType(ctx *ChannelTypeContext) {}

// ExitChannelType is called when production ChannelType is exited.
func (s *BaseFGGListener) ExitChannelType(ctx *ChannelTypeContext) {}

// EnterTyps is called when production typs is entered.
func (s *BaseFGGListener) EnterTyps(ctx *TypsContext) {}

// ExitTyps is called when production typs is exited.
func (s *BaseFGGListener) ExitTyps(ctx *TypsContext) {}

// EnterTypeFormals is called when production typeFormals is entered.
func (s *BaseFGGListener) EnterTypeFormals(ctx *TypeFormalsContext) {}

// ExitTypeFormals is called when production typeFormals is exited.
func (s *BaseFGGListener) ExitTypeFormals(ctx *TypeFormalsContext) {}

// EnterTypeFDecls is called when production typeFDecls is entered.
func (s *BaseFGGListener) EnterTypeFDecls(ctx *TypeFDeclsContext) {}

// ExitTypeFDecls is called when production typeFDecls is exited.
func (s *BaseFGGListener) ExitTypeFDecls(ctx *TypeFDeclsContext) {}

// EnterTypeFDecl is called when production typeFDecl is entered.
func (s *BaseFGGListener) EnterTypeFDecl(ctx *TypeFDeclContext) {}

// ExitTypeFDecl is called when production typeFDecl is exited.
func (s *BaseFGGListener) ExitTypeFDecl(ctx *TypeFDeclContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseFGGListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseFGGListener) ExitProgram(ctx *ProgramContext) {}

// EnterDecls is called when production decls is entered.
func (s *BaseFGGListener) EnterDecls(ctx *DeclsContext) {}

// ExitDecls is called when production decls is exited.
func (s *BaseFGGListener) ExitDecls(ctx *DeclsContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseFGGListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseFGGListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterMethReturn is called when production MethReturn is entered.
func (s *BaseFGGListener) EnterMethReturn(ctx *MethReturnContext) {}

// ExitMethReturn is called when production MethReturn is exited.
func (s *BaseFGGListener) ExitMethReturn(ctx *MethReturnContext) {}

// EnterSequence is called when production Sequence is entered.
func (s *BaseFGGListener) EnterSequence(ctx *SequenceContext) {}

// ExitSequence is called when production Sequence is exited.
func (s *BaseFGGListener) ExitSequence(ctx *SequenceContext) {}

// EnterSprintf is called when production Sprintf is entered.
func (s *BaseFGGListener) EnterSprintf(ctx *SprintfContext) {}

// ExitSprintf is called when production Sprintf is exited.
func (s *BaseFGGListener) ExitSprintf(ctx *SprintfContext) {}

// EnterCaseSelect is called when production CaseSelect is entered.
func (s *BaseFGGListener) EnterCaseSelect(ctx *CaseSelectContext) {}

// ExitCaseSelect is called when production CaseSelect is exited.
func (s *BaseFGGListener) ExitCaseSelect(ctx *CaseSelectContext) {}

// EnterChClose is called when production ChClose is entered.
func (s *BaseFGGListener) EnterChClose(ctx *ChCloseContext) {}

// ExitChClose is called when production ChClose is exited.
func (s *BaseFGGListener) ExitChClose(ctx *ChCloseContext) {}

// EnterChDispatch is called when production ChDispatch is entered.
func (s *BaseFGGListener) EnterChDispatch(ctx *ChDispatchContext) {}

// ExitChDispatch is called when production ChDispatch is exited.
func (s *BaseFGGListener) ExitChDispatch(ctx *ChDispatchContext) {}

// EnterGoroutine is called when production Goroutine is entered.
func (s *BaseFGGListener) EnterGoroutine(ctx *GoroutineContext) {}

// ExitGoroutine is called when production Goroutine is exited.
func (s *BaseFGGListener) ExitGoroutine(ctx *GoroutineContext) {}

// EnterAssignment is called when production Assignment is entered.
func (s *BaseFGGListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production Assignment is exited.
func (s *BaseFGGListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterMethCallPrime is called when production MethCallPrime is entered.
func (s *BaseFGGListener) EnterMethCallPrime(ctx *MethCallPrimeContext) {}

// ExitMethCallPrime is called when production MethCallPrime is exited.
func (s *BaseFGGListener) ExitMethCallPrime(ctx *MethCallPrimeContext) {}

// EnterGuardDispatch is called when production GuardDispatch is entered.
func (s *BaseFGGListener) EnterGuardDispatch(ctx *GuardDispatchContext) {}

// ExitGuardDispatch is called when production GuardDispatch is exited.
func (s *BaseFGGListener) ExitGuardDispatch(ctx *GuardDispatchContext) {}

// EnterGuardReceive is called when production GuardReceive is entered.
func (s *BaseFGGListener) EnterGuardReceive(ctx *GuardReceiveContext) {}

// ExitGuardReceive is called when production GuardReceive is exited.
func (s *BaseFGGListener) ExitGuardReceive(ctx *GuardReceiveContext) {}

// EnterMethDecl is called when production methDecl is entered.
func (s *BaseFGGListener) EnterMethDecl(ctx *MethDeclContext) {}

// ExitMethDecl is called when production methDecl is exited.
func (s *BaseFGGListener) ExitMethDecl(ctx *MethDeclContext) {}

// EnterStructTypeLit is called when production StructTypeLit is entered.
func (s *BaseFGGListener) EnterStructTypeLit(ctx *StructTypeLitContext) {}

// ExitStructTypeLit is called when production StructTypeLit is exited.
func (s *BaseFGGListener) ExitStructTypeLit(ctx *StructTypeLitContext) {}

// EnterInterfaceTypeLit is called when production InterfaceTypeLit is entered.
func (s *BaseFGGListener) EnterInterfaceTypeLit(ctx *InterfaceTypeLitContext) {}

// ExitInterfaceTypeLit is called when production InterfaceTypeLit is exited.
func (s *BaseFGGListener) ExitInterfaceTypeLit(ctx *InterfaceTypeLitContext) {}

// EnterFieldDecls is called when production fieldDecls is entered.
func (s *BaseFGGListener) EnterFieldDecls(ctx *FieldDeclsContext) {}

// ExitFieldDecls is called when production fieldDecls is exited.
func (s *BaseFGGListener) ExitFieldDecls(ctx *FieldDeclsContext) {}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *BaseFGGListener) EnterFieldDecl(ctx *FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *BaseFGGListener) ExitFieldDecl(ctx *FieldDeclContext) {}

// EnterSpecs is called when production specs is entered.
func (s *BaseFGGListener) EnterSpecs(ctx *SpecsContext) {}

// ExitSpecs is called when production specs is exited.
func (s *BaseFGGListener) ExitSpecs(ctx *SpecsContext) {}

// EnterSigSpec is called when production SigSpec is entered.
func (s *BaseFGGListener) EnterSigSpec(ctx *SigSpecContext) {}

// ExitSigSpec is called when production SigSpec is exited.
func (s *BaseFGGListener) ExitSigSpec(ctx *SigSpecContext) {}

// EnterInterfaceSpec is called when production InterfaceSpec is entered.
func (s *BaseFGGListener) EnterInterfaceSpec(ctx *InterfaceSpecContext) {}

// ExitInterfaceSpec is called when production InterfaceSpec is exited.
func (s *BaseFGGListener) ExitInterfaceSpec(ctx *InterfaceSpecContext) {}

// EnterSig is called when production sig is entered.
func (s *BaseFGGListener) EnterSig(ctx *SigContext) {}

// ExitSig is called when production sig is exited.
func (s *BaseFGGListener) ExitSig(ctx *SigContext) {}

// EnterParams is called when production params is entered.
func (s *BaseFGGListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseFGGListener) ExitParams(ctx *ParamsContext) {}

// EnterParamDecl is called when production paramDecl is entered.
func (s *BaseFGGListener) EnterParamDecl(ctx *ParamDeclContext) {}

// ExitParamDecl is called when production paramDecl is exited.
func (s *BaseFGGListener) ExitParamDecl(ctx *ParamDeclContext) {}

// EnterCall is called when production Call is entered.
func (s *BaseFGGListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production Call is exited.
func (s *BaseFGGListener) ExitCall(ctx *CallContext) {}

// EnterVariable is called when production Variable is entered.
func (s *BaseFGGListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production Variable is exited.
func (s *BaseFGGListener) ExitVariable(ctx *VariableContext) {}

// EnterAssert is called when production Assert is entered.
func (s *BaseFGGListener) EnterAssert(ctx *AssertContext) {}

// ExitAssert is called when production Assert is exited.
func (s *BaseFGGListener) ExitAssert(ctx *AssertContext) {}

// EnterSelect is called when production Select is entered.
func (s *BaseFGGListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production Select is exited.
func (s *BaseFGGListener) ExitSelect(ctx *SelectContext) {}

// EnterChanRecv is called when production ChanRecv is entered.
func (s *BaseFGGListener) EnterChanRecv(ctx *ChanRecvContext) {}

// ExitChanRecv is called when production ChanRecv is exited.
func (s *BaseFGGListener) ExitChanRecv(ctx *ChanRecvContext) {}

// EnterStructLit is called when production StructLit is entered.
func (s *BaseFGGListener) EnterStructLit(ctx *StructLitContext) {}

// ExitStructLit is called when production StructLit is exited.
func (s *BaseFGGListener) ExitStructLit(ctx *StructLitContext) {}

// EnterMakeChan is called when production MakeChan is entered.
func (s *BaseFGGListener) EnterMakeChan(ctx *MakeChanContext) {}

// ExitMakeChan is called when production MakeChan is exited.
func (s *BaseFGGListener) ExitMakeChan(ctx *MakeChanContext) {}

// EnterExprs is called when production exprs is entered.
func (s *BaseFGGListener) EnterExprs(ctx *ExprsContext) {}

// ExitExprs is called when production exprs is exited.
func (s *BaseFGGListener) ExitExprs(ctx *ExprsContext) {}
