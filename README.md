<img src="https://raw.githubusercontent.com/vutratenko/pig/main/pig.png" alt="PIG" width="200"/>

# PIG - a dummy app

### Build/Test/Run

To build project hit:

`go build`

To run tests hit:

`go test ./...`

To run the app hit:

`go run main.go`



or run compiled binary



`./pig`


### Environment variables

| Variable | Description |
|----------|----------|
| PORT     | A port that the application is run on. Default is 8000   | 
| DATABASE | A dbstring for PostgreSQL DB. If not specified, service runs without DB | 



### Notes

Be also aware of ./resources directory which contains index HTML page. This should be copied alongside with ./pig binary as well with the same structure




