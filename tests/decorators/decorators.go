package decorators

import . "github.com/onsi/ginkgo/v2"

var (
	// SIGs
	SigCompute           = []interface{}{Label("sig-compute")}
	SigOperator          = []interface{}{Label("sig-operator")}
	SigNetwork           = []interface{}{Label("sig-network")}
	SigStorage           = []interface{}{Label("sig-storage")}
	SigComputeRealtime   = []interface{}{Label("sig-compute-realtime")}
	SigComputeMigrations = []interface{}{Label("sig-compute-migrations")}
	SigMonitoring        = []interface{}{Label("sig-monitoring")}

	// HW
	GPU        = []interface{}{Label("GPU")}
	VGPU       = []interface{}{Label("VGPU")}
	SRIOV      = []interface{}{Label("SRIOV")}
	StorageReq = []interface{}{Label("storage-req")}
	Multus     = []interface{}{Label("Multus")}
	Macvtap    = []interface{}{Label("Macvtap")}
	Invtsc     = []interface{}{Label("Invtsc")}

	// Features
	Sysprep          = []interface{}{Label("Sysprep")}
	Windows          = []interface{}{Label("Windows")}
	Networking       = []interface{}{Label("Networking")}
	VMIlifecycle     = []interface{}{Label("VMIlifecycle")}
	Expose           = []interface{}{Label("Expose")}
	NonRoot          = []interface{}{Label("verify-non-root")}
	NativeSsh        = []interface{}{Label("native-ssh")}
	ExcludeNativeSsh = []interface{}{Label("exclude-native-ssh")}
	Reenlightenment  = []interface{}{Label("Reenlightenment")}
	PasstGate        = []interface{}{Label("PasstGate")}
	Upgrade          = []interface{}{Label("Upgrade")}
)
