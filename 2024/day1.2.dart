import 'util.dart';

void main() {
  var fileContent = readFile('./files/day1.txt');
  var lines = fileContent.split('\n');
  var total = 0;

  for (final line in lines) {
    var count = 0;
    var items = line.split('   ');
    var left = int.parse(items[0]);

    for (var i = 0; i < lines.length; i++) {
      if (left == int.parse(lines[i].split('   ')[1])) count++;
    }
    total += left * count;
  }

  print(total);
}
