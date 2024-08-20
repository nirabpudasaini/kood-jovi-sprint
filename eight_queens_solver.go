package sprint

import (
	"strconv"
	"strings"
)

func EightQueensSolver() string {
	var solutions []string
	var result strings.Builder

	// Boolean arrays to keep track of the attack positions
	columns := [8]bool{} // Tracks if a column is under attack
	diag1 := [15]bool{}  // Tracks the main diagonal (row - col)
	diag2 := [15]bool{}  // Tracks the secondary diagonal (row + col)
	queens := [8]int{}   // Holds the queen's positions

	// Start solving the problem from the first row
	solve(0, &queens, &columns, &diag1, &diag2, &solutions)

	// Combine all solutions into the final output string
	for _, solution := range solutions {
		result.WriteString(solution)
		result.WriteString("\n")
	}

	return result.String()
}

// Helper function to solve the eight queens problem recursively
func solve(row int, queens *[8]int, columns *[8]bool, diag1, diag2 *[15]bool, solutions *[]string) {
	if row == 8 {
		// Construct the solution string from the queens' positions
		var solution strings.Builder
		for i := 0; i < 8; i++ {
			solution.WriteString(strconv.Itoa(queens[i] + 1))
		}
		*solutions = append(*solutions, solution.String())
		return
	}

	for col := 0; col < 8; col++ {
		if !columns[col] && !diag1[row-col+7] && !diag2[row+col] {
			// Place the queen
			queens[row] = col
			columns[col] = true
			diag1[row-col+7] = true
			diag2[row+col] = true

			// Move to the next row
			solve(row+1, queens, columns, diag1, diag2, solutions)

			// Backtrack and remove the queen
			columns[col] = false
			diag1[row-col+7] = false
			diag2[row+col] = false
		}
	}
}
