package model

type Payment struct {
	TableNumber  int    `db:"table_number"`
	OrderName    string `db:"person"`
	FoodCode     string `db:"food_code"`
	FoodName     string
	FoodPrice    int `db:"price"`
	TotalInvoice int
}

func (p *Payment) GetTableNumber() int {
	return p.TableNumber
}

func (p *Payment) GetOrderName() string {
	return p.OrderName
}

func (p *Payment) GetFoodCode() string {
	return p.FoodCode
}

func (p *Payment) GetFoodName() string {
	return p.FoodName
}

func (p *Payment) GetFoodPrice() int {
	return p.FoodPrice
}

func (p *Payment) GetTotalInvoice() int {
	return p.TotalInvoice
}

func (p *Payment) SetTableNumber(code int) {
	p.TableNumber = code
}

func (p *Payment) SetPerson(code string) {
	p.OrderName = code
}

func (p *Payment) SetFoodCode(code string) {
	p.FoodCode = code
}

func (p *Payment) SetFoodName(code string) {
	p.FoodName = code
}

func (p *Payment) SetFoodPrice(code int) {
	p.FoodPrice = code
}

func (p *Payment) SetTotalInvoice(code int) {
	p.TotalInvoice = code
}

func NewPayment(tableNumber int) Payment {
	return Payment{
		TableNumber: tableNumber,
	}
}
