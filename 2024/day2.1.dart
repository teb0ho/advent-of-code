import 'util.dart';

void main() {
  var fileContent = readFile('./files/day2.txt');
  var lines = fileContent.split('\n');
  var total = 0;

  for (final line in lines) {
    var safe = true;
    var positive = false;

    var items = line.split(' ');

    for (var i = 0; i < items.length - 1; i++) {
      var answer = int.parse(items[i]) - int.parse(items[i + 1]);

      if (i == 0) {
        positive = answer > 0 || answer == 0;
      } else if (positive && answer < 0) {
        safe = false;
        break;
      } else if (!positive && answer >= 0) {
        safe = false;
        break;
      }

      if (positive && (answer < 1 || answer > 3)) {
        safe = false;
        break;
      }

      if (!positive && (answer.abs() < 1 || answer.abs() > 3)) {
        safe = false;
        break;
      }
    }
    if (safe) total++;
  }

  print(total);
}
