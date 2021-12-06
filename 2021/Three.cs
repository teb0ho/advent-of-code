namespace AllDays;

public class Three
{
    public static int GetEpsilonByGammeRate()
    {
        string[] text = System.IO.File.ReadAllLines("files/day3.txt");
        Dictionary<int, int> epilson = new Dictionary<int, int>();

        for (int i = 0; i < 12; i++) 
        {
            int zero = 0, one = 0;
            foreach (var item in text)
            {
                if (item[i] == '0')
                {
                    zero++;
                }
                else 
                {
                    one++;
                }
            }

            if (zero > one)
            {
                epilson[i] = 0;
            }
            else 
            {
                epilson[i] = 1;
            }
        }

        string epilsonBinary = "";
        string gammaBinary = "";

        foreach (KeyValuePair<int, int> item in epilson)
        {
            epilsonBinary += item.Value.ToString();
        }

        foreach (var item in epilsonBinary)
        {
            if (item == '0')
            {
                gammaBinary += "1";
            }
            else 
            {
                gammaBinary += "0";
            }
        }

        return Convert.ToInt32(epilsonBinary, 2) * Convert.ToInt32(gammaBinary, 2);
    }
}