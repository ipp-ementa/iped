# iped
IPED (ipp-ementa distributor) component

## Documentation

IPED documentation can be found at [here](https://github.com/ipp-ementa/iped-documentation)

## Dependencies

|Name|Used for|
|----|--------|
|[gorm](https://github.com/jinzhu/gorm)|Map models into database and perform database operations|
|[echo](https://github.com/labstack/echo)|Launch REST API on a web server as well provide middleware functions handle|

## Deploy

The following environment variables are required to deploy IPED

|Variable|Description|
|--------|-----------|
|IPED_PORT|The port which the web server will be available|
|IPEW_CONNECTION_STRING|Absolute path to IPEW SQLite database file|