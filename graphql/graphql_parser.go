// Generated from GraphQL.g4 by ANTLR 4.6.

package parser // GraphQL

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 24, 247, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 18,
	4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 4,
	24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 29,
	9, 29, 4, 30, 9, 30, 3, 2, 6, 2, 62, 10, 2, 13, 2, 14, 2, 63, 3, 3, 3,
	3, 5, 3, 68, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 74, 10, 4, 3, 4, 5, 4,
	77, 10, 4, 3, 4, 3, 4, 5, 4, 81, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 86, 10,
	5, 3, 5, 7, 5, 89, 10, 5, 12, 5, 14, 5, 92, 11, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 5, 7, 101, 10, 7, 3, 8, 3, 8, 5, 8, 105, 10, 8, 3,
	8, 5, 8, 108, 10, 8, 3, 8, 5, 8, 111, 10, 8, 3, 9, 3, 9, 5, 9, 115, 10,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 125,
	10, 11, 12, 11, 14, 11, 128, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 5, 13, 139, 10, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 5, 14, 145, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	5, 15, 154, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 6, 17, 161, 10,
	17, 13, 17, 14, 17, 162, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 177, 10, 18, 3, 19, 3, 19, 3,
	20, 3, 20, 3, 20, 3, 20, 7, 20, 185, 10, 20, 12, 20, 14, 20, 188, 11, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 196, 10, 21, 3, 22, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 5, 24, 206, 10, 24, 3, 25,
	3, 25, 3, 25, 3, 25, 5, 25, 212, 10, 25, 3, 26, 3, 26, 5, 26, 216, 10,
	26, 3, 26, 3, 26, 5, 26, 220, 10, 26, 5, 26, 222, 10, 26, 3, 27, 3, 27,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 7,
	30, 236, 10, 30, 12, 30, 14, 30, 239, 11, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	5, 30, 245, 10, 30, 3, 30, 2, 2, 31, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 2, 3, 3, 2, 6, 7, 248, 2, 61, 3, 2, 2, 2, 4, 67, 3, 2, 2, 2, 6, 80,
	3, 2, 2, 2, 8, 82, 3, 2, 2, 2, 10, 95, 3, 2, 2, 2, 12, 100, 3, 2, 2, 2,
	14, 102, 3, 2, 2, 2, 16, 114, 3, 2, 2, 2, 18, 116, 3, 2, 2, 2, 20, 120,
	3, 2, 2, 2, 22, 131, 3, 2, 2, 2, 24, 135, 3, 2, 2, 2, 26, 140, 3, 2, 2,
	2, 28, 148, 3, 2, 2, 2, 30, 157, 3, 2, 2, 2, 32, 160, 3, 2, 2, 2, 34, 176,
	3, 2, 2, 2, 36, 178, 3, 2, 2, 2, 38, 180, 3, 2, 2, 2, 40, 191, 3, 2, 2,
	2, 42, 197, 3, 2, 2, 2, 44, 200, 3, 2, 2, 2, 46, 205, 3, 2, 2, 2, 48, 211,
	3, 2, 2, 2, 50, 221, 3, 2, 2, 2, 52, 223, 3, 2, 2, 2, 54, 225, 3, 2, 2,
	2, 56, 229, 3, 2, 2, 2, 58, 244, 3, 2, 2, 2, 60, 62, 5, 4, 3, 2, 61, 60,
	3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2,
	64, 3, 3, 2, 2, 2, 65, 68, 5, 6, 4, 2, 66, 68, 5, 28, 15, 2, 67, 65, 3,
	2, 2, 2, 67, 66, 3, 2, 2, 2, 68, 5, 3, 2, 2, 2, 69, 81, 5, 8, 5, 2, 70,
	71, 5, 10, 6, 2, 71, 73, 7, 20, 2, 2, 72, 74, 5, 38, 20, 2, 73, 72, 3,
	2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 77, 5, 32, 17, 2,
	76, 75, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 5,
	8, 5, 2, 79, 81, 3, 2, 2, 2, 80, 69, 3, 2, 2, 2, 80, 70, 3, 2, 2, 2, 81,
	7, 3, 2, 2, 2, 82, 83, 7, 3, 2, 2, 83, 90, 5, 12, 7, 2, 84, 86, 7, 4, 2,
	2, 85, 84, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89,
	5, 12, 7, 2, 88, 85, 3, 2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2,
	90, 91, 3, 2, 2, 2, 91, 93, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 94, 7,
	5, 2, 2, 94, 9, 3, 2, 2, 2, 95, 96, 9, 2, 2, 2, 96, 11, 3, 2, 2, 2, 97,
	101, 5, 14, 8, 2, 98, 101, 5, 24, 13, 2, 99, 101, 5, 26, 14, 2, 100, 97,
	3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 99, 3, 2, 2, 2, 101, 13, 3, 2, 2,
	2, 102, 104, 5, 16, 9, 2, 103, 105, 5, 20, 11, 2, 104, 103, 3, 2, 2, 2,
	104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2, 106, 108, 5, 32, 17, 2, 107,
	106, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 110, 3, 2, 2, 2, 109, 111,
	5, 8, 5, 2, 110, 109, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 15, 3, 2,
	2, 2, 112, 115, 5, 18, 10, 2, 113, 115, 7, 20, 2, 2, 114, 112, 3, 2, 2,
	2, 114, 113, 3, 2, 2, 2, 115, 17, 3, 2, 2, 2, 116, 117, 7, 20, 2, 2, 117,
	118, 7, 8, 2, 2, 118, 119, 7, 20, 2, 2, 119, 19, 3, 2, 2, 2, 120, 121,
	7, 9, 2, 2, 121, 126, 5, 22, 12, 2, 122, 123, 7, 4, 2, 2, 123, 125, 5,
	22, 12, 2, 124, 122, 3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2,
	2, 2, 126, 127, 3, 2, 2, 2, 127, 129, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2,
	129, 130, 7, 10, 2, 2, 130, 21, 3, 2, 2, 2, 131, 132, 7, 20, 2, 2, 132,
	133, 7, 8, 2, 2, 133, 134, 5, 46, 24, 2, 134, 23, 3, 2, 2, 2, 135, 136,
	7, 11, 2, 2, 136, 138, 5, 30, 16, 2, 137, 139, 5, 32, 17, 2, 138, 137,
	3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 25, 3, 2, 2, 2, 140, 141, 7, 11,
	2, 2, 141, 142, 7, 12, 2, 2, 142, 144, 5, 36, 19, 2, 143, 145, 5, 32, 17,
	2, 144, 143, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146,
	147, 5, 8, 5, 2, 147, 27, 3, 2, 2, 2, 148, 149, 7, 13, 2, 2, 149, 150,
	5, 30, 16, 2, 150, 151, 7, 12, 2, 2, 151, 153, 5, 36, 19, 2, 152, 154,
	5, 32, 17, 2, 153, 152, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 3,
	2, 2, 2, 155, 156, 5, 8, 5, 2, 156, 29, 3, 2, 2, 2, 157, 158, 7, 20, 2,
	2, 158, 31, 3, 2, 2, 2, 159, 161, 5, 34, 18, 2, 160, 159, 3, 2, 2, 2, 161,
	162, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 33, 3,
	2, 2, 2, 164, 165, 7, 14, 2, 2, 165, 166, 7, 20, 2, 2, 166, 167, 7, 8,
	2, 2, 167, 177, 5, 46, 24, 2, 168, 169, 7, 14, 2, 2, 169, 177, 7, 20, 2,
	2, 170, 171, 7, 14, 2, 2, 171, 172, 7, 20, 2, 2, 172, 173, 7, 9, 2, 2,
	173, 174, 5, 22, 12, 2, 174, 175, 7, 10, 2, 2, 175, 177, 3, 2, 2, 2, 176,
	164, 3, 2, 2, 2, 176, 168, 3, 2, 2, 2, 176, 170, 3, 2, 2, 2, 177, 35, 3,
	2, 2, 2, 178, 179, 5, 52, 27, 2, 179, 37, 3, 2, 2, 2, 180, 181, 7, 9, 2,
	2, 181, 186, 5, 40, 21, 2, 182, 183, 7, 4, 2, 2, 183, 185, 5, 40, 21, 2,
	184, 182, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186,
	187, 3, 2, 2, 2, 187, 189, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 189, 190,
	7, 10, 2, 2, 190, 39, 3, 2, 2, 2, 191, 192, 5, 42, 22, 2, 192, 193, 7,
	8, 2, 2, 193, 195, 5, 50, 26, 2, 194, 196, 5, 44, 23, 2, 195, 194, 3, 2,
	2, 2, 195, 196, 3, 2, 2, 2, 196, 41, 3, 2, 2, 2, 197, 198, 7, 15, 2, 2,
	198, 199, 7, 20, 2, 2, 199, 43, 3, 2, 2, 2, 200, 201, 7, 16, 2, 2, 201,
	202, 5, 48, 25, 2, 202, 45, 3, 2, 2, 2, 203, 206, 5, 48, 25, 2, 204, 206,
	5, 42, 22, 2, 205, 203, 3, 2, 2, 2, 205, 204, 3, 2, 2, 2, 206, 47, 3, 2,
	2, 2, 207, 212, 7, 21, 2, 2, 208, 212, 7, 23, 2, 2, 209, 212, 7, 22, 2,
	2, 210, 212, 5, 58, 30, 2, 211, 207, 3, 2, 2, 2, 211, 208, 3, 2, 2, 2,
	211, 209, 3, 2, 2, 2, 211, 210, 3, 2, 2, 2, 212, 49, 3, 2, 2, 2, 213, 215,
	5, 52, 27, 2, 214, 216, 5, 56, 29, 2, 215, 214, 3, 2, 2, 2, 215, 216, 3,
	2, 2, 2, 216, 222, 3, 2, 2, 2, 217, 219, 5, 54, 28, 2, 218, 220, 5, 56,
	29, 2, 219, 218, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 222, 3, 2, 2, 2,
	221, 213, 3, 2, 2, 2, 221, 217, 3, 2, 2, 2, 222, 51, 3, 2, 2, 2, 223, 224,
	7, 20, 2, 2, 224, 53, 3, 2, 2, 2, 225, 226, 7, 17, 2, 2, 226, 227, 5, 50,
	26, 2, 227, 228, 7, 18, 2, 2, 228, 55, 3, 2, 2, 2, 229, 230, 7, 19, 2,
	2, 230, 57, 3, 2, 2, 2, 231, 232, 7, 17, 2, 2, 232, 237, 5, 48, 25, 2,
	233, 234, 7, 4, 2, 2, 234, 236, 5, 48, 25, 2, 235, 233, 3, 2, 2, 2, 236,
	239, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 240,
	3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 240, 241, 7, 18, 2, 2, 241, 245, 3, 2,
	2, 2, 242, 243, 7, 17, 2, 2, 243, 245, 7, 18, 2, 2, 244, 231, 3, 2, 2,
	2, 244, 242, 3, 2, 2, 2, 245, 59, 3, 2, 2, 2, 29, 63, 67, 73, 76, 80, 85,
	90, 100, 104, 107, 110, 114, 126, 138, 144, 153, 162, 176, 186, 195, 205,
	211, 215, 219, 221, 237, 244,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "','", "'}'", "'query'", "'mutation'", "':'", "'('", "')'",
	"'...'", "'on'", "'fragment'", "'@'", "'$'", "'='", "'['", "']'", "'!'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"NAME", "STRING", "BOOLEAN", "NUMBER", "WS",
}

