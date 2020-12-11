with open("input.txt", 'r') as f:
    grid = [[i.split("\n")[0] for i in row] for row in f.readlines()]


def simulate(grid,treeshold, count_function):
    updated = True
    while updated:
        new_grid = []
        updated = False
        for y in range(len(grid)):
            new_grid.append([])
            for x in range(len(grid[y])):
                new_grid[-1].append(grid[y][x])
                if grid == ".":
                    new_grid[-1][-1] = "."
                    continue
                occupied = count_function(x, y, grid)
                if occupied == 0 and grid[y][x] == "L":
                    updated = True
                    new_grid[-1][-1] = "#"
                elif occupied >= treeshold and grid[y][x] == "#":
                    updated = True
                    new_grid[-1][-1] = "L"
        grid = new_grid
    return sum([sum([1 for i in row if i=="#"]) for row in grid])

def count(x, y, grid):
    occupied = 0
    for i in range(max(0, y-1), min(y+2, len(grid))):
        for j in range(max(0, x-1), min(len(grid[i]), x+2)):
            if i==y and x==j:
                continue
            if grid[i][j] == "#":
                occupied += 1
    return occupied

def count_in_sight(x,y, grid):
    occupied = 0
    movement = {(1,0), (0,1), (-1,0),(-1,-1),(0,-1), (1,1),(-1,1),(1,-1)}
    for dx,dy in movement:
        ddx,ddy = x+dx,y+dy
        while True:
            #Faster then checking first
            try:
                if grid[ddy][ddx] == "#":
                    occupied += 1
                    break
                elif grid[ddy][ddx] == "L":
                    break
            except:
                break
            ddx, ddy = ddx+dx, ddy+dy
    return occupied


def print_grid(grid):
    for i in grid:
        print("".join(i))

print(f"Part 1:{simulate(grid, 4, count)}")
print(f"Part 2:{simulate(grid, 5, count_in_sight)}")