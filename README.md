# Money Exchanger

It's a simple microservice that convert money from one currency to another by actual exchange rate.

## How to run?

run next command:

```shell
make docker/up
```

> Make sure you have install dependencies.  
> If not, you should use `make deps`

## Features

- [x] Use Memcached for key-value database.
- [x] If data expires in database, fetch it from external API - **[currencyfreaks.com](https://currencyfreaks.com/)**
