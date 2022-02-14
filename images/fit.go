package images

type Fit int

const (
	// fit directory
	Fill Fit = 0
	// fit innner with the same aspect ratio
	Contain Fit = 1
	// fit to the maximusize evenif it strcks out with the same aspect ratio
	Cover Fit = 2
)
