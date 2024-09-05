package mysql

import (
	"github.com/may-fly/go-sqlparser/sqlstmt"

	"github.com/may-fly/go-sqlparser/base"
	mysqlparser "github.com/may-fly/go-sqlparser/mysql/antlr4"

	"github.com/antlr4-go/antlr/v4"
)

func GetMysqlParserTree(baseLine int, statement string) (antlr.ParseTree, *antlr.CommonTokenStream, error) {
	lexer := mysqlparser.NewMySqlLexer(antlr.NewInputStream(statement))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := mysqlparser.NewMySqlParser(stream)

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

type MysqlParser struct {
}

func (*MysqlParser) Parse(stmt string) ([]sqlstmt.Stmt, error) {
	tree, _, err := GetMysqlParserTree(1, stmt)
	if err != nil {
		return nil, err
	}

	return tree.Accept(new(MysqlVisitor)).([]sqlstmt.Stmt), nil
}
