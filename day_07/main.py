import queue


bags = {}
with open("input.txt", "r") as file:
    for line in file.readlines():
        line = line.replace(".","").replace("bags", "").replace("bag","")
        bag, content = line.split("  contain ")
        bags[bag] = {}
        for i in content.split(","):
            b = " ".join(i.split()[1:])
            if b == "other":
                continue 
            q = int(i.split()[0])
            bags[bag][b] = q

def valid_outmost(valid_bag, bags):
    containt_bag = set()
    len_before = 0
    for bag, contents in bags.items():
        if valid_bag in contents:
            containt_bag.add(bag)

    while len_before!=len(containt_bag):
        len_before = len(containt_bag)
        for bag, content in bags.items():
            if bag in containt_bag:
                continue
            for i in containt_bag.copy():
                if i in content:
                    containt_bag.add(bag)
    return containt_bag

def bag_capacity(bag_name, bags):
    capacity = 0
    q = queue.Queue()
    q.put((bag_name,1))
    while not q.empty():
        top_bag, mult = q.get()
        for bag, cap in bags[top_bag].items():
            capacity += cap*mult
            for b,c in bags[bag].items():
                capacity += mult*cap*c
                q.put((b,mult*cap*c))
    return capacity

print(f"Part 1: {len(valid_outmost('shiny gold', bags))}")
print(f"Part 2: {bag_capacity('shiny gold',bags)}")