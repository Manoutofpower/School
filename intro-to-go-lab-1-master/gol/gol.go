package main

var (
	ALive byte = 255
	Dead  byte = 0
)

func calculateNextState(p golParams, world [][]byte) [][]byte {
	w := copyWorld(world)
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			numOfLive, _ := calculateNeighbours(x, y, w)
			v := world[y][x]
			if v == ALive {
				switch {
				case numOfLive < 2:
					world[y][x] = Dead
				case numOfLive > 3:
					world[y][x] = Dead
				}
			} else if v == Dead {
				if numOfLive == 3 {
					world[y][x] = ALive
				}
			}
		}
	}
	return world
}

func copyWorld(world [][]byte) [][]byte {
	w := make([][]byte, len(world), len(world[0]))
	for k, v := range world {
		w[k] = make([]byte, len(v))
		copy(w[k], world[k])
	}
	return w
}

func calculateNeighbours(x1 int, y1 int, world [][]byte) (numOfLive int, numOfDead int) {
	for y := y1 - 1; y <= y1+1; y++ {
		for x := x1 - 1; x <= x1+1; x++ {
			if (x >= 0 && y >= 0) &&
				(y < len(world) && x < len(world[0])) &&
				!(x == x1 && y == y1) {
				if world[y][x] == ALive {
					numOfLive = numOfLive + 1
				}
				if world[y][x] == Dead {
					numOfDead = numOfDead + 1
				}
			}
		}
	}
	return numOfLive, numOfDead
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	aliveCells := []cell{}
	for i, v := range world {
		for j, r := range v {
			if r == 255 {
				aliveCells = append(aliveCells, cell{j, i})
			}
		}
	}
	return aliveCells
}
