package versioning

import (
	"encoding/json"
	"fmt"
	"time"
)

// Version represents a single version of content
type Version struct {
	ID          string    `json:"id"`
	ContentID   string    `json:"content_id"`
	VersionNum  int       `json:"version_num"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	FilePath    string    `json:"file_path"`
	FileSize    int64     `json:"file_size"`
	Metadata    Metadata  `json:"metadata"`
}

// Metadata stores additional version information
type Metadata struct {
	Changes    []string          `json:"changes"`
	Tags       []string          `json:"tags"`
	CustomData map[string]string `json:"custom_data"`
}

// VersionManager handles version operations
type VersionManager struct {
	versions map[string][]Version
}

// NewVersionManager creates a new version manager
func NewVersionManager() *VersionManager {
	return &VersionManager{
		versions: make(map[string][]Version),
	}
}

// CreateVersion creates a new version for content
func (vm *VersionManager) CreateVersion(contentID, filePath, createdBy, description string, fileSize int64) (*Version, error) {
	versions := vm.versions[contentID]
	versionNum := len(versions) + 1

	version := &Version{
		ID:          fmt.Sprintf("%s-v%d", contentID, versionNum),
		ContentID:   contentID,
		VersionNum:  versionNum,
		Description: description,
		CreatedAt:   time.Now(),
		CreatedBy:   createdBy,
		FilePath:    filePath,
		FileSize:    fileSize,
		Metadata: Metadata{
			Changes:    []string{},
			Tags:       []string{},
			CustomData: make(map[string]string),
		},
	}

	vm.versions[contentID] = append(vm.versions[contentID], *version)
	return version, nil
}

// GetVersion retrieves a specific version
func (vm *VersionManager) GetVersion(contentID string, versionNum int) (*Version, error) {
	versions, exists := vm.versions[contentID]
	if !exists {
		return nil, fmt.Errorf("content not found: %s", contentID)
	}

	if versionNum < 1 || versionNum > len(versions) {
		return nil, fmt.Errorf("invalid version number: %d", versionNum)
	}

	return &versions[versionNum-1], nil
}

// GetLatestVersion retrieves the most recent version
func (vm *VersionManager) GetLatestVersion(contentID string) (*Version, error) {
	versions, exists := vm.versions[contentID]
	if !exists || len(versions) == 0 {
		return nil, fmt.Errorf("no versions found for content: %s", contentID)
	}

	return &versions[len(versions)-1], nil
}

// ListVersions returns all versions for a content ID
func (vm *VersionManager) ListVersions(contentID string) ([]Version, error) {
	versions, exists := vm.versions[contentID]
	if !exists {
		return nil, fmt.Errorf("content not found: %s", contentID)
	}

	return versions, nil
}

// DeleteVersion removes a specific version
func (vm *VersionManager) DeleteVersion(contentID string, versionNum int) error {
	versions, exists := vm.versions[contentID]
	if !exists {
		return fmt.Errorf("content not found: %s", contentID)
	}

	if versionNum < 1 || versionNum > len(versions) {
		return fmt.Errorf("invalid version number: %d", versionNum)
	}

	// Remove the version from the slice
	vm.versions[contentID] = append(versions[:versionNum-1], versions[versionNum:]...)
	return nil
}

// CompareVersions returns the differences between two versions
func (vm *VersionManager) CompareVersions(contentID string, v1, v2 int) (map[string]interface{}, error) {
	version1, err := vm.GetVersion(contentID, v1)
	if err != nil {
		return nil, err
	}

	version2, err := vm.GetVersion(contentID, v2)
	if err != nil {
		return nil, err
	}

	diff := map[string]interface{}{
		"version1":      v1,
		"version2":      v2,
		"size_diff":     version2.FileSize - version1.FileSize,
		"time_diff":     version2.CreatedAt.Sub(version1.CreatedAt),
		"created_by_v1": version1.CreatedBy,
		"created_by_v2": version2.CreatedBy,
	}

	return diff, nil
}

// ToJSON converts a version to JSON
func (v *Version) ToJSON() (string, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// AddMetadata adds metadata to a version
func (v *Version) AddMetadata(key, value string) {
	if v.Metadata.CustomData == nil {
		v.Metadata.CustomData = make(map[string]string)
	}
	v.Metadata.CustomData[key] = value
}

// AddTag adds a tag to a version
func (v *Version) AddTag(tag string) {
	v.Metadata.Tags = append(v.Metadata.Tags, tag)
}

// AddChange records a change description
func (v *Version) AddChange(change string) {
	v.Metadata.Changes = append(v.Metadata.Changes, change)
}
