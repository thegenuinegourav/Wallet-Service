package config

type Config struct {
	App 		App 		`json:"app"`
	Database 	Database 	`json:"database"`
	Cache 		Cache 		`json:"cache"`
}

type App struct {
	Port string `json:"port"`
}

type Database struct {
	Engine      string  `json:"engine"`
	Server		string	`json:"server"`
	Port 		string	`json:"port"`
	User 		string 	`json:"user"`
	Password 	string 	`json:"password"`
	Name 		string 	`json:"name"`
}

type Cache struct {
	Server		string		`json:"server"`
	Port 		string		`json:"port"`
	Password 	string 		`json:"password"`
	Wallet		Property 	`json:"wallet"`
	Idempotent 	Property 	`json:"idempotency"`
}

type Property struct {
	Db		int		`json:"db"`
	Expiry 	int	`json:"expiryInHours"`
}