var ruleNames = []string{
	"document", "definition", "operationDefinition", "selectionSet", "operationType",
	"selection", "field", "fieldName", "alias", "arguments", "argument", "fragmentSpread",
	"inlineFragment", "fragmentDefinition", "fragmentName", "directives", "directive",
	"typeCondition", "variableDefinitions", "variableDefinition", "variable",
	"defaultValue", "valueOrVariable", "value", "type_", "typeName", "listType",
	"nonNullType", "array",
}

type GraphQLParser struct {
	*antlr.BaseParser
}

func NewGraphQLParser(input antlr.TokenStream) *GraphQLParser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(GraphQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GraphQL.g4"

	return this
}

// GraphQLParser tokens.
const (
	GraphQLParserEOF     = antlr.TokenEOF
	GraphQLParserT__0    = 1
	GraphQLParserT__1    = 2
	GraphQLParserT__2    = 3
	GraphQLParserT__3    = 4
	GraphQLParserT__4    = 5
	GraphQLParserT__5    = 6
	GraphQLParserT__6    = 7
	GraphQLParserT__7    = 8
	GraphQLParserT__8    = 9
	GraphQLParserT__9    = 10
	GraphQLParserT__10   = 11
	GraphQLParserT__11   = 12
	GraphQLParserT__12   = 13
	GraphQLParserT__13   = 14
	GraphQLParserT__14   = 15
	GraphQLParserT__15   = 16
	GraphQLParserT__16   = 17
	GraphQLParserNAME    = 18
	GraphQLParserSTRING  = 19
	GraphQLParserBOOLEAN = 20
	GraphQLParserNUMBER  = 21
	GraphQLParserWS      = 22
)

