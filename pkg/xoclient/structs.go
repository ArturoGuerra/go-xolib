package xoclient

type (
	// VDIRef is a string reference to a VDI
	VDIRef string
	// VBDRef is a string reference to a VBD
	VBDRef string
	// VMRef is a string reference to a VM
	VMRef string
	// SRRef is a string reference to an SR
	SRRef string

	// SR is a Storage Repository used to store VDIs
	SR struct {
		Type      string   `json:"type" mapstructure:"type"`
		ID        string   `json:"id" mapstructure:"id"`
		Pool      string   `json:"$pool" mapstructure:"$pool"`
		PoolID    string   `json:"$poolId" mapstructure:"$poolId"`
		UUID      string   `json:"uuid" mapstructure:"uuid"`
		SRType    string   `json:"SR_type"`
		VDIs      []string `json:"VDIs"`
		PBDs      []string `json:"$PBDs"`
		NameLabel string   `json:"name_label"`
		Size      int64    `json:"size"`
		Usage     int64    `json:"usage"`
	}

	// VDI is a Virtual Disk Image that is stored in a Storage Repository or SR
	VDI struct {
		Type      string   `json:"type" mapstructure:"type"`
		ID        VDIRef   `json:"id" mapstructure:"id"`
		Pool      string   `json:"$pool" mapstructure:"$pool"`
		PoolID    string   `json:"$poolId" mapstructure:"$poolId"`
		UUID      VDIRef   `json:"uuid" mapstructure:"uuid"`
		SR        SRRef    `json:"$SR" mapstructure:"$SR"`
		VBDs      []VBDRef `json:"$VBDs" mapstructure:"$VBDs"`
		Size      int64    `json:"size" mapstructure:"size"`
		Usage     int64    `json:"usage" mapstructure:"usage"`
		NameLabel string   `json:"name_label" mapstructure:"name_label"`
	}

	// VBD is a Virtual Device Block that is attached to a VM
	VBD struct {
		Type     string `json:"type" mapstructure:"type"`
		ID       string `json:"id" mapstructure:"id"`
		Pool     string `json:"$pool" mapstructure:"$pool"`
		PoolID   string `json:"$poolId" mapstructure:"$poolId"`
		UUID     string `json:"uuid" mapstructure:"uuid"`
		VDI      string `json:"VDI" mapstructure:"VDI"`
		VM       string `json:"VM" mapstructure:"VM"`
		Attached bool   `json:"attached" mapstructure:"attached"`
		Device   string `json:"device" mapstructure:"device"`
	}

	// VM is a VM Ref
	VM struct {
		Type      string   `json:"type" mapstructure:"type"`
		ID        string   `json:"id" mapstructure:"id"`
		Pool      string   `json:"$pool" mapstructure:"$pool"`
		PoolID    string   `json:"$poolId" mapstructure:"$poolId"`
		UUID      string   `json:"uuid" mapstructure:"uuid"`
		VBDs      []string `json:"$VBDs" mapstructure:"$VBDs"`
		NameLabel string   `json:"name_label" mapstructure:"name_label"`
	}

	// Host is an XCP-ng host or server connected to Xen Orchestra
	Host struct {
		Type      string   `mapstructure:"type"`
		ID        string   `mapstructure:"id"`
		Pool      string   `mapstructure:"$pool"`
		PoolID    string   `mapstructure:"$poolId"`
		UUID      string   `mapstructure:"uuid"`
		Version   string   `mapstructure:"version"`
		NameLabel string   `mapstructure:"name_label"`
		Tags      []string `mapstructure:"tags"`
		Hostname  string   `mapstructure:"hostname"`
		Enabled   bool     `mapstructure:"enabled"`
	}
)
