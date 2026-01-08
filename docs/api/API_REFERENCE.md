# API Documentation

This document describes the REST API endpoints for the video editor platform.

## Base URL
```
Development: http://localhost:3000/api/v1
Production: https://api.videdit.com/api/v1
```

## Authentication

All authenticated endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <token>
```

### Authentication Endpoints

#### POST /auth/register
Register a new user account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securePassword123",
  "username": "johndoe"
}
```

**Response:** 201 Created
```json
{
  "user": {
    "id": "user-uuid",
    "email": "user@example.com",
    "username": "johndoe"
  },
  "token": "jwt-token"
}
```

#### POST /auth/login
Login to existing account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

**Response:** 200 OK
```json
{
  "user": {
    "id": "user-uuid",
    "email": "user@example.com",
    "username": "johndoe"
  },
  "token": "jwt-token"
}
```

## Project Endpoints

#### GET /projects
List all projects for authenticated user.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20)
- `sort` (optional): Sort field (default: updated_at)

**Response:** 200 OK
```json
{
  "projects": [
    {
      "id": "project-uuid",
      "name": "My Video Project",
      "description": "Project description",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-02T00:00:00Z",
      "thumbnail": "https://cdn.videdit.com/thumbnails/project-uuid.jpg"
    }
  ],
  "pagination": {
    "total": 42,
    "page": 1,
    "pages": 3
  }
}
```

#### POST /projects
Create a new project.

**Request Body:**
```json
{
  "name": "My New Project",
  "description": "Project description",
  "settings": {
    "resolution": "1920x1080",
    "framerate": 30,
    "aspectRatio": "16:9"
  }
}
```

**Response:** 201 Created

#### GET /projects/:id
Get project details.

**Response:** 200 OK
```json
{
  "id": "project-uuid",
  "name": "My Video Project",
  "description": "Project description",
  "settings": {
    "resolution": "1920x1080",
    "framerate": 30,
    "aspectRatio": "16:9"
  },
  "timeline": {
    "tracks": [...]
  },
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-02T00:00:00Z"
}
```

#### PUT /projects/:id
Update project details.

#### DELETE /projects/:id
Delete a project.

## Asset Endpoints

#### GET /projects/:projectId/assets
List all assets in a project.

#### POST /projects/:projectId/assets/upload
Request upload URL for new asset.

**Request Body:**
```json
{
  "filename": "video.mp4",
  "filesize": 104857600,
  "mimetype": "video/mp4"
}
```

**Response:** 200 OK
```json
{
  "assetId": "asset-uuid",
  "uploadUrl": "https://storage.videdit.com/upload/signed-url",
  "expiresAt": "2024-01-01T01:00:00Z"
}
```

#### POST /projects/:projectId/assets/:assetId/complete
Notify that upload is complete.

#### GET /projects/:projectId/assets/:assetId
Get asset details.

#### DELETE /projects/:projectId/assets/:assetId
Delete an asset.

## Version Control Endpoints

#### GET /projects/:projectId/commits
List commit history.

**Response:** 200 OK
```json
{
  "commits": [
    {
      "id": "commit-hash",
      "message": "Added intro sequence",
      "author": {
        "id": "user-uuid",
        "username": "johndoe"
      },
      "timestamp": "2024-01-01T12:00:00Z",
      "parent": "parent-commit-hash"
    }
  ]
}
```

#### POST /projects/:projectId/commits
Create a new commit.

**Request Body:**
```json
{
  "message": "Added intro sequence",
  "timeline": {
    "tracks": [...]
  }
}
```

#### GET /projects/:projectId/branches
List all branches.

#### POST /projects/:projectId/branches
Create a new branch.

**Request Body:**
```json
{
  "name": "experimental-edit",
  "fromCommit": "commit-hash"
}
```

#### POST /projects/:projectId/merge
Merge two branches.

**Request Body:**
```json
{
  "source": "experimental-edit",
  "target": "main",
  "strategy": "auto"
}
```

**Response:** 200 OK or 409 Conflict
```json
{
  "status": "success|conflict",
  "conflicts": [
    {
      "type": "clip-overlap",
      "track": 1,
      "timeRange": [10.5, 15.2]
    }
  ]
}
```

## AI Service Endpoints

#### POST /ai/analyze
Request AI analysis of video.

**Request Body:**
```json
{
  "assetId": "asset-uuid",
  "analyses": ["scene-detection", "object-recognition", "transcription"]
}
```

**Response:** 202 Accepted
```json
{
  "jobId": "job-uuid",
  "status": "queued"
}
```

#### GET /ai/jobs/:jobId
Get AI job status.

#### POST /ai/effects/apply
Apply AI effect to video.

#### POST /ai/recommend
Get AI recommendations.

## Collaboration Endpoints

#### GET /projects/:projectId/collaborators
List project collaborators.

#### POST /projects/:projectId/collaborators
Add collaborator to project.

#### GET /projects/:projectId/comments
Get all comments on project.

#### POST /projects/:projectId/comments
Add a comment.

## WebSocket Events

Connect to: `ws://localhost:3000/collaborate/:projectId`

### Client → Server Events

**join-session**
```json
{
  "event": "join-session",
  "projectId": "project-uuid"
}
```

**timeline-update**
```json
{
  "event": "timeline-update",
  "delta": {
    "operation": "add-clip",
    "track": 1,
    "clip": {...}
  }
}
```

**cursor-move**
```json
{
  "event": "cursor-move",
  "position": {
    "x": 100,
    "y": 200
  }
}
```

### Server → Client Events

**user-joined**
```json
{
  "event": "user-joined",
  "user": {
    "id": "user-uuid",
    "username": "johndoe",
    "color": "#FF5733"
  }
}
```

**timeline-updated**
```json
{
  "event": "timeline-updated",
  "userId": "user-uuid",
  "delta": {...}
}
```

**cursor-updated**
```json
{
  "event": "cursor-updated",
  "userId": "user-uuid",
  "position": {...}
}
```

## Error Responses

All errors follow this format:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable error message",
    "details": {}
  }
}
```

### Common Error Codes
- `400` Bad Request - Invalid input
- `401` Unauthorized - Missing or invalid token
- `403` Forbidden - Insufficient permissions
- `404` Not Found - Resource not found
- `409` Conflict - Version conflict or merge conflict
- `429` Too Many Requests - Rate limit exceeded
- `500` Internal Server Error - Server error

## Rate Limits

- **Free Tier**: 100 requests/hour
- **Pro Tier**: 1000 requests/hour
- **Enterprise**: Unlimited

Rate limit headers:
```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1609459200
```
