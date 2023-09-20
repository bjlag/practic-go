package label

type Label struct {
	start int
	end   int
	s     string
}

func label(str string, labels []Label) string {
	if len(labels) == 0 {
		return str
	}

	for _, label := range labels {
		str = str[:label.start-1] + label.s + str[label.end:]
	}

	return str
}
