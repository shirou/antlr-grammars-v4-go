// Generated from GraphQL.g4 by ANTLR 4.6.

package parser // GraphQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGraphQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGraphQLVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitOperationDefinition(ctx *OperationDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitSelectionSet(ctx *SelectionSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitOperationType(ctx *OperationTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitSelection(ctx *SelectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitFieldName(ctx *FieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitFragmentSpread(ctx *FragmentSpreadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitInlineFragment(ctx *InlineFragmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitFragmentDefinition(ctx *FragmentDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitFragmentName(ctx *FragmentNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDirectives(ctx *DirectivesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDirective(ctx *DirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitTypeCondition(ctx *TypeConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitVariableDefinitions(ctx *VariableDefinitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitVariableDefinition(ctx *VariableDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDefaultValue(ctx *DefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitValueOrVariable(ctx *ValueOrVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitStringValue(ctx *StringValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitNumberValue(ctx *NumberValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitArrayValue(ctx *ArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitListType(ctx *ListTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitNonNullType(ctx *NonNullTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}
