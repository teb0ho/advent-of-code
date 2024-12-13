import 'dart:collection';

import 'util.dart';

bool horizontal(int columnIndex, int rowIndex, List<String> lines, int columns,
    String xmas, String samx) {
  if (!(columns - columnIndex > 3)) return false;

  String substring = lines[rowIndex].substring(columnIndex, columnIndex + 3);
  return xmas == substring || samx == substring;
}

bool vertical(int columnIndex, int rowIndex, List<String> lines, int columns,
    int rows, String xmas, String samx) {
  var result = false;
  if ((rowIndex + 3 < rows)) {}

  String substring = lines[rowIndex].substring(columnIndex, columnIndex + 3);
  return xmas == substring || samx == substring;
}

void main() {
  var fileContent = readFile('./files/day4.txt');

  var lines = fileContent.split('\n');

  final HashSet<String> words = HashSet();

  var count = 0;
  var xmas = 'XMAS';
  var samx = 'SAMX';

  var columns = lines[0].length;
  var rows = lines.length;

  for (var i = 0; i < lines.length; i++) {
    for (var j = 0; j < columns; j++) {
      if (horizontal(j, i, lines, columns, xmas, samx))
        ("${j},${i}-horizontal");
    }
  }
}
