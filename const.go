package ray

// Predefined ray identifiers
const (
	RootID = "ROOT"
	BootID = "BOOT"
)

// ROOT ray. Other rays are forked from this one
var ROOT Interface

// BOOT ray. To be used during startup/shutdown process
var BOOT Interface

func init() {
	// Building ROOT ray
	rc := create(nil)
	rc.id = RootID
	rc.depth = 0
	rc.name = RootID

	ROOT = rc

	// Building BOOT ray
	bc := create(rc)
	bc.id = BootID
	bc.name = BootID

	BOOT = bc
}
