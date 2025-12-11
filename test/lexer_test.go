package lexer_test

import (
	"github.com/VarunSharma3520/varC/internal/lexer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenTokens_SingleKeyword(t *testing.T) {
	src := `owner`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"type": "keyword"`)
	assert.Contains(t, out, `"value": "owner"`)
}

func TestFlattenTokens_IdentifierNotKeyword(t *testing.T) {
	src := `ownerX`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"type": "identifier"`)
	assert.Contains(t, out, `"value": "ownerX"`)
}

func TestFlattenTokens_IntegerLiteral(t *testing.T) {
	src := `123`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"value": "123"`)
}

func TestFlattenTokens_HexLiteral(t *testing.T) {
	src := `0xDEADBEEF`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"value": "0xDEADBEEF"`)
}

func TestFlattenTokens_BinLiteral(t *testing.T) {
	src := `0b101010`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"value": "0b101010"`)
}

func TestFlattenTokens_OctLiteral(t *testing.T) {
	src := `0755`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"value": "0755"`)
}

func TestFlattenTokens_FloatLiteral(t *testing.T) {
	src := `1.23f`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"value": "1.23f"`)
}

func TestFlattenTokens_ExpFloatLiteral(t *testing.T) {
	src := `1.2e-3`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"value": "1.2e-3"`)
}

func TestFlattenTokens_StringLiteral(t *testing.T) {
	src := `"hello\nworld"`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"type": "literal"`)
	assert.Contains(t, out, `"hello`)
}

func TestFlattenTokens_CharLiteral(t *testing.T) {
	src := `'\n'`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
  assert.Contains(t, out, `"value": "'\\n'"`)
}


func TestFlattenTokens_SingleLineComment(t *testing.T) {
	src := `// hello`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"type": "comment"`)
	assert.Contains(t, out, `hello`)
}

func TestFlattenTokens_BlockComment(t *testing.T) {
	src := `/* multi line */`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"type": "comment"`)
	assert.Contains(t, out, `multi`)
}

func TestFlattenTokens_PreprocessorDefine(t *testing.T) {
	src := `#define X 10`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `"type": "preprocessor"`)
	assert.Contains(t, out, `#define X 10`)
}

func TestFlattenTokens_PreprocessorInclude(t *testing.T) {
	src := `#include <stdio.h>`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)
	assert.Contains(t, out, `#include <stdio.h>`)
}

func TestFlattenTokens_ArithmeticOperators(t *testing.T) {
	src := `a+b-c*d/e%f`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)

	assert.Contains(t, out, `"value": "+"`)
	assert.Contains(t, out, `"value": "-"`)
	assert.Contains(t, out, `"value": "*"`)
	assert.Contains(t, out, `"value": "/"`)
	assert.Contains(t, out, `"value": "%"`)
}

func TestFlattenTokens_ComparisonOperators(t *testing.T) {
	src := `a==b a!=b a<=b a>=b a<b a>b`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)

	assert.Contains(t, out, `"value": "=="`)
	assert.Contains(t, out, `"value": "!="`)
	assert.Contains(t, out, `"value": "<="`)
	assert.Contains(t, out, `"value": ">="`)
	assert.Contains(t, out, `"value": "<"`)
	assert.Contains(t, out, `"value": ">"`)
}

func TestFlattenTokens_LogicalOperators(t *testing.T) {
	src := `a&&b || !c`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)

	assert.Contains(t, out, `"value": "&&"`)
	assert.Contains(t, out, `"value": "||"`)
	assert.Contains(t, out, `"value": "!"`)
}

func TestFlattenTokens_ShiftOperators(t *testing.T) {
	src := `a<<2 b>>1`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)

	assert.Contains(t, out, `"value": "<<"`)
	assert.Contains(t, out, `"value": ">>"`)
}

func TestFlattenTokens_TernaryOperator(t *testing.T) {
	src := `a ? b : c`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)

	assert.Contains(t, out, `"value": "?"`)
	assert.Contains(t, out, `"value": ":"`)
}

func TestFlattenTokens_IncrementDecrement(t *testing.T) {
	src := `a++ --b`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)

	assert.Contains(t, out, `"value": "++"`)
	assert.Contains(t, out, `"value": "--"`)
}

func TestFlattenTokens_PointerSyntax(t *testing.T) {
	src := `int *p = &x;`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)

	assert.Contains(t, out, `"value": "*"`)
	assert.Contains(t, out, `"value": "&"`)
}

func TestFlattenTokens_ArraySyntax(t *testing.T) {
	src := `int a[10];`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)

	assert.Contains(t, out, `"value": "["`)
	assert.Contains(t, out, `"value": "]"`)
}

func TestFlattenTokens_FunctionDeclaration(t *testing.T) {
	src := `int add(int a, int b) { return a+b; }`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)

	assert.Contains(t, out, `"value": "add"`)
	assert.Contains(t, out, `"value": "return"`)
}

func TestFlattenTokens_MoveKeyword(t *testing.T) {
	src := `move(x)`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)

	assert.Contains(t, out, `"value": "move"`)
}

func TestFlattenTokens_UnsafeBlock(t *testing.T) {
	src := `unsafe { raw int* p; }`
	out, err := lexer.FlattenTokens(src)
	assert.NoError(t, err)

	assert.Contains(t, out, `"value": "unsafe"`)
	assert.Contains(t, out, `"value": "raw"`)
}

func TestFlattenTokens_InvalidInput(t *testing.T) {
	_, err := lexer.FlattenTokens("$$$")
	assert.Error(t, err)
}
