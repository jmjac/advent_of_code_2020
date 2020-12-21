count = {}
allergens = set() 
with open("input.txt", "r") as f:
    for i in f.readlines():
        foods, allerg = i.split("(contains ")
        allerg = allerg.replace(")","").replace("\n","")
        for food in foods.split():
            if food not in count:
                    count[food] = {"dishes": 0}
            count[food]["dishes"] +=1
            for a in allerg.split(", "):
                if a not in count[food]:
                    count[food][a] = 0
                count[food][a]+=1
                allergens.add(a)


def find_allergens(count, allergens):
    most_likely = {a:(0,set()) for a in allergens}
    for food in count:
        for a in count[food]:
            if a == "dishes":
                continue
            if count[food][a] > most_likely[a][0]:
                most_likely[a] = (count[food][a], {food})
            elif count[food][a] == most_likely[a][0]:
                most_likely[a][1].add(food)

    found_allergens = set()
    for i in most_likely.values():
        for j in i[1]:
            found_allergens.add(j)
    return found_allergens, most_likely

def count_non_allergens(count, found_allergens):
    total = 0
    for i in count:
        if i not in found_allergens:
            total+=count[i]["dishes"]
    return total

def sort_allergens( most_likely):
    food_to_allergen = {}
    allergen_possible = {k:v[1] for k,v in most_likely.items()}
    while len(allergen_possible)!=len(food_to_allergen):
        for a in allergen_possible:
            allergen_possible[a] -= set(food_to_allergen)
            if len(allergen_possible[a]) == 1:
                food_to_allergen[allergen_possible[a].pop()] = a
    return ",".join([i for i in sorted(food_to_allergen, key=food_to_allergen.get)])

found_allergens, most_likely = find_allergens(count, allergens)
print(f"Part 1: {count_non_allergens(count, found_allergens)}")
print(f"Part 2 {sort_allergens(most_likely)}")