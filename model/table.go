package model

type Table struct {
	Number int  `db:"id"`
	Status bool `db:"available_status"`
}

func (t *Table) GetTableNumber() int {
	return t.Number
}

func (t *Table) GetTableStatus() bool {
	return t.Status
}

func (t *Table) SetTableNumber(code int) {
	t.Number = code
}

func (t *Table) SetTableStatus(code bool) {
	t.Status = code
}
