module github.com/nlively/gosnake/client

go 1.23.2


require (
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/nlively/gosnake/common v0.0.0
    github.com/nlively/gosnake/server v0.0.0
)

replace github.com/nlively/gosnake/common => ../common
replace github.com/nlively/gosnake/server => ../server
