import re
f = open('input.txt')

stringToNumberMap = {
    'zero': 0,
    'one': 1,
    'two': 2,
    'three': 3,
    'four' : 4,
    'five' : 5,
    'six' : 6,
    'seven' : 7,
    'eight' : 8,
    'nine' : 9
}

stringToNumberMatch = {
    'z': {
        'ero': '0'
    },
    'o': {
        'ne': '1'
    },
    't': {
        'wo': '2',
        'hree': '3'
    },
    'f': {
        'our': '4',
        'ive': '5'
    },
    's': {
        'ix': '6',
        'even': '7'
    },
    'e': {
        'ight': '8',
    },
    'n': {
        'ine': '9'
    }
}

# iterate over the characters in the string
# and convert them to numbers in sequence
def convertStringNumbersToNumbersInLine(line):
    stack = []
    i = 0
    max = len(line)
    while i < max:
        stack.append(line[i])
        char = line[i]
        if char in stringToNumberMatch:
            for key in stringToNumberMatch[char]:
                if line[i+1:i+1+len(key)] == key:
                    stack.pop()
                    stack.append(stringToNumberMatch[char][key])
                    i += len(key) -1
        i += 1
    return stack

total = 0
for line in f:
    print(line)
    line = ''.join(convertStringNumbersToNumbersInLine(line))
    firstAndLastDigitInLineRegex = r'(\d)'
    firstAndLastDigitInLine = re.findall(r'(\d)', line)
    print(firstAndLastDigitInLine)
    print(firstAndLastDigitInLine[0] + firstAndLastDigitInLine[-1])
    total += int(firstAndLastDigitInLine[0] + firstAndLastDigitInLine[-1])

print(total)
