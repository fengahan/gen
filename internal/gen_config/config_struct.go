package gen_config

//@ConfigStruct(key="###")
type ProjectItem struct {
	AddrPort string `yaml:"AddrPort"`
	GenBuildPath string `yaml:"GenBuildPath"`
	GenBuildPackage string `yaml:"GenBuildPackage"`
	WireGenFile []string `yaml:"WireGenFile"`
	Main string `yaml:"Main"`
}
//@ConfigStruct()
type GenSystemConfig struct {
	WirePath string `yaml:"WirePath"`
	ProjectList []ProjectItem `yaml:"ProjectList"`
}

