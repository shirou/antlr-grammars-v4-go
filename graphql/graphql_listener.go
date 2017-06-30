// Generated from GraphQL.g4 by ANTLR 4.6.

package parser // GraphQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GraphQLListener is a complete listener for a parse tree produced by GraphQLParser.
type GraphQLListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterOperationDefinition is called when entering the operationDefinition production.
	EnterOperationDefinition(c *OperationDefinitionContext)

	// EnterSelectionSet is called when entering the selectionSet production.
	EnterSelectionSet(c *SelectionSetContext)

	// EnterOperationType is called when entering the operationType production.
	EnterOperationType(c *OperationTypeContext)

	// EnterSelection is called when entering the selection production.
	EnterSelection(c *SelectionContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterFragmentSpread is called when entering the fragmentSpread production.
	EnterFragmentSpread(c *FragmentSpreadContext)

	// EnterInlineFragment is called when entering the inlineFragment production.
	EnterInlineFragment(c *InlineFragmentContext)

	// EnterFragmentDefinition is called when entering the fragmentDefinition production.
	EnterFragmentDefinition(c *FragmentDefinitionContext)

	// EnterFragmentName is called when entering the fragmentName production.
	EnterFragmentName(c *FragmentNameContext)

	// EnterDirectives is called when entering the directives production.
	EnterDirectives(c *DirectivesContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterTypeCondition is called when entering the typeCondition production.
	EnterTypeCondition(c *TypeConditionContext)

	// EnterVariableDefinitions is called when entering the variableDefinitions production.
	EnterVariableDefinitions(c *VariableDefinitionsContext)

	// EnterVariableDefinition is called when entering the variableDefinition production.
	EnterVariableDefinition(c *VariableDefinitionContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterDefaultValue is called when entering the defaultValue production.
	EnterDefaultValue(c *DefaultValueContext)

	// EnterValueOrVariable is called when entering the valueOrVariable production.
	EnterValueOrVariable(c *ValueOrVariableContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterNumberValue is called when entering the numberValue production.
	EnterNumberValue(c *NumberValueContext)

	// EnterBooleanValue is called when entering the booleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// EnterArrayValue is called when entering the arrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterListType is called when entering the listType production.
	EnterListType(c *ListTypeContext)

	// EnterNonNullType is called when entering the nonNullType production.
	EnterNonNullType(c *NonNullTypeContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitOperationDefinition is called when exiting the operationDefinition production.
	ExitOperationDefinition(c *OperationDefinitionContext)

	// ExitSelectionSet is called when exiting the selectionSet production.
	ExitSelectionSet(c *SelectionSetContext)

	// ExitOperationType is called when exiting the operationType production.
	ExitOperationType(c *OperationTypeContext)

	// ExitSelection is called when exiting the selection production.
	ExitSelection(c *SelectionContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitFragmentSpread is called when exiting the fragmentSpread production.
	ExitFragmentSpread(c *FragmentSpreadContext)

	// ExitInlineFragment is called when exiting the inlineFragment production.
	ExitInlineFragment(c *InlineFragmentContext)

	// ExitFragmentDefinition is called when exiting the fragmentDefinition production.
	ExitFragmentDefinition(c *FragmentDefinitionContext)

	// ExitFragmentName is called when exiting the fragmentName production.
	ExitFragmentName(c *FragmentNameContext)

	// ExitDirectives is called when exiting the directives production.
	ExitDirectives(c *DirectivesContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitTypeCondition is called when exiting the typeCondition production.
	ExitTypeCondition(c *TypeConditionContext)

	// ExitVariableDefinitions is called when exiting the variableDefinitions production.
	ExitVariableDefinitions(c *VariableDefinitionsContext)

	// ExitVariableDefinition is called when exiting the variableDefinition production.
	ExitVariableDefinition(c *VariableDefinitionContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitDefaultValue is called when exiting the defaultValue production.
	ExitDefaultValue(c *DefaultValueContext)

	// ExitValueOrVariable is called when exiting the valueOrVariable production.
	ExitValueOrVariable(c *ValueOrVariableContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitNumberValue is called when exiting the numberValue production.
	ExitNumberValue(c *NumberValueContext)

	// ExitBooleanValue is called when exiting the booleanValue production.
	ExitBooleanValue(c *BooleanValueContext)

	// ExitArrayValue is called when exiting the arrayValue production.
	ExitArrayValue(c *ArrayValueContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitListType is called when exiting the listType production.
	ExitListType(c *ListTypeContext)

	// ExitNonNullType is called when exiting the nonNullType production.
	ExitNonNullType(c *NonNullTypeContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)
}
