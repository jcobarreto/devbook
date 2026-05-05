# Devbook

A social media platform built with Go, featuring user authentication, posts, and social networking capabilities.

## Features

- **User Management**: Create, read, update, and delete user accounts
- **Authentication**: JWT-based authentication for secure API access
- **Posts**: Create, read, update, and delete posts with like functionality
- **Social Features**: Follow/unfollow users, view followers and following lists
- **Account Management**: Change password and delete account functionality
- **Password Security**: Hashed passwords using bcrypt for secure storage

## Architecture

Devbook consists of two main components:

- **API** (`/api`): RESTful API server built with Go and Gorilla Mux
- **Webapp** (`/webapp`): Web frontend for interacting with the API

## Tech Stack

### API
- **Go 1.26.0**
- **Gorilla Mux**: HTTP router and URL matcher
- **MySQL**: Database
- **JWT**: Authentication (dgrijalva/jwt-go)
- **Crypto**: Password hashing (golang.org/x/crypto)
- **Email Validation**: badoux/checkmail
- **Environment Configuration**: joho/godotenv

### Webapp
- **Go 1.26.0**
- **Gorilla Mux**: HTTP router
- **HTML Templates**: Server-side rendering
- **Environment Configuration**: joho/godotenv

## Prerequisites

- Go 1.26.0+ (for webapp) and Go 1.26.0+ (for api)
- MySQL 5.7+
- Git

## Installation

### 1. Clone the repository

```bash
git clone <repository-url>
cd devbook
```

### 2. Set up MySQL Database

```bash
# Start your MySQL server
mysql -u root -p < api/sql/sql.sql

# If needed, load seed data
mysql -u root -p devbook < api/sql/data.sql
```

Alternatively, if you have Docker:

```bash
docker run --name mysql-devbook -e MYSQL_ROOT_PASSWORD=root -d -p 3306:3306 mysql:latest
mysql -h 127.0.0.1 -u root -p -e "source api/sql/sql.sql"
```

### 3. Configure API

Create or update `/api/.env` with your configuration:

```env
DB_USERNAME=golang
DB_PASSWORD=golang
DB_NAME=devbook
API_PORT=5001
SECRET_KEY=<your-secret-key>
```

The default SECRET_KEY is provided in `.env`, but for production, generate a new one:

```bash
# Generate a secure random key (run from project root)
# Then base64 encode it and add to .env
```

### 4. Configure Webapp

Create or update `/webapp/.env` with your configuration:

```env
API_URL=http://localhost:5001
APP_PORT=3000
HASH_KEY=<your-hash-key>
BLOCK_KEY=<your-block-key>
```

The `HASH_KEY` and `BLOCK_KEY` are used for secure cookie encryption. You can generate random keys:

```bash
# Use any secure random key generation method and base64 encode
```

## Running the Application

### Running the API

```bash
cd api
go mod download
go run main.go
```

The API will listen on `http://localhost:5001`

### Running the Webapp

```bash
cd webapp
go mod download
go run main.go
```

The webapp will listen on `http://localhost:3000`

### Running Both Simultaneously

In separate terminal windows:

```bash
# Terminal 1 - API
cd api && go run main.go

# Terminal 2 - Webapp
cd webapp && go run main.go
```

## API Endpoints

### Authentication
- `POST /login` - Authenticate user and receive JWT token

### Users
- `POST /users` - Create a new user
- `GET /users` - Get all users (requires auth)
- `GET /users/{userId}` - Get user details (requires auth)
- `PUT /users/{userId}` - Update user profile (requires auth)
- `DELETE /users/{userId}` - Delete user account (requires auth)
- `POST /users/{userId}/update-password` - Change user password (requires auth)

### Social Features
- `POST /users/{userId}/follow` - Follow a user (requires auth)
- `POST /users/{userId}/unfollow` - Unfollow a user (requires auth)
- `GET /users/{userId}/followers` - Get user's followers (requires auth)
- `GET /users/{userId}/following` - Get users the user is following (requires auth)

