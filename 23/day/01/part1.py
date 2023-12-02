import re
f = open('input.txt')

total = 0
for line in f:
    firstAndLastDigitInLineRegex = r'(\d)'
    firstAndLastDigitInLine = re.findall(r'(\d)', line)
    total += int(firstAndLastDigitInLine[0] + firstAndLastDigitInLine[-1])
    print(firstAndLastDigitInLine)

print(total)
