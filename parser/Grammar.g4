grammar Grammar;

expression: handler EOF;

handler: ID R_BRACKET_L arguments? R_BRACKET_R;

arguments: argument ( R_COMA argument)*;

argument: handler | ID operator value;

value: typeString | typeBool | typeInt | typeFloat | typeHandler;

typeHandler: handler;
typeString: TypeString;
TypeString: '"' ( '\\' (~'\\' | '\\') | ~[\\"])* '"';
typeBool: ( 'true' | 'false');
typeInt: T_INTEIRO;
typeFloat: T_INTEIRO R_DOT T_INTEIRO;
operator:
	R_NEGATE operator
	| R_EQ
	| R_CONTAINS
	| R_START_WITH
	| R_END_WITH
	| R_IN
	| R_GT
	| R_GTE
	| R_LT
	| R_LTE
	| R_BETWEEN
	| R_NOT_BETWEEN;

R_NEGATE: '!';
R_SIGLE_QUOTA: '\'';
R_BRACE_L: '{';
R_BRACE_R: '}';
R_BRACKET_L: '(';
R_BRACKET_R: ')';
R_PIPE: '|';
R_EQ: '=';
R_CONTAINS: '~=';
R_START_WITH: '|=';
R_END_WITH: '=|';
R_IN: 'in';
R_GT: '>';
R_GTE: '>=';
R_LT: '<';
R_LTE: '<=';
R_BETWEEN: '><';
R_NOT_BETWEEN: '<>';
R_DOT: '.';
R_COMA: ',';
R_COLON: ':';
T_INTEIRO: DIGITO+;
ID: CHARACTER ( CHARACTER | DIGITO)*;
fragment STRING_CHARS: STRING_CHAR+;
fragment STRING_CHAR: ~["];

fragment DIGITO: [0-9];
fragment CHARACTER: [a-zA-Z_];
WS: [ \t\r\n]+ -> skip;
