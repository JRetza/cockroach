// Code generated by optgen; DO NOT EDIT.

package opt

const (
	UnknownOp Operator = iota

	// Scalar Operators
	SubqueryOp
	VariableOp
	ConstOp
	TrueOp
	FalseOp
	PlaceholderOp
	TupleOp
	ProjectionsOp
	AggregationsOp
	GroupingsOp
	FiltersOp
	ExistsOp
	AndOp
	OrOp
	NotOp
	EqOp
	LtOp
	GtOp
	LeOp
	GeOp
	NeOp
	InOp
	NotInOp
	LikeOp
	NotLikeOp
	ILikeOp
	NotILikeOp
	SimilarToOp
	NotSimilarToOp
	RegMatchOp
	NotRegMatchOp
	RegIMatchOp
	NotRegIMatchOp
	IsOp
	IsNotOp
	ContainsOp
	ContainedByOp
	BitandOp
	BitorOp
	BitxorOp
	PlusOp
	MinusOp
	MultOp
	DivOp
	FloorDivOp
	ModOp
	PowOp
	ConcatOp
	LShiftOp
	RShiftOp
	FetchValOp
	FetchTextOp
	FetchValPathOp
	FetchTextPathOp
	UnaryPlusOp
	UnaryMinusOp
	UnaryComplementOp
	FunctionOp

	// Relational Operators
	ScanOp
	ValuesOp
	SelectOp
	ProjectOp
	InnerJoinOp
	LeftJoinOp
	RightJoinOp
	FullJoinOp
	SemiJoinOp
	AntiJoinOp
	InnerJoinApplyOp
	LeftJoinApplyOp
	RightJoinApplyOp
	FullJoinApplyOp
	SemiJoinApplyOp
	AntiJoinApplyOp
	GroupByOp
	UnionOp
	IntersectOp
	ExceptOp

	// Enforcer Operators
	SortOp
	PresentOp

	// NumOperators tracks the total count of operators.
	NumOperators
)

const opNames = "unknownsubqueryvariableconsttruefalseplaceholdertupleprojectionsaggregationsgroupingsfiltersexistsandornoteqltgtlegeneinnot-inlikenot-likei-likenot-i-likesimilar-tonot-similar-toreg-matchnot-reg-matchreg-i-matchnot-reg-i-matchisis-notcontainscontained-bybitandbitorbitxorplusminusmultdivfloor-divmodpowconcatl-shiftr-shiftfetch-valfetch-textfetch-val-pathfetch-text-pathunary-plusunary-minusunary-complementfunctionscanvaluesselectprojectinner-joinleft-joinright-joinfull-joinsemi-joinanti-joininner-join-applyleft-join-applyright-join-applyfull-join-applysemi-join-applyanti-join-applygroup-byunionintersectexceptsortpresent"

var opIndexes = [...]uint32{0, 7, 15, 23, 28, 32, 37, 48, 53, 64, 76, 85, 92, 98, 101, 103, 106, 108, 110, 112, 114, 116, 118, 120, 126, 130, 138, 144, 154, 164, 178, 187, 200, 211, 226, 228, 234, 242, 254, 260, 265, 271, 275, 280, 284, 287, 296, 299, 302, 308, 315, 322, 331, 341, 355, 370, 380, 391, 407, 415, 419, 425, 431, 438, 448, 457, 467, 476, 485, 494, 510, 525, 541, 556, 571, 586, 594, 599, 608, 614, 618, 625}

var ScalarOperators = [...]Operator{
	SubqueryOp,
	VariableOp,
	ConstOp,
	TrueOp,
	FalseOp,
	PlaceholderOp,
	TupleOp,
	ProjectionsOp,
	AggregationsOp,
	GroupingsOp,
	FiltersOp,
	ExistsOp,
	AndOp,
	OrOp,
	NotOp,
	EqOp,
	LtOp,
	GtOp,
	LeOp,
	GeOp,
	NeOp,
	InOp,
	NotInOp,
	LikeOp,
	NotLikeOp,
	ILikeOp,
	NotILikeOp,
	SimilarToOp,
	NotSimilarToOp,
	RegMatchOp,
	NotRegMatchOp,
	RegIMatchOp,
	NotRegIMatchOp,
	IsOp,
	IsNotOp,
	ContainsOp,
	ContainedByOp,
	BitandOp,
	BitorOp,
	BitxorOp,
	PlusOp,
	MinusOp,
	MultOp,
	DivOp,
	FloorDivOp,
	ModOp,
	PowOp,
	ConcatOp,
	LShiftOp,
	RShiftOp,
	FetchValOp,
	FetchTextOp,
	FetchValPathOp,
	FetchTextPathOp,
	UnaryPlusOp,
	UnaryMinusOp,
	UnaryComplementOp,
	FunctionOp,
}

var BooleanOperators = [...]Operator{
	TrueOp,
	FalseOp,
	AndOp,
	OrOp,
	NotOp,
}

var ComparisonOperators = [...]Operator{
	EqOp,
	LtOp,
	GtOp,
	LeOp,
	GeOp,
	NeOp,
	InOp,
	NotInOp,
	LikeOp,
	NotLikeOp,
	ILikeOp,
	NotILikeOp,
	SimilarToOp,
	NotSimilarToOp,
	RegMatchOp,
	NotRegMatchOp,
	RegIMatchOp,
	NotRegIMatchOp,
	IsOp,
	IsNotOp,
	ContainsOp,
	ContainedByOp,
}

var BinaryOperators = [...]Operator{
	BitandOp,
	BitorOp,
	BitxorOp,
	PlusOp,
	MinusOp,
	MultOp,
	DivOp,
	FloorDivOp,
	ModOp,
	PowOp,
	ConcatOp,
	LShiftOp,
	RShiftOp,
	FetchValOp,
	FetchTextOp,
	FetchValPathOp,
	FetchTextPathOp,
}

var UnaryOperators = [...]Operator{
	UnaryPlusOp,
	UnaryMinusOp,
	UnaryComplementOp,
}

var RelationalOperators = [...]Operator{
	ScanOp,
	ValuesOp,
	SelectOp,
	ProjectOp,
	InnerJoinOp,
	LeftJoinOp,
	RightJoinOp,
	FullJoinOp,
	SemiJoinOp,
	AntiJoinOp,
	InnerJoinApplyOp,
	LeftJoinApplyOp,
	RightJoinApplyOp,
	FullJoinApplyOp,
	SemiJoinApplyOp,
	AntiJoinApplyOp,
	GroupByOp,
	UnionOp,
	IntersectOp,
	ExceptOp,
}

var JoinOperators = [...]Operator{
	InnerJoinOp,
	LeftJoinOp,
	RightJoinOp,
	FullJoinOp,
	SemiJoinOp,
	AntiJoinOp,
	InnerJoinApplyOp,
	LeftJoinApplyOp,
	RightJoinApplyOp,
	FullJoinApplyOp,
	SemiJoinApplyOp,
	AntiJoinApplyOp,
}

var JoinApplyOperators = [...]Operator{
	InnerJoinApplyOp,
	LeftJoinApplyOp,
	RightJoinApplyOp,
	FullJoinApplyOp,
	SemiJoinApplyOp,
	AntiJoinApplyOp,
}

var EnforcerOperators = [...]Operator{
	SortOp,
	PresentOp,
}