with open('input.txt') as f:
    adapters = [int(i) for i in f.readlines()]


def chain_differences(chain):
    diff = {1:0, 2: 0, 3:0}
    for i in range(1, len(chain)):
        diff[chain[i] - chain[i-1]]+=1
    return diff[3]*diff[1]

#Original slow solution
def part2old(adapters):
    start = 0
    total = 1
    for i in range(len(adapters)):
        if adapters[i] - adapters[i-1]==3:
            total *= all_possible_arrangments([adapters[start]], adapters[start+1:i])
            start = i
    return total

def all_possible_arrangments(chain, adapters):
    if len(adapters)==0:
        return 1
    if adapters[0] - chain[-1] >3:
        return 0
    chains = []
    i = 0
    while i<len(adapters) and adapters[i] - chain[-1]<=3:
        chains.append((chain + [adapters[i]],i))
        i+=1
    return sum([all_possible_arrangments(new_chain, adapters[i+1:]) for new_chain,i in chains])

def part2fast(i, limit, adapters, mem):
    if i == limit:
        return 1
    res = 0
    for diff in [1,2,3]:
        if i+diff in mem:
            res += mem[i+diff]
        elif i+diff in adapters:
            mem[i+diff] = part2fast(i+diff,limit, adapters, mem)
            res += mem[i+diff]
    return res


adapters.append(0)
adapters.sort()
adapters.append(max(adapters)+3)

print(f"Part 1:{chain_differences(adapters)}")
print(f"Part 2A:{part2fast(0, adapters[-1], set(adapters), {})}")