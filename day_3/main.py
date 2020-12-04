data = []
with open("input.txt", "r") as file:
    for line in file.readlines():
        line = line.replace("\n", "")
        data.append(line)

def countTrees(data, step_right, step_down):
    check_index = step_right
    count = 0
    for i in range(step_down, len(data), step_down):
        if data[i][check_index%len(data[i])] == "#":
            count+=1
        check_index += step_right
    return count

def problem1(data):
    return countTrees(data, 3, 1)

def problem2(data):
    total = 1
    for step in [1,3,5,7]:
        total *= countTrees(data,step, 1)
    total *= countTrees(data, 1, 2)
    return total


print(f"Problem 1: {problem1(data)}")
print(f"Problem 2: {problem2(data)}")