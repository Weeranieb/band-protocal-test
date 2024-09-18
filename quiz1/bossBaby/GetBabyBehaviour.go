package bossbaby

const (
	Bad  = "Bad boy"
	Good = "Good boy"
)

func GetBabyBehaviour(action string) string {
	// simply check
	if len(action) < 1 || action[0] == 'R' || action[len(action)-1] == 'S' {
		return Bad
	}

	// Shot is +, Revenge is- = Good, < 0 means Bad
	counter := 0
	isRevengeMode := false

	for _, char := range action {
		// check if boss baby finish revenge
		if char == 'S' && isRevengeMode {
			// boss baby revenge more than he got shot
			if counter < 0 {
				return Bad
			}
			isRevengeMode = false
		}

		if char == 'S' {
			counter++
		} else {
			if !isRevengeMode {
				isRevengeMode = true
			}
			counter--
		}
	}

	// boss baby still in not finish revenge his neighbour
	if counter > 0 {
		return Bad
	}

	return Good
}
