//go:build windows && 386

package main

// On 386, __stdcall (WINAPI) exports are name-decorated as _Name@bytes, but CNG resolves
// GetKeyStorageInterface by its plain, undecorated name. --kill-at strips the @bytes suffix so
// the export matches what CNG looks up. This is a no-op on amd64, where stdcall exports are
// already undecorated, so it is scoped to 386 to leave the amd64 build untouched.

/*
#cgo LDFLAGS: -Wl,--kill-at
*/
import "C"