// GraphQLParser rules.
const (
	GraphQLParserRULE_document            = 0
	GraphQLParserRULE_definition          = 1
	GraphQLParserRULE_operationDefinition = 2
	GraphQLParserRULE_selectionSet        = 3
	GraphQLParserRULE_operationType       = 4
	GraphQLParserRULE_selection           = 5
	GraphQLParserRULE_field               = 6
	GraphQLParserRULE_fieldName           = 7
	GraphQLParserRULE_alias               = 8
	GraphQLParserRULE_arguments           = 9
	GraphQLParserRULE_argument            = 10
	GraphQLParserRULE_fragmentSpread      = 11
	GraphQLParserRULE_inlineFragment      = 12
	GraphQLParserRULE_fragmentDefinition  = 13
	GraphQLParserRULE_fragmentName        = 14
	GraphQLParserRULE_directives          = 15
	GraphQLParserRULE_directive           = 16
	GraphQLParserRULE_typeCondition       = 17
	GraphQLParserRULE_variableDefinitions = 18
	GraphQLParserRULE_variableDefinition  = 19
	GraphQLParserRULE_variable            = 20
	GraphQLParserRULE_defaultValue        = 21
	GraphQLParserRULE_valueOrVariable     = 22
	GraphQLParserRULE_value               = 23
	GraphQLParserRULE_type_               = 24
	GraphQLParserRULE_typeName            = 25
	GraphQLParserRULE_listType            = 26
	GraphQLParserRULE_nonNullType         = 27
	GraphQLParserRULE_array               = 28
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *DocumentContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *GraphQLParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraphQLParserRULE_document)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__10))) != 0) {
		{
			p.SetState(58)
			p.Definition()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) OperationDefinition() IOperationDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationDefinitionContext)
}

