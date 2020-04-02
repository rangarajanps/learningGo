//dog package provides method to convert human years to dog years
package dog

// Years returns the converted dog years for a given human year
func Years(humanYear uint) uint {
	return humanYear * 7
}
