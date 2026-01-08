# Project Structure Overview

This document provides a visual overview of the VideEdit project structure.

## Directory Tree

```
ambitious_project/
│
├── 📱 frontend/                    # Web Application (React/Vue)
│   ├── src/
│   │   ├── components/
│   │   │   ├── editor/            # Main video editor UI
│   │   │   ├── timeline/          # Timeline editing components
│   │   │   ├── preview/           # Video preview player
│   │   │   └── collaboration/     # Real-time collaboration UI
│   │   ├── pages/                 # Application pages/routes
│   │   ├── hooks/                 # Custom React hooks
│   │   ├── store/                 # State management (Redux/Zustand)
│   │   └── utils/                 # Frontend utilities
│   └── public/                    # Static assets
│
├── ⚙️ backend/                     # Backend API Services (Node.js/Go)
│   ├── api/
│   │   ├── routes/                # API route definitions
│   │   └── controllers/           # Request handlers
│   ├── services/
│   │   ├── versioning/           # Git-like version control
│   │   ├── ai-processing/        # AI integration layer
│   │   └── collaboration/        # Real-time features
│   ├── models/                    # Database models
│   ├── database/                  # Migrations & schemas
│   ├── middleware/                # Express middleware
│   └── utils/                     # Shared utilities
│
├── 🤖 ai-services/                 # AI/ML Microservices (Python)
│   ├── video-analysis/            # Scene detection, object recognition
│   ├── effects/                   # AI-powered effects (color, bg removal)
│   ├── recommendations/           # Smart editing suggestions
│   ├── transcription/             # Speech-to-text, captions
│   └── models/                    # ML model storage
│
├── 🔄 versioning/                  # Version Control Engine
│   ├── git-engine/                # Core VCS logic
│   ├── diff/                      # Timeline diff algorithms
│   ├── merge/                     # Merge conflict resolution
│   └── branches/                  # Branch management
│
├── 💾 storage/                     # File Storage
│   ├── videos/                    # Original & proxy videos
│   ├── thumbnails/                # Preview thumbnails
│   ├── temp/                      # Temporary processing files
│   └── cache/                     # Cached data
│
├── 🏗️ infrastructure/              # DevOps & Deployment
│   ├── docker/                    # Docker configurations
│   ├── kubernetes/                # K8s manifests
│   ├── nginx/                     # Nginx configs
│   └── scripts/                   # Deployment scripts
│
├── 📚 docs/                        # Documentation
│   ├── api/                       # API reference
│   ├── architecture/              # System architecture
│   ├── user-guide/                # User documentation
│   └── developer-guide/           # Developer docs
│
└── 🧪 tests/                       # Test Suites
    ├── backend/                   # Backend tests
    ├── frontend/                  # Frontend tests
    ├── ai-services/               # AI service tests
    ├── integration/               # Integration tests
    └── e2e/                       # End-to-end tests
```

## Component Interactions

```
┌─────────────────────────────────────────────────────────────┐
│                         User Browser                        │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐  │
│  │ Timeline │  │  Preview │  │  Collab  │  │    AI    │  │
│  │  Editor  │  │  Player  │  │   UI     │  │  Panel   │  │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  └────┬─────┘  │
└───────┼─────────────┼─────────────┼─────────────┼─────────┘
        │             │             │             │
        └─────────────┴─────────────┴─────────────┘
                      │
              ┌───────┴────────┐
              │  API Gateway   │
              │   (Backend)    │
              └───────┬────────┘
                      │
        ┌─────────────┼─────────────┬─────────────┐
        │             │             │             │
        ▼             ▼             ▼             ▼
┌──────────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐
│ Versioning   │ │  Collab  │ │    AI    │ │ Storage  │
│   Service    │ │ Service  │ │ Services │ │ Service  │
└──────┬───────┘ └────┬─────┘ └────┬─────┘ └────┬─────┘
       │              │            │            │
       └──────────────┴────────────┴────────────┘
                      │
              ┌───────┴────────┐
              │   Data Layer   │
              │ PostgreSQL     │
              │ MongoDB        │
              │ Redis          │
              │ RabbitMQ       │
              └────────────────┘
```

