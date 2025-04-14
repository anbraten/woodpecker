package wasm

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// RegistryConfig holds configuration for the WASM registry service
type RegistryConfig struct {
	RegistryURLs     []string      `json:"registry_urls"`
	CachePath        string        `json:"cache_path"`
	RefreshInterval  time.Duration `json:"refresh_interval"`
	Timeout          time.Duration `json:"timeout"`
	RetryCount       uint          `json:"retry_count"`
	TrustedPublicKey []string      `json:"trusted_public_keys"`
}

// Registry represents the registry index file
type Registry struct {
	Modules   []Module  `json:"modules"`   // module ID -> module
	Generated time.Time `json:"generated"` // when the index was generated
	Version   string    `json:"version"`   // version of the index format
}

// registryService manages the WASM registry
type registryService struct {
	config         RegistryConfig
	moduleIndex    map[string]Module
	binaryCacheMux sync.RWMutex
	httpClient     *http.Client
	lastUpdate     time.Time
}

// NewWasmRegistryService creates a new WASM registry service
func NewRegistryService(config RegistryConfig) (*registryService, error) {
	if len(config.RegistryURLs) == 0 {
		return nil, errors.New("no registry URLs provided")
	}

	if config.CachePath == "" {
		config.CachePath = filepath.Join(os.TempDir(), "wasm-registry-cache")
	}

	if config.RefreshInterval == 0 {
		config.RefreshInterval = 24 * time.Hour
	}

	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	// Create cache directory if it doesn't exist
	if err := os.MkdirAll(config.CachePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	service := &registryService{
		config:      config,
		moduleIndex: make(map[string]Module),
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}

	// Initial update of the registry
	if err := service.UpdateIndex(context.Background(), true); err != nil {
		return nil, fmt.Errorf("failed to update registry index: %w", err)
	}

	return service, nil
}

// UpdateIndex updates the registry index
func (w *registryService) UpdateIndex(ctx context.Context, force bool) error {
	if !force && time.Since(w.lastUpdate) < w.config.RefreshInterval {
		return nil
	}

	w.binaryCacheMux.Lock()
	defer w.binaryCacheMux.Unlock()

	var lastErr error
	for _, url := range w.config.RegistryURLs {
		registry, err := w.fetchRegistry(ctx, url)
		if err != nil {
			lastErr = err
			continue
		}

		if err := w.verifyRegistry(registry, url); err != nil {
			lastErr = fmt.Errorf("failed to verify registry from %s: %w", url, err)
			continue
		}

		for _, module := range registry.Modules {
			if _, ok := w.moduleIndex[module.ID]; ok {
				log.Printf("module %s already exists in the index, skipping", module.ID)
				continue
			}
			w.moduleIndex[module.ID] = module
		}
	}

	if len(w.moduleIndex) == 0 && lastErr != nil {
		return fmt.Errorf("failed to update any registry: %w", lastErr)
	}

	w.lastUpdate = time.Now()
	return nil
}

// fetchRegistry fetches the registry index from the given URL
func (w *registryService) fetchRegistry(ctx context.Context, url string) (*Registry, error) {
	var registry *Registry

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := w.httpClient.Do(req)
	if err != nil {
		return registry, fmt.Errorf("failed to fetch registry from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return registry, fmt.Errorf("unexpected status code %d from %s", resp.StatusCode, url)
	}

	if err := json.NewDecoder(resp.Body).Decode(&registry); err != nil {
		return registry, fmt.Errorf("failed to parse registry from %s: %w", url, err)
	}

	return registry, nil
}

// verifyRegistry verifies the integrity and authenticity of the registry
func (w *registryService) verifyRegistry(registry *Registry, url string) error {
	// Verify that the registry is not too old
	if time.Since(registry.Generated) > 30*24*time.Hour {
		return errors.New("registry is too old")
	}

	// Verify registry signature if we have trusted keys
	if len(w.config.TrustedPublicKey) > 0 {
		// TODO: check the signature of the registry using the trusted public keys
		// foundValidSignature := false
		// for keyID, signature := range registry.Signatures {
		// 	if signature != "" && keyID != "" {
		// 		foundValidSignature = true
		// 		break
		// 	}
		// }

		// if !foundValidSignature {
		// 	return errors.New("no valid signature found for registry")
		// }
	}

	return nil
}

// GetModuleInfo returns the module info for the given module ID
func (w *registryService) GetModuleInfo(moduleID string) (*Module, error) {
	w.binaryCacheMux.RLock()
	defer w.binaryCacheMux.RUnlock()

	// Search for the module in all registries
	if module, ok := w.moduleIndex[moduleID]; ok {
		return &module, nil
	}

	return nil, fmt.Errorf("module %s not found in any registry", moduleID)
}

// GetModuleVersion returns the artifact URL for the given module and version
func (w *registryService) GetModuleVersion(moduleID, version string) (*Artifact, error) {
	module, err := w.GetModuleInfo(moduleID)
	if err != nil {
		return nil, err
	}

	// If version is a tag, resolve it to a version
	if tagVersion, ok := module.Tags[version]; ok {
		version = tagVersion
	}

	// Check if the version exists
	artifactMetadata, ok := module.Versions[version]
	if !ok {
		return nil, fmt.Errorf("version %s not found for module %s", version, moduleID)
	}

	return artifactMetadata, nil
}

// ListModules returns all modules in the registry
func (w *registryService) ListModules() ([]Module, error) {
	w.binaryCacheMux.RLock()
	defer w.binaryCacheMux.RUnlock()

	var modules []Module

	for _, module := range w.moduleIndex {
		modules = append(modules, module)
	}

	return modules, nil
}

// getCachePath returns the cache path for the given module ID and version
func (w *registryService) getCachePath(module *Module, version string) string {
	return filepath.Join(w.config.CachePath, module.ID, version, "module.wasm")
}

// downloadModule downloads the WASM module binary
func (w *registryService) downloadModule(ctx context.Context, artifact *Artifact) ([]byte, error) {
	var data []byte
	var err error

	for i := uint(0); i <= w.config.RetryCount; i++ {
		// Attempt to download the module
		data, err = w.tryDownloadModule(ctx, artifact)
		if err == nil {
			break
		}

		// If we've reached the retry limit, give up
		if i == w.config.RetryCount {
			return nil, fmt.Errorf("failed to download module after %d attempts: %w", w.config.RetryCount+1, err)
		}

		// Wait before retrying
		time.Sleep(time.Duration(i+1) * time.Second)
	}

	// Verify the SHA256 checksum
	hash := sha256.Sum256(data)
	calculatedHash := hex.EncodeToString(hash[:])
	if calculatedHash != artifact.Digest {
		return nil, fmt.Errorf("checksum mismatch: expected %s, got %s", artifact.Digest, calculatedHash)
	}

	// Verify the signature if we have trusted keys
	// if len(w.TrustedPublicKey) > 0 && artifact.Signature != "" {
	// 	// TODO: verify the signature using the trusted public keys
	// }

	return data, nil
}

// tryDownloadModule attempts to download the WASM module binary
func (w *registryService) tryDownloadModule(ctx context.Context, artifact *Artifact) ([]byte, error) {
	// TODO: limit download size to 250KB
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, artifact.URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := w.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// GetBinary returns the WASM binary for the given module ID and version
func (w *registryService) GetBinary(ctx context.Context, module *Module, version string) ([]byte, error) {
	artifact, ok := module.Versions[version]
	if !ok {
		return nil, fmt.Errorf("version %s not found for module %s", version, module.ID)
	}

	// Check if the binary is in the cache
	cachePath := w.getCachePath(module, version)
	if data, err := os.ReadFile(cachePath); err == nil {
		// Verify the SHA256 checksum
		hash := sha256.Sum256(data)
		calculatedHash := hex.EncodeToString(hash[:])
		if calculatedHash == artifact.Digest {
			return data, nil
		}
	}

	data, err := w.downloadModule(ctx, artifact)
	if err != nil {
		return nil, err
	}

	// Save the binary to the cache
	if err := os.MkdirAll(filepath.Dir(cachePath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}
	if err := os.WriteFile(cachePath, data, 0644); err != nil {
		return nil, fmt.Errorf("failed to save binary to cache: %w", err)
	}

	return data, nil
}
