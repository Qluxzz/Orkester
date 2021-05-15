package indexFiles

import "database/sql"

func CreateValidNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func CreateValidNullInt(n int) sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(n),
		Valid: true,
	}
}
