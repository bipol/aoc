import re
import functools 

f = open('input.txt')

powers = []
validGames = []
for line in f:
    [game, selections] = line.split(':')
    invalidGame = False
    trials = selections.split(';')
    # red, green, blue
    minCube = [0, 0, 0]
    for trial in trials:
        picks = trial.split(',')
        for pick in picks:
            [count, color] = pick.strip().split(' ')
            if color == 'red':
                if int(count) > minCube[0]:
                    minCube[0] = int(count)
            if color == 'green':
                if int(count) > minCube[1]:
                    minCube[1] = int(count)
            if color == 'blue':
                if int(count) > minCube[2]:
                    minCube[2] = int(count)
    powers.append(functools.reduce(lambda x, b: x * b, minCube, 1))

total = functools.reduce(lambda x, b: int(b) + x, powers, 0)
print(total)
