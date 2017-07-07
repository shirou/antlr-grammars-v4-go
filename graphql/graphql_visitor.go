// Generated from GraphQL.g4 by ANTLR 4.6.

package parser // GraphQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GraphQLParser.
type GraphQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GraphQLParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by GraphQLParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#operationDefinition.
	VisitOperationDefinition(ctx *OperationDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#selectionSet.
	VisitSelectionSet(ctx *SelectionSetContext) interface{}

	// Visit a parse tree produced by GraphQLParser#operationType.
	VisitOperationType(ctx *OperationTypeContext) interface{}

	// Visit a parse tree produced by GraphQLParser#selection.
	VisitSelection(ctx *SelectionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by GraphQLParser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by GraphQLParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by GraphQLParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by GraphQLParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by GraphQLParser#fragmentSpread.
	VisitFragmentSpread(ctx *FragmentSpreadContext) interface{}

	// Visit a parse tree produced by GraphQLParser#inlineFragment.
	VisitInlineFragment(ctx *InlineFragmentContext) interface{}

	// Visit a parse tree produced by GraphQLParser#fragmentDefinition.
	VisitFragmentDefinition(ctx *FragmentDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#fragmentName.
	VisitFragmentName(ctx *FragmentNameContext) interface{}

	// Visit a parse tree produced by GraphQLParser#directives.
	VisitDirectives(ctx *DirectivesContext) interface{}

	// Visit a parse tree produced by GraphQLParser#directive.
	VisitDirective(ctx *DirectiveContext) interface{}

	// Visit a parse tree produced by GraphQLParser#typeCondition.
	VisitTypeCondition(ctx *TypeConditionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#variableDefinitions.
	VisitVariableDefinitions(ctx *VariableDefinitionsContext) interface{}

	// Visit a parse tree produced by GraphQLParser#variableDefinition.
	VisitVariableDefinition(ctx *VariableDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by GraphQLParser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#valueOrVariable.
	VisitValueOrVariable(ctx *ValueOrVariableContext) interface{}

	// Visit a parse tree produced by GraphQLParser#stringValue.
	VisitStringValue(ctx *StringValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#numberValue.
	VisitNumberValue(ctx *NumberValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#booleanValue.
	VisitBooleanValue(ctx *BooleanValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#arrayValue.
	VisitArrayValue(ctx *ArrayValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by GraphQLParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by GraphQLParser#listType.
	VisitListType(ctx *ListTypeContext) interface{}

	// Visit a parse tree produced by GraphQLParser#nonNullType.
	VisitNonNullType(ctx *NonNullTypeContext) interface{}

	// Visit a parse tree produced by GraphQLParser#array.
	VisitArray(ctx *ArrayContext) interface{}
}