### Posts
- `POST /posts` - Create a new post (requires auth)
- `GET /posts` - Get all posts (requires auth)
- `GET /posts/{postId}` - Get post details (requires auth)
- `PUT /posts/{postId}` - Update a post (requires auth)
- `DELETE /posts/{postId}` - Delete a post (requires auth)
- `GET /users/{userId}/posts` - Get posts by a specific user (requires auth)
- `POST /posts/{postId}/like` - Like a post (requires auth)
- `POST /posts/{postId}/unlike` - Unlike a post (requires auth)

## Database Schema

### Users Table
```sql
- id: Primary key (auto-increment)
- name: User's full name
- nick: Unique username
- email: Unique email address
- password: Hashed password
- created_at: Account creation timestamp
```

### Posts Table
```sql
- id: Primary key (auto-increment)
- title: Post title
- content: Post content
- author_id: Foreign key to users table
- likes: Number of likes
- created_at: Post creation timestamp
```

### Followers Table
```sql
- user_id: User being followed
- follower_id: User who is following
- Primary key: (user_id, follower_id)
```

## Deployment

### Docker Deployment

Create a `Dockerfile` for the API:

```dockerfile
FROM golang:1.26.0-alpine AS builder
WORKDIR /app
COPY . .
WORKDIR /app/api
RUN go build -o api .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api/api .
COPY api/.env .
CMD ["./api"]
```

Build and run:

```bash
docker build -t devbook-api .
docker run -p 5001:5001 --env-file api/.env devbook-api
```

### Production Considerations

1. **Environment Variables**: Use secure secret management (AWS Secrets Manager, HashiCorp Vault, etc.)
2. **Database**: Use managed MySQL service (AWS RDS, Google Cloud SQL, etc.)
3. **HTTPS**: Use a reverse proxy (Nginx, Caddy) with SSL certificates
4. **Logging**: Implement structured logging for production monitoring
5. **Error Handling**: Set appropriate error logging and alerting
6. **Rate Limiting**: Consider implementing rate limiting on API endpoints
7. **CORS**: Configure CORS appropriately for your domain

## Development

### Project Structure

```
devbook/
├── api/
│   ├── src/
│   │   ├── config/        # Configuration loading
│   │   ├── controllers/   # Request handlers
│   │   ├── models/        # Data structures
│   │   ├── repositories/  # Database operations
│   │   ├── router/        # Route definitions
│   │   ├── middlewares/   # Authentication & logging
│   │   ├── security/      # Security utilities
│   │   ├── db/            # Database connection
│   │   └── responses/     # Response utilities
│   ├── sql/               # Database setup
│   ├── main.go
│   └── go.mod
│
├── webapp/
│   ├── src/
│   │   ├── config/        # Configuration loading
│   │   ├── router/        # Route definitions
│   │   ├── templates/     # HTML templates
│   │   ├── cookies/       # Cookie management
│   │   └── utils/         # Utility functions
│   ├── main.go
│   └── go.mod
│
├── README.md
└── .gitignore
```

### Running Tests

```bash
cd api
go test ./...

cd ../webapp
go test ./...
```

### Making API Requests

Example using curl:

```bash
# Create a user
curl -X POST http://localhost:5001/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","nick":"johndoe","email":"john@example.com","password":"secure123"}'

# Login
curl -X POST http://localhost:5001/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"secure123"}'

# Use the returned token in subsequent requests
curl -X GET http://localhost:5001/users \
  -H "Authorization: Bearer <token>"
```

## Troubleshooting

### Database Connection Issues
- Verify MySQL is running and accessible
- Check credentials in `.env` file match your MySQL setup
- Ensure the database exists and tables are created

### Port Already in Use
```bash
# Find process using port 5001 (API)
lsof -i :5001

# Find process using port 3000 (Webapp)
lsof -i :3000

# Kill the process (macOS/Linux)
kill -9 <PID>
```

### Authentication Errors
- Ensure SECRET_KEY in `/api/.env` is set and matches between API instances
- Check that JWT tokens aren't expired
- Verify token is being sent in Authorization header correctly

## License

[Add your license information here]

## Contributing

[Add contribution guidelines here]
