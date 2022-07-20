//go:build !dev
// +build !dev

package appmod

func IsDev() bool {
	return false
}
