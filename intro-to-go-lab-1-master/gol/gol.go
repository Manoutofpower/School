package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	aliveCells := []cell{}
	for i, v := range world {
		for j, r := range v {
			if r == 255 {
				aliveCells = append(aliveCells, cell{i, j})
			}
		}
	}
	return aliveCells
}
