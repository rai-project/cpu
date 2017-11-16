// +build amd64

package cpuid

import (
	"github.com/intel-go/cpuid"
)

func SupportsAVX() bool {
	return cpuid.EnabledAVX && cpuid.HasFeature(cpuid.AVX)
}

func SupportsAVX2() bool {
	return cpuid.EnabledAVX && cpuid.HasFeature(cpuid.AVX2)
}

func SupportsAVX512() bool {
	return cpuid.EnabledAVX512 && cpuid.HasFeature(cpuid.AVX512F)
}
