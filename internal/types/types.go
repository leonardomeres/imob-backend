package types

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	ListingLink string `json:"listingLink"`
	Notes       string `json:"notes"`
}
