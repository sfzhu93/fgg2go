//$ antlr4 -Dlanguage=Go -o parser/fg parser/FG.g4


// FG.g4
grammar FG;


/* Keywords */

FUNC      : 'func' ;
INTERFACE : 'interface' ;
MAIN      : 'main' ;
PACKAGE   : 'package' ;
RETURN    : 'return' ;
STRUCT    : 'struct' ;
TYPE      : 'type' ;

IMPORT    : 'import' ;
FMT       : 'fmt' ;
PRINTF    : 'Printf' ;
SPRINTF   : 'Sprintf' ;

CLOSE     : 'close';
GORTN     : 'go'; //goroutine
MAKE      : 'make';
SELECT    : 'select';
DEFAULT   : 'default';
CASE      : 'case';
/* Tokens */

fragment LETTER : ('a' .. 'z') | ('A' .. 'Z') | '\u03b1' | '\u03b2' ;
fragment DIGIT  : ('0' .. '9') ;
//fragment HACK   : 'ᐸ' | 'ᐳ' ;  // Doesn't seem to work?
fragment MONOM_HACK   : '\u1438' | '\u1433' | '\u1428' ;  // Hack for monom output
NAME            : (LETTER | '_' | MONOM_HACK) (LETTER | '_' | DIGIT | MONOM_HACK)* ;
WHITESPACE      : [ \r\n\t]+ -> skip ;
COMMENT         : '/*' .*? '*/' -> channel(HIDDEN) ;
LINE_COMMENT    : '//' ~[\r\n]* -> channel(HIDDEN) ;
STRING          : '"' (LETTER | DIGIT | ' ' | '.' | ',' | '_' | '%' | '#' | '(' | ')' | '+' | '-')* '"' ;


/* Rules */

// Conventions:
// "tag=" to distinguish repeat productions within a rule: comes out in
// field/getter names.
// "#tag" for cases within a rule: comes out as Context names (i.e., types).
// "plurals", e.g., decls, used for sequences: comes out as "helper" Contexts,
// nodes that group up actual children underneath -- makes "adapting" easier.

program    : PACKAGE MAIN ';'
             (IMPORT STRING ';')?
             decls? FUNC MAIN '(' ')' '{'
             ('_' '=' expr | FMT '.' PRINTF '(' '"%#v"' ',' expr ')')
             '}' EOF ;
decls      : ((typeDecl | methDecl) ';')+ ;
typeDecl   : TYPE NAME typeLit ;  // TODO: tag id=NAME, better for adapting (vs., index constants)

methBody   : RETURN expr                            # MethReturn
           | stmt ';' methBody                      # Sequence
           | FMT '.' SPRINTF '(' (STRING | '"%#v"') (',' | expr)* ')'  # Sprintf
           | SELECT '{' (CASE caseGuard ':' methBody)* (DEFAULT ':' methBody)? '}' #CaseSelect
           ;

methDecl   : FUNC '(' paramDecl ')' sig '{' methBody '}' ; //TODO: update for method body
typeLit    : STRUCT '{' fieldDecls? '}'             # StructTypeLit
           | INTERFACE '{' specs? '}'               # InterfaceTypeLit 
           ;

channelTypeLit       : 'chan' '<-'| '<-' 'chan' | 'chan' ;  // The order is necessary for matching
//TODO: maybe add parenthesis to allow more flexible association?
type_ : tyName=NAME | channelTypeLit type_;
           
fieldDecls : fieldDecl (';' fieldDecl)* ;
fieldDecl  : field=NAME typ=type_;
specs      : spec (';' spec)* ;
spec       : sig                                    # SigSpec
           | NAME                                   # InterfaceSpec
           ;
sig        : meth=NAME '(' params? ')' ret=NAME ;
params     : paramDecl (',' paramDecl)* ;
paramDecl  : vari=NAME typ=type_ ;

methCall_  : '.' NAME '(' args=exprs? ')'           # MethCallPrime //dirty hack for mutual left recursive
           ;

stmt       : CLOSE '(' expr ')'                     # ChClose
           | expr '<-' expr                         # ChDispatch
           | GORTN recv=expr methCall_              # Goroutine
           | lvalue=NAME ':=' expr                  # Assignment
           ;

caseGuard  : expr '<-' expr                         # GuardDispatch
           | lvalue=NAME ':=' '<-' expr             # GuardReceive
           ;

expr       : NAME                                   # Variable
           | NAME '{' exprs? '}'                    # StructLit
           | expr '.' NAME                          # Select
           | recv=expr methCall_                    # Call
           | expr '.' '(' NAME ')'                  # Assert
           | MAKE '(' type_ ')'                     # MakeChan
           | '<-' expr                              # ChanRecv
           ;//TODO: may need <assoc=right>
exprs      : expr (',' expr)* ;

