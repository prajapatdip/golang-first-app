package main

type (
	Customer struct {
		FirstName string
		LastName  string
	}
)

func getCustomerData() (customers []Customer) {
	dip := Customer{FirstName: "Dip", LastName: "Prajapat"}

	customers = append(customers, dip,
		Customer{FirstName: "Nirav", LastName: "Prajapat"},
		Customer{FirstName: "Aniket", LastName: "Vegda"},
		Customer{FirstName: "Dhananjay", LastName: "Patil"},
	)

	return customers
}
