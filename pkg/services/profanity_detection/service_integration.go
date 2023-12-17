package profanity_detection

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

var (
	_ servicemesh.Service                  = &Service{}
	_ servicemesh.HasLogger                = &Service{}
	_ servicemesh.HasDependencies          = &Service{}
	_ config_file_manager.HasDefaultConfig = &Service{}
	_ DetectsProfanity                     = &Service{}
)

type Dependency = DetectsProfanity

type DetectsProfanity interface {
	IsProfane(string) bool
	ExtractProfanity(string) string
	Censor(string) string
}
