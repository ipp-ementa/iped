# iped
IPED (ipp-ementa distributor) component

[![codecov](https://codecov.io/gh/ipp-ementa/iped/branch/master/graph/badge.svg)](https://codecov.io/gh/ipp-ementa/iped)

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

### Changelog

Version: `1.0`

- All REST API functionalities implemented (as of 2019/09/07)
- 100% unit tested models
- 75.53% unit tested controllers