package opennextcdk


type BaseSiteEnvironmentOutputsInfo struct {
	EnvironmentOutputs *map[string]*string `field:"required" json:"environmentOutputs" yaml:"environmentOutputs"`
	Path *string `field:"required" json:"path" yaml:"path"`
	Stack *string `field:"required" json:"stack" yaml:"stack"`
}
