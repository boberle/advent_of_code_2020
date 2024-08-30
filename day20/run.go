package day20

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Side []int

type Content []Point

type Tile struct {
	id             int
	sideLen        int
	contentSideLen int
	topSide        Side
	bottomSide     Side
	leftSide       Side
	rightSide      Side
	content        Content
	topTile        *Tile
	bottomTile     *Tile
	leftTile       *Tile
	rightTile      *Tile
}

type Image struct {
	content Content
	sideLen int
}

func Run(infile string) {

	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	tiles := parseFile(fh)
	tilePtrs := []*Tile{}
	for i := range tiles {
		tilePtrs = append(tilePtrs, &tiles[i])
	}
	assembleTiles(tilePtrs)
	corners := findCorners(tiles)
	cornerMultiplication := 1
	for _, corner := range corners {
		cornerMultiplication *= corner.id
	}
	fmt.Printf("Ids of the corners multiplied together (part 1): %d\n", cornerMultiplication)

	monster := Content{
		{18, 0},
		{0, 1},
		{5, 1},
		{6, 1},
		{11, 1},
		{12, 1},
		{17, 1},
		{18, 1},
		{19, 1},
		{1, 2},
		{4, 2},
		{7, 2},
		{10, 2},
		{13, 2},
		{16, 2},
	}

	image := joinTiles(tiles)
	monsterCount := image.findMonster(monster)
	fmt.Printf("Found %d monsters, so there are %d - %d = %d points in the image that are not sea monsters (part 2)\n",
		monsterCount,
		len(image.content),
		monsterCount*len(monster),
		len(image.content)-monsterCount*len(monster),
	)

}

func parseFile(reader io.Reader) []Tile {
	scanner := bufio.NewScanner(reader)

	tiles := []Tile{}
	for {
		if tile, found := parseTile(scanner); found {
			tiles = append(tiles, tile)
		} else {
			break
		}
	}

	return tiles
}

func parseTile(scanner *bufio.Scanner) (Tile, bool) {
	tile := Tile{}
	if !scanner.Scan() {
		return Tile{}, false
	}
	tile.id = parseTileId(scanner.Text())

	lines := []string{}
	for scanner.Scan() && scanner.Text() != "" {
		if len(lines) == 0 {
			tile.sideLen = len(scanner.Text())
			tile.contentSideLen = len(scanner.Text()) - 2
		}
		lines = append(lines, scanner.Text())
	}

	parseTileContent(lines, &tile)
	return tile, true
}

func parseTileId(line string) int {
	pat, err := regexp.Compile("Tile (\\d+):")
	if err != nil {
		panic(nil)
	}

	matches := pat.FindStringSubmatch(line)
	if matches == nil || len(matches) != 2 {
		panic("can't find tile id")
	}

	id, err := strconv.Atoi(matches[1])
	return id
}

func parseSide(line string) Side {
	side := Side{}
	for i, char := range line {
		if char == '#' {
			side = append(side, i)
		}
	}
	return side
}

func parseTileContent(lines []string, tile *Tile) {
	for y, line := range lines {
		if y == 0 {
			tile.topSide = parseSide(line)
			if line[0] == '#' {
				tile.leftSide = append(tile.leftSide, 0)
			}
			if line[len(line)-1] == '#' {
				tile.rightSide = append(tile.rightSide, 0)
			}
		} else if y == tile.sideLen-1 {
			tile.bottomSide = parseSide(line)
			if line[0] == '#' {
				tile.leftSide = append(tile.leftSide, len(line)-1)
			}
			if line[len(line)-1] == '#' {
				tile.rightSide = append(tile.rightSide, len(line)-1)
			}
		} else {
			for x, char := range line {
				if char == '#' {
					if x == 0 {
						tile.leftSide = append(tile.leftSide, y)
					} else if x == len(lines)-1 {
						tile.rightSide = append(tile.rightSide, y)
					} else {
						tile.content = append(tile.content, Point{x - 1, y - 1})
					}
				}
			}
		}
	}
}

func (s *Side) reverse(size int) Side {
	newSide := Side{}
	for i := len(*s) - 1; i >= 0; i-- {
		newSide = append(newSide, size-(*s)[i]-1)
	}
	return newSide
}

func (c *Content) rotate90(size int) Content {
	rv := make([]Point, len(*c))
	for i, p := range *c {
		rv[i].x = size - p.y - 1
		rv[i].y = p.x
	}
	return rv
}

func (c *Content) rotate180(size int) Content {
	rv := make([]Point, len(*c))
	for i, p := range *c {
		rv[i].x = size - p.x - 1
		rv[i].y = size - p.y - 1
	}
	return rv
}

