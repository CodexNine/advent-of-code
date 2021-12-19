module github.com/CodexNine/advent-of-code/day-1/src

go 1.17

require github.com/CodexNine/advent-of-code/day-1/src/pkg/readFile v0.0.0-20211218213141-e529de4c1d1a

require (
	checkNums v0.0.0-00010101000000-000000000000 // indirect
	readFile v0.0.0-00010101000000-000000000000 // indirect
)

replace readFile => ./pkg/readFile

replace checkNums => ./pkg/checkNums
