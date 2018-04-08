// Package seeds is responsible for listing seeds.
package seeds

// Seeds returns network seeds from http://grin-tech.org/seeds.txt.
func Seeds() []string {
	return []string{
		"192.241.160.172",
		"109.74.202.16",
		"198.245.50.26",
		"46.4.91.48",
	}
}
