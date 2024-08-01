package main



// func heuristicFloodFill(snapshot GameSnapshot) int {
// 	// Create a copy of the board to avoid modifying the original
// 	tempBoard := make([][]int, snapshot.Height())
// 	for i := range tempBoard {
// 		tempBoard[i] = make([]int, snapshot.Width())
// 		copy(tempBoard[i], snapshot.Board()[i])
// 	}

// 	// Set the snake's body as obstacles
// 	for _, s := range snapshot.Snake() {
// 		tempBoard[s[0]][s[1]] = -1
// 	}

// 	// Set the head as the starting point
// 	tempBoard[snapshot.Head()[0]][snapshot.Head()[1]] = 1

// 	// Perform flood fill
// 	queue := [][]int{snapshot.Head()}
// 	accessibleSpaces := 0
// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]
// 		accessibleSpaces++

// 		// Explore adjacent cells
// 		for _, direction := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
// 			nextX := current[0] + direction[0]
// 			nextY := current[1] + direction[1]

// 			// Check if the cell is within the board boundaries and empty
// 			if nextX >= 0 && nextX < snapshot.Height() && nextY >= 0 && nextY < snapshot.Width() && tempBoard[nextX][nextY] == 0 {
// 				// Mark the cell as visited
// 				tempBoard[nextX][nextY] = 1
// 				// Add the cell to the queue
// 				queue = append(queue, []int{nextX, nextY})
// 			}
// 		}
// 	}

// 	return accessibleSpaces
// }