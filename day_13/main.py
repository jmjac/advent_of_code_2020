with open("input.txt", "r") as f:
    depart_time = int(f.readline())
    buses = [i for i in f.readline().split(',')]


def part1(depart_time, buses):
    valid_buses = [int(i) for i in buses if i != 'x']
    times_ids = [(int(depart_time/i+1)*i - depart_time, i)
                 for i in valid_buses]
    time, bus_id = min(times_ids)
    return time*bus_id


def part2(buses):
    # Using Chinese Remainder Theorem
    N = 1
    for i in buses:
        if i != 'x':
            N *= int(i)

    ans = 0
    for i in range(len(buses)):
        if buses[i] != 'x':
            x = 0
            bus_id = int(buses[i])
            while ((N*x//bus_id)) % bus_id != 1:
                x += 1
            ans += (bus_id - i)*x*N//bus_id
    return ans % N


print(part1(depart_time, buses))
print(part2(buses))
