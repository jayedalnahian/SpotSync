# SpotSync

SpotSync is a Go-based backend API for managing smart parking and EV charging reservations. It supports user authentication, parking zone management, and reservation flows with role-based access control and capacity-safe booking logic.

## Features

- User registration and login with JWT authentication
- Role-based access for drivers and admins
- Parking zone creation and browsing
- Reservation creation with concurrency-safe capacity checks
- My reservations lookup and cancellation
- Admin access to all reservations
- Clean architecture with separate handler, service, repository, DTO, and model layers

## 🌐 Live URLs

| Service | URL |
|---------|-----|

| Backend | [https://spotsync-dcxl.onrender.com/](https://spotsync-dcxl.onrender.com/) |



## Tech Stack

- Go 1.26+
- Echo v5
- GORM
- PostgreSQL
- Validator
- JWT
- bcrypt

## Project Structure

```text
cmd/                 # Application entrypoint
internal/
  auth/              # JWT helpers
  config/            # Environment and database configuration
  domain/
    user/            # Authentication and user domain
    parking_zones/   # Parking zone domain
    reservations/    # Reservation domain
  middlewares/       # Auth middleware
  server/            # HTTP server bootstrap
```

## Architecture

The project follows a layered structure:

- Handler: HTTP request handling and response formatting
- Service: Business rules and orchestration
- Repository: Database access and transactions
- DTO: Request and response payloads
- Model: GORM entities

This separation keeps the API layer clean and makes the reservation logic easier to maintain.

## Environment Variables

Create a `.env` file in the project root with the following values:

```env
PORT=8080
DSN=postgres://username:password@host:5432/database
JWT_SECRET=your-secret-key
```

## Clone This Project

```bash
git clone https://github.com/jayedalnahian/SpotSync.git
```

## Running Locally

1. Install Go dependencies:

```bash
go mod download
```

2. Start the server:

```bash
go run ./cmd/main.go
```

3. Health check:

```bash
curl http://localhost:8080/health
```

## API Endpoints

### Authentication

- POST /api/v1/auth/register
- POST /api/v1/auth/login

### Parking Zones

- GET /api/v1/zones
- GET /api/v1/zones/:id
- POST /api/v1/zones (admin only)

### Reservations

- POST /api/v1/reservations
- GET /api/v1/reservations/my-reservations
- DELETE /api/v1/reservations/:id
- GET /api/v1/reservations (admin only)

## Example Requests

### Register a user

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john.doe@example.com","password":"securePassword123","role":"driver"}'
```

### Create a parking zone (admin)

```bash
curl -X POST http://localhost:8080/api/v1/zones \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name":"Terminal 1 EV Charging","type":"ev_charging","total_capacity":20,"price_per_hour":5.5}'
```

### Reserve a spot

```bash
curl -X POST http://localhost:8080/api/v1/reservations \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"zone_id":1,"license_plate":"ABC-1234"}'
```

## Notes

The reservation flow uses a database transaction with row-level locking to prevent overbooking when multiple users try to reserve the last available spot at the same time.
