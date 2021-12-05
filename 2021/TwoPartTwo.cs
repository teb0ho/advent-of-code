namespace AllDays;

public class TwoPartTwo
{

    public static int GetHorizontalByDepthProduct()
    {
        string[] text = System.IO.File.ReadAllLines("files/day2.txt");
        Positions positions = new Positions
        {
            Vertical = 0,
            Horizontal = 0,
            Aim = 0
        };


        foreach (var item in text)
        {
            var splitText = item.Split(" ");
            var direction = splitText[0];
            var movement = int.Parse(splitText[1]);

            if (direction == "forward")
            {
                positions.Horizontal += movement;
                if (positions.Aim != 0)
                {
                    positions.Vertical += positions.Aim * movement;
                }
            }
            else if (direction == "up")
            {
                positions.Aim -= movement;
            }
            else
            {
                positions.Aim += movement;
            }
        }

        return positions.Horizontal * positions.Vertical;
    }
}