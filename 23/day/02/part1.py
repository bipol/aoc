import re
import functools 

f = open('input.txt')

whatsInTheBag = {
    'red': 12,
    'green': 13,
    'blue': 14
}

validGames = []
for line in f:
    [game, selections] = line.split(':')
    invalidGame = False
    trials = selections.split(';')
    for trial in trials:
        picks = trial.split(',')
        for pick in picks:
            [count, color] = pick.strip().split(' ')
            if whatsInTheBag[color] < int(count):
                invalidGame = True
                break;
        if invalidGame:
            break;
    if not invalidGame:
        validGames.append(game)


stringNums = map(lambda x: re.findall(r'(\d+)', x)[0], validGames)
total = functools.reduce(lambda x, b: int(b) + x, stringNums, 0)
print(total)
