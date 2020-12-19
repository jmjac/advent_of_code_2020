with open("input.txt", "r") as f:
    rules_data, messages_data = f.read().split("\n\n")
    rules = {}
    basic_rules = set()
    for rule in rules_data.split("\n"):
        n, r = rule.split(": ")
        n = int(n)
        rules[n] = []
        for i in r.split(" | "):
            if '"' in i:
                rules[n].append(i.replace('"', "").split()[0])
                basic_rules.add(n)
            else:
                rules[n].append([int(j) for j in i.split()])

    messages = messages_data.split("\n")


def check_messages(messages, rules, depth):
    total = 0
    #The rule 0 is build from 8 and 11, they both are build from 42 and 31 so we just generate those
    valid_31 = set()
    valid_42 = set()
    for rule_set in convert_rule(31, rules):
        for rule in rule_set:
            valid_31.add(rule)

    for rule_set in convert_rule(42, rules):
        for rule in rule_set:
            valid_42.add(rule)

    num_to_valid = {31: valid_31, 42: valid_42}
    comb = []
    temp = []
    # Generate the possible combinations without storing them whole
    for i in range(1, depth+1):
        temp = [42] + temp + [31]
        for j in range(1, depth+1):
            comb.append([42]*j+temp)

    for message in messages:
        for valid in comb:
            start = 0
            for i in valid:
                len_valid = len(next(iter(num_to_valid[i])))
                if message[start:start+len_valid] not in num_to_valid[i]:
                    break
                start += len_valid
            else:
                if start == len(message):
                    total += 1
                    break

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

print(f"Part 1: {check_messages(messages, rules,1)}")
#Increase the depth for longer inputs
print(f"Part 2: {check_messages(messages, rules,10)}")