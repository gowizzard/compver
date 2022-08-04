package version_core

func Compare(blocks []Block) string {

	for _, value := range blocks {

		switch {
		case value.Number1 > value.Number2:
			return value.Name + " update"
		case value.Number1 == value.Number2:
			break
		case value.Number1 < value.Number2:
			return value.Name + " downgrade"
		}

	}

	return "no changes"

}
