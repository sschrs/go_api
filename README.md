# Redis Installation for MacOS
```bash
brew install redis

redis-server
```

# Run API
Edit config.yaml file if necessary
```yaml
Api:
  Host: 'localhost'
  Port: 5000
  Prefork: true # to increase performance
Database:
  DatabaseName: 'data.db'
  PrepareStmt: true # to increase query performance
# change if necessary
Redis: 
  Host: 'localhost'
  Port: 6379
  Password: ''
  DB: 0
```

Install Go packages
```bash
go mod tidy
```

Run app
```bash
go run main.go
```

# Test Endpoints
## Generate Random Test Data
```http
POST http://127.0.0.1:5000/api/v1/generate_customers/10000000
```
This endpoint generates 10 million random record and inserts into database.

## Get Records
```http
GET http://127.0.0.1:5000/api/v1/customer?customer_id=CUST03683&date_start=2024-01-01&date_end=2024-12-31
```
This endpoint returns the customer data for given customer_id, date_start and date_end.

## Get Transaction Sum
```http
GET http://127.0.0.1:5000/api/v1/customer_transaction_summary?customer_id=CUST03683&date_start=2024-01-01&date_end=2024-12-31
```
This endpoint returns total sum of transaction_sum for given customer_id, date_start and date_end.



# go_api
