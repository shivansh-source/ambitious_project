# AI-Powered Agentic Version Control (Python)

This directory contains an AI-powered version control system adapted for video editing projects, using LLM agents for intelligent versioning operations.

## Overview

Unlike traditional Git, this agentic version control system uses Large Language Models (LLMs) to:
- **Intelligently analyze** timeline changes and generate semantic commit messages
- **Auto-detect conflicts** and suggest resolution strategies
- **Understand context** from video content and editing patterns
- **Provide insights** on editing history and collaboration patterns

## Structure

### `/agentic_engine.py`
Core AI-powered version control engine:
- Repository initialization
- Commit creation with AI-generated messages
- Branch operations
- AI-assisted merge strategies
- Intelligent diff analysis

### `/diff`
AI-enhanced diff algorithms:
- Timeline diff (added/removed/modified clips)
- Semantic change detection
- Effect parameter analysis
- Audio track modifications
- AI-generated diff summaries

### `/merge`
AI-assisted merge strategies:
- LLM-powered conflict detection
- Intelligent three-way merge
- Context-aware conflict resolution
- Automatic merge suggestions
- Interactive merge with AI guidance

### `/branches`
Branch management:
- Branch creation and deletion
- Branch switching
- Branch comparison with AI insights
- Parallel editing workflows
- Feature branch isolation

## Version Control Concepts for Video

### What Gets Versioned
- **Timeline Structure**: Clips, their positions, and durations
- **Effects**: Applied effects and their parameters
- **Transitions**: Transition types and settings
- **Audio**: Audio tracks, levels, and effects
- **Metadata**: Project settings, export configurations
- **Assets**: References to video/audio/image files (not the files themselves)

### Diff Visualization
```
Timeline Track 1:
  [Clip A] [Clip B] [Clip C]
+ [Effect: Color Grade]
- [Transition: Fade]
+ [Transition: Cross Dissolve]

Audio Track:
+ [Clip D: background_music.mp3]
  [Modified: Volume -6dB -> -3dB]
```

### Merge Scenarios

#### Non-Conflicting Merge
- User A adds clips to track 1
- User B adds clips to track 2
- Auto-merge successful

#### Conflicting Merge
- User A modifies clip duration at 0:10
- User B adds effect to same clip
- Requires manual resolution

## Commands (Git-like)

```bash
# Initialize a video project repository
videdit init

# Save current state
videdit commit -m "Added intro sequence"

# Create a new branch for experimentation
videdit branch experimental-edit

# Switch to a branch
videdit checkout experimental-edit

# Merge changes back
videdit merge main

# View history
videdit log

# Compare versions
videdit diff main experimental-edit

# Revert to previous version
videdit checkout <commit-hash>
```

## Technology Stack
- **Language**: Python 3.9+
- **AI/LLM**: OpenAI API / Anthropic Claude / Local LLMs
- **Framework**: FastAPI (for versioning API service)
- **Storage**: JSON/MessagePack for project snapshots
- **Hashing**: SHA-256 for commit identification
- **Compression**: LZ4/Zstd for efficient storage
- **Database**: PostgreSQL (metadata) + MongoDB (timeline snapshots)

## AI-Powered Features

### Semantic Commit Messages
Instead of manual messages, the AI analyzes changes and generates descriptions:
```
✅ "Add intro sequence with logo animation and fade-in effect"
✅ "Adjust color grading in outdoor scenes (increased saturation by 15%)"
✅ "Add background music (upbeat electronic) and balance audio levels"
```

### Intelligent Merge
The AI considers:
- Temporal relationships between clips
- Semantic compatibility of effects
- User intent from commit history
- Video editing best practices

Example merge resolution:
```python
# AI detects: Both users added clips at 0:10
# AI suggests: Offset second clip to 0:15 (maintains both edits)
# AI explains: "Both clips show the same subject, staggering creates better flow"
```

### Content-Aware Diffs
Traditional diff shows raw changes. AI diff provides context:
```
Traditional: "Modified clip_3 position from 10.5 to 12.0"
AI-Enhanced: "Moved reaction shot 1.5s later to align with dialogue peak"
```
