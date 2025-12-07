# ecom-RTGP

A full-stack e-commerce platform with authentication and admin dashboard built with React, Tailwind CSS, Golang, and PostgreSQL.

## Features

- 🔐 User Authentication (Register/Login)
- 👤 User Profile Management
- 🔒 Protected Routes
- 👑 Admin Dashboard
- 📊 User Management
- 🎨 Modern UI with Tailwind CSS
- 🚀 Fast and Responsive

## Tech Stack

### Frontend
- React 18 with TypeScript
- Vite (Build Tool)
- React Router v6
- Tailwind CSS
- Axios for API calls

### Backend
- Golang
- Gin Web Framework
- PostgreSQL Database
- JWT Authentication
- Bcrypt Password Hashing

## Project Structure

```
ecom-RTGP/
├── frontend/               # React frontend application
│   ├── src/
│   │   ├── components/     # React components
│   │   ├── contexts/       # React contexts (Auth)
│   │   ├── pages/          # Page components
│   │   └── services/       # API services
│   └── package.json
├── backend/                # Golang backend application
│   ├── cmd/api/           # Application entry point
│   ├── config/            # Configuration files
│   ├── db/                # Database connection and queries
│   ├── handlers/          # HTTP request handlers
│   ├── middleware/        # HTTP middlewares
│   ├── models/            # Data models
│   ├── routes/            # Route definitions
│   └── utils/             # Utility functions
└── docker-compose.yml     # PostgreSQL Docker setup
```

## Prerequisites

- Node.js (v18 or higher)
- Go (v1.21 or higher)
- Docker and Docker Compose
- PostgreSQL (or use Docker Compose)

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/tchskk/ecom-RTGP.git
cd ecom-RTGP
```

### 2. Start PostgreSQL Database

```bash
docker-compose up -d
```

This will start a PostgreSQL database on `localhost:5432` with:
- Database: `ecom_rtgp`
- User: `postgres`
- Password: `postgres`

### 3. Setup Backend

```bash
cd backend

# Copy environment file
cp .env.example .env

# Install dependencies
go mod download

# Build the application
go build -o bin/api ./cmd/api

# Run the server
./bin/api
```

The backend will start on `http://localhost:8080`

### 4. Setup Frontend

```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

The frontend will start on `http://localhost:5173`

## Environment Variables

### Backend (.env)

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ecom_rtgp
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
PORT=8080
```

## API Endpoints

### Authentication
- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login user

### Protected Routes (Requires Authentication)
- `GET /api/profile` - Get user profile

### Admin Routes (Requires Admin Role)
- `GET /api/admin/users` - Get all users

## Usage

### Register a New User

1. Open the frontend at `http://localhost:5173`
2. Click "Register here" on the login page
3. Fill in your details:
   - First Name
   - Last Name
   - Email
   - Password (min 6 characters)
4. Click "Register"

### Login

1. Enter your email and password
2. Click "Login"
3. You'll be redirected to your dashboard

### Access Admin Dashboard

To access the admin dashboard, you need to manually update a user's role to "admin" in the database:

```sql
-- Connect to your PostgreSQL database
docker exec -it ecom_rtgp_db psql -U postgres -d ecom_rtgp

-- Update user role to admin
UPDATE users SET role = 'admin' WHERE email = 'your-email@example.com';
```

After updating the role, logout and login again to see the admin features.

## Development

### Frontend Development

```bash
cd frontend
npm run dev       # Start development server
npm run build     # Build for production
npm run preview   # Preview production build
```

### Backend Development

```bash
cd backend
go run cmd/api/main.go    # Run with hot reload (requires air)
go build -o bin/api ./cmd/api  # Build binary
go test ./...               # Run tests
```

## Database Schema

### Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    role VARCHAR(20) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Future Enhancements

- [ ] Product catalog
- [ ] Shopping cart
- [ ] Order management
- [ ] Payment integration
- [ ] Product search and filtering
- [ ] User reviews and ratings
- [ ] Email notifications
- [ ] Password reset functionality
- [ ] Profile image upload
- [ ] Multi-language support

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

## Contact

For questions or support, please open an issue on GitHub.