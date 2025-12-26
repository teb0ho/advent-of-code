func printingDepartment() {
	file, err := os.Open("./files/day3.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	total := 0
	lineCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			character := line[i]

		}
	}
}