func (s *DefinitionContext) FragmentDefinition() IFragmentDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentDefinitionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *GraphQLParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraphQLParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__0, GraphQLParserT__3, GraphQLParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.OperationDefinition()
		}

	case GraphQLParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.FragmentDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationDefinitionContext is an interface to support dynamic dispatch.
type IOperationDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationDefinitionContext differentiates from other interfaces.
	IsOperationDefinitionContext()
}

type OperationDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationDefinitionContext() *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationDefinition
	return p
}

func (*OperationDefinitionContext) IsOperationDefinitionContext() {}

func NewOperationDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationDefinition

	return p
}

func (s *OperationDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationDefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *OperationDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *OperationDefinitionContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, 0)
}

func (s *OperationDefinitionContext) VariableDefinitions() IVariableDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionsContext)
}

func (s *OperationDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *OperationDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationDefinition(s)
	}
}

func (s *OperationDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationDefinition(s)
	}
}

func (p *GraphQLParser) OperationDefinition() (localctx IOperationDefinitionContext) {
	localctx = NewOperationDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GraphQLParserRULE_operationDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.SelectionSet()
		}

	case GraphQLParserT__3, GraphQLParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.OperationType()
		}
		{
			p.SetState(69)
			p.Match(GraphQLParserNAME)
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__6 {
			{
				p.SetState(70)
				p.VariableDefinitions()
			}

		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__11 {
			{
				p.SetState(73)
				p.Directives()
			}

		}
		{
			p.SetState(76)
			p.SelectionSet()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISelectionSetContext is an interface to support dynamic dispatch.
type ISelectionSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionSetContext differentiates from other interfaces.
	IsSelectionSetContext()
}

type SelectionSetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionSetContext() *SelectionSetContext {
	var p = new(SelectionSetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_selectionSet
	return p
}

func (*SelectionSetContext) IsSelectionSetContext() {}

func NewSelectionSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionSetContext {
	var p = new(SelectionSetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_selectionSet

	return p
}

func (s *SelectionSetContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionSetContext) AllSelection() []ISelectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectionContext)(nil)).Elem())
	var tst = make([]ISelectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectionContext)
		}
	}

	return tst
}

func (s *SelectionSetContext) Selection(i int) ISelectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectionContext)
}

func (s *SelectionSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSelectionSet(s)
	}
}

func (s *SelectionSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSelectionSet(s)
	}
}

func (p *GraphQLParser) SelectionSet() (localctx ISelectionSetContext) {
	localctx = NewSelectionSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GraphQLParserRULE_selectionSet)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(GraphQLParserT__0)
	}
	{
		p.SetState(81)
		p.Selection()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__1)|(1<<GraphQLParserT__8)|(1<<GraphQLParserNAME))) != 0 {
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__1 {
			{
				p.SetState(82)
				p.Match(GraphQLParserT__1)
			}

		}
		{
			p.SetState(85)
			p.Selection()
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
		p.Match(GraphQLParserT__2)
	}

	return localctx
}

// IOperationTypeContext is an interface to support dynamic dispatch.
type IOperationTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationTypeContext differentiates from other interfaces.
	IsOperationTypeContext()
}

type OperationTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTypeContext() *OperationTypeContext {
	var p = new(OperationTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationType
	return p
}

func (*OperationTypeContext) IsOperationTypeContext() {}

func NewOperationTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTypeContext {
	var p = new(OperationTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationType

	return p
}

func (s *OperationTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *OperationTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationType(s)
	}
}

func (s *OperationTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationType(s)
	}
}

