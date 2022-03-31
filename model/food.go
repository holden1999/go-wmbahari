package model

type Food struct {
	FoodCode  string `db:"code"`
	FoodName  string `db:"name"`
	FoodPrice int    `db:"price"`
}

func (f *Food) SetFoodCode(code string) {
	f.FoodCode = code
}

func (f *Food) SetFoodName(name string) {
	f.FoodName = name
}
func (f *Food) SetFoodPrice(price int) {
	f.FoodPrice = price
}
func (f *Food) GetFoodCode() string {
	return f.FoodCode
}

func (f *Food) GetFoodName() string {
	return f.FoodName
}

func (f *Food) GetFoodPrice() int {
	return f.FoodPrice
}
