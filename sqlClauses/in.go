package sqlClauses

import "strings"

func CreateInClause(values []string) string {
	sqlInStart := "IN ("
	sqlInClose := ")"
	return sqlInStart + strings.TrimSuffix(strings.Join(values, ""), ",") + sqlInClose
}