func (c *Content) rotate270(size int) Content {
	rv := make([]Point, len(*c))
	for i, p := range *c {
		rv[i].x = p.y
		rv[i].y = size - p.x - 1
	}
	return rv
}

func (c *Content) flipH(size int) Content {
	rv := Content{}
	for _, pos := range *c {
		rv = append(rv, Point{size - pos.x - 1, pos.y})
	}
	return rv
}

func (c *Content) flipV(size int) Content {
	rv := Content{}
	for _, pos := range *c {
		rv = append(rv, Point{pos.x, size - pos.y - 1})
	}
	return rv
}

func (t *Tile) isFree() bool {
	return t.topTile == nil && t.bottomTile == nil && t.rightTile == nil && t.leftTile == nil
}

func (t *Tile) rotate90() {
	if !t.isFree() {
		log.Fatalf("tile %d is not free to rotate", t.id)
	}
	t.content = t.content.rotate90(t.contentSideLen)
	t.rightSide, t.leftSide, t.topSide, t.bottomSide = t.topSide, t.bottomSide, t.leftSide.reverse(t.sideLen), t.rightSide.reverse(t.sideLen)
}

func (t *Tile) rotate180() {
	if !t.isFree() {
		log.Fatalf("tile %d is not free to rotate", t.id)
	}
	t.content = t.content.rotate180(t.contentSideLen)
	t.rightSide, t.leftSide, t.topSide, t.bottomSide = t.leftSide.reverse(t.sideLen), t.rightSide.reverse(t.sideLen), t.bottomSide.reverse(t.sideLen), t.topSide.reverse(t.sideLen)
}

func (t *Tile) rotate270() {
	if !t.isFree() {
		log.Fatalf("tile %d is not free to rotate", t.id)
	}
	t.content = t.content.rotate270(t.contentSideLen)
	t.rightSide, t.leftSide, t.topSide, t.bottomSide = t.bottomSide.reverse(t.sideLen), t.topSide.reverse(t.sideLen), t.rightSide, t.leftSide
}

func (t *Tile) flipH() {
	if !t.isFree() {
		log.Fatalf("tile %d is not free to rotate", t.id)
	}
	t.content = t.content.flipH(t.contentSideLen)
	t.rightSide, t.leftSide, t.topSide, t.bottomSide = t.leftSide, t.rightSide, t.topSide.reverse(t.sideLen), t.bottomSide.reverse(t.sideLen)
}

func (t *Tile) flipV() {
	if !t.isFree() {
		log.Fatalf("tile %d is not free to rotate", t.id)
	}
	t.content = t.content.flipV(t.contentSideLen)
	t.rightSide, t.leftSide, t.topSide, t.bottomSide = t.rightSide.reverse(t.sideLen), t.leftSide.reverse(t.sideLen), t.bottomSide, t.topSide
}

func findCorners(tiles []Tile) [4]Tile {
	corners := [4]Tile{}
	i := 0
	for _, tile := range tiles {
		sideCount := 0
		if tile.topTile == nil {
			sideCount++
		}
		if tile.bottomTile == nil {
			sideCount++
		}
		if tile.leftTile == nil {
			sideCount++
		}
		if tile.rightTile == nil {
			sideCount++
		}
		if sideCount == 2 {
			if !(i < 4) {
				log.Fatalln("more than 4 corners")
			}
			corners[i] = tile
			i++
		}
	}
	if i != 4 {
		log.Fatalf("corners found: %d\n", i)
	}

	return corners
}

func findTopLeftCorner(tiles []Tile) Tile {
	corners := findCorners(tiles)
	for _, corner := range corners {
		if corner.topTile == nil && corner.leftTile == nil {
			return corner
		}
	}
	panic("no top left corner")
}

func joinTiles(tiles []Tile) Image {
	rv := Image{}
	topLeftCorner := findTopLeftCorner(tiles)

	currentRowStart := &topLeftCorner
	sideLen := 0
	for y := 0; currentRowStart != nil; y++ {
		current := currentRowStart
		for x := 0; current != nil; x++ {
			for _, point := range current.content {
				rv.content = append(rv.content, Point{point.x + x*current.contentSideLen, point.y + y*current.contentSideLen})
			}
			current = current.rightTile
		}
		sideLen += currentRowStart.contentSideLen
		currentRowStart = currentRowStart.bottomTile
	}
	rv.sideLen = sideLen

	return rv
}

func (img *Image) findMonster(monster Content) int {
	rv := 0
	if n := img.content.findMonsterInAllRotations(img.sideLen, monster); n > rv {
		rv = n
	}
	if n := img.content.flipH(img.sideLen).findMonsterInAllRotations(img.sideLen, monster); n > rv {
		rv = n
	}
	if n := img.content.flipV(img.sideLen).findMonsterInAllRotations(img.sideLen, monster); n > rv {
		rv = n
	}
	return rv
}

