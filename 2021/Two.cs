namespace AllDays;

public class Positions
{
    public int Vertical { get; set; }
    public int Horizontal { get; set; }
    public int Aim { get; set; }
}

public class Two
{

    public static int GetHorizontalByDepthProduct() 
    {
        string[] text = System.IO.File.ReadAllLines("files/day2.txt");
        Positions positions = new Positions 
        {
            Vertical = 0,
            Horizontal = 0
        };


        foreach (var item in text)
        {
            var splitText = item.Split(" ");
            var direction = splitText[0];
            var movement = int.Parse(splitText[1]);

            if (direction == "forward")
            {
                positions.Horizontal += movement;      
            } 
            else if (direction == "up")
            {
                positions.Vertical -= movement;
            }
            else 
            {
                positions.Vertical += movement;
            }
        }

        return positions.Vertical * positions.Horizontal;
    }
}