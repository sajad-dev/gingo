package config

type AppConfig struct {
	APP_NAME string `defualt:"GOLANG_APP"`
	DESCRIPTION string `defualt:""`
	AUTHOR  string `defualt:"Sajad pourajam"`
	DEBUG string `defualt:"true"`
	JWT string `defualt:""`
	
	DATABASE_NAME string `defualt:"GOLANG_APP"`
	DATABASE_USER string `defualt:"root"`
	DATABASE_PASSWORD string `defualt:"root"`
	DATABASE_PORT string `defualt:"3306"`
	DATABASE_ADDRESS string `defualt:"127.0.0.1"`

	STORAGE_PATH string `defualt:"../../storage/file"`
}
