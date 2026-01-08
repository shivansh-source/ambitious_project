# Database Schema

This document outlines the database schema for VideEdit.

## PostgreSQL Schema (Relational Data)

### Users Table
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    display_name VARCHAR(100),
    avatar_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
```

### Projects Table
```sql
CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    owner_id UUID REFERENCES users(id) ON DELETE CASCADE,
    settings JSONB DEFAULT '{}',
    current_branch VARCHAR(100) DEFAULT 'main',
    thumbnail_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_projects_owner ON projects(owner_id);
CREATE INDEX idx_projects_updated ON projects(updated_at DESC);
```

### Commits Table
```sql
CREATE TABLE commits (
    id VARCHAR(64) PRIMARY KEY, -- SHA-256 hash
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    parent_id VARCHAR(64) REFERENCES commits(id),
    author_id UUID REFERENCES users(id),
    message TEXT NOT NULL,
    timeline_ref VARCHAR(64) NOT NULL, -- Reference to MongoDB document
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_commits_project ON commits(project_id);
CREATE INDEX idx_commits_parent ON commits(parent_id);
CREATE INDEX idx_commits_created ON commits(created_at DESC);
```

### Branches Table
```sql
CREATE TABLE branches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    head_commit_id VARCHAR(64) REFERENCES commits(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(project_id, name)
);

CREATE INDEX idx_branches_project ON branches(project_id);
```

### Assets Table
```sql
CREATE TABLE assets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    filename VARCHAR(255) NOT NULL,
    original_filename VARCHAR(255) NOT NULL,
    mimetype VARCHAR(100) NOT NULL,
    filesize BIGINT NOT NULL,
    storage_path TEXT NOT NULL,
    thumbnail_url TEXT,
    duration FLOAT, -- For video/audio files
    width INTEGER, -- For video/image files
    height INTEGER, -- For video/image files
    metadata JSONB DEFAULT '{}',
    status VARCHAR(50) DEFAULT 'processing', -- processing, ready, failed
    uploaded_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_assets_project ON assets(project_id);
CREATE INDEX idx_assets_status ON assets(status);
```

### Collaborators Table
```sql
CREATE TABLE collaborators (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(20) NOT NULL, -- viewer, editor, admin
    invited_by UUID REFERENCES users(id),
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(project_id, user_id)
);

CREATE INDEX idx_collaborators_project ON collaborators(project_id);
CREATE INDEX idx_collaborators_user ON collaborators(user_id);
```

### Comments Table
```sql
CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    author_id UUID REFERENCES users(id),
    commit_id VARCHAR(64) REFERENCES commits(id),
    content TEXT NOT NULL,
    timeline_position FLOAT, -- Position in timeline (seconds)
    track_id INTEGER, -- Track number
    parent_comment_id UUID REFERENCES comments(id), -- For threaded comments
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_comments_project ON comments(project_id);
CREATE INDEX idx_comments_commit ON comments(commit_id);
```

### AI Jobs Table
```sql
CREATE TABLE ai_jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    asset_id UUID REFERENCES assets(id) ON DELETE CASCADE,
    job_type VARCHAR(100) NOT NULL, -- scene-detection, transcription, etc.
    status VARCHAR(50) DEFAULT 'queued', -- queued, processing, completed, failed
    progress INTEGER DEFAULT 0, -- 0-100
    result JSONB,
    error_message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    started_at TIMESTAMP,
    completed_at TIMESTAMP
);

