# Template Files

This directory contains template configuration files for the VideEdit project.

## How to Use Templates

Template files (with `.template` extension) need to be copied and renamed before use:

### Backend Setup
```bash
cd backend
cp package.json.template package.json
# Remove the comment on line 1 that says "# Package.json Template for Backend Service"
npm install
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
- After copying, **remove the comment headers** from JSON files
- Customize the dependencies based on your specific needs
- Template files help you understand what dependencies are needed for each service

## Environment Files

Similarly, create `.env` files from examples:

### Backend `.env`
```bash
cd backend
cp .env.example .env  # (when .env.example is created)
# Edit .env with your actual credentials
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
2. Install dependencies (`npm install` or `pip install`)
3. Configure environment variables
4. Start development servers

See [Getting Started Guide](docs/developer-guide/GETTING_STARTED.md) for detailed instructions.
