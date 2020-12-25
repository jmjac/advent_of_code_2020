with open("input.txt", "r") as file:
    tiles_to_flip = [i for i in file.readlines()]

movement = {"w": (-1, 0),
            "e": (1, 0),
            "S": (1, -1),
            "W": (0, -1),
            "E": (0, 1),
            "N": (-1, 1)}


black_tiles = set()
for tile in tiles_to_flip:
    tile = tile.replace("se", "S").replace("sw", "W").replace(
        "nw", "N").replace("ne", "E").replace("\n", "")
    x,y = 0,0
    for i in tile:
        dx,dy = movement[i]
        x+=dx
        y+=dy
    if (x,y) not in black_tiles:
        black_tiles.add((x,y))
    else:
        black_tiles.remove((x,y))

def part2(black_tiles):
    for _ in range(100):
        black_neighbours = {}
        for tile in black_tiles:
            x,y = tile
            if (x,y) not in black_neighbours:
                black_neighbours[(x,y)] = 0
            for m in movement.values():
                dx, dy = m
                dx = dx+x
                dy = dy+y
                if (dx, dy) not in black_neighbours:
                        black_neighbours[(dx, dy)] = 0
                black_neighbours[(dx, dy)]+=1
        new_black = set()
        for k,v in black_neighbours.items():
            x,y = k
            if v==2 and (x, y) not in black_tiles:
                new_black.add((x,y))
            elif (x,y) in black_tiles:
                if v>0 and v<=2:
                    new_black.add((x,y))
        black_tiles = new_black
    return len(black_tiles)

print(f"Part 1: {len(black_tiles)}")
print(f"Part 2: {part2(black_tiles)}")