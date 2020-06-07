package xoclient

type (
	// IDRef is the generic id reference
	IDRef string

	// Filter is used to filter generic objects
	Filter string

	// Generic is used as a generic return type
	Generic map[string]interface{}

	// Server xen server host details
	Server struct {
	}

	// VDI Virtual Disk Image
	VDI struct {
	}

	// VBD Virtual Block Device ie. Device the VM can see
	VBD struct {
	}
)
