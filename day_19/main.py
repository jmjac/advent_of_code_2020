with open("input.txt", "r") as f:
    rules_data, messages_data = f.read().split("\n\n")
    rules = {}
    basic_rules = set()
    for rule in rules_data.split("\n"):
        n,r = rule.split(": ")
        n = int(n)
        rules[n] = []
        for i in r.split(" | "):
            if '"' in i:
                rules[n].append(i.replace('"',"").split()[0])
                basic_rules.add(n)
            else:
                rules[n].append([int(j) for j in i.split()])

    messages = messages_data.split("\n")



def part1(messages, rules):
    total = 0
    valid = set([i for i in convert_rule(0, rules)[0]])
    for message in messages:
        if message in valid:
            total+=1
    return total


def convert_rule(rule, rules):
    converted_rules = [] 
    for rule_set in rules[rule]:
        converted = [""]
        for r in rule_set:
            if r in rules:
                new_converted = []
                for k in convert_rule(r, rules):
                    for j in k:
                        for base in converted:
                            new_converted.append(base + j)
                converted = new_converted
            else:
                for i in range(len(converted)):
                    converted[i] += r
        converted_rules.append(converted)
    return converted_rules

print(f"Part 1: {part1(messages, rules)}")