func (c Content) findMonsterInAllRotations(sideLen int, monster Content) int {
	rv := 0
	if n := c.findMonster(monster); n > rv {
		rv = n
	}
	if n := c.rotate90(sideLen).findMonster(monster); n > rv {
		rv = n
	}
	if n := c.rotate180(sideLen).findMonster(monster); n > rv {
		rv = n
	}
	if n := c.rotate270(sideLen).findMonster(monster); n > rv {
		rv = n
	}
	return rv
}

func (c Content) findMonster(monster Content) int {
	// limitation: this will find overlapping monsters, if any
	var maxX, maxY int
	for _, pos := range c {
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}

	used := Content{}

	rv := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			contained := true
			u := Content{}
			for _, pos := range monster {
				if !slices.Contains(c, Point{pos.x + x, pos.y + y}) {
					contained = false
					break
				} else {
					u = append(u, Point{pos.x + x, pos.y + y})
				}
			}
			if contained {
				rv++
				used = append(used, u...)
			}
		}
	}
	//if rv > 0 {
	//	c.printMonster(used, 96)
	//	fmt.Println("==================================================")
	//}
	return rv
}

func (c *Content) printMonster(monster Content, sideLen int) {
	mat := make([][]string, sideLen)
	for i := 0; i < sideLen; i++ {
		mat[i] = make([]string, sideLen)
	}
	for i := 0; i < sideLen; i++ {
		for j := 0; j < sideLen; j++ {
			mat[i][j] = "."
		}
	}
	for _, point := range *c {
		if slices.Contains(monster, point) {
			mat[point.y][point.x] = "O"
		} else {
			mat[point.y][point.x] = "#"
		}
	}
	for y := 0; y < sideLen; y++ {
		fmt.Println(strings.Join(mat[y], ""))
	}
}

func (s Side) equal(o Side) bool {
	if o == nil {
		return false
	}
	if len(s) != len(o) {
		return false
	}
	for i, v := range s {
		if (o)[i] != v {
			return false
		}
	}
	return true
}

func (t *Tile) findRight(tiles []*Tile) {
	for _, tile := range tiles {
		if found := t.pairRight(tile); found {
			break
		}
	}
}

func (t *Tile) pairRight(tile *Tile) bool {
	side := t.rightSide
	revSide := side.reverse(t.sideLen)
	found := false
	if tile.leftTile == nil {
		if tile.leftSide.equal(side) {
			found = true
		} else if tile.leftSide.equal(revSide) {
			tile.flipV()
			found = true
		}
	}
	if !found && tile.topTile == nil {
		if tile.topSide.equal(side) {
			tile.rotate270()
			tile.flipV()
			found = true
		} else if tile.topSide.equal(revSide) {
			tile.rotate270()
			found = true
		}
	}
	if !found && tile.rightTile == nil {
		if tile.rightSide.equal(side) {
			tile.rotate180()
			tile.flipV()
			found = true
		} else if tile.rightSide.equal(revSide) {
			tile.rotate180()
			found = true
		}
	}
	if !found && tile.bottomTile == nil {
		if tile.bottomSide.equal(side) {
			tile.rotate90()
			found = true
		} else if tile.bottomSide.equal(revSide) {
			tile.rotate90()
			tile.flipV()
			found = true
		}
	}
	if found {
		t.rightTile = tile
		tile.leftTile = t
	}
	return found
}

func (t *Tile) findLeft(tiles []*Tile) {
	for _, tile := range tiles {
		if found := t.pairLeft(tile); found {
			break
		}
	}
}

func (t *Tile) pairLeft(tile *Tile) bool {
	side := t.leftSide
	revSide := side.reverse(t.sideLen)
	found := false
	if tile.leftTile == nil {
		if tile.leftSide.equal(side) {
			tile.rotate180()
			tile.flipV()
			found = true
		} else if tile.leftSide.equal(revSide) {
			tile.rotate180()
			found = true
		}
	}
	if !found && tile.topTile == nil {
		if tile.topSide.equal(side) {
			tile.rotate90()
			found = true
		} else if tile.topSide.equal(revSide) {
			tile.rotate90()
			tile.flipV()
			found = true
		}
	}
	if !found && tile.rightTile == nil {
		if tile.rightSide.equal(side) {
			found = true
		} else if tile.rightSide.equal(revSide) {
			tile.flipV()
			found = true
		}
	}
	if !found && tile.bottomTile == nil {
		if tile.bottomSide.equal(side) {
			tile.rotate270()
			tile.flipV()
			found = true
		} else if tile.bottomSide.equal(revSide) {
			tile.rotate270()
			found = true
		}
	}
	if found {
		t.leftTile = tile
		tile.rightTile = t
	}
	return found
}

