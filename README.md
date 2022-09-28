# wb-internship
## WB level 0 task

## Before getting started:

1. Make sure your database and database user configured correctly
2. Run database migrations
   
   ```
   migrate -path ./migrations -database {POSTGRES_URL} up
   ```
3. Make sure `nats-streaming` server is running
   
   ```
    docker run -p 4222:4222 -p 8222:8222 nats-streaming
   ```
4. Check `Makefile` for general commands


## Run service
```
make dev
```
