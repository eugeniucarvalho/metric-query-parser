/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
 /**
 TODO:
 SwitchCASE incompleto
 */

// : 
grammar Grammar; 

expression
    : handler EOF
    ;

handler
    : ID R_BRACKET_L arguments? R_BRACKET_R
    ;

arguments
    : argument ( R_COMA argument )*
    ;

argument
    : handler 
    | ID operator value
    ;

value
    : typeString
    | typeBool
    | typeInt
    | typeFloat
    ;
// : '"' (~'"'|'"')* '"'
typeString
    : R_SIGLE_QUOTA .*? R_SIGLE_QUOTA
    ;
typeBool
    : ( 'true'| 'false')
    ;
typeInt
    : T_INTEIRO
    ;
typeFloat
    : T_INTEIRO R_DOT T_INTEIRO
    ;
operator
    : R_NEGATE operator
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
    | R_NOT_BETWEEN
    ;

R_NEGATE                : '!';
R_SIGLE_QUOTA           : '\'';
R_BRACE_L               : '{';
R_BRACE_R               : '}';
R_BRACKET_L             : '(';
R_BRACKET_R             : ')';
R_PIPE                  : '|';
R_EQ                    : '=';
R_CONTAINS              : '~=';
R_START_WITH            : '|=';
R_END_WITH              : '=|';
R_IN                    : 'in';
R_GT                    : '>';
R_GTE                   : '>=';
R_LT                    : '<';
R_LTE                   : '<=';
R_BETWEEN               : '><';
R_NOT_BETWEEN           : '<>';
R_DOT                   : '.';
R_COMA                  : ',';
R_COLON                 : ':';
T_INTEIRO               : DIGITO+;

fragment    DIGITO      : [0-9] ;
fragment    CHARACTER   : [a-zA-Z_] ;
fragment    CHARACTER_UP: [A-Z];

ID                      : CHARACTER ( CHARACTER | DIGITO )* ;

R_LINE_COMMENT          : '//' .*? '\r'? '\n' -> skip ;  // Match "#" stuff '\n'
R_WS                    : ([\t\r\n]|' ')+ -> skip ;
R_COMMENT               : '/*' .*? '*/' -> skip ; // Match "/*" stuff "*/"