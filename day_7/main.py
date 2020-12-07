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
print(bags)
containt_gold = set()
updated = True
for bag, contents in bags.items():
    if "shiny gold" in contents:
        containt_gold.add(bag)


while updated:
    updated = False
    for bag, content in bags.items():
        for i in containt_gold.copy():
            if i in content and bag not in containt_gold:
                containt_gold.add(bag)
                updated = True
print(f"Part 1: {len(containt_gold)}")

