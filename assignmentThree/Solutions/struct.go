package Solutions

type Details struct {
	Name        string
	Age         int
	PhoneNumber int
	Address     Address
}

type Address struct {
	City    string
	State   string
	Pincode int
}
