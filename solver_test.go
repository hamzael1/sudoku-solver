package sudoku


import "testing"
func TestSolve(t *testing.T) {
	solved, _ := solve(InitGrid());
	if solved == false {
		t.Error("Error Solving solvable puzzle");
	}

	unsolvable := [SIZE][SIZE]int {
		{ 5, 1, 6, 8, 4, 9, 7, 3, 2 },
		{ 3, 0, 7, 6, 0, 5, 0, 0, 0 },
		{ 8, 0, 9, 7, 0, 0, 0, 6, 5 },
		{ 1, 3, 5, 0, 6, 0, 9, 0, 7 },
		{ 4, 7, 2, 5, 9, 1, 0, 0, 6 },
		{ 9, 6, 8, 3, 7, 0, 0, 5, 0 },
		{ 2, 5, 3, 1, 8, 6, 0, 7, 4 },
		{ 6, 8, 4, 2, 0, 7, 5, 0, 0 },
    	{ 7, 9, 1, 0, 5, 0, 6, 0, 8 },
	};

	solved, _ = solve(unsolvable);
	if solved {
		t.Error("Error solving unsolvable puzzle")
	}
}