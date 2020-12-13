from math import cos, sin, radians

with open("input.txt", 'r') as f:
    actions = [i for i in f.readlines()]

# change in (x,y,angle)
movement = {
    "N" : lambda x,angle: (0,x,0),
    "S" : lambda x,angle: (0,-x,0),
    "W" : lambda x,angle: (-x,0,0),
    "E" : lambda x,angle: (x,0,0),
    "L" : lambda x,angle: (0,0,x),
    "R" : lambda x,angle: (0,0,-x),
    "F" : lambda x,angle: (int(x*sin(angle)), int(x*cos(angle)),0)
}

def part1(actions):
    x,y,angle = 0,0,0
    for i in actions:
        dx,dy,dangle = movement[i[0]](int(i[1:]),angle)
        x += dx
        y += dy
        angle += radians(dangle)
    return abs(x)+abs(y)

def part2(actions):
    w_x,w_y = 10,1
    s_x, s_y = 0,0
    for i in actions:
        a = i[0]
        v = int(i[1:])
        if a == "F":
            s_x += w_x*v
            s_y += w_y*v
        else:
            dx,dy,angle = movement[a](v,0)
            if angle!=0:
                turn_angle = radians(angle)
                w_x, w_y = w_x*int(cos(turn_angle))-w_y*int(sin(turn_angle)), w_x*int(sin(turn_angle))+w_y*int(cos(turn_angle))
            else:
                w_x += dx
                w_y += dy
    return abs(s_x)+abs(s_y)

print(part1(actions))
print(part2(actions))
