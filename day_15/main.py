import time
with open("input.txt", "r") as f:
    numbers = [int(i) for i in f.readline().split(",")]


def part1(numbers, nth):
    mem = {}
    turn = 1
    num = 0
    for i in numbers:
        if i in mem:
            num = turn - mem[i]
        else:
            num = 0
        mem[i] = turn
        turn += 1

    while turn != nth:
        if num in mem:
            next_num = turn - mem[num]
        else:
            next_num = 0
        mem[num] = turn
        num = next_num
        turn += 1
    return num


print(f"Part 1: {part1(numbers, 2020)}")
print(f"Part 2: {part1(numbers, 30000000)}")