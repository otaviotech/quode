# Quode

Yet another quotes app. (this one is about coding quotes)

I created this project just to practice Go. With this I aim to apply (in Go) some of the principles that are language agnostic, like clean code, solid, hexagonal architecture, but in an idiomatic "Go way".

What exaclty I want to practice?
- Go syntax (in an idiomatic way)
- Go modules
- Concurrency (of course)
- Contexts (of course)
- Error handling
- Testing
- Best practices (clean code, solid, etc)
- Clean / Hexagonal architecture
- Domain-driven design
- Databases (migrations, transactions, etc)

## Tools

| Tool | Purpose |
| - | - |
| Viper | Configuration |
| Testify | Testing |
| Sqlc | SQL queries code generation |
| golang-migrate | Database migration |
| Wire | Dependency injection |
| Chi | HTTP routing |
| uuid | UUID v4 generation |

## Running

To run the app, you need to setup some environment variables. The file `.env.example` is as a template.

Then you can use docker-compose to spin up a database and then run the app:

```bash
cp .env.example .env
docker-compose up -d db
make migrate
make start 
```

## Testing

To run the tests, you need to have a database running (because we have some integration tests at the infra layer). 
You can use docker-compose to spin up a test database and then run the tests with or without coverage:

```bash
docker-compose up -d db_test
make migratetest 
make test
# or
make testcoverage
```

## Todo
- [ ] Add a continuous integration pipeline using Github Actions
- [ ] Add OpenAPI/Swagger docs
- [ ] Add a gRPC server
- [ ] Add a GraphQL server