func (p *GraphQLParser) OperationType() (localctx IOperationTypeContext) {
	localctx = NewOperationTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GraphQLParserRULE_operationType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	_la = p.GetTokenStream().LA(1)

	if !(_la == GraphQLParserT__3 || _la == GraphQLParserT__4) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ISelectionContext is an interface to support dynamic dispatch.
type ISelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionContext differentiates from other interfaces.
	IsSelectionContext()
}

type SelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionContext() *SelectionContext {
	var p = new(SelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_selection
	return p
}

func (*SelectionContext) IsSelectionContext() {}

func NewSelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionContext {
	var p = new(SelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_selection

	return p
}

func (s *SelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *SelectionContext) FragmentSpread() IFragmentSpreadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentSpreadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentSpreadContext)
}

func (s *SelectionContext) InlineFragment() IInlineFragmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineFragmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineFragmentContext)
}

func (s *SelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSelection(s)
	}
}

func (s *SelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSelection(s)
	}
}

func (p *GraphQLParser) Selection() (localctx ISelectionContext) {
	localctx = NewSelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GraphQLParserRULE_selection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.FragmentSpread()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)
			p.InlineFragment()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FieldContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FieldContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *GraphQLParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GraphQLParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.FieldName()
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__6 {
		{
			p.SetState(101)
			p.Arguments()
		}

	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__11 {
		{
			p.SetState(104)
			p.Directives()
		}

	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__0 {
		{
			p.SetState(107)
			p.SelectionSet()
		}

	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *FieldNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *GraphQLParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GraphQLParserRULE_fieldName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Alias()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Match(GraphQLParserNAME)
		}

	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(GraphQLParserNAME)
}

func (s *AliasContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, i)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (p *GraphQLParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GraphQLParserRULE_alias)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(GraphQLParserNAME)
	}
	{
		p.SetState(115)
		p.Match(GraphQLParserT__5)
	}
	{
		p.SetState(116)
		p.Match(GraphQLParserNAME)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *GraphQLParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GraphQLParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(GraphQLParserT__6)
	}
	{
		p.SetState(119)
		p.Argument()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLParserT__1 {
		{
			p.SetState(120)
			p.Match(GraphQLParserT__1)
		}
		{
			p.SetState(121)
			p.Argument()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(127)
		p.Match(GraphQLParserT__7)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, 0)
}

func (s *ArgumentContext) ValueOrVariable() IValueOrVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueOrVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueOrVariableContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *GraphQLParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GraphQLParserRULE_argument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(GraphQLParserNAME)
	}
	{
		p.SetState(130)
		p.Match(GraphQLParserT__5)
	}
	{
		p.SetState(131)
		p.ValueOrVariable()
	}

	return localctx
}

// IFragmentSpreadContext is an interface to support dynamic dispatch.
type IFragmentSpreadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentSpreadContext differentiates from other interfaces.
	IsFragmentSpreadContext()
}

type FragmentSpreadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentSpreadContext() *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentSpread
	return p
}

func (*FragmentSpreadContext) IsFragmentSpreadContext() {}

func NewFragmentSpreadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentSpread

	return p
}

func (s *FragmentSpreadContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentSpreadContext) FragmentName() IFragmentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentSpreadContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentSpreadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentSpreadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentSpreadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentSpread(s)
	}
}

func (s *FragmentSpreadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentSpread(s)
	}
}

func (p *GraphQLParser) FragmentSpread() (localctx IFragmentSpreadContext) {
	localctx = NewFragmentSpreadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GraphQLParserRULE_fragmentSpread)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(GraphQLParserT__8)
	}
	{
		p.SetState(134)
		p.FragmentName()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__11 {
		{
			p.SetState(135)
			p.Directives()
		}

	}

	return localctx
}

// IInlineFragmentContext is an interface to support dynamic dispatch.
type IInlineFragmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineFragmentContext differentiates from other interfaces.
	IsInlineFragmentContext()
}

type InlineFragmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineFragmentContext() *InlineFragmentContext {
	var p = new(InlineFragmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inlineFragment
	return p
}

func (*InlineFragmentContext) IsInlineFragmentContext() {}

func NewInlineFragmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineFragmentContext {
	var p = new(InlineFragmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inlineFragment

	return p
}

func (s *InlineFragmentContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineFragmentContext) TypeCondition() ITypeConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *InlineFragmentContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *InlineFragmentContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InlineFragmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineFragmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineFragmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInlineFragment(s)
	}
}

func (s *InlineFragmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInlineFragment(s)
	}
}

