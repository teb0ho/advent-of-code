import 'util.dart';

void main() {
  var fileContent = readFile('./files/day3.txt')
      .replaceAll(RegExp(r"don't\(\).*?do\(\)", dotAll: true), ' ')
      .replaceAll(RegExp(r"don't\(\).*"), ' ');

  RegExp regex = RegExp(r'mul\((\d+),(\d+)\)');

  var matches = regex.allMatches(fileContent);
  var total = 0;

  for (final match in matches) {
    total += int.parse(match.group(1) ?? '') * int.parse(match.group(2) ?? '');
  }

  print(total);
}
