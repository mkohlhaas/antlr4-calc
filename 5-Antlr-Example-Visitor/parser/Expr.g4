grammar Expr;

prog : (decl | expr)+ EOF;
decl : IDENT ':' INT_TYPE '=' NUM;
expr : expr '*' expr
     | expr '+' expr
     | IDENT
     | NUM;

IDENT    : [a-z] [a-zA-Z0-9]*;
NUM      : '0' | '-'? [1-9] [0-9]*;
INT_TYPE : 'INT';
COMMENT  : '--' ~ [\r\n]* -> skip;
WS       : [ \t\n]+ -> skip;
