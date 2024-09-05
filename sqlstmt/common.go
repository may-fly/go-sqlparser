package sqlstmt

type (
	IExpr interface {
		INode

		isExpr()
	}

	Expr struct {
		*Node
	}

	LogicalExpr struct {
		Expr

		Operator string
		Exprs    []IExpr
	}

	PredicateExpr struct {
		Expr
	}
)

func (*Expr) isExpr() {}

type (
	IPredicate interface {
		INode

		isPredicate()
	}

	Predicate struct {
		*Node
	}

	BinaryComparisonPredicate struct {
		Predicate

		Left               IPredicate
		Right              IPredicate
		ComparisonOperator string
	}

	InPredicate struct {
		*Node

		Predicate  IPredicate
		Exprs      []IExpr
		SelectStmt ISelectStmt
	}

	ExprAtomPredicate struct {
		Predicate

		ExprAtom IExprAtom
	}
)

func (*Predicate) isPredicate() {}

type (
	IExprAtom interface {
		INode

		isExprAtom()
	}

	ExprAtom struct {
		*Node
	}

	ExprAtomFunctionCall struct {
		*Node
	}

	ExprAtomConstant struct {
		ExprAtom

		Constant *Constant
	}
)

func (*ExprAtom) isExprAtom() {}

type (
	ITableSource interface {
		INode

		isTableSource()
	}

	TableSource struct {
		*Node
	}

	TableSources struct {
		*Node

		TableSources []ITableSource
	}

	TableSourceBase struct {
		TableSource

		TableSourceItem ITableSourceItem
		JoinParts       []IJoinPart
	}

	ITableSourceItem interface {
		INode
	}

	TableSourceItem struct {
		*Node
	}

	AtomTableItem struct {
		TableSourceItem

		TableName *TableName // 表名
		Alias     string     // 别名
	}
)

func (*TableSource) isTableSource() {}

type (
	Constant struct {
		*Node

		Value string
	}

	FullId struct {
		*Node

		Uids []string
	}
)

type ColumnName struct {
	*Node

	Owner             string
	Identifier        *IdentifierValue
	NestedObjectAttrs []string
}

type TableName struct {
	*Node

	Owner      string
	Identifier *IdentifierValue
}

type (
	FuncCall interface {
		INode
	}
)
