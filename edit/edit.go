package edit

type Disk string

const (
	Local Disk = "local"
	Mount Disk = "mount"
)

type Edit struct {
	Timeline *Timeline     `json:"timeline"`
	Output   *Output       `json:"output"`
	Merge    *[]MergeField `json:"merge,omitempty"`
	Callback string        `json:"callback,omitempty"`
	Disk     Disk          `json:"disk,omitempty"`
}

func NewEdit() *Edit {
	e := new(Edit)
	return e

}

func (e *Edit) SetTimeline(timeline *Timeline) *Edit {
	e.Timeline = timeline
	return e
}

func (e *Edit) SetOutput(output *Output) *Edit {
	e.Output = output
	return e
}

func (e *Edit) SetMerges(merge *[]MergeField) *Edit {
	e.Merge = merge
	return e
}

func (e *Edit) SetCallback(callback string) *Edit {
	e.Callback = callback
	return e
}

func (e *Edit) SetDisk(disk Disk) *Edit {
	e.Disk = disk
	return e
}
