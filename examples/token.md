# I will use the token categories

KEYWORD: varC keyword (owner, borrow, borrow_mut, move, unsafe, opt, raw)
KEYWORD_C: C keyword (int, float, if, for, return, etc.)
IDENT: identifier
INT_LIT: integer literal
FLOAT_LIT: floating literal
CHAR_LIT: char literal
STR_LIT: string literal
OP: operator (+ - * / % ... )
PUNCT: punctuation (; , ( ) { } [ ] -> .)
PREPROC: preprocessor directive (#define ...)
COMMENT: comments (ignored by parser, but tokenizable)

# Expected Token Stream (Partial view, first 150+ tokens)

```c
PREPROC: #define
IDENT: MACRO
PUNCT: (
IDENT: x
PUNCT: ,
IDENT: y
PUNCT: )
PUNCT: (
IDENT: x
PUNCT: )
OP: +
PUNCT: ;
PREPROC: #define
IDENT: STR
PUNCT: "
IDENT: string
OP: \t
IDENT: with
OP: \n
IDENT: escapes
PUNCT: \"
PUNCT: "
PREPROC: #define
IDENT: HEX
INT_LIT: 0xDEADBEEF
PREPROC: #define
IDENT: BIN
INT_LIT: 0b101010
PREPROC: #define
IDENT: OCT
INT_LIT: 0755
PREPROC: #include
PUNCT: <
IDENT: stdio
PUNCT: >
PREPROC: #include
PUNCT: <
IDENT: stdlib
PUNCT: >
PREPROC: #include
PUNCT: <
IDENT: stdint
PUNCT: >

KEYWORD_C: struct
IDENT: Node
PUNCT: {
KEYWORD_C: int
IDENT: value
PUNCT: ;
KEYWORD: opt
KEYWORD: owner
KEYWORD_C: struct
IDENT: Node
PUNCT: *
IDENT: next
PUNCT: ;
PUNCT: }

KEYWORD: owner
KEYWORD_C: int
PUNCT: *
IDENT: make_value
PUNCT: (
KEYWORD: opt
KEYWORD_C: int
IDENT: seed
PUNCT: )
PUNCT: {

KEYWORD: owner
KEYWORD_C: int
PUNCT: *
IDENT: p
OP: =
PUNCT: (
KEYWORD: owner
KEYWORD_C: int
PUNCT: *)
IDENT: malloc
PUNCT: (
KEYWORD_C: sizeof
PUNCT: (
KEYWORD_C: int
PUNCT: )
PUNCT: )
PUNCT: ;

PUNCT: *
IDENT: p
OP: =
IDENT: seed
OP: ?
IDENT: seed
OP: :
INT_LIT: 123
PUNCT: ;

KEYWORD_C: return
IDENT: p
PUNCT: ;
PUNCT: }

KEYWORD_C: void
IDENT: print_value
PUNCT: (
KEYWORD: borrow
KEYWORD_C: int
PUNCT: *
IDENT: p
PUNCT: )
PUNCT: {
IDENT: printf
PUNCT: (
STR_LIT: "%d\n"
PUNCT: ,
IDENT: *
IDENT: p
PUNCT: )
PUNCT: ;
PUNCT: }

KEYWORD_C: void
IDENT: mutate
PUNCT: (
KEYWORD: borrow_mut
KEYWORD_C: int
PUNCT: *
IDENT: p
PUNCT: )
PUNCT: {
OP: ++
PUNCT: (
OP: *
IDENT: p
PUNCT: )
PUNCT: ;
PUNCT: }

KEYWORD_C: void
IDENT: write_raw
PUNCT: (
KEYWORD: raw
KEYWORD_C: int
PUNCT: *
IDENT: p
PUNCT: )
PUNCT: {
PUNCT: *
IDENT: p
OP: =
INT_LIT: 999
PUNCT: ;
PUNCT: }

KEYWORD_C: void
IDENT: consume
PUNCT: (
KEYWORD: owner
KEYWORD_C: int
PUNCT: *
IDENT: p
PUNCT: )
PUNCT: {
IDENT: free
PUNCT: (
IDENT: p
PUNCT: )
PUNCT: ;
PUNCT: }

KEYWORD_C: int
IDENT: operators
PUNCT: (
KEYWORD_C: int
IDENT: a
PUNCT: ,
KEYWORD_C: int
IDENT: b
PUNCT: )
PUNCT: {
KEYWORD_C: int
IDENT: r
OP: =
INT_LIT: 0
PUNCT: ;
IDENT: r
OP: +=
IDENT: a
OP: +
IDENT: b
PUNCT: ;
IDENT: r
OP: -=
IDENT: a
OP: -
IDENT: b
PUNCT: ;
```


