# Backend Services

This directory contains the backend services for the AI-powered video editor with Git-like versioning.

## Structure

### `/api`
REST API endpoints for:
- User authentication and authorization
- Video project management
- Version control operations (commit, push, pull, merge)
- Collaboration features (comments, reviews, permissions)
- AI service integration endpoints

### `/services`

#### `/versioning`
Git-like versioning engine for video projects:
- Track changes to video timelines, effects, and assets
- Branch management for experimental edits
- Merge conflict resolution for collaborative editing
- Version history and rollback functionality

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
- Version control metadata
- Collaboration data structures

### `/database`
Database migrations and seeders:
- Schema definitions
- Migration scripts
- Sample data for development

### `/middleware`
Express/API middleware:
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
- **Runtime**: Node.js / Go (choose based on performance needs)
- **API Framework**: Express.js / Gin / FastAPI
- **Database**: PostgreSQL (relational data) + MongoDB (video metadata)
- **Cache**: Redis (sessions, real-time data)
- **Message Queue**: RabbitMQ / Kafka (async processing)
