package metrics

var defaultTags string

func SetDefaultTags(tags AnyTags) {
	defaultTags = tags.String()
}

func SetDefaultGCEnabled(v bool) {}
