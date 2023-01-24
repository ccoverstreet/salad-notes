package databse

type DatabaseHandle interface {
	GetRecordUID(key string) any
}
