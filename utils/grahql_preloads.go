package utils

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

func GetPreloads(ctx context.Context, satisfies []string) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, satisfies),
		satisfies,
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, satisfies []string, prefix string) (preloads []string) {
	for _, column := range fields {
		fmt.Println([]interface{}{
			"name:",
			column.Name,
			"Type:",
			column.Definition.Type.Name(),
			"Object Name:",
			column.ObjectDefinition.Name,
			"One of:",
			column.ObjectDefinition.OneOf(satisfies...),
		})
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, satisfies), satisfies, prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}
