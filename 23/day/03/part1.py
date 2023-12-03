import re
import functools

f = open('input.txt')

symbolIndex = []
possibleParts = {}
parts = []

i = 0
lineLength = 0
for line in f:
    partGroups = re.finditer(r'(\d+)', line)
    for partGroup in partGroups:
        if i not in possibleParts:
            possibleParts[i] = []
        possibleParts[i].append( [ partGroup.span(), partGroup.group() ])
    charIndex = 0
    for char in line:
        if char != '.' and char != '\n' and not char.isdigit():
            symbolIndex.append([i, charIndex])
        charIndex += 1
    i += 1

debugPartList = []
# given a symbol, like 1, 3, what is adjacent to it?
# 0,2 - 0, 3 - 0, 4
# 1,2 - 1, 4
# 2,2 - 2, 3 - 2, 4
for symbol in symbolIndex:
    possibleAdjRows = range(symbol[0] - 1, symbol[0] + 2)
    possibleAdjCols = range(symbol[1] - 1, symbol[1] + 2)
    possibleAdjIndexes = zip(possibleAdjRows, possibleAdjCols)
    for possibleAdjIndex in possibleAdjIndexes:
        if possibleAdjIndex[0] in possibleParts:
            # this is the row, where we have a list of parts
            # part = [[(0,3), '467]]
            # so we need to determine if 0,3 is in the range of possibleAdjIndex
            # if single digit, it'd be like (0,1)
            # our symbol is 1,3
            notFound = []
            for part in possibleParts[possibleAdjIndex[0]]:
                intersection = set(possibleAdjCols).intersection(range(part[0][0], part[0][1]))
                #print("part:", part)
                #print("symbol:", symbol)
                #print("possibleAdjIndex:", possibleAdjIndex)
                #print("intersection:", intersection)
                if len(intersection) > 0:
                    parts.append(part[1])
                    debugPartList.append([symbol, part[1]])
                else:
                    notFound.append(part)
            possibleParts[possibleAdjIndex[0]] = notFound
for part in debugPartList:
    print(part)

print(functools.reduce(lambda x, y: x + int(y), parts, 0))

# 316874 is invalid
# 519504 is too low
# 517137 is too low
