# Template Files

This directory contains template configuration files for the VideEdit project.

## How to Use Templates

Template files (with `.template` extension) need to be copied and renamed before use:

### Backend Setup (Go)
```bash
cd backend
cp go.mod.template go.mod
# Remove the comment header lines (line 1-2)
cp Makefile.template Makefile
go mod download
```

### Frontend Setup
```bash
cd frontend
cp package.json.template package.json
# Remove the comment on line 1 that says "# Package.json Template for Frontend"
npm install
```

### AI Services Setup
```bash
cd ai-services
cp requirements.txt.template requirements.txt
# Remove the comment on line 1 that says "# Requirements.txt Template for AI Services"
pip install -r requirements.txt
```

## Important Notes

- **Template files are not meant to be used directly**
- After copying, **remove the comment headers** from template files
- For Go templates (go.mod.template), remove the first 2 comment lines
- For JSON templates, remove the comment header line
- For Python requirements, remove the comment header line
- Customize the dependencies based on your specific needs
- Template files help you understand what dependencies are needed for each service

## Environment Files

Similarly, create `.env` files for configuration:

### Backend `.env` (Go)
```bash
cd backend
cat > .env << 'EOF'
ENV=development
PORT=3000
DATABASE_URL=postgresql://user:pass@localhost:5432/videdit
MONGODB_URL=mongodb://user:pass@localhost:27017/videdit
REDIS_URL=redis://localhost:6379
JWT_SECRET=your-secret-key
EOF
```

### Frontend `.env`
```bash
cd frontend
cp .env.example .env  # (when .env.example is created)
# Edit .env with your API URLs
```

## Next Steps

After setting up from templates:
1. Review and customize dependencies
2. Install dependencies:
   - Backend: `go mod download`
   - Frontend: `npm install`
   - AI/Versioning: `pip install -r requirements.txt`
3. Configure environment variables
4. Start development servers

See [Getting Started Guide](developer-guide/GETTING_STARTED.md) for detailed instructions.
