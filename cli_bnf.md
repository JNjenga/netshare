<expression> ::= <command> <args>
<command>    ::= "ls" | "cd" | "cp"
<args>       ::= <arg> | <arg> EOF
<arg>        ::= <string>
