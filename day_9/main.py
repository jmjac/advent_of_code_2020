import queue

with open("input.txt", "r") as f:
    data = [int(i) for i in f.readlines()]


def find_invalid(data, keep):
    q = queue.Queue()
    mem = set()
    for i in range(keep):
        q.put(data[i])
        mem.add(data[i])

    for k in range(keep, len(data)):
        check = data[k]
        for i in mem:
            if check-i in mem:
                break
        else:
            return check
        mem.remove(q.get())
        q.put(check)
        mem.add(check)
    raise ValueError("No invalid number")

def sum_to_invalid(data, invalid):
    start = 0
    running_sum = 0
    for i in range(len(data)):
        running_sum+=data[i]
        
        while running_sum>invalid:
            running_sum -= data[start]
            start+=1

        if running_sum==invalid:
            return data[start]+data[i]
            
    raise ValueError("Sum not found")

invalid = find_invalid(data, 25)
print(f"Part 1: {invalid}")
print(f"Part 2: {sum_to_invalid(data, invalid)}")