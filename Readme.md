## Getting started

### 1. Create environment file

  Create ```postgres/postgres.env``` using ```postgres/postgres.env.sample``` as template

  Optional: Create ```pgadmin/pgadmin.env``` using ```pgadmin/pgadmin.env.sample``` as template

### 2. Create config file

  Create ```server/config.yml``` using ```server/config.yml.sample``` as a template. Use the values defined in ```postgres/postgres.env```

### 3. Starting PostgreSQL

  ```
  docker compose up -d postgres
  ``` 

### 4. Run migrations

  ```
  migrate -database "postgres://user:password@localhost:5432/database?sslmode=disable" -path db/migrations up
  ```
