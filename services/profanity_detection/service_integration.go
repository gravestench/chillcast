package profanity_detection

import (
	"github.com/gravestench/runtime"

	"github.com/gravestench/chillcast/services/config_file_manager"
)

var (
	_ runtime.Service                      = &Service{}
	_ runtime.HasLogger                    = &Service{}
	_ runtime.HasDependencies              = &Service{}
	_ config_file_manager.HasDefaultConfig = &Service{}
	_ DetectsProfanity                     = &Service{}
)

type Dependency = DetectsProfanity

type DetectsProfanity interface {
	IsProfane(string) bool
	ExtractProfanity(string) string
	Censor(string) string
}
