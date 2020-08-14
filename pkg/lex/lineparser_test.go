package lex

import "testing"

func TestDefaultLineParser(t *testing.T) {
	// code snippet with no whitespace
	noWhiteSpaceCode := "a:=5;a+=1;a<b;"
	// same code snippet with spaces and tab characters as whitespace
	whiteSpacedCode := "  a  :=	5;	a   +=  1  ; a <	b	;	"

	code1, code2 := DefaultLineParser(noWhiteSpaceCode), DefaultLineParser(whiteSpacedCode)

	if len(code1) != len(code2) {
		t.Error("default line parser did not ignore whitespace properly(1)")
		t.FailNow()
	}

	// test that types and values match
	for i := range code1 {
		v1, v2 := code1[i], code2[i]
		if v1.Value != v2.Value {
			t.Error("default line parser did not parse token values properly")
			t.FailNow()
		}
		if v1.Type != v2.Type {
			t.Error("default line parser did not parse token types properly")
			t.FailNow()
		}
	}
}
