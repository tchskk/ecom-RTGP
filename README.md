# ecom-RTGP

## Backend - GoLang
- Gin: handles HTTP requests (routes, controllers)
- GORM (Object-Relational Mapping): interacts with the database in an ORM style
- dotenv: loads environment variables (DB URL, JWT secret)
- JWT: handles secure login sessions (tokens instead of cookies)

email: user1@test.com
password: secret123


## PSQL Commands Used

```
CREATE DATABASE myapp;

CREATE USER myapp_user WITH ENCRYPTED PASSWORD 'mypassword';

GRANT ALL PRIVILEGES ON DATABASE myapp TO myapp_user;

GRANT USAGE, CREATE ON SCHEMA public TO myapp_user;
```
