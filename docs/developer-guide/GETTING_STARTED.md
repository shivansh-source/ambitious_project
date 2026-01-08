# Getting Started - Developer Guide

This guide will help you set up the development environment and start contributing to the AI-powered collaborative video editor.

## Prerequisites

### Required Software
- **Node.js** (v18+) and npm/yarn
- **Python** (3.9+) for AI services
- **Docker** and Docker Compose
- **Git**
- **PostgreSQL** (14+)
- **Redis** (6+)

### Recommended
- **Go** (1.19+) if using Go for backend
- **FFmpeg** (4.4+) for video processing
- **CUDA** (optional, for GPU acceleration)

## Initial Setup

### 1. Clone the Repository
```bash
git clone https://github.com/shivansh-source/ambitious_project.git
cd ambitious_project
```

### 2. Install Dependencies

#### Backend
```bash
cd backend
npm install
# or if using Go:
# go mod download
```

#### Frontend
```bash
cd frontend
npm install
```

#### AI Services
```bash
cd ai-services
pip install -r requirements.txt
# or using poetry:
# poetry install
```

### 3. Environment Configuration

Create `.env` files in each service directory:

#### Backend `.env`
```env
NODE_ENV=development
PORT=3000
DATABASE_URL=postgresql://user:password@localhost:5432/videdit
REDIS_URL=redis://localhost:6379
JWT_SECRET=your-secret-key
AWS_ACCESS_KEY_ID=your-key
AWS_SECRET_ACCESS_KEY=your-secret
S3_BUCKET=videdit-storage
```

#### Frontend `.env`
```env
VITE_API_URL=http://localhost:3000
VITE_WS_URL=ws://localhost:3000
```

#### AI Services `.env`
```env
MODEL_PATH=/path/to/models
REDIS_URL=redis://localhost:6379
GPU_ENABLED=false
```

### 4. Database Setup

```bash
# Start PostgreSQL (via Docker)
docker run -d \
  --name videdit-postgres \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=videdit \
  -p 5432:5432 \
  postgres:14

# Run migrations
cd backend
npm run migrate
# or: go run cmd/migrate/main.go
```

### 5. Start Development Servers

#### Using Docker Compose (Recommended)
```bash
# Start all services
docker-compose up

# Or start specific services
docker-compose up backend frontend
```

#### Manual Start

Terminal 1 - Backend:
```bash
cd backend
npm run dev
```

Terminal 2 - Frontend:
```bash
cd frontend
npm run dev
```

Terminal 3 - AI Services:
```bash
cd ai-services
python -m uvicorn main:app --reload
```

Terminal 4 - Redis:
```bash
docker run -p 6379:6379 redis:6
```

## Development Workflow

### Making Changes

1. **Create a feature branch**
```bash
git checkout -b feature/your-feature-name
```

2. **Make your changes**
```bash
# Edit files
# Add tests
# Update documentation
```

3. **Test your changes**
```bash
# Backend tests
cd backend && npm test

# Frontend tests
cd frontend && npm test

# AI services tests
cd ai-services && pytest
```

4. **Commit and push**
```bash
git add .
git commit -m "feat: add your feature description"
git push origin feature/your-feature-name
```

5. **Create pull request**
- Go to GitHub
- Create PR from your branch to main
- Request review

### Code Style

#### JavaScript/TypeScript
- Use ESLint and Prettier
- Run: `npm run lint`
- Auto-fix: `npm run lint:fix`

#### Python
- Use Black and Flake8
- Run: `black . && flake8`

#### Go
- Use gofmt and golint
- Run: `go fmt ./... && golint ./...`

### Testing

#### Unit Tests
```bash
# Backend
npm test

# Frontend
npm test

# AI Services
pytest tests/
```

#### Integration Tests
```bash
npm run test:integration
```

#### E2E Tests
```bash
cd tests/e2e
npm run test:e2e
```

## Project Structure Quick Reference

```
ambitious_project/
├── backend/              # Backend API services
│   ├── api/             # REST endpoints
│   ├── services/        # Business logic
│   └── models/          # Data models
├── frontend/            # Web application
│   └── src/
│       ├── components/  # React components
│       ├── pages/       # Application pages
│       └── store/       # State management
├── ai-services/         # AI/ML services
│   ├── video-analysis/  # Video analysis
│   ├── effects/         # AI effects
│   └── models/          # ML models
├── versioning/          # Version control engine
├── storage/             # File storage
├── infrastructure/      # Deployment configs
└── docs/               # Documentation
```

## Common Tasks

### Adding a New API Endpoint
1. Create route in `backend/api/routes/`
2. Add controller in `backend/api/controllers/`
3. Add service logic in `backend/services/`
4. Add tests in `tests/backend/`
5. Update API documentation

### Adding a New UI Component
1. Create component in `frontend/src/components/`
2. Add styles
3. Add to Storybook (if applicable)
4. Add tests
5. Use in relevant pages

### Adding a New AI Feature
1. Create service in `ai-services/`
2. Train/download ML model
3. Create API endpoint
4. Add to processing queue
5. Integrate with backend
6. Add UI controls

### Running Specific Tests
```bash
# Single test file
npm test -- video.test.js

# With coverage
npm test -- --coverage

# Watch mode
npm test -- --watch
```

## Debugging

### Backend Debugging
```bash
# Node.js
node --inspect-brk server.js

# Attach VSCode debugger
# Use launch.json configuration
```

### Frontend Debugging
- Use React DevTools
- Use Redux DevTools
- Browser DevTools (Sources tab)

### AI Services Debugging
```bash
# Python debugger
python -m pdb service.py

# Or use IPython
ipython
```

## Troubleshooting

### Port Already in Use
```bash
# Find process using port
lsof -i :3000

# Kill process
kill -9 <PID>
```

### Database Connection Issues
```bash
# Check if PostgreSQL is running
docker ps | grep postgres

# Check logs
docker logs videdit-postgres
```

### Video Processing Failures
```bash
# Check FFmpeg installation
ffmpeg -version

# Check logs
tail -f logs/video-processing.log
```

## Additional Resources

- [API Documentation](../api/)
- [Architecture Overview](../architecture/SYSTEM_ARCHITECTURE.md)
- [User Guide](../user-guide/)
- [Contributing Guidelines](../../CONTRIBUTING.md)

## Getting Help

- Create an issue on GitHub
- Join our Discord server
- Check existing documentation
- Ask in pull request comments
