package models

type PgExceptionStatusCode = string

const (
	NotFound         PgExceptionStatusCode = "not_found"
	Unknown          PgExceptionStatusCode = "unknown"
	UnsupportedModel PgExceptionStatusCode = "unsuported_model"
)
