namespace AllDays
{
    public class One
    {
        public static int GetDifferences() 
        {
            string text = System.IO.File.ReadAllText("day1-1.txt");
            var lines = text.Split('\n');

            int count = 0;

            for (int i = 1; i < lines.Length; i++)
            {
                int a = int.Parse(lines[i]);
                int b = int.Parse(lines[i - 1]);

                if ((a - b) > 0) 
                {
                    count++;
                } 
            }
            return count;
        }
    }
}