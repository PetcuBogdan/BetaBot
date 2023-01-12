package main


type jsonResponse struct {
	Error 	bool 	`json:"error"`
	Message string 	`json:"message"`
}

type ShopJson struct {
	Shop	string	`json:"shop"`
}

type ProductSupremeJson struct {
	Name 		string	`json:"productName"`
	Category 	string	`json:"category"`
	Colors		string  `json:"colors"`
} 

type ProductPalaceJson struct{
	Name 		string	`json:"productName"`
	Category 	string	`json:"category"`
}

type IdCard struct{
	Id	int	 	`json:"id"`
}

type IdProfile struct{
	Id	int	 	`json:"id"`
}

type IdTask struct{
	Id	int	 	`json:"id"`
}

type Card struct {
	Id 				int
    NameCard 		string
    CardNumber 		string
    ExpirationDate 	string
    Cvv 			string
}

type Profile struct {
	Id		   int	
	Lname      string
	Fname	   string	 
	Email      string
	Phone      string
	Address    string
	Address2   string
	Postcode   string
	City       string
	County	   string
	Country    string
}

type Task struct {
	Profile 		int
	Card 			int
	Id 				int
	Name 			string
	Shop 			string
	ProductName 	string
	Category 		string
	Size 			string
	Color 			string
	Status 			string
}


type CardJson struct {
	Id 				int    `json:"id"`
    NameCard 		string `json:"nameCard"`
    CardNumber 		string `json:"cardNumber"`
    ExpirationDate 	string `json:"expirationDate"`
    Cvv 			string `json:"cvv"`
}

type ProfileJson struct {
	Id		   int		`json:"id"`
	Lname      string	`json:"lname"`
	Fname	   string	`json:"fname"`
	Email      string	`json:"email"`
	Phone      string	`json:"phone"`
	Address    string	`json:"address"`
	Address2   string	`json:"address2"`
	Postcode   string	`json:"postcode"`
	City       string	`json:"city"`
	County	   string	`json:"county"`
	Country    string	`json:"country"`
}

type TaskJson struct {
	Id 				int				`json:"id"`
	Name 			string			`json:"name"`
	Shop 			string			`json:"shop"`
	ProductName 	string			`json:"productName"`
	Category 		string			`json:"category"`
	Size 			string			`json:"size"`
	Color 			string			`json:"color"`
	Profile 		int				`json:"profile"`
	Card 			int				`json:"card"`
	Status 			string			`json:"status"`
}

type ReadyTask struct {
	Task 		Task
	Profile 	Profile
	Card 		Card
}