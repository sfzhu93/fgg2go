
grammar FGG;

/* Keywords */

FUNC: 'func';
INTERFACE: 'interface';
MAIN: 'main';
PACKAGE: 'package';
RETURN: 'return';
STRUCT: 'struct';
TYPE: 'type';

IMPORT: 'import';
FMT: 'fmt';
PRINTF: 'Printf';
SPRINTF: 'Sprintf';

CLOSE     : 'close';
GORTN     : 'go'; //goroutine
MAKE      : 'make';
SELECT    : 'select';
DEFAULT   : 'default';
CASE      : 'case';
/* Tokens */

fragment LETTER: ('a' .. 'z')
	| ('A' .. 'Z')
	| 'α' // For FGR deserialization
	| 'β';
fragment DIGIT: ('0' .. '9');
NAME: (LETTER | '_') (LETTER | '_' | DIGIT)*;
WHITESPACE: [ \r\n\t]+ -> skip;
COMMENT: '/*' .*? '*/' -> channel(HIDDEN);
LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
STRING:
	'"' (
		LETTER
		| DIGIT
		| ' '
		| '.'
		| ','
		| '_'
		| '%'
		| '#'
		| '('
		| ')'
		| '+'
		| '-'
	)* '"';

/* Rules */

// Conventions: "tag=" to distinguish repeat productions within a rule: comes out in field/getter
// names. "#tag" for cases within a rule: comes out as Context names (i.e., types). "plurals", e.g.,
// decls, used for sequences: comes out as "helper" Contexts, nodes that group up actual children
// underneath -- makes "adapting" easier.

channelTypeLit       : 'chan' '<-'| '<-' 'chan' | 'chan' ;  // The order is necessary for matching

typ: NAME # TypeParam | NAME '[' typs? ']' # TypeName | channelTypeLit typ # ChannelType;
typs: typ (',' typ)*;
typeFormals:
	'[' TYPE typeFDecls? ']'; // Refactored "(...)" into here
typeFDecls: typeFDecl (',' typeFDecl)*;
typeFDecl: NAME typ; // CHECKME: #TypeName ?
program:
	PACKAGE MAIN ';' (IMPORT STRING ';')? decls? FUNC MAIN '(' ')' '{' (
		'_' '=' expr
		| FMT '.' PRINTF '(' '"%#v"' ',' expr ')'
	) '}' EOF;
decls: ((typeDecl | methDecl) ';')+;
typeDecl:
	TYPE NAME typeFormals typeLit; // TODO: tag id=NAME, better for adapting (vs., index constants)

methBody   : RETURN expr                            # MethReturn
           | stmt ';' methBody                      # Sequence
           | FMT '.' SPRINTF '(' (STRING | '"%#v"') (',' | expr)* ')'  # Sprintf
           | SELECT '{' (CASE caseGuard ':' methBody)* (DEFAULT ':' methBody)? '}' #CaseSelect
           ;

stmt       : CLOSE '(' expr ')'                     # ChClose
           | expr '<-' expr                         # ChDispatch
           | GORTN recv=expr methCall_              # Goroutine
           | lvalue=NAME ':=' expr                  # Assignment
           ;

methCall_  : '.' NAME '[' targs = typs? ']' '(' args = exprs? ')'           # MethCallPrime //dirty hack for mutual left recursive
           ;

caseGuard  : expr '<-' expr                         # GuardDispatch
           | lvalue=NAME ':=' '<-' expr             # GuardReceive
           ;

methDecl:
	FUNC '(' recv = NAME typn = NAME typeFormals ')' sig '{' methBody '}';
typeLit:
	STRUCT '{' fieldDecls? '}'	# StructTypeLit
	| INTERFACE '{' specs? '}'	# InterfaceTypeLit;
fieldDecls: fieldDecl (';' fieldDecl)*;
fieldDecl: field = NAME typ;
specs: spec (';' spec)*;
spec:
	sig		# SigSpec
	| typ	# InterfaceSpec; // Must be a #TypeName, \tau_I -- refactor?
sig: meth = NAME typeFormals '(' params? ')' typ;
params: paramDecl (',' paramDecl)*;
paramDecl: vari = NAME typ;
expr:
	NAME																# Variable
	| typ '{' exprs? '}' /* typ is #TypeName, \tau_S */					# StructLit
	| expr '.' NAME														# Select
	| recv = expr methCall_                                             # Call
	| expr '.' '(' typ ')'												# Assert
	| MAKE '(' typ ')'                                                  # MakeChan
	| '<-' expr                                                         # ChanRecv
	;
exprs: expr (',' expr)*;

