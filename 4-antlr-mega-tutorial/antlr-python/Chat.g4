grammar Chat;

chat     : line+ EOF ;
line     : name command message NEWLINE ;
message  : (emoticon | link | color | mention | WORD | WHITESPACE)+ ;
name     : WORD WHITESPACE;
command  : (SAYS | SHOUTS) ':' WHITESPACE ;
emoticon : ':' '-'? ')' | ':' '-'? '(' ;
link     : TEXT TEXT ;
color    : '/' WORD '/' message '/';
mention  : '@' WORD ;

fragment A         : [aA] ;
fragment S         : [sS] ;
fragment Y         : [yY] ;
fragment H         : [hH] ;
fragment O         : [oO] ;
fragment U         : [uU] ;
fragment T         : [tT] ;
fragment LOWERCASE : [a-z] ;
fragment UPPERCASE : [A-Z] ;

SAYS       : S A Y S ;
SHOUTS     : S H O U T S ;
WORD       : (LOWERCASE | UPPERCASE | '_')+ ;
WHITESPACE : (' ' | '\t')+ ;
NEWLINE    : ('\r'? '\n' | '\r')+ ;
TEXT       : ('[' | '(') ~[\])]+ (']' | ')');
