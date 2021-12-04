namespace AllDays;
public class OnePartTwo
{
    public static int GetDifferences()
    {
        string text = System.IO.File.ReadAllText("files/day1-1.txt");
        var lines = text.Split('\n');

        int count = 0;
        int sum = 0;
        List<int> sums = new List<int>();

        for (int i = 0; i < lines.Length - 2; i++)
        {
            sum += int.Parse(lines[i]);
            sum += int.Parse(lines[i + 1]);
            sum += int.Parse(lines[i + 2]);

            sums.Add(sum);
            sum = 0;
        }

        for (int i = 1; i < sums.Count; i++)
        {
            int a = sums[i - 1];
            int b = sums[i];
            if ((b - a) > 0)
            {
                count++;
            }
        }

        return count;
    }
}
