data = []
with open("input.txt", "r") as file:
    for line in file.readlines():
        data.append(int(line))

def problem1(data:list) -> int:
    data_set = set(data)
    for i in data:
        if 2020-i in data_set:
            return (2020-i)*i


def problem2(data:list) -> int:
    data_set = set(data)
    for i in range(len(data)):
        for j in range(i, len(data)):
            if 2020 -data[i] - data[j] in data_set:
                return (2020-data[i]-data[j])*data[i]*data[j]

    
print(problem1(data))
print(problem2(data))