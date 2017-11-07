// +build amd64

package cpuid

import (
	"github.com/intel-go/cpuid"
)

func SupportsAVX() bool {
	return cpuid.EnabledAVX && cpuid.HasFeature(cpuid.AVX)
}
