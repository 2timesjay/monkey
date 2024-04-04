package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

// func (p *Parser) ParseProgram() *ast.Program {
// 	program = newProgramASTNode()

// 	advanceTokens()

// 	// Until the lexer reaches EOF, parse statements and add them to the program
// 	// If a parsed token is not a statement, do not add to program
// 	for currentToken() != EOF_TOKEN {
// 		statement = null

// 		if currentToken() == LET_TOKEN {
// 			statement = parseLetStatement()
// 		} else if currentToken() == RETURN_TOKEN {
// 			statement = parseReturnStatement()
// 		} else if currentToken == IF_TOKEN {
// 			statement = parseIfStatement()
// 		}

// 		if statement != null {
// 			program.Statements.push(statement)
// 		}

// 		advanceTokens()
// 	}

// 	return program
// }

// func parseLetStatement() {
// 	advanceTokens()

// }
