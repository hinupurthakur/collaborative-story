package api

import (
	"fmt"
	"strings"
)

/*
generateOrderByClause joins column
with specified order
to generate Order by clause for DB query
*/
func genrateOrderByClause(column []string, order []string) string {
	condition := ""
	for i, j := 0, 0; i < len(column) || j < len(order); i, j = i+1, j+1 {
		if i < len(column) {
			condition += column[i] + " "
		}
		if j < len(order) {
			condition += order[j] + ","
		}
	}
	condition = strings.Trim(condition, ",")
	condition = strings.Trim(condition, " ")
	if len(condition) == 0 {
		return condition
	}
	return fmt.Sprintf("ORDER BY %s", condition)
}
