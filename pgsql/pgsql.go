package pgsql

import (
	pgparser "github.com/may-fly/go-sqlparser/pgsql/antlr4"

	"github.com/may-fly/go-sqlparser/base"
	"github.com/may-fly/go-sqlparser/sqlstmt"

	"github.com/antlr4-go/antlr/v4"
)

func GetPgsqlParserTree(baseLine int, statement string) (antlr.ParseTree, *antlr.CommonTokenStream, error) {
	lexer := pgparser.NewPostgreSQLLexer(antlr.NewInputStream(statement))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := pgparser.NewPostgreSQLParser(stream)

	lexerErrorListener := &base.ParseErrorListener{
		BaseLine: baseLine,
	}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrorListener)

	parserErrorListener := &base.ParseErrorListener{
		BaseLine: baseLine,
	}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrorListener)
	parser.BuildParseTrees = true

	tree := parser.Root()

	if lexerErrorListener.Err != nil {
		return nil, nil, lexerErrorListener.Err
	}

	if parserErrorListener.Err != nil {
		return nil, nil, parserErrorListener.Err
	}

	return tree, stream, nil
}

type PgsqlParser struct {
}

func (*PgsqlParser) Parse(stmt string) ([]sqlstmt.Stmt, error) {
	tree, _, err := GetPgsqlParserTree(1, stmt)
	if err != nil {
		return nil, err
	}

	return tree.Accept(new(PgsqlVisitor)).([]sqlstmt.Stmt), nil
}
