package parser

import (
	"github.com/samitcheema/Monkey/ast"
	"github.com/samitcheema/Monkey/lexer"
	"github.com/samitcheema/Monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
