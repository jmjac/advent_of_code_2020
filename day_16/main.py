def part1(req, tickets):
    all_req = set()
    for i in req.values():
        all_req |= i
    invalid = set()
    total = 0
    for i in range(len(tickets)):
        ticket = tickets[i]
        for j in ticket:
            if j not in all_req:
                total += j
                invalid.add(i)
                break
    return total, invalid


def part2(req, tickets, my_ticket, invalid):
    tickets = [tickets[i] for i in range(len(tickets)) if i not in invalid]
    tickets.append(my_ticket)
    field_index_map = {}

    for i in req:
        field_index_map[i] = set()
        for index in range(len(req)):
            for ticket in tickets:
                if ticket[index] not in req[i]:
                    break
            else:
                field_index_map[i].add(index)
    total = 1
    assigned = set()
    for k, _ in sorted(field_index_map.items(), key=lambda x: len(x[1])):
        field_index_map[k] -= assigned
        field_index_map[k] = field_index_map[k].pop()
        assigned.add(field_index_map[k])
    for i in field_index_map:
        if "departure" in i:
            total *= my_ticket[field_index_map[i]]
    return total


req = {}
my_ticket = []
tickets = []

with open('input.txt', 'r') as f:
    data = req
    for line in f.readlines():
        line = line.strip("\n")
        if line == "":
            continue
        if line == "your ticket:":
            data = my_ticket
            continue
        elif line == "nearby tickets:":
            data = tickets
            continue

        if data is req:
            name, values = line.split(":")
            left, right = values.split(" or ")
            l_left, h_left = left.split("-")
            l_right, h_right = right.split("-")
            req[name] = set()
            for j in range(int(l_left), int(h_left)+1):
                req[name].add(j)
            for j in range(int(l_right), int(h_right)+1):
                req[name].add(j)
        else:
            data.append([int(j) for j in line.split(",")])
    my_ticket = my_ticket[0]


total_invalid, invalid = part1(req, tickets)
print(f"Part 1: {total_invalid}")
print(f"Part 2: {part2(req, tickets, my_ticket, invalid)}")
