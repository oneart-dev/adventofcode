package star18

func Fun(input []string) interface{} {
	x := 0
	y := 0
	runes := make([][]rune, 0, len(input))

	for i, line := range input {
		if len(line) == 0 {
			continue
		}

		runes = append(runes, []rune(line))

		for ii, r := range line {
			if r == 'S' {
				x = ii
				y = i
			}
		}
	}

	steps := 0
	px, py := 0, 0
	for {
		steps++
		nx, ny := findNextMove(runes, x, y, px, py)
		px, py = x, y
		x, y = nx, ny
		if input[y][x] == 'S' {
			break
		}
	}

	return steps / 2
}

type Dir struct {
	x, y int
	dir  rune
}

var up = []Dir{{0, -1, '|'}, {0, -1, '7'}, {0, -1, 'F'}}
var down = []Dir{{0, 1, '|'}, {0, 1, 'J'}, {0, 1, 'L'}}
var left = []Dir{{-1, 0, '-'}, {-1, 0, 'L'}, {-1, 0, 'F'}}
var right = []Dir{{1, 0, '-'}, {1, 0, 'J'}, {1, 0, '7'}}

var allowed = map[rune][]Dir{
	'S': append(append(append(up, down...), right...), left...),
	'|': append(up, down...),
	'-': append(left, right...),
	'7': append(left, down...),
	'J': append(up, left...),
	'L': append(up, right...),
	'F': append(down, right...),
}

func findNextMove(nums [][]rune, x, y, px, py int) (int, int) {
	m := allowed[nums[y][x]]
	for _, d := range m {
		if v := checkMove(nums, x, y, px, py, d); v != nil {
			return x + d.x, y + d.y
		}
	}

	panic("should not reach here")
}

func checkMove(nums [][]rune, x, y, px, py int, d Dir) *rune {
	ty := y + d.y
	tx := x + d.x

	if tx < 0 || ty < 0 || tx >= len(nums[y]) || ty >= len(nums) {
		return nil
	}

	if tx == px && ty == py {
		return nil
	}

	if nums[ty][tx] == 'S' {
		s := 'S'
		return &s
	}

	if d.dir == nums[ty][tx] {
		return &nums[ty][tx]
	}

	return nil
}
