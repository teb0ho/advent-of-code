import sys

def get_lines():
    lines = []
    with open('input.txt', 'r') as reader:
        line = reader.readline()
        while line != '':
            lines.append(line.strip())
            line = reader.readline()

    return lines


sample = get_lines()
results = {}
results_final = []
word_num = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9
}



for line in sample:
    numbers = []
    for idx,character in enumerate(line):
        if character.isnumeric():
            if results.get(int(character)) != None:
                results[int(character)].append(idx)
            else:
                results[int(character)] = [idx]
    for key, value in word_num.items():
        index = line.find(key)
        if index != -1:
            if results.get(value) != None:
                results[value].append(index)
            else: 
                results[value] = [index]
    min = sys.maxsize
    min_key = 0
    max = 0
    max_key = 0
    for key, value in results.items(): 
        if value != None:
            for i in value:
                if min > i:
                    min = i
                    min_key = key
                if max < i:
                    max = i
                    max_key = key

                results_final.append(int(f'{results.get(min_key)}{results.get(max_key)}'))

print(sum(results))


