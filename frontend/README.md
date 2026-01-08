# Frontend - Video Editor Interface

This directory contains the web-based video editor interface with collaborative editing features.

## Structure

### `/src/components`

#### `/editor`
Main video editor components:
- Canvas and rendering engine
- Tool panels (cut, trim, effects)
- Properties inspector
- Asset library browser

#### `/timeline`
Timeline components:
- Multi-track timeline view
- Clip manipulation (drag, resize, split)
- Keyframe editor
- Audio waveform visualization

#### `/preview`
Video preview components:
- Real-time playback
- Quality settings
- Export preview
- Multi-camera view

#### `/collaboration`
Collaboration UI components:
- Live cursor tracking
- Comment threads
- Review mode
- User presence indicators

### `/src/pages`
Application pages:
- Dashboard (project listing)
- Editor (main editing interface)
- Version history
- Project settings
- User profile

### `/src/hooks`
Custom React hooks:
- `useVideoPlayer` - Video playback control
- `useTimeline` - Timeline state management
- `useCollaboration` - Real-time collaboration
- `useVersionControl` - Git operations

### `/src/store`
State management (Redux/Zustand):
- Editor state (current timeline, selection)
- Project state (assets, settings)
- Collaboration state (users, comments)
- Version control state (branches, commits)

### `/src/utils`
Frontend utilities:
- Video format conversion
- Timeline calculations
- WebSocket helpers
- API clients

### `/public`
Static assets:
- Icons and images
- Fonts
- Service worker
- HTML template

## Technology Stack
- **Framework**: React / Vue.js / Svelte
- **State Management**: Redux Toolkit / Zustand
- **Video Processing**: FFmpeg.wasm / WebCodecs API
- **Canvas Rendering**: Fabric.js / Konva.js
- **Real-time**: Socket.io / WebRTC
- **Build Tool**: Vite / Webpack