## Data Flow Example: Committing Changes

```
1. User clicks "Commit" in UI
   └─> Frontend/Timeline Component
       │
2. Send commit request to backend
   └─> POST /api/v1/projects/:id/commits
       │
3. Backend validates and processes
   └─> Backend/API/Controllers/versionController
       │
4. Calculate timeline diff
   └─> Versioning/diff/calculator
       │
5. Create commit object
   └─> Versioning/git-engine/commit
       │
6. Save to database
   └─> Database (PostgreSQL + MongoDB)
       │
7. Return commit info to user
   └─> Frontend receives response
       │
8. Update UI with new commit
   └─> Timeline shows saved state
```

## Technology Stack by Component

### Frontend Stack
```
React/Vue.js ─┬─> Components
              ├─> Redux/Zustand (State)
              ├─> Socket.io (WebSocket)
              ├─> Fabric.js (Canvas)
              ├─> FFmpeg.wasm (Video)
              └─> Axios (HTTP)
```

### Backend Stack
```
Node.js/Go ─┬─> Express/Gin (API)
            ├─> PostgreSQL (Data)
            ├─> MongoDB (Documents)
            ├─> Redis (Cache)
            ├─> RabbitMQ (Queue)
            └─> Socket.io (WebSocket)
```

### AI Services Stack
```
Python ─┬─> FastAPI (API)
        ├─> TensorFlow/PyTorch (ML)
        ├─> OpenCV (Video)
        ├─> FFmpeg (Processing)
        └─> Celery (Tasks)
```

## File Naming Conventions

### Backend (Node.js)
- Routes: `{resource}.routes.js` (e.g., `version.routes.js`)
- Controllers: `{resource}.controller.js`
- Services: `{resource}.service.js`
- Models: `{resource}.model.js`

### Frontend (React)
- Components: `{ComponentName}.jsx`
- Hooks: `use{HookName}.js`
- Pages: `{PageName}Page.jsx`
- Utilities: `{utility}.util.js`

### AI Services (Python)
- Services: `{service}_service.py`
- Models: `{model}_model.py`
- Utils: `{utility}_utils.py`

## Key Features by Directory

### `/backend/services/versioning`
- Git-like commit system
- Branch management
- Merge conflict resolution
- Timeline diffing

### `/frontend/src/components/timeline`
- Multi-track timeline
- Drag-and-drop clips
- Keyframe editing
- Real-time collaboration cursors

### `/ai-services/video-analysis`
- Scene detection
- Object recognition
- Audio analysis
- Quality assessment

### `/versioning/merge`
- Three-way merge algorithm
- Conflict detection
- Resolution strategies
- Timeline reconciliation

## Getting Started Paths

### For Frontend Developers
```
1. Read: docs/developer-guide/GETTING_STARTED.md
2. Explore: frontend/src/components/
3. Start with: frontend/src/components/timeline/
4. Run: npm run dev
```

### For Backend Developers
```
1. Read: backend/README.md
2. Explore: backend/api/routes/
3. Start with: backend/services/versioning/
4. Run: npm run dev
```

### For ML Engineers
```
1. Read: ai-services/README.md
2. Explore: ai-services/video-analysis/
3. Start with: ai-services/models/
4. Run: uvicorn main:app --reload
```

### For DevOps
```
1. Read: infrastructure/README.md
2. Explore: infrastructure/docker/
3. Start with: docker-compose.yml
4. Run: docker-compose up
```

---

This structure is designed to be:
- **Modular**: Each component is independent
- **Scalable**: Easy to add new features
- **Maintainable**: Clear separation of concerns
- **Collaborative**: Multiple teams can work in parallel
