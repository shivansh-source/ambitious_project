#!/bin/bash

# VideEdit Quick Start Script
# This script sets up the development environment and starts all services

set -e

echo "🎬 VideEdit - AI-Powered Collaborative Video Editor"
echo "===================================================="
echo ""

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker first."
    echo "   Visit: https://docs.docker.com/get-docker/"
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose is not installed. Please install Docker Compose first."
    echo "   Visit: https://docs.docker.com/compose/install/"
    exit 1
fi

echo "✅ Docker and Docker Compose are installed"
echo ""

# Navigate to infrastructure directory
cd infrastructure/docker

echo "🚀 Starting VideEdit services..."
echo ""
echo "This will start:"
echo "  - PostgreSQL database"
echo "  - MongoDB database"
echo "  - Redis cache"
echo "  - RabbitMQ message queue"
echo "  - MinIO (S3-compatible storage)"
echo "  - Backend API (when ready)"
echo "  - Frontend app (when ready)"
echo "  - AI services (when ready)"
echo ""

# Start services
docker-compose up -d

echo ""
echo "✅ Services are starting up..."
echo ""
echo "🌐 Access the services at:"
echo "  - Frontend:         http://localhost:5173"
echo "  - Backend API:      http://localhost:3000"
echo "  - AI Services:      http://localhost:8000"
echo "  - RabbitMQ Admin:   http://localhost:15672 (user: videdit, pass: videdit_dev_password)"
echo "  - MinIO Console:    http://localhost:9001 (user: videdit, pass: videdit_dev_password)"
echo ""
echo "📊 Check service status:"
echo "  docker-compose ps"
echo ""
echo "📋 View logs:"
echo "  docker-compose logs -f [service-name]"
echo ""
echo "🛑 Stop services:"
echo "  docker-compose down"
echo ""
echo "📚 For more information, read:"
echo "  - docs/developer-guide/GETTING_STARTED.md"
echo "  - docs/PROJECT_STRUCTURE.md"
echo ""
echo "Happy editing! 🎥"
