package sqlClauses

import "testing"

func Test_CreateInClause(t *testing.T) {
	values := []string{"'a',", "'b',", "'c',"}
	sql := CreateInClause(values)
	if sql != "IN ('a','b','c')" {
		t.Error("Expected 'IN ('a','b','c')', got ", sql)
	}
}
