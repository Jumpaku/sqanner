package expr

import (
	"cloud.google.com/go/civil"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/node/types"
	"math/big"
	"time"
)

type ScalarNode interface {
	node.Node
	Type() types.TypeCode
	Int64Value() int64
	Float64Value() float64
	BoolValue() bool
	StringValue() string
	BytesValue() []byte
	TimestampValue() time.Time
	DateValue() civil.Date
	JSONValue() any
	NumericValue() *big.Rat
}

func AcceptInt64(s *parse.ParseState, v int64) *scalarNode {
	return &scalarNode{
		NodeBase:   node.NewNodeBase(s.Begin(), s.End()),
		typeCode:   types.TypeCodeInt64,
		int64Value: v,
	}
}

func AcceptFloat64(s *parse.ParseState, v float64) *scalarNode {
	return &scalarNode{
		NodeBase:     node.NewNodeBase(s.Begin(), s.End()),
		typeCode:     types.TypeCodeFloat64,
		float64Value: v,
	}
}

func AcceptBool(s *parse.ParseState, v bool) *scalarNode {
	return &scalarNode{
		NodeBase:  node.NewNodeBase(s.Begin(), s.End()),
		typeCode:  types.TypeCodeBool,
		boolValue: v,
	}
}

func AcceptString(s *parse.ParseState, v string) *scalarNode {
	return &scalarNode{
		NodeBase:    node.NewNodeBase(s.Begin(), s.End()),
		typeCode:    types.TypeCodeString,
		stringValue: v,
	}
}

func AcceptBytes(s *parse.ParseState, v []byte) *scalarNode {
	return &scalarNode{
		NodeBase:   node.NewNodeBase(s.Begin(), s.End()),
		typeCode:   types.TypeCodeBytes,
		bytesValue: v,
	}
}

func AcceptTimestamp(s *parse.ParseState, v time.Time) *scalarNode {
	return &scalarNode{
		NodeBase:       node.NewNodeBase(s.Begin(), s.End()),
		typeCode:       types.TypeCodeTimestamp,
		timestampValue: v,
	}
}

func AcceptDate(s *parse.ParseState, v civil.Date) *scalarNode {
	return &scalarNode{
		NodeBase:  node.NewNodeBase(s.Begin(), s.End()),
		typeCode:  types.TypeCodeDate,
		dateValue: v,
	}
}

func AcceptJSON(s *parse.ParseState, v any) *scalarNode {
	return &scalarNode{
		NodeBase:  node.NewNodeBase(s.Begin(), s.End()),
		typeCode:  types.TypeCodeJSON,
		jsonValue: v,
	}
}

func AcceptNumeric(s *parse.ParseState, v *big.Rat) *scalarNode {
	return &scalarNode{
		NodeBase:     node.NewNodeBase(s.Begin(), s.End()),
		typeCode:     types.TypeCodeNumeric,
		numericValue: v,
	}
}

func NewInt64(v int64) *scalarNode {
	return &scalarNode{
		typeCode:   types.TypeCodeInt64,
		int64Value: v,
	}
}

func NewFloat64(v float64) *scalarNode {
	return &scalarNode{
		typeCode:     types.TypeCodeFloat64,
		float64Value: v,
	}
}

func NewBool(v bool) *scalarNode {
	return &scalarNode{
		typeCode:  types.TypeCodeBool,
		boolValue: v,
	}
}

func NewString(v string) *scalarNode {
	return &scalarNode{
		typeCode:    types.TypeCodeString,
		stringValue: v,
	}
}

func NewBytes(v []byte) *scalarNode {
	return &scalarNode{
		typeCode:   types.TypeCodeBytes,
		bytesValue: v,
	}
}

func NewTimestamp(v time.Time) *scalarNode {
	return &scalarNode{
		typeCode:       types.TypeCodeTimestamp,
		timestampValue: v,
	}
}

func NewDate(v civil.Date) *scalarNode {
	return &scalarNode{
		typeCode:  types.TypeCodeDate,
		dateValue: v,
	}
}

func NewJSON(v any) *scalarNode {
	return &scalarNode{
		typeCode:  types.TypeCodeJSON,
		jsonValue: v,
	}
}

func NewNumeric(v *big.Rat) *scalarNode {
	return &scalarNode{
		typeCode:     types.TypeCodeNumeric,
		numericValue: v,
	}
}

type scalarNode struct {
	node.NodeBase
	typeCode       types.TypeCode
	int64Value     int64
	float64Value   float64
	boolValue      bool
	stringValue    string
	bytesValue     []byte
	timestampValue time.Time
	dateValue      civil.Date
	jsonValue      any
	numericValue   *big.Rat
}

var _ ScalarNode = (*scalarNode)(nil)

func (n *scalarNode) Children() []node.Node {
	return nil
}

func (n *scalarNode) Type() types.TypeCode {
	return n.typeCode
}

func (n *scalarNode) Int64Value() int64 {
	return n.int64Value
}

func (n *scalarNode) Float64Value() float64 {
	return n.float64Value
}

func (n *scalarNode) BoolValue() bool {
	return n.boolValue
}

func (n *scalarNode) StringValue() string {
	return n.stringValue
}

func (n *scalarNode) BytesValue() []byte {
	return n.bytesValue
}

func (n *scalarNode) TimestampValue() time.Time {
	return n.timestampValue
}

func (n *scalarNode) DateValue() civil.Date {
	return n.dateValue
}

func (n *scalarNode) JSONValue() any {
	return n.jsonValue
}

func (n *scalarNode) NumericValue() *big.Rat {
	return n.numericValue
}
