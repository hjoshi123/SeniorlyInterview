# Seniorly Interview

## Running

* Create an .env file in home directory with the following variables POSTGRES_USER,POSTGRES_PASSWORD,POSTGRES_DB
* `source .env`
* `docker-compose up -d --build`

Note that when first started, the `server` container might not start (it might be stopped) the reason benig the database takes some time to startup. this can be avoided using healthchecks but for brevity I have not done it

## Endpoints
* `/buy_pizza` - placing an order takes in mobile number, pizza type. Note that once the order is in preparing stage you cannot place another order  till its done.
* `/track_pizza` -  tracking the status through user's mobile number
* `/update_status` - updating status from a chefk through order id and status