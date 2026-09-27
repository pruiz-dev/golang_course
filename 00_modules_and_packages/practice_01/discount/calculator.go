package discount

const baseDiscount = 10

func Calculate(price int) int {
	return price - baseDiscount
}
