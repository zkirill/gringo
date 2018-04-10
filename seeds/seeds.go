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

// InitialDifficulty is the initial block difficulty for testnet 2.
// TODO: Move this to the appropriate package.
const InitialDifficulty uint64 = 1000
