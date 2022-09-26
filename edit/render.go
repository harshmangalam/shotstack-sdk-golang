package edit

type Timeline struct {
}

type Output struct {
}

type Merge struct {
}

type Edit struct {
	timeline Timeline
	output   Output
	merge    []Merge
	callback string
	disk     string // "local" | "mount"
}
