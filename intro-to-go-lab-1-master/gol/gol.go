package main


var (
	ALive byte = 255
	Dead byte = 0
)

func calculateNextState(p golParams, world [][]byte) [][]byte {
	for i := 0 ; i < p.imageHeight ; i++ {
		for j := 0 ; j < p.imageWidth ; j++ {
			v := world[i][j]
			if v == ALive {

			} else if (v == Dead){

			}
		}
	}
	return world
}

func calculateNeighbours(i int ,j int, world [][]byte) (numOfLive int , numOfDead int){
	if i - 1 > 0 {
		
	}
	return 0 ,0
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
