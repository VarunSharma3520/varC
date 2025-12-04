package lexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
)

type Token struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
}

var keywords = map[string]bool{
	"owner": true, "borrow": true, "borrow_mut": true, "move": true,
	"unsafe": true, "opt": true, "raw": true,
	"typedef": true, "struct": true, "union": true, "enum": true,
	"static": true, "extern": true, "inline": true, "const": true,
	"volatile": true, "restrict": true, "unsigned": true,
	"long": true, "int": true, "float": true, "double": true,
	"char": true, "void": true, "if": true, "else": true,
	"for": true, "while": true, "do": true, "switch": true,
	"case": true, "break": true, "continue": true,
	"return": true, "goto": true, "sizeof": true,
}

type rule struct {
	Type string
	Re   *regexp.Regexp
}

var rules = []rule{
	{"whitespace", regexp.MustCompile(`\A[ \t\r\n]+`)},
	{"comment", regexp.MustCompile(`\A//[^\n]*`)},
	{"comment", regexp.MustCompile(`\A/\*[\s\S]*?\*/`)},
	{"preprocessor", regexp.MustCompile(`\A#[^\n]*`)},
	{"string", regexp.MustCompile(`\A"([^"\\]|\\.)*"`)},
	{"char", regexp.MustCompile(`\A'([^'\\]|\\.)*'`)},
	{"float", regexp.MustCompile(`\A(\d+\.\d*|\.\d+)([eE][+-]?\d+)?[fF]?|\A\d+[eE][+-]?\d+[fF]?`)},
	{"hex", regexp.MustCompile(`\A0x[0-9A-Fa-f]+`)},
	{"bin", regexp.MustCompile(`\A0b[01]+`)},
	{"oct", regexp.MustCompile(`\A0[0-7]+`)},
	{"int", regexp.MustCompile(`\A\d+`)},
	{"operator", regexp.MustCompile(`\A(\+\+|--|==|!=|<=|>=|&&|\|\||<<|>>|[-+*/%&|^=<>!~?:])`)},
	{"punctuation", regexp.MustCompile(`\A[()\[\]{};,\.]`)},
	{"identifier", regexp.MustCompile(`\A[a-zA-Z_][a-zA-Z0-9_]*`)},
}

func FlattenTokens(source string) (string, error) {
	var tokens []Token
	input := source
	line, col := 1, 1

	for len(input) > 0 {
		matched := false

		for _, r := range rules {
			loc := r.Re.FindStringIndex(input)
			if loc == nil {
				continue
			}

			val := input[loc[0]:loc[1]]
			input = input[loc[1]:]
			matched = true

			if r.Type == "whitespace" {
				// update line and column for whitespace
				for _, ch := range val {
					if ch == '\n' {
						line++
						col = 1
					} else {
						col++
					}
				}
				break
			}

			tType := r.Type
			if tType == "identifier" && keywords[val] {
				tType = "keyword"
			}
			if tType == "string" || tType == "char" ||
				tType == "int" || tType == "float" ||
				tType == "hex" || tType == "bin" || tType == "oct" {
				tType = "literal"
			}

			tokens = append(tokens, Token{
				Type:   tType,
				Value:  val,
				Line:   line,
				Column: col,
			})

			// update line and column after consuming token
			for _, ch := range val {
				if ch == '\n' {
					line++
					col = 1
				} else {
					col++
				}
			}
			break
		}

		if !matched {
			return "", fmt.Errorf("lexer error: unrecognized token near: %.20q", input)
		}
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)

	if err := enc.Encode(tokens); err != nil {
		return "", err
	}
	return buf.String(), nil
}
