package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	expected_statements := 3
	if len(program.Statements) != expected_statements {
		t.Fatalf("program.Statements Does not contain %d statements. got=%d", expected_statements, (program.Statements))
	}

	test := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range test {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for i, msg := range errors {
		t.Errorf("parser error %d: %q", i, msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		// %q is a verb for formatting strings safely escapes them
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value no '%s'. got=%s",
			name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return add(10, 10);
	return 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	expected_statements := 3
	if len(program.Statements) != expected_statements {
		t.Fatalf("program.Statements Does not contain %d statements. got=%d", expected_statements, (program.Statements))
	}

	// TODO: Implement and update
	test := []struct {
		expectedReturnValue ast.Expression
	}{
		{nil},
		{nil},
		{nil},
	}

	for i, tt := range test {
		stmt := program.Statements[i]
		if !testReturnStatement(t, stmt, tt.expectedReturnValue) {
			return
		}
	}
}

func testReturnStatement(t *testing.T, s ast.Statement, returnValue ast.Expression) bool {
	if s.TokenLiteral() != "return" {
		// %q is a verb for formatting strings safely escapes them
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
	}

	returnStmt, ok := s.(*ast.ReturnStatement)
	if !ok {
		t.Errorf("s not *ast.ReturnStatement. got=%T", s)
		return false
	}

	if returnStmt.ReturnValue != returnValue {
		t.Errorf("returnStmt.ReturnValue not '%s'. got=%s",
			returnStmt, returnValue)
		return false
	}

	return true
}
