import queue
bags = {}
with open("input.txt", "r") as file:
    for line in file.readlines():
        line = line.replace("\n", "").replace(".","").replace("bags", "").replace("bag","")
        bag, content = line.split("  contain ")
        bags[bag] = {}
        for i in content.split(","):
            b = " ".join(i.split()[1:])
            if b == "other":
                continue 
            q = int(i.split()[0])
            bags[bag][b] = q

def valid_outmost(bags):
    containt_gold = set()
    updated = True
    for bag, contents in bags.items():
        if "shiny gold" in contents:
            containt_gold.add(bag)

    while updated:
        updated = False
        for bag, content in bags.items():
            if bag in containt_gold:
                continue
            for i in containt_gold.copy():
                if i in content:
                    containt_gold.add(bag)
                    updated = True
    return containt_gold

def shiny_capacity(bags):
    capacity = 0
    q = queue.Queue()
    q.put(("shiny gold",1))
    while not q.empty():
        top_bag, mult = q.get()
        for bag, cap in bags[top_bag].items():
            capacity += cap*mult
            for b,c in bags[bag].items():
                capacity += mult*cap*c
                q.put((b,mult*cap*c))
    return capacity

print(f"Part 1: {len(valid_outmost(bags))}")
print(f"Part 2: {shiny_capacity(bags)}")