import 'dart:collection';

import 'util.dart';

bool horizontal(int columnIndex, int rowIndex, List<String> lines, int columns,
    String xmas, String samx) {
  if (!((columnIndex + 3) < columns)) return false;

  String substring = lines[rowIndex].substring(columnIndex, columnIndex + 4);
  return xmas == substring || samx == substring;
}

bool vertical(int columnIndex, int rowIndex, List<String> lines, int columns,
    int rows, String xmas, String samx) {
  if (!(rowIndex + 3 < rows)) return false;

  var substring = "";
  substring += lines[rowIndex][columnIndex];
  substring += lines[rowIndex + 1][columnIndex];
  substring += lines[rowIndex + 2][columnIndex];
  substring += lines[rowIndex + 3][columnIndex];

  return xmas == substring || samx == substring;
}

bool crossRight(int columnIndex, int rowIndex, List<String> lines, int columns,
    int rows, String xmas, String samx) {
  if (!(rowIndex + 3 < rows && columnIndex + 3 < columns)) return false;

  var substring = "";
  substring += lines[rowIndex][columnIndex];
  substring += lines[rowIndex + 1][columnIndex + 1];
  substring += lines[rowIndex + 2][columnIndex + 2];
  substring += lines[rowIndex + 3][columnIndex + 3];

  return xmas == substring || samx == substring;
}

bool crossLeft(int columnIndex, int rowIndex, List<String> lines, int columns,
    int rows, String xmas, String samx) {
  if (!(rowIndex + 3 < rows && columnIndex - 3 > -1)) return false;

  var substring = "";
  substring += lines[rowIndex][columnIndex];
  substring += lines[rowIndex + 1][columnIndex - 1];
  substring += lines[rowIndex + 2][columnIndex - 2];
  substring += lines[rowIndex + 3][columnIndex - 3];

  return xmas == substring || samx == substring;
}

void main() {
  var fileContent = readFile('./files/day4.txt');

  var lines = fileContent.split('\n');

  final HashSet<String> words = HashSet();

  var xmas = 'XMAS';
  var samx = 'SAMX';

  var columns = lines[0].length;
  var rows = lines.length;

  for (var i = 0; i < lines.length; i++) {
    for (var j = 0; j < columns; j++) {
      if (horizontal(j, i, lines, columns, xmas, samx))
        words.add("$i,$j-horizontal");
      if (vertical(j, i, lines, columns, rows, xmas, samx))
        words.add("$i,$j-vertical");
      if (crossLeft(j, i, lines, columns, rows, xmas, samx))
        words.add("$i,$j-crossLeft");
      if (crossRight(j, i, lines, columns, rows, xmas, samx))
        words.add("$i,$j-crossRight");
    }
  }

  print(words.length);
}
