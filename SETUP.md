# Complete Setup Guide

This guide walks you through setting up VideEdit from scratch.

## Quick Start (Recommended)

The fastest way to get started is using the quick-start script:

```bash
# Clone the repository
git clone https://github.com/shivansh-source/ambitious_project.git
cd ambitious_project

# Run the quick-start script
./quick-start.sh
```

This will start all infrastructure services using Docker Compose.

## Manual Setup

If you prefer to set up manually or want to understand each step:

### 1. Prerequisites

Install the required software:

- **Go** 1.21+ ([Download](https://go.dev/dl/))
- **Python** 3.9+ ([Download](https://www.python.org/downloads/))
- **Docker** & Docker Compose ([Download](https://docs.docker.com/get-docker/))
- **Git** ([Download](https://git-scm.com/downloads))

### 2. Clone Repository

```bash
git clone https://github.com/shivansh-source/ambitious_project.git
cd ambitious_project
```

### 3. Setup Backend (Go)

```bash
cd backend

# Copy template and remove the comment header
cp go.mod.template go.mod
# Edit go.mod and remove line 1-2: "# Go Module Template for Backend Service"

# Copy Makefile template
cp Makefile.template Makefile

# Download dependencies
go mod download

# Create environment file
cat > .env << 'EOF'
ENV=development
PORT=3000
DATABASE_URL=postgresql://videdit:videdit_dev_password@localhost:5432/videdit
MONGODB_URL=mongodb://videdit:videdit_dev_password@localhost:27017/videdit
REDIS_URL=redis://localhost:6379
RABBITMQ_URL=amqp://videdit:videdit_dev_password@localhost:5672
JWT_SECRET=your-secret-key-change-in-production
EOF

cd ..
```

### 4. Setup Frontend

```bash
cd frontend

# Copy template and remove the comment header
cp package.json.template package.json
# Edit package.json and remove line 1: "# Package.json Template for Frontend"

# Install dependencies
npm install

# Create environment file
cat > .env << 'EOF'
VITE_API_URL=http://localhost:3000
VITE_WS_URL=ws://localhost:3000
EOF

cd ..
```

### 5. Setup AI Services & Versioning

```bash
# Setup AI Services
cd ai-services

# Copy template and remove the comment header
cp requirements.txt.template requirements.txt
# Edit requirements.txt and remove line 1: "# Requirements.txt Template for AI Services"

# Create virtual environment (recommended)
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate

# Install dependencies
pip install -r requirements.txt

# Create environment file
cat > .env << 'EOF'
REDIS_URL=redis://localhost:6379
RABBITMQ_URL=amqp://videdit:videdit_dev_password@localhost:5672
MODEL_PATH=/path/to/models
GPU_ENABLED=false
EOF

cd ..

# Setup Agentic Versioning (Python)
cd versioning

# Use same virtual environment or create new one
source ../ai-services/venv/bin/activate

# Install additional dependencies for versioning
pip install openai anthropic  # For LLM integration

# Create environment file
cat > .env << 'EOF'
OPENAI_API_KEY=your-openai-key-here
# OR
ANTHROPIC_API_KEY=your-anthropic-key-here
REDIS_URL=redis://localhost:6379
DATABASE_URL=postgresql://videdit:videdit_dev_password@localhost:5432/videdit
EOF

cd ..
```

### 6. Start Infrastructure Services

```bash
cd infrastructure/docker

# Start databases and message queue
docker-compose up -d postgres mongodb redis rabbitmq minio

# Wait for services to be ready (about 30 seconds)
sleep 30

# Verify services are running
docker-compose ps
```

### 7. Initialize Database

```bash
cd ../../backend

# Run database migrations (when implemented)
# npm run migrate

# For now, you can connect to PostgreSQL and create the schema manually:
# psql postgresql://videdit:videdit_dev_password@localhost:5432/videdit
# Then run the SQL from docs/DATABASE_SCHEMA.md
```

### 8. Start Development Servers

Open separate terminal windows for each service:

**Terminal 1 - Backend (Go):**
```bash
cd backend
make run
# OR: go run ./cmd/server/main.go
# Server should start on http://localhost:3000
```

**Terminal 2 - Frontend:**
```bash
cd frontend
npm run dev
# Server should start on http://localhost:5173
```

**Terminal 3 - AI Services:**
```bash
cd ai-services
source venv/bin/activate  # If using virtual environment
uvicorn main:app --reload --port 8000
# Server should start on http://localhost:8000
```

**Terminal 4 - Versioning Service (Python):**
```bash
cd versioning
source venv/bin/activate  # If using virtual environment
python -m uvicorn api:app --reload --port 8001
# Server should start on http://localhost:8001
```

### 9. Verify Setup

Open your browser and navigate to:
- Frontend: http://localhost:5173
- Backend API: http://localhost:3000/health (when implemented)
- AI Services: http://localhost:8000/docs (FastAPI auto-docs)

## Development Workflow

### Making Changes

1. Create a feature branch:
```bash
git checkout -b feature/your-feature
```

2. Make your changes in the appropriate directory:
   - Backend (Go): `backend/`
   - Frontend (React): `frontend/`
   - AI Services (Python): `ai-services/`
   - Versioning (Python): `versioning/`
   - Docs: `docs/`

3. Test your changes locally

4. Commit and push:
```bash
git add .
git commit -m "feat: your feature description"
git push origin feature/your-feature
```

### Running Tests

```bash
# Backend tests (Go)
cd backend
make test
# OR: go test ./...

# Frontend tests
cd frontend
npm test

# AI services tests
cd ai-services
pytest

# Versioning tests
cd versioning
pytest
```
cd frontend
npm test

# AI services tests
cd ai-services
pytest
```

### Stopping Services

Stop all Docker services:
```bash
cd infrastructure/docker
docker-compose down
```

Stop development servers:
- Press `Ctrl+C` in each terminal window

## Troubleshooting

### Port Already in Use

If you get "port already in use" errors:

```bash
# Find process using port 3000 (backend)
lsof -i :3000
kill -9 <PID>

# Find process using port 5173 (frontend)
lsof -i :5173
kill -9 <PID>
```

### Docker Services Not Starting

```bash
# Check Docker logs
docker-compose logs postgres
docker-compose logs mongodb
docker-compose logs redis

# Restart a specific service
docker-compose restart postgres
```

### Database Connection Failed

Make sure the database is running:
```bash
docker-compose ps
```

If PostgreSQL is not running:
```bash
docker-compose up -d postgres
```

### Module Not Found

Make sure you installed dependencies:
```bash
# Backend (Go)
cd backend && go mod download

# Frontend
cd frontend && npm install

# AI Services & Versioning
cd ai-services && pip install -r requirements.txt
cd versioning && pip install openai anthropic
```

## Next Steps

Now that you have the development environment set up:

1. **Read the architecture docs**: [docs/architecture/SYSTEM_ARCHITECTURE.md](docs/architecture/SYSTEM_ARCHITECTURE.md)
2. **Explore the codebase**: Start with the component you're interested in
3. **Check the API docs**: [docs/api/API_REFERENCE.md](docs/api/API_REFERENCE.md)
4. **Read contributing guidelines**: [CONTRIBUTING.md](CONTRIBUTING.md)

## Additional Resources

- [Project Structure](docs/PROJECT_STRUCTURE.md)
- [Database Schema](docs/DATABASE_SCHEMA.md)
- [Template Files Guide](docs/TEMPLATES.md)
- [Developer Guide](docs/developer-guide/GETTING_STARTED.md)

## Getting Help

- **Issues**: Report bugs on [GitHub Issues](https://github.com/shivansh-source/ambitious_project/issues)
- **Discussions**: Ask questions in [GitHub Discussions](https://github.com/shivansh-source/ambitious_project/discussions)

Happy coding! 🚀