func (p *GraphQLParser) InlineFragment() (localctx IInlineFragmentContext) {
	localctx = NewInlineFragmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GraphQLParserRULE_inlineFragment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(GraphQLParserT__8)
	}
	{
		p.SetState(139)
		p.Match(GraphQLParserT__9)
	}
	{
		p.SetState(140)
		p.TypeCondition()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__11 {
		{
			p.SetState(141)
			p.Directives()
		}

	}
	{
		p.SetState(144)
		p.SelectionSet()
	}

	return localctx
}

// IFragmentDefinitionContext is an interface to support dynamic dispatch.
type IFragmentDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentDefinitionContext differentiates from other interfaces.
	IsFragmentDefinitionContext()
}

type FragmentDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentDefinitionContext() *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentDefinition
	return p
}

func (*FragmentDefinitionContext) IsFragmentDefinitionContext() {}

func NewFragmentDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentDefinition

	return p
}

func (s *FragmentDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentDefinitionContext) FragmentName() IFragmentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentDefinitionContext) TypeCondition() ITypeConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *FragmentDefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FragmentDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentDefinition(s)
	}
}

func (s *FragmentDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentDefinition(s)
	}
}

func (p *GraphQLParser) FragmentDefinition() (localctx IFragmentDefinitionContext) {
	localctx = NewFragmentDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GraphQLParserRULE_fragmentDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(GraphQLParserT__10)
	}
	{
		p.SetState(147)
		p.FragmentName()
	}
	{
		p.SetState(148)
		p.Match(GraphQLParserT__9)
	}
	{
		p.SetState(149)
		p.TypeCondition()
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__11 {
		{
			p.SetState(150)
			p.Directives()
		}

	}
	{
		p.SetState(153)
		p.SelectionSet()
	}

	return localctx
}

// IFragmentNameContext is an interface to support dynamic dispatch.
type IFragmentNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentNameContext differentiates from other interfaces.
	IsFragmentNameContext()
}

type FragmentNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentNameContext() *FragmentNameContext {
	var p = new(FragmentNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentName
	return p
}

func (*FragmentNameContext) IsFragmentNameContext() {}

func NewFragmentNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentNameContext {
	var p = new(FragmentNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentName

	return p
}

func (s *FragmentNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, 0)
}

func (s *FragmentNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentName(s)
	}
}

func (s *FragmentNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentName(s)
	}
}

func (p *GraphQLParser) FragmentName() (localctx IFragmentNameContext) {
	localctx = NewFragmentNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GraphQLParserRULE_fragmentName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(GraphQLParserNAME)
	}

	return localctx
}

// IDirectivesContext is an interface to support dynamic dispatch.
type IDirectivesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectivesContext differentiates from other interfaces.
	IsDirectivesContext()
}

type DirectivesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectivesContext() *DirectivesContext {
	var p = new(DirectivesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directives
	return p
}

func (*DirectivesContext) IsDirectivesContext() {}

func NewDirectivesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectivesContext {
	var p = new(DirectivesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directives

	return p
}

func (s *DirectivesContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectivesContext) AllDirective() []IDirectiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveContext)(nil)).Elem())
	var tst = make([]IDirectiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveContext)
		}
	}

	return tst
}

func (s *DirectivesContext) Directive(i int) IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *DirectivesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectivesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectivesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectives(s)
	}
}

func (s *DirectivesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectives(s)
	}
}

func (p *GraphQLParser) Directives() (localctx IDirectivesContext) {
	localctx = NewDirectivesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GraphQLParserRULE_directives)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserT__11 {
		{
			p.SetState(157)
			p.Directive()
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, 0)
}

func (s *DirectiveContext) ValueOrVariable() IValueOrVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueOrVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueOrVariableContext)
}

func (s *DirectiveContext) Argument() IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (p *GraphQLParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GraphQLParserRULE_directive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(GraphQLParserT__11)
		}
		{
			p.SetState(163)
			p.Match(GraphQLParserNAME)
		}
		{
			p.SetState(164)
			p.Match(GraphQLParserT__5)
		}
		{
			p.SetState(165)
			p.ValueOrVariable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Match(GraphQLParserT__11)
		}
		{
			p.SetState(167)
			p.Match(GraphQLParserNAME)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.Match(GraphQLParserT__11)
		}
		{
			p.SetState(169)
			p.Match(GraphQLParserNAME)
		}
		{
			p.SetState(170)
			p.Match(GraphQLParserT__6)
		}
		{
			p.SetState(171)
			p.Argument()
		}
		{
			p.SetState(172)
			p.Match(GraphQLParserT__7)
		}

	}

	return localctx
}

// ITypeConditionContext is an interface to support dynamic dispatch.
type ITypeConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeConditionContext differentiates from other interfaces.
	IsTypeConditionContext()
}

type TypeConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeConditionContext() *TypeConditionContext {
	var p = new(TypeConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeCondition
	return p
}

func (*TypeConditionContext) IsTypeConditionContext() {}

func NewTypeConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeConditionContext {
	var p = new(TypeConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeCondition

	return p
}

func (s *TypeConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeConditionContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TypeConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeCondition(s)
	}
}

func (s *TypeConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeCondition(s)
	}
}

func (p *GraphQLParser) TypeCondition() (localctx ITypeConditionContext) {
	localctx = NewTypeConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GraphQLParserRULE_typeCondition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.TypeName()
	}

	return localctx
}

// IVariableDefinitionsContext is an interface to support dynamic dispatch.
type IVariableDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDefinitionsContext differentiates from other interfaces.
	IsVariableDefinitionsContext()
}

type VariableDefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionsContext() *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinitions
	return p
}

func (*VariableDefinitionsContext) IsVariableDefinitionsContext() {}

func NewVariableDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variableDefinitions

	return p
}

func (s *VariableDefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionsContext) AllVariableDefinition() []IVariableDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDefinitionContext)(nil)).Elem())
	var tst = make([]IVariableDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDefinitionContext)
		}
	}

	return tst
}

func (s *VariableDefinitionsContext) VariableDefinition(i int) IVariableDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionContext)
}

func (s *VariableDefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariableDefinitions(s)
	}
}

func (s *VariableDefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariableDefinitions(s)
	}
}

func (p *GraphQLParser) VariableDefinitions() (localctx IVariableDefinitionsContext) {
	localctx = NewVariableDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GraphQLParserRULE_variableDefinitions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(GraphQLParserT__6)
	}
	{
		p.SetState(179)
		p.VariableDefinition()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLParserT__1 {
		{
			p.SetState(180)
			p.Match(GraphQLParserT__1)
		}
		{
			p.SetState(181)
			p.VariableDefinition()
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(187)
		p.Match(GraphQLParserT__7)
	}

	return localctx
}

// IVariableDefinitionContext is an interface to support dynamic dispatch.
type IVariableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDefinitionContext differentiates from other interfaces.
	IsVariableDefinitionContext()
}

type VariableDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionContext() *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinition
	return p
}

func (*VariableDefinitionContext) IsVariableDefinitionContext() {}

func NewVariableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variableDefinition

	return p
}

func (s *VariableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableDefinitionContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VariableDefinitionContext) DefaultValue() IDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *VariableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariableDefinition(s)
	}
}

func (p *GraphQLParser) VariableDefinition() (localctx IVariableDefinitionContext) {
	localctx = NewVariableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GraphQLParserRULE_variableDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Variable()
	}
	{
		p.SetState(190)
		p.Match(GraphQLParserT__5)
	}
	{
		p.SetState(191)
		p.Type_()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__13 {
		{
			p.SetState(192)
			p.DefaultValue()
		}

	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *GraphQLParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GraphQLParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(GraphQLParserT__12)
	}
	{
		p.SetState(196)
		p.Match(GraphQLParserNAME)
	}

	return localctx
}

// IDefaultValueContext is an interface to support dynamic dispatch.
type IDefaultValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultValueContext differentiates from other interfaces.
	IsDefaultValueContext()
}

type DefaultValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValueContext() *DefaultValueContext {
	var p = new(DefaultValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_defaultValue
	return p
}

func (*DefaultValueContext) IsDefaultValueContext() {}

func NewDefaultValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValueContext {
	var p = new(DefaultValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_defaultValue

	return p
}

func (s *DefaultValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValueContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DefaultValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDefaultValue(s)
	}
}

func (s *DefaultValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDefaultValue(s)
	}
}

func (p *GraphQLParser) DefaultValue() (localctx IDefaultValueContext) {
	localctx = NewDefaultValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GraphQLParserRULE_defaultValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(GraphQLParserT__13)
	}
	{
		p.SetState(199)
		p.Value()
	}

	return localctx
}

// IValueOrVariableContext is an interface to support dynamic dispatch.
type IValueOrVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueOrVariableContext differentiates from other interfaces.
	IsValueOrVariableContext()
}

type ValueOrVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueOrVariableContext() *ValueOrVariableContext {
	var p = new(ValueOrVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_valueOrVariable
	return p
}

func (*ValueOrVariableContext) IsValueOrVariableContext() {}

func NewValueOrVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueOrVariableContext {
	var p = new(ValueOrVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_valueOrVariable

	return p
}

func (s *ValueOrVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueOrVariableContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueOrVariableContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ValueOrVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueOrVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueOrVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterValueOrVariable(s)
	}
}

func (s *ValueOrVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitValueOrVariable(s)
	}
}

