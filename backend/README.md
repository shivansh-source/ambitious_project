# Backend Services (Go)

This directory contains the backend API services for the AI-powered video editor with Git-like versioning.

## Structure

### `/api`
REST API endpoints for:
- User authentication and authorization
- Video project management
- Version control operations (commit, push, pull, merge)
- Collaboration features (comments, reviews, permissions)
- AI service integration endpoints

### `/services`

#### `/ai-processing`
Integration layer for AI services:
- Video analysis orchestration
- Effect application pipeline
- Real-time AI suggestions
- Background processing queue

#### `/collaboration`
Real-time collaboration features:
- WebSocket server for live editing sessions
- Presence tracking and cursor synchronization
- Comment and review system
- Conflict detection and resolution

### `/models`
Database models and schemas:
- User and authentication models
- Project and video models
- Collaboration data structures

### `/database`
Database migrations and seeders:
- Schema definitions
- Migration scripts
- Sample data for development

### `/middleware`
Gin middleware:
- Authentication and authorization
- Rate limiting
- Request validation
- Error handling

### `/utils`
Shared utilities and helpers:
- Video processing utilities
- File system operations
- Logging and monitoring
- Configuration management

## Technology Stack
- **Runtime**: Go 1.21+
- **Framework**: Gin
- **Database**: PostgreSQL (relational data) + MongoDB (video metadata)
- **Cache**: Redis (sessions, real-time data)
- **Message Queue**: RabbitMQ / Kafka (async processing)
- **WebSocket**: Gorilla WebSocket

## Note on Version Control
The Git-like version control engine has been moved to `/versioning` and is implemented in Python as an AI-powered agentic system. The backend API provides HTTP endpoints to interact with the versioning service.
