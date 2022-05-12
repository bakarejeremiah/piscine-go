package piscine

func ConcatParams(args []string) string {
	var jeremiah string = args[0]
	for _, s := range args[1:] {
		jeremiah += "\n" + s
	}

	return jeremiah
}
