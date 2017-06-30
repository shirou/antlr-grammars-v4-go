// Generated from GraphQL.g4 by ANTLR 4.6.

package parser // GraphQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGraphQLListener is a complete listener for a parse tree produced by GraphQLParser.
type BaseGraphQLListener struct{}

var _ GraphQLListener = &BaseGraphQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGraphQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGraphQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGraphQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGraphQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseGraphQLListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseGraphQLListener) ExitDocument(ctx *DocumentContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseGraphQLListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseGraphQLListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterOperationDefinition is called when production operationDefinition is entered.
func (s *BaseGraphQLListener) EnterOperationDefinition(ctx *OperationDefinitionContext) {}

// ExitOperationDefinition is called when production operationDefinition is exited.
func (s *BaseGraphQLListener) ExitOperationDefinition(ctx *OperationDefinitionContext) {}

// EnterSelectionSet is called when production selectionSet is entered.
func (s *BaseGraphQLListener) EnterSelectionSet(ctx *SelectionSetContext) {}

// ExitSelectionSet is called when production selectionSet is exited.
func (s *BaseGraphQLListener) ExitSelectionSet(ctx *SelectionSetContext) {}

// EnterOperationType is called when production operationType is entered.
func (s *BaseGraphQLListener) EnterOperationType(ctx *OperationTypeContext) {}

// ExitOperationType is called when production operationType is exited.
func (s *BaseGraphQLListener) ExitOperationType(ctx *OperationTypeContext) {}

// EnterSelection is called when production selection is entered.
func (s *BaseGraphQLListener) EnterSelection(ctx *SelectionContext) {}

// ExitSelection is called when production selection is exited.
func (s *BaseGraphQLListener) ExitSelection(ctx *SelectionContext) {}

// EnterField is called when production field is entered.
func (s *BaseGraphQLListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseGraphQLListener) ExitField(ctx *FieldContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseGraphQLListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseGraphQLListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseGraphQLListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseGraphQLListener) ExitAlias(ctx *AliasContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGraphQLListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGraphQLListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseGraphQLListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseGraphQLListener) ExitArgument(ctx *ArgumentContext) {}

// EnterFragmentSpread is called when production fragmentSpread is entered.
func (s *BaseGraphQLListener) EnterFragmentSpread(ctx *FragmentSpreadContext) {}

// ExitFragmentSpread is called when production fragmentSpread is exited.
func (s *BaseGraphQLListener) ExitFragmentSpread(ctx *FragmentSpreadContext) {}

// EnterInlineFragment is called when production inlineFragment is entered.
func (s *BaseGraphQLListener) EnterInlineFragment(ctx *InlineFragmentContext) {}

// ExitInlineFragment is called when production inlineFragment is exited.
func (s *BaseGraphQLListener) ExitInlineFragment(ctx *InlineFragmentContext) {}

// EnterFragmentDefinition is called when production fragmentDefinition is entered.
func (s *BaseGraphQLListener) EnterFragmentDefinition(ctx *FragmentDefinitionContext) {}

// ExitFragmentDefinition is called when production fragmentDefinition is exited.
func (s *BaseGraphQLListener) ExitFragmentDefinition(ctx *FragmentDefinitionContext) {}

// EnterFragmentName is called when production fragmentName is entered.
func (s *BaseGraphQLListener) EnterFragmentName(ctx *FragmentNameContext) {}

// ExitFragmentName is called when production fragmentName is exited.
func (s *BaseGraphQLListener) ExitFragmentName(ctx *FragmentNameContext) {}

// EnterDirectives is called when production directives is entered.
func (s *BaseGraphQLListener) EnterDirectives(ctx *DirectivesContext) {}

// ExitDirectives is called when production directives is exited.
func (s *BaseGraphQLListener) ExitDirectives(ctx *DirectivesContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseGraphQLListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseGraphQLListener) ExitDirective(ctx *DirectiveContext) {}

// EnterTypeCondition is called when production typeCondition is entered.
func (s *BaseGraphQLListener) EnterTypeCondition(ctx *TypeConditionContext) {}

// ExitTypeCondition is called when production typeCondition is exited.
func (s *BaseGraphQLListener) ExitTypeCondition(ctx *TypeConditionContext) {}

// EnterVariableDefinitions is called when production variableDefinitions is entered.
func (s *BaseGraphQLListener) EnterVariableDefinitions(ctx *VariableDefinitionsContext) {}

// ExitVariableDefinitions is called when production variableDefinitions is exited.
func (s *BaseGraphQLListener) ExitVariableDefinitions(ctx *VariableDefinitionsContext) {}

// EnterVariableDefinition is called when production variableDefinition is entered.
func (s *BaseGraphQLListener) EnterVariableDefinition(ctx *VariableDefinitionContext) {}

// ExitVariableDefinition is called when production variableDefinition is exited.
func (s *BaseGraphQLListener) ExitVariableDefinition(ctx *VariableDefinitionContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseGraphQLListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseGraphQLListener) ExitVariable(ctx *VariableContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseGraphQLListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseGraphQLListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterValueOrVariable is called when production valueOrVariable is entered.
func (s *BaseGraphQLListener) EnterValueOrVariable(ctx *ValueOrVariableContext) {}

// ExitValueOrVariable is called when production valueOrVariable is exited.
func (s *BaseGraphQLListener) ExitValueOrVariable(ctx *ValueOrVariableContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseGraphQLListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseGraphQLListener) ExitStringValue(ctx *StringValueContext) {}

// EnterNumberValue is called when production numberValue is entered.
func (s *BaseGraphQLListener) EnterNumberValue(ctx *NumberValueContext) {}

// ExitNumberValue is called when production numberValue is exited.
func (s *BaseGraphQLListener) ExitNumberValue(ctx *NumberValueContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseGraphQLListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseGraphQLListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterArrayValue is called when production arrayValue is entered.
func (s *BaseGraphQLListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production arrayValue is exited.
func (s *BaseGraphQLListener) ExitArrayValue(ctx *ArrayValueContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseGraphQLListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseGraphQLListener) ExitType_(ctx *Type_Context) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseGraphQLListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseGraphQLListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseGraphQLListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseGraphQLListener) ExitListType(ctx *ListTypeContext) {}

// EnterNonNullType is called when production nonNullType is entered.
func (s *BaseGraphQLListener) EnterNonNullType(ctx *NonNullTypeContext) {}

// ExitNonNullType is called when production nonNullType is exited.
func (s *BaseGraphQLListener) ExitNonNullType(ctx *NonNullTypeContext) {}

// EnterArray is called when production array is entered.
func (s *BaseGraphQLListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseGraphQLListener) ExitArray(ctx *ArrayContext) {}
