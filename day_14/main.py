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
        elif mask[i] == "1":
            for j in range(len(adresses)):
                adresses[j] += "1"
        else:
            for j in range(len(adresses)):
                if i >= len(mask) - len(adress):
                    adresses[j] += adress[-(len(mask) - (i))]
                else:

                    adresses[j] += "0"

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
            masks = ["", ""]
            for i in value:
                if i == "X":
                    masks[0] += "0"
                    masks[1] += "1"
                else:
                    masks[0] += i
                    masks[1] += i
            masks[0] = int(masks[0], 2)
            masks[1] = int(masks[1], 2)
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
