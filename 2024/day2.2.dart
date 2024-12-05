import 'util.dart';

void main() {
  var fileContent = readFile('./files/day2.txt');
  var lines = fileContent.split('\n');
  var total = 0;

  for (final line in lines) {
    var safe = true;
    var positive = false;

    var items = line.split(' ');
    var removedIndex = 0;

    var modifiedArray = List.from(items);
    void modifyList() {
      var copy = List.from(items);
      copy.removeAt(removedIndex++);
      modifiedArray = copy;
    }

    var i = 0;
    while (i < modifiedArray.length - 1) {
      var answer =
          int.parse(modifiedArray[i]) - int.parse(modifiedArray[i + 1]);

      if (i == 0) {
        positive = answer > 0 || answer == 0;
      } else if (positive && answer < 0) {
        if (removedIndex == items.length) {
          safe = false;
          break;
        }
        modifyList();
        i = 0;
        continue;
      } else if (!positive && answer >= 0) {
        if (removedIndex == items.length) {
          safe = false;
          break;
        }
        modifyList();
        i = 0;
        continue;
      } else if (positive && answer == 0) {
        if (removedIndex == items.length) {
          safe = false;
          break;
        }
        modifyList();
        i = 0;
        continue;
      }

      if (positive && (answer < 1 || answer > 3)) {
        if (removedIndex == items.length) {
          safe = false;
          break;
        }
        modifyList();
        i = 0;
        continue;
      }

      if (!positive && (answer.abs() < 1 || answer.abs() > 3)) {
        if (removedIndex == items.length) {
          safe = false;
          break;
        }
        modifyList();
        i = 0;
        continue;
      }

      i++;
    }
    if (safe) {
      total++;
      print(line);
    }
  }

  print(total);
}
