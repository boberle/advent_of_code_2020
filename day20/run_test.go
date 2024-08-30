package day20

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func Test_parseTileContent(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  Tile
	}{
		{"2x2 tile", []string{"#..#", "##.#", "..##", "##.#"}, Tile{
			sideLen:        4,
			contentSideLen: 2,
			topSide:        Side{0, 3},
			bottomSide:     Side{0, 1, 3},
			leftSide:       Side{0, 1, 3},
			rightSide:      Side{0, 1, 2, 3},
			content:        Content{{0, 0}, {1, 1}},
		}},
		{"3x3 tile", []string{"..##.", "##..#", "#...#", "####.", "##.##"}, Tile{
			sideLen:        5,
			contentSideLen: 3,
			topSide:        Side{2, 3},
			bottomSide:     Side{0, 1, 3, 4},
			leftSide:       Side{1, 2, 3, 4},
			rightSide:      Side{1, 2, 4},
			content:        Content{{0, 0}, {0, 2}, {1, 2}, {2, 2}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tile{sideLen: tt.want.sideLen, contentSideLen: tt.want.contentSideLen}
			parseTileContent(tt.input, &got)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseFile(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Tile
	}{
		{"2x 2x2", "Tile 123:\n#..#\n##.#\n..##\n##.#\n\nTile 456:\n#...\n##.#\n...#\n##.#", []Tile{
			{
				id:             123,
				sideLen:        4,
				contentSideLen: 2,
				topSide:        Side{0, 3},
				bottomSide:     Side{0, 1, 3},
				leftSide:       Side{0, 1, 3},
				rightSide:      Side{0, 1, 2, 3},
				content:        Content{{0, 0}, {1, 1}},
			},
			{
				id:             456,
				sideLen:        4,
				contentSideLen: 2,
				topSide:        Side{0},
				bottomSide:     Side{0, 1, 3},
				leftSide:       Side{0, 1, 3},
				rightSide:      Side{1, 2, 3},
				content:        Content{{0, 0}},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			assert.Equalf(t, tt.want, parseFile(reader), "parseFile(%v)", tt.input)
		})
	}
}

func TestSide_reverse(t *testing.T) {
	side := Side{1, 3}
	assert.Equal(t, Side{0, 2}, side.reverse(4))
}

func getContent3() Content {
	return Content{{1, 0}, {0, 1}, {2, 1}, {2, 2}}
}

func checkContent(t *testing.T, want, got Content) {
	wantMap := map[Point]struct{}{}
	for _, point := range want {
		wantMap[point] = struct{}{}
	}
	gotMap := map[Point]struct{}{}
	for _, point := range got {
		gotMap[point] = struct{}{}
	}
	assert.Equal(t, wantMap, gotMap)
}

func TestContent_rotate90(t *testing.T) {
	want := Content{{1, 0}, {2, 1}, {0, 2}, {1, 2}}
	content := getContent3()
	checkContent(t, want, content.rotate90(3))
}

func TestContent_rotate180(t *testing.T) {
	want := Content{{0, 0}, {0, 1}, {2, 1}, {1, 2}}
	content := getContent3()
	checkContent(t, want, content.rotate180(3))
}

func TestContent_rotate270(t *testing.T) {
	want := Content{{1, 0}, {2, 0}, {0, 1}, {1, 2}}
	content := getContent3()
	checkContent(t, want, content.rotate270(3))
}

func TestContent_flipH(t *testing.T) {
	want := Content{{1, 0}, {0, 1}, {2, 1}, {0, 2}}
	content := getContent3()
	checkContent(t, want, content.flipH(3))
}

func TestContent_flipV(t *testing.T) {
	want := Content{{1, 2}, {0, 1}, {2, 1}, {2, 0}}
	content := getContent3()
	checkContent(t, want, content.flipV(3))
}

func getTile3() Tile {
	tile := Tile{
		sideLen:        5,
		contentSideLen: 3,
		topSide:        Side{0, 3},
		bottomSide:     Side{0, 2, 4},
		leftSide:       Side{0, 3, 4},
		rightSide:      Side{1, 2, 4},
		content:        Content{{1, 0}, {0, 1}, {2, 1}, {2, 2}},
	}
	return tile
}

func checkTile(t *testing.T, want, got Tile) {
	checkContent(t, want.content, got.content)
	assert.Equal(t, want.topSide, got.topSide)
	assert.Equal(t, want.bottomSide, got.bottomSide)
	assert.Equal(t, want.rightSide, got.rightSide)
	assert.Equal(t, want.leftSide, got.leftSide)
	assert.Equal(t, want.sideLen, got.sideLen)
	assert.Equal(t, want.contentSideLen, got.contentSideLen)
}

func TestTile_rotate90(t *testing.T) {
	want := Tile{
		sideLen:        5,
		contentSideLen: 3,
		topSide:        Side{0, 1, 4},
		bottomSide:     Side{0, 2, 3},
		leftSide:       Side{0, 2, 4},
		rightSide:      Side{0, 3},
		content:        Content{{1, 0}, {2, 1}, {0, 2}, {1, 2}},
	}
	tile := getTile3()
	tile.rotate90()
	checkTile(t, want, tile)
}

func TestTile_rotate180(t *testing.T) {
	want := Tile{
		sideLen:        5,
		contentSideLen: 3,
		topSide:        Side{0, 2, 4},
		bottomSide:     Side{1, 4},
		leftSide:       Side{0, 2, 3},
		rightSide:      Side{0, 1, 4},
		content:        Content{{0, 0}, {0, 1}, {2, 1}, {1, 2}},
	}
	tile := getTile3()
	tile.rotate180()
	checkTile(t, want, tile)
}

func TestTile_rotate270(t *testing.T) {
	want := Tile{
		sideLen:        5,
		contentSideLen: 3,
		topSide:        Side{1, 2, 4},
		bottomSide:     Side{0, 3, 4},
		leftSide:       Side{1, 4},
		rightSide:      Side{0, 2, 4},
		content:        Content{{1, 0}, {2, 0}, {0, 1}, {1, 2}},
	}
	tile := getTile3()
	tile.rotate270()
	checkTile(t, want, tile)
}

func TestTile_flipH(t *testing.T) {
	want := Tile{
		sideLen:        5,
		contentSideLen: 3,
		topSide:        Side{1, 4},
		bottomSide:     Side{0, 2, 4},
		leftSide:       Side{1, 2, 4},
		rightSide:      Side{0, 3, 4},
		content:        Content{{1, 0}, {0, 1}, {2, 1}, {0, 2}},
	}
	tile := getTile3()
	tile.flipH()
	checkTile(t, want, tile)
}

func TestTile_flipV(t *testing.T) {
	want := Tile{
		sideLen:        5,
		contentSideLen: 3,
		topSide:        Side{0, 2, 4},
		bottomSide:     Side{0, 3},
		leftSide:       Side{0, 1, 4},
		rightSide:      Side{0, 2, 3},
		content:        Content{{1, 2}, {0, 1}, {2, 1}, {2, 0}},
	}
	tile := getTile3()
	tile.flipV()
	checkTile(t, want, tile)
}

func Test_joinTiles(t *testing.T) {
	tile1 := Tile{contentSideLen: 2, content: Content{{0, 0}, {1, 0}, {0, 1}}}
	tile2 := Tile{contentSideLen: 2, content: Content{{0, 1}, {1, 1}}}
	tile3 := Tile{contentSideLen: 2, content: Content{{1, 1}}}
	tile4 := Tile{contentSideLen: 2, content: Content{{0, 0}, {0, 1}}}
	tile1.rightTile = &tile2
	tile2.leftTile = &tile1
	tile1.bottomTile = &tile3
	tile3.topTile = &tile1
	tile3.rightTile = &tile4
	tile4.leftTile = &tile3
	tile2.bottomTile = &tile4
	tile4.topTile = &tile2
	tiles := []Tile{tile1, tile2, tile3, tile4}

	want := Image{
		content: Content{{0, 0}, {1, 0}, {0, 1}, {2, 1}, {3, 1}, {2, 2}, {1, 3}, {2, 3}},
		sideLen: 4,
	}
	got := joinTiles(tiles)
	assert.Equal(t, want.sideLen, got.sideLen)
	checkContent(t, want.content, got.content)
}

func TestContent_findMonster(t *testing.T) {
	content := Content{
		{1, 0},
		{2, 0},
		{4, 0},
		{2, 1},
		{3, 1},
		{2, 2},
		{3, 2},
		{0, 3},
		{1, 3},
		{3, 3},
	}
	tests := []struct {
		name    string
		monster Content
		want    int
	}{
		{"no monster", Content{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {3, 0}, {3, 1}}, 0},
		{"one monster", Content{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {3, 0}}, 1},
		{"two monsters", Content{{0, 1}, {1, 1}, {2, 0}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, content.findMonster(tt.monster), "findMonster(%v)", tt.monster)
		})
	}
}

func TestImage_findMonster(t *testing.T) {
	content := Content{
		{1, 0},
		{2, 0},
		{4, 0},
		{2, 1},
		{3, 1},
		{2, 2},
		{3, 2},
		{0, 3},
		{1, 3},
		{3, 3},
	}
	content = content.rotate90(4)
	content.flipV(4)
	image := Image{sideLen: 4, content: content}
	tests := []struct {
		name    string
		monster Content
		want    int
	}{
		{"no monster", Content{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {3, 0}, {3, 1}}, 0},
		{"one monster", Content{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {3, 0}}, 1},
		// 3 because overlaps
		{"two monsters", Content{{0, 1}, {1, 1}, {2, 0}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, image.findMonster(tt.monster), "findMonster(%v)", tt.monster)
		})
	}
}

func TestTile_pairRight_withLeft(t *testing.T) {
	tile1 := Tile{
		sideLen:        4,
		contentSideLen: 2,
		rightSide:      Side{0, 1},
	}
	tile2 := Tile{
		sideLen:        4,
		contentSideLen: 2,
		leftSide:       Side{0, 1},
	}
	found := tile1.pairRight(&tile2)
	assert.Equal(t, true, found)
	assert.Equal(t, &tile2, tile1.rightTile)
	assert.Equal(t, &tile1, tile2.leftTile)
}

func TestTile_pairRight_withTop(t *testing.T) {
	tile1 := Tile{
		sideLen:        4,
		contentSideLen: 2,
		rightSide:      Side{0, 1},
	}
	tile2 := Tile{
		sideLen:        4,
		contentSideLen: 2,
		topSide:        Side{0, 1},
		rightSide:      Side{2, 3},
	}
	found := tile1.pairRight(&tile2)
	assert.Equal(t, true, found)
	assert.Equal(t, &tile2, tile1.rightTile)
	assert.Equal(t, &tile1, tile2.leftTile)
	assert.Equal(t, Side{2, 3}, tile2.bottomSide)
}

func TestTile_findRight(t *testing.T) {
	tile1 := Tile{
		sideLen:        4,
		contentSideLen: 2,
		rightSide:      Side{0, 1},
	}
	tile2 := Tile{
		sideLen:        4,
		contentSideLen: 2,
		topSide:        Side{0, 1},
		rightSide:      Side{2, 3},
	}
	tile3 := Tile{
		sideLen:        4,
		contentSideLen: 2,
		bottomSide:     Side{0, 2},
		rightSide:      Side{1, 2, 3},
	}
	tiles := []*Tile{&tile3, &tile2}
	tile1.findRight(tiles)
	assert.Equal(t, &tile2, tile1.rightTile)
	assert.Equal(t, &tile1, tile2.leftTile)
	assert.Equal(t, Side{2, 3}, tile2.bottomSide)
}

func Test_assembleTiles_noRotation(t *testing.T) {
	tile1 := Tile{
		id:             1,
		sideLen:        4,
		contentSideLen: 2,
		rightSide:      Side{0},
		bottomSide:     Side{0, 1, 2},
	}
	tile2 := Tile{
		id:             2,
		sideLen:        4,
		contentSideLen: 2,
		leftSide:       Side{0},
		bottomSide:     Side{0, 1, 2, 3},
	}
	tile3 := Tile{
		id:             3,
		sideLen:        4,
		contentSideLen: 2,
		topSide:        Side{0, 1, 2},
		rightSide:      Side{0, 1},
	}
	tile4 := Tile{
		id:             4,
		sideLen:        4,
		contentSideLen: 2,
		topSide:        Side{0, 1, 2, 3},
		leftSide:       Side{0, 1},
	}
	tiles := []*Tile{&tile2, &tile4, &tile3, &tile1}
	assembleTiles(tiles)
	assert.Equal(t, &tile2, tile1.rightTile)
	assert.Equal(t, &tile3, tile1.bottomTile)
	assert.Equal(t, &tile4, tile2.bottomTile)
	assert.Equal(t, &tile1, tile2.leftTile)
	assert.Equal(t, &tile1, tile3.topTile)
	assert.Equal(t, &tile4, tile3.rightTile)
	assert.Equal(t, &tile2, tile4.topTile)
	assert.Equal(t, &tile3, tile4.leftTile)

	assert.Nil(t, tile1.topTile)
	assert.Nil(t, tile1.leftTile)
	assert.Nil(t, tile2.topTile)
	assert.Nil(t, tile2.rightTile)
	assert.Nil(t, tile3.bottomTile)
	assert.Nil(t, tile3.leftTile)
	assert.Nil(t, tile4.rightTile)
	assert.Nil(t, tile4.bottomTile)
}

func Test_assembleTiles_withRotation(t *testing.T) {
	tile1 := Tile{
		id:             1,
		sideLen:        4,
		contentSideLen: 2,
		rightSide:      Side{0},
		bottomSide:     Side{0, 1, 2},
	}
	tile2 := Tile{
		id:             2,
		sideLen:        4,
		contentSideLen: 2,
		topSide:        Side{3},
		leftSide:       Side{0, 1, 2, 3},
	}
	tile3 := Tile{
		id:             3,
		sideLen:        4,
		contentSideLen: 2,
		bottomSide:     Side{0, 1, 2},
		rightSide:      Side{2, 3},
	}
	tile4 := Tile{
		id:             4,
		sideLen:        4,
		contentSideLen: 2,
		leftSide:       Side{0, 1, 2, 3},
		topSide:        Side{0, 1},
		// limitations:
		// the following won't work because you can't rotate or flip a tile twice, and
		// here the top side is symmetric, so it can serve two orientations. Depending
		// on the order, the system has no way to now which orientation is the correct
		// one, and since there can't be any correction later (no rotation/flipping twice)
		// you might end up stuck
		//topSide:   Side{0, 1, 2, 3},
		//rightSide: Side{0, 1},
	}
	tiles := []*Tile{&tile1, &tile2, &tile3, &tile4}
	assembleTiles(tiles)
	assert.Equal(t, &tile2, tile1.rightTile)
	assert.Equal(t, &tile3, tile1.bottomTile)
	assert.Equal(t, &tile4, tile2.bottomTile)
	assert.Equal(t, &tile1, tile2.leftTile)
	assert.Equal(t, &tile1, tile3.topTile)
	assert.Equal(t, &tile4, tile3.rightTile)
	assert.Equal(t, &tile2, tile4.topTile)
	assert.Equal(t, &tile3, tile4.leftTile)

	assert.Nil(t, tile1.topTile)
	assert.Nil(t, tile1.leftTile)
	assert.Nil(t, tile2.topTile)
	assert.Nil(t, tile2.rightTile)
	assert.Nil(t, tile3.bottomTile)
	assert.Nil(t, tile3.leftTile)
	assert.Nil(t, tile4.rightTile)
	assert.Nil(t, tile4.bottomTile)
}

func Test_assembleTile_withCompleteDataset(t *testing.T) {
	fh, err := os.Open("input_test")
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	tileList := parseFile(fh)
	tiles := map[int]*Tile{}
	for i, tile := range tileList {
		tiles[tile.id] = &tileList[i]
	}
	availableTiles := []*Tile{tiles[2311], tiles[1951], tiles[1171], tiles[1427], tiles[1489], tiles[2473], tiles[2971], tiles[2729], tiles[3079]}
	assembleTiles(availableTiles)

}
