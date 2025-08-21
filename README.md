# wallet-ex1

- Link Google Drive: 

## Config
- Check configuration in here: conf/config.yml
- Database in root directory with filename klikcair.sql

## Install
- go mod tidy
- Don't forget check database connection.
- Make sure you in root directory
- Run command: ``make run``
- Open postmen
- Import postman JSON: ``klikcair.postman_collection.json``
- And you can testing the endpoints

## API(s)

```
[GET] /api/v1/wallets/balance/{user_id}


[POST] /api/v1/transactions/withdraw
{
    "code": "qwerty",
    "amount": 9000
}

```