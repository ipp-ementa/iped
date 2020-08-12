# iped
IPED (ipp-ementa distributor) component

[![codecov](https://codecov.io/gh/ipp-ementa/iped/branch/master/graph/badge.svg)](https://codecov.io/gh/ipp-ementa/iped)

## Documentation

IPED documentation can be found [here](https://github.com/ipp-ementa/iped-documentation)

## Requirements

- Deno `v1.2.2` installed (Version which the service was developed, may work with other versions)

## Dependencies

|Name|Used for|
|----|--------|
|[monads](https://deno.land/x/monads@v0.3.4)|Functional monads to get rid of nasty exceptions and strenghen functions output|
|[mongo](https://deno.land/x/mongo@v0.10.0/mod.ts)|ORM like for MongoDB|
|[oak](https://deno.land/x/oak/)|Serve web api|
|[uuid](https://deno.land/std/uuid/mod.ts)|Generate UUID v4 tokens|

To install the dependencies run the following command:

`deno cache --unstable deps.ts`

## Run

The following environment variables are required to run and deploy IPED:

|Variable|Description|
|--------|-----------|
|PORT|The port which the web server will be available|
|MONGO_DB_CONNECTION_STRING|Connection String of MongoDB database|

Then execute the following command:

`deno run --allow-net --allow-write --allow-read --allow-plugin --allow-env --unstable mod.ts --port=${PORT}`

### Changelog

Version: `1.2`

- API collections are now preceeded by `api/`
- Canteen now references its geographical location
- 100% unit tested models
- 75.53% unit tested controllers