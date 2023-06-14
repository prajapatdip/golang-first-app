package main

import "fmt"

func main() {
	customers := getCustomerData()

	for _, customer := range customers {

		fmt.Println(customer)
	}
}

// I/O function syntax
// func getData(inputs) (outputs){
// }

func getData() (customers []string) {
	// Slice inisializtion
	// customers = [] string
	customers = []string{"Dip", "Nirav"}

	customers = append(customers, "Chaitanya")
	customers = append(customers, "Sahil")
	customers = append(customers, "Ajit")

	// for i := 0; i < len(customers); i++ {
	// 	fmt.Println(customers[i])
	// }

	// for each loop
	for i, customer := range customers {
		customer = customers[i]
		fmt.Println(customer)
	}

	return customers
}
