
def get_lines():
    lines = []
    with open('input.txt', 'r') as reader:
        line = reader.readline()
        while line != '':
            lines.append(line.strip())
            line = reader.readline()

    return lines


sample = get_lines()
results = []

for line in sample:
    numbers = []
    for character in line:
        if character.isnumeric():
            numbers.append(int(character))
    if len(numbers) > 1:
        results.append(int(f'{numbers[0]}{numbers[len(numbers) -1]}'))
    else:
        results.append(int(f'{numbers[0]}{numbers[0]}'))

print(sum(results))


