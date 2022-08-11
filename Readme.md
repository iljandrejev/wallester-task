# Wallester Task

Solution made by author

### Requirements
* Docker

### Installation
1. Copy `.env.example` to `.env` and fill it with own values
2. Run command in the terminal
```shell
docker-compose up -d
```
3. Browse application `localhost:8085`

### Database
In the `sql` folder located 2 sql scripts. 
* `migrate_up.sql` - will create customers table
* `seed.sql` - will fill customers table with the seed data