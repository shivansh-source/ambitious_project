# Storage

This directory contains storage management for video files, assets, and cached data.

## Structure

### `/videos`
Raw video file storage:
- Original uploaded videos
- Rendered/exported videos
- Proxy videos (lower resolution for editing)
- Organized by project and user

### `/thumbnails`
Generated thumbnail images:
- Video clip thumbnails
- Timeline preview thumbnails
- Project cover images
- Auto-generated at multiple sizes

### `/temp`
Temporary working files:
- In-progress renders
- Temporary processing files
- Upload chunks
- Automatically cleaned up

### `/cache`
Cached data for performance:
- Waveform data
- Preview frames
- Metadata cache
- AI analysis results

## Storage Strategy

### File Organization
```
storage/
├── videos/
│   ├── {user_id}/
│   │   ├── {project_id}/
│   │   │   ├── originals/
│   │   │   │   └── {asset_id}.mp4
│   │   │   ├── proxies/
│   │   │   │   └── {asset_id}_proxy.mp4
│   │   │   └── exports/
│   │   │       └── {export_id}.mp4
├── thumbnails/
│   ├── {project_id}/
│   │   ├── {asset_id}_thumb.jpg
│   │   └── {asset_id}_sprite.jpg
```

### Storage Backends

#### Local Storage
- Development environment
- Small deployments
- Fast access for testing

#### Cloud Storage (Production)
- **AWS S3** / **Google Cloud Storage** / **Azure Blob**
- Scalable and reliable
- CDN integration
- Lifecycle policies for old files

### Optimization Techniques

#### Video Proxies
- Generate lower resolution versions for smooth editing
- Original files used only for final export
- Reduces bandwidth and processing requirements

#### Chunked Uploads
- Large file upload support
- Resume capability
- Progress tracking

#### Lazy Loading
- Load video segments on-demand
- Reduce initial load time
- Stream video data as needed

#### Cleanup Policies
- Auto-delete temporary files after 24 hours
- Archive old project files
- User quota management

## Technology Stack
- **Local**: File system operations
- **Cloud**: AWS SDK / Google Cloud SDK
- **CDN**: CloudFront / Cloud CDN
- **Database**: PostgreSQL (file metadata)
