import 'dart:io';

void main() {
  var file = File('./files/day1.txt');
  try {
    var fileContent = file.readAsStringSync();

    List<String> lines = fileContent.split('\n');
    List<int> left = [];
    List<int> right = [];
    List<int> results = [];

    for (final line in lines) {
      var items = line.split('   ');
      left.add(int.parse(items[0]));
      right.add(int.parse(items[1]));
    }

    left.sort();
    right.sort();

    while (left.length != 0) {
      var result = left[0] - right[0];
      results.add(result.abs());
      left.removeAt(0);
      right.removeAt(0);
    }

    print(results.reduce((a, b) => a + b));
  } catch (e) {
    print('error reading file: $e');
  }
}
