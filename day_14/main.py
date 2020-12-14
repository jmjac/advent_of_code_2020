def write(adress, value, mem, masks):
    mem[adress] = masks[0] | (value & masks[1])


def find_adresses(adress, mask):
    adress = bin(int(adress))[2:]
    adresses = [""]
    for i in range(len(mask)):
        if mask[i] == "X":
            new_adresses = []
            for j in range(len(adresses)):
                new_adresses.append(adresses[j] + "1")
                new_adresses.append(adresses[j] + "0")
            adresses = new_adresses
            continue

        if mask[i] == "1":
            next_bit = "1"
        else:
            if i >= len(mask) - len(adress):
                next_bit = adress[-(len(mask) - (i))]
            else:
                next_bit = "0"

        for j in range(len(adresses)):
            adresses[j] += next_bit

    for j in range(len(adresses)):
        adresses[j] = int(adresses[j], 2)
    return adresses


mem_part1 = {}
mem_part2 = {}
with open("input.txt", "r") as f:
    for i in f.readlines():
        instr, value = i.strip("\n").split(" = ")
        if instr == "mask":
            mask = value
            masks = [int(value.replace("X", "0"), 2),
                     int(value.replace("X", "1"), 2)]
        else:
            value = int(value)
            adress = instr.split("[")[1][:-1]
            # Part 1
            write(adress, value, mem_part1, masks)
            # Part 2
            for a in find_adresses(adress, mask):
                mem_part2[a] = value

print(f"Part 1: {sum(mem_part1.values())}")
print(f"Part 2: {sum(mem_part2.values())}")
