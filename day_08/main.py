program = []

with open("input.txt", "r") as file:
    for line in file.readlines():
        program.append(line.split())


def execute(op, arg, ic, acc):
    # return ic change, acc change
    if op == "nop":
        return ic+1, acc
    if op == "jmp":
        return ic+int(arg), acc
    if op == "acc":
        return ic+1, acc+int(arg)
    raise ValueError(f"No operation: {op}")


def find_loop(program):
    ic = 0
    acc = 0
    executed = set()
    while ic < len(program):
        executed.add(ic)
        op, arg = program[ic]
        new_ic, new_acc = execute(op, arg, ic, acc)
        if new_ic in executed:
            #Found loop
            return (ic, acc)
        ic, acc = new_ic, new_acc
    return (ic, acc)

def part2(program):
    for i in range(len(program)):
        original = program[i][0]
        if program[i][0] == "jmp":
            program[i][0] = "nop"
        elif program[i][0] == "nop":
            program[i][0] == "jmp"
        else:
            continue

        ic, acc = find_loop(program)
        if ic>=len(program):
            return acc
        else:
            program[i][0] = original

print(f"Part 1: {find_loop(program)[1]}")
print(f"Part 2: {part2(program)}")