func (t *Tile) findTop(tiles []*Tile) {
	for _, tile := range tiles {
		if found := t.pairTop(tile); found {
			break
		}
	}
}

func (t *Tile) pairTop(tile *Tile) bool {
	side := t.topSide
	revSide := side.reverse(t.sideLen)
	found := false
	if tile.leftTile == nil {
		if tile.leftSide.equal(side) {
			tile.rotate270()
			found = true
		} else if tile.leftSide.equal(revSide) {
			tile.rotate270()
			tile.flipH()
			found = true
		}
	}
	if !found && tile.topTile == nil {
		if tile.topSide.equal(side) {
			tile.rotate180()
			tile.flipH()
			found = true
		} else if tile.topSide.equal(revSide) {
			tile.rotate180()
			found = true
		}
	}
	if !found && tile.rightTile == nil {
		if tile.rightSide.equal(side) {
			tile.rotate90()
			tile.flipH()
			found = true
		} else if tile.rightSide.equal(revSide) {
			tile.rotate90()
			found = true
		}
	}
	if !found && tile.bottomTile == nil {
		if tile.bottomSide.equal(side) {
			found = true
		} else if tile.bottomSide.equal(revSide) {
			tile.flipH()
			found = true
		}
	}
	if found {
		t.topTile = tile
		tile.bottomTile = t
	}
	return found
}

func (t *Tile) findBottom(tiles []*Tile) {
	for _, tile := range tiles {
		if found := t.pairBottom(tile); found {
			break
		}
	}
}

func (t *Tile) pairBottom(tile *Tile) bool {
	side := t.bottomSide
	revSide := side.reverse(t.sideLen)
	found := false
	if tile.leftTile == nil {
		if tile.leftSide.equal(side) {
			tile.rotate90()
			tile.flipH()
			found = true
		} else if tile.leftSide.equal(revSide) {
			tile.rotate90()
			found = true
		}
	}
	if !found && tile.topTile == nil {
		if tile.topSide.equal(side) {
			found = true
		} else if tile.topSide.equal(revSide) {
			tile.flipH()
			found = true
		}
	}
	if !found && tile.rightTile == nil {
		if tile.rightSide.equal(side) {
			tile.rotate270()
			found = true
		} else if tile.rightSide.equal(revSide) {
			tile.rotate270()
			tile.flipH()
			found = true
		}
	}
	if !found && tile.bottomTile == nil {
		if tile.bottomSide.equal(side) {
			tile.rotate180()
			tile.flipH()
			found = true
		} else if tile.bottomSide.equal(revSide) {
			tile.rotate180()
			found = true
		}
	}
	if found {
		t.bottomTile = tile
		tile.topTile = t
	}
	return found
}

func (t *Tile) findNeighbors(tiles []*Tile, processed map[*Tile]struct{}) {
	if _, found := processed[t]; found {
		return
	}
	processed[t] = struct{}{}
	newTiles := []*Tile{}
	for _, tile := range tiles {
		if t != tile {
			newTiles = append(newTiles, tile)
		}
	}
	if t.topSide != nil && len(t.topSide) > 0 && t.topTile == nil {
		t.findTop(newTiles)
		if t.topTile != nil {
			t.topTile.findNeighbors(tiles, processed)
		}
	}
	if t.bottomSide != nil && len(t.bottomSide) > 0 && t.bottomTile == nil {
		t.findBottom(newTiles)
		if t.bottomTile != nil {
			t.bottomTile.findNeighbors(tiles, processed)
		}
	}
	if t.leftSide != nil && len(t.leftSide) > 0 && t.leftTile == nil {
		t.findLeft(newTiles)
		if t.leftTile != nil {
			t.leftTile.findNeighbors(tiles, processed)
		}
	}
	if t.rightSide != nil && len(t.rightSide) > 0 && t.rightTile == nil {
		t.findRight(newTiles)
		if t.rightTile != nil {
			t.rightTile.findNeighbors(tiles, processed)
		}
	}
}

func assembleTiles(tiles []*Tile) {
	processed := map[*Tile]struct{}{}
	for _, tile := range tiles {
		tile.findNeighbors(tiles, processed)
	}
	emptySides := 0
	for _, tile := range tiles {
		if tile.topTile == nil {
			emptySides++
		}
		if tile.bottomTile == nil {
			emptySides++
		}
		if tile.leftTile == nil {
			emptySides++
		}
		if tile.rightTile == nil {
			emptySides++
		}
	}
	wantEmptySides := int(math.Sqrt(float64(len(tiles)))) * 4
	if emptySides != wantEmptySides {
		log.Fatalf("expected %d empty sides, got %d\n", wantEmptySides, emptySides)
	}
}
