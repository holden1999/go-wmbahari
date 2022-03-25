package model

type TableList struct {
	Number int  `db:"id"`
	Status bool `db:"available_status"`
}

func (t *TableList) GetTableNumber() int {
	return t.Number
}

func (t *TableList) GetTableStatus() bool {
	return t.Status
}

func (t *TableList) SetTableNumber(code int) {
	t.Number = code
}

func (t *TableList) SetTableStatus(code bool) {
	t.Status = code
}
