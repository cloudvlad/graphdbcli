package license

// LicenseData represents the name of the license
// and the information related to it, referred as note.
type LicenseData struct {
    Name string
    Note string
}

// GetName returns the name of the license.
func (l LicenseData) GetName() string {
    return l.Name
}

// GetNote returns the note associated with the license.
func (l LicenseData) GetNote() string {
    return l.Note
}