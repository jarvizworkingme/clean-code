# Structure Simple Project For clean code

## Requirements
This project implements the following functionalities:

- **GetOutstanding**: Returns the current outstanding amount on a loan, or 0 if there is no outstanding amount (e.g., the loan is closed).
- **IsDelinquent**: Determines if there has been more than 2 weeks of non-payment on the loan amount.
- **MakePayment**: Processes a payment of a specified amount on the loan.

## API Endpoints

### Create Customer
Creates a new customer with the specified email.

```bash
curl --location 'localhost:7778/api/v1/customer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@dummy.dummy"
}'
```

### Create Loan and Other Operations
Below are the steps for creating a loan and managing its lifecycle through different operations:

#### Create Loan
Creates a loan associated with the provided email, interest rate, amount, and tenor.
```bash
curl --location 'localhost:7778/api/v1/loan' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "1@dummy.dummy",
    "interest_rate": 10,
    "amount": 5000000,
    "tenor": 50
}'
```

#### Get Outstanding Balance
i'm following requirements to get the outstanding balance of the loan
```bash
Retrieves the current outstanding balance on a loan, returning 0 if no outstanding balance exists:

```bash
curl --location 'localhost:7778/api/v1/out-standing'
```

#### Get Payment Schedule
Retrieves the payment schedule based on user email. The `is_delinquent` parameter defaults to `false`. If set to `true`, it returns only the payments that are overdue by more than 2 weeks(i'm following requirements IsDelinquent):
```bash
curl --location 'localhost:7778/api/v1/payment/schedule?email=optional&is_delinquent=false'
```

#### Make a Payment
Processes a payment of the specified amount on the loan using the loan ID obtained from the payment schedule:
i'm following requirements to make a payment
```bash
curl --location 'localhost:7778/api/v1/payment' \
--header 'Content-Type: application/json' \
--data '{
    "loan_id": 1,
    "amount": 220000
}'
```
## Getting Started

### Running the Server

The server runs on port 7778 by default. For the database, I am using PostgreSQL. To start the database, navigate to the `docker` folder and run:

```sh
docker compose up -d
```

### Database Migration

Set the `.env` variable `MIGRATE_DB` to `true` to drop and create the database schema:

```sh
MIGRATE_DB=true
```

### Starting the Application

Run the following commands to start the application:

```sh
go mod init
go mod tidy
go run main.go
```

### Unit Tests

I have pooled the unit tests into a single suite located in the `service/suite_test.go`  folder.
```sh
go test -v ./api/service
```
