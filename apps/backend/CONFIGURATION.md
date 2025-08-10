# Configuration Guide

The Taskify backend application requires several environment variables to be set. Create a `.env` file in the `apps/backend/` directory with the following variables:

## Required Environment Variables

### Primary Configuration
```bash
TASKER_PRIMARY_ENV=development
```

### Server Configuration
```bash
TASKER_SERVER_PORT=8080
TASKER_SERVER_READ_TIMEOUT=30
TASKER_SERVER_WRITE_TIMEOUT=30
TASKER_SERVER_IDLE_TIMEOUT=60
TASKER_SERVER_CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173
```

### Database Configuration
```bash
TASKER_DATABASE_HOST=localhost
TASKER_DATABASE_PORT=5432
TASKER_DATABASE_USER=postgres
TASKER_DATABASE_PASSWORD=password
TASKER_DATABASE_NAME=taskify
TASKER_DATABASE_SSL_MODE=disable
TASKER_DATABASE_MAX_OPEN_CONNS=25
TASKER_DATABASE_MAX_IDLE_CONNS=5
TASKER_DATABASE_CONN_MAX_LIFETIME=300
TASKER_DATABASE_CONN_MAX_IDLE_TIME=60
```

### Redis Configuration
```bash
TASKER_REDIS_ADDRESS=localhost:6379
```

### Auth Configuration
```bash
TASKER_AUTH_SECRET_KEY=your-secret-key-here-change-in-production
```

### Integration Configuration
```bash
TASKER_INTEGRATION_RESEND_API_KEY=your-resend-api-key-here
```

### Database DSN (for migrations)
```bash
TASKER_DB_DSN=postgres://postgres:password@localhost:5432/taskify?sslmode=disable
```

## Quick Setup

1. Copy the above variables into a `.env` file in the `apps/backend/` directory
2. Update the values according to your local setup
3. Make sure you have PostgreSQL and Redis running locally
4. Run the application with `task run`

## Production Notes

- Change `TASKER_AUTH_SECRET_KEY` to a strong, random secret
- Update database credentials for production
- Set `TASKER_PRIMARY_ENV=production`
- Configure proper CORS origins for your frontend domain
- Use a proper Resend API key for email functionality 