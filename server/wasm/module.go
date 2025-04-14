package wasm

// Module represents a WASM module in the registry
type Module struct {
	ID          string               `json:"id"` // unique identifier
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Author      string               `json:"author,omitempty"`
	URL         string               `json:"url,omitempty"`
	Versions    map[string]*Artifact `json:"versions"` // version -> artifact URL
	Tags        map[string]string    `json:"tags"`     // tag -> version mapping
}

// Artifact represents a specific version of a WASM module
type Artifact struct {
	URL       string `json:"url"`                 // download url
	Digest    string `json:"digest"`              // sha256 checksum
	Size      int64  `json:"size"`                // size of the binary in bytes
	Signature string `json:"signature,omitempty"` // signature of the binary
	SignedBy  string `json:"signed_by,omitempty"` // key ID that signed the binary
}
