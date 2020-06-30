package server

type Mount struct {
	Type          Type
	Source        string
	Target        string
	ReadOnly      bool
	Consistency   Consistency
	BindOptions   BindOptions
	VolumeOptions VolumeOptions
	TmpfsOptions  TmpfsOptions
}