CREATE INDEX idx_ai_jobs_project ON ai_jobs(project_id);
CREATE INDEX idx_ai_jobs_status ON ai_jobs(status);
```

## MongoDB Schema (Document Data)

### timelines Collection
```javascript
{
  _id: "timeline-uuid",
  commitId: "commit-hash",
  projectId: "project-uuid",
  version: 1,
  settings: {
    resolution: "1920x1080",
    framerate: 30,
    sampleRate: 48000,
    aspectRatio: "16:9"
  },
  tracks: [
    {
      id: "track-1",
      type: "video", // video, audio, text
      name: "Video Track 1",
      locked: false,
      muted: false,
      clips: [
        {
          id: "clip-1",
          assetId: "asset-uuid",
          start: 0.0, // Position in timeline (seconds)
          end: 5.5,
          trimStart: 0.0, // Trim from source asset
          trimEnd: 5.5,
          speed: 1.0,
          volume: 1.0,
          effects: [
            {
              type: "color-correction",
              params: {
                brightness: 1.2,
                contrast: 1.1,
                saturation: 1.0
              }
            }
          ],
          transitions: {
            in: {
              type: "fade",
              duration: 0.5
            },
            out: {
              type: "cross-dissolve",
              duration: 1.0
            }
          }
        }
      ]
    }
  ],
  metadata: {
    totalDuration: 120.5,
    clipCount: 15
  },
  createdAt: ISODate("2024-01-01T00:00:00Z")
}
```

### ai_analysis Collection
```javascript
{
  _id: "analysis-uuid",
  assetId: "asset-uuid",
  analysisType: "scene-detection",
  result: {
    scenes: [
      {
        start: 0.0,
        end: 5.2,
        confidence: 0.95,
        description: "Indoor office scene"
      },
      {
        start: 5.2,
        end: 12.8,
        confidence: 0.89,
        description: "Outdoor park scene"
      }
    ]
  },
  metadata: {
    modelVersion: "v1.2.0",
    processingTime: 12.5
  },
  createdAt: ISODate("2024-01-01T00:00:00Z")
}
```

### transcriptions Collection
```javascript
{
  _id: "transcription-uuid",
  assetId: "asset-uuid",
  language: "en",
  segments: [
    {
      start: 0.5,
      end: 3.2,
      text: "Welcome to this video tutorial",
      confidence: 0.98,
      speaker: "Speaker 1"
    },
    {
      start: 3.5,
      end: 7.1,
      text: "Today we'll learn about video editing",
      confidence: 0.95,
      speaker: "Speaker 1"
    }
  ],
  fullText: "Welcome to this video tutorial. Today we'll learn about video editing.",
  metadata: {
    modelVersion: "whisper-v3",
    duration: 120.5
  },
  createdAt: ISODate("2024-01-01T00:00:00Z")
}
```

## Redis Schema (Cache & Sessions)

### Session Keys
```
session:{userId}:{sessionId} -> JSON session data
TTL: 24 hours
```

### Presence Keys (Real-time Collaboration)
```
presence:{projectId}:{userId} -> JSON user presence data
TTL: 5 minutes (refreshed periodically)

presence:users:{projectId} -> Set of active user IDs
```

### Cache Keys
```
cache:project:{projectId} -> JSON project data
cache:timeline:{commitId} -> JSON timeline data
cache:asset:{assetId} -> JSON asset metadata
TTL: 1 hour
```

### Rate Limiting
```
ratelimit:{userId}:{endpoint} -> Request count
TTL: 1 hour
```

## Relationships

```
users (1) ──< collaborators >── (N) projects
  │                                   │
  │                                   ├──< (N) commits
  │                                   │      │
  │                                   │      └──> (1) timeline (MongoDB)
  └──> (N) commits                    │
                                      ├──< (N) branches
                                      │
                                      ├──< (N) assets
                                      │      │
                                      │      └──> (N) ai_jobs
                                      │            │
                                      │            └──> ai_analysis (MongoDB)
                                      │
                                      └──< (N) comments
```

## Indexes for Performance

### PostgreSQL
- User lookups: email, username
- Project listings: owner_id, updated_at
- Commit history: project_id, created_at
- Assets: project_id, status

### MongoDB
```javascript
db.timelines.createIndex({ commitId: 1 })
db.timelines.createIndex({ projectId: 1 })
db.ai_analysis.createIndex({ assetId: 1, analysisType: 1 })
db.transcriptions.createIndex({ assetId: 1 })
```

## Data Retention

- **Projects**: Kept indefinitely unless deleted by owner
- **Commits**: Kept for 90 days after project deletion
- **Assets**: Original files kept for 90 days after deletion
- **AI Results**: Kept with assets
- **Logs**: 30 days
- **Temp Files**: 24 hours