func (p *GraphQLParser) ValueOrVariable() (localctx IValueOrVariableContext) {
	localctx = NewValueOrVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GraphQLParserRULE_valueOrVariable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__14, GraphQLParserSTRING, GraphQLParserBOOLEAN, GraphQLParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.Value()
		}

	case GraphQLParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Variable()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringValueContext struct {
	*ValueContext
}

func NewStringValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValueContext {
	var p = new(StringValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(GraphQLParserSTRING, 0)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitStringValue(s)
	}
}

type BooleanValueContext struct {
	*ValueContext
}

func NewBooleanValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanValueContext {
	var p = new(BooleanValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *BooleanValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(GraphQLParserBOOLEAN, 0)
}

func (s *BooleanValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterBooleanValue(s)
	}
}

func (s *BooleanValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitBooleanValue(s)
	}
}

type NumberValueContext struct {
	*ValueContext
}

func NewNumberValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberValueContext {
	var p = new(NumberValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NumberValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNUMBER, 0)
}

func (s *NumberValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterNumberValue(s)
	}
}

func (s *NumberValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitNumberValue(s)
	}
}

type ArrayValueContext struct {
	*ValueContext
}

func NewArrayValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayValueContext {
	var p = new(ArrayValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

func (p *GraphQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GraphQLParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(209)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserSTRING:
		localctx = NewStringValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(GraphQLParserSTRING)
		}

	case GraphQLParserNUMBER:
		localctx = NewNumberValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(GraphQLParserNUMBER)
		}

	case GraphQLParserBOOLEAN:
		localctx = NewBooleanValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(207)
			p.Match(GraphQLParserBOOLEAN)
		}

	case GraphQLParserT__14:
		localctx = NewArrayValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(208)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *Type_Context) NonNullType() INonNullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonNullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonNullTypeContext)
}

func (s *Type_Context) ListType() IListTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *GraphQLParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GraphQLParserRULE_type_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.TypeName()
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__16 {
			{
				p.SetState(212)
				p.NonNullType()
			}

		}

	case GraphQLParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.ListType()
		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__16 {
			{
				p.SetState(216)
				p.NonNullType()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (p *GraphQLParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GraphQLParserRULE_typeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(GraphQLParserNAME)
	}

	return localctx
}

// IListTypeContext is an interface to support dynamic dispatch.
type IListTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListTypeContext differentiates from other interfaces.
	IsListTypeContext()
}

type ListTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListTypeContext() *ListTypeContext {
	var p = new(ListTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_listType
	return p
}

func (*ListTypeContext) IsListTypeContext() {}

func NewListTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeContext {
	var p = new(ListTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_listType

	return p
}

func (s *ListTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitListType(s)
	}
}

func (p *GraphQLParser) ListType() (localctx IListTypeContext) {
	localctx = NewListTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GraphQLParserRULE_listType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(GraphQLParserT__14)
	}
	{
		p.SetState(224)
		p.Type_()
	}
	{
		p.SetState(225)
		p.Match(GraphQLParserT__15)
	}

	return localctx
}

// INonNullTypeContext is an interface to support dynamic dispatch.
type INonNullTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNonNullTypeContext differentiates from other interfaces.
	IsNonNullTypeContext()
}

type NonNullTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonNullTypeContext() *NonNullTypeContext {
	var p = new(NonNullTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_nonNullType
	return p
}

func (*NonNullTypeContext) IsNonNullTypeContext() {}

func NewNonNullTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonNullTypeContext {
	var p = new(NonNullTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_nonNullType

	return p
}

func (s *NonNullTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *NonNullTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonNullTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonNullTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterNonNullType(s)
	}
}

func (s *NonNullTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitNonNullType(s)
	}
}

func (p *GraphQLParser) NonNullType() (localctx INonNullTypeContext) {
	localctx = NewNonNullTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GraphQLParserRULE_nonNullType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(GraphQLParserT__16)
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *GraphQLParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GraphQLParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Match(GraphQLParserT__14)
		}
		{
			p.SetState(230)
			p.Value()
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraphQLParserT__1 {
			{
				p.SetState(231)
				p.Match(GraphQLParserT__1)
			}
			{
				p.SetState(232)
				p.Value()
			}

			p.SetState(237)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(238)
			p.Match(GraphQLParserT__15)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(240)
			p.Match(GraphQLParserT__14)
		}
		{
			p.SetState(241)
			p.Match(GraphQLParserT__15)
		}

	}

	return localctx
}
