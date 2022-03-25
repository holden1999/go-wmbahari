package model

type FoodList struct {
	FoodCode  string `db:"code"`
	FoodName  string `db:"name"`
	FoodPrice int    `db:"price"`
}

func (f *FoodList) SetFoodCode(code string) {
	f.FoodCode = code
}

func (f *FoodList) SetFoodName(name string) {
	f.FoodName = name
}
func (f *FoodList) SetFoodPrice(price int) {
	f.FoodPrice = price
}
func (f *FoodList) GetFoodCode() string {
	return f.FoodCode
}

func (f *FoodList) GetFoodName() string {
	return f.FoodName
}

func (f *FoodList) GetFoodPrice() int {
	return f.FoodPrice
}
