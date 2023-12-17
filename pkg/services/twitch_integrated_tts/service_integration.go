package twitch_integrated_tts

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

var (
	_ servicemesh.Service                  = &Service{}
	_ servicemesh.HasLogger                = &Service{}
	_ servicemesh.HasDependencies          = &Service{}
	_ config_file_manager.HasDefaultConfig = &Service{}
)
