# Versioning Engine

This directory contains the Git-like version control system adapted for video editing projects.

## Structure

### `/git-engine`
Core version control engine:
- Repository initialization
- Commit creation and management
- Branch operations
- Tag management
- Remote repository sync

### `/diff`
Video project diff algorithms:
- Timeline diff (added/removed/modified clips)
- Effect parameter changes
- Asset reference changes
- Audio track modifications
- Metadata changes

### `/merge`
Merge strategies for video projects:
- Three-way merge for timeline conflicts
- Effect conflict resolution
- Asset conflict handling
- Automatic merge when possible
- Interactive merge tool for conflicts

### `/branches`
Branch management:
- Branch creation and deletion
- Branch switching
- Branch comparison
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
- **Core Logic**: Custom implementation inspired by Git internals
- **Storage**: JSON/MessagePack for project snapshots
- **Hashing**: SHA-256 for commit identification
- **Compression**: LZ4/Zstd for efficient storage
