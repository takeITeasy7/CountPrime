namespace CountPrime
{
    internal class Program
    {
        const int MAX = 1000000;

        static void Main(string[] args)
        {
            DateTime start = DateTime.Now;

            // 直列
            PrimeTest();

            // 並列
            // PrimeTestParallel();

            DateTime end = DateTime.Now;
            TimeSpan delta = end - start;
            Console.WriteLine($"かかった時間は{delta.TotalSeconds}秒です");
        }

        // MAXまでの素数を求める（並列）
        static void PrimeTestParallel()
        {
            Parallel.For(2, MAX + 1, (i) =>
            {
                JudgePrime(i);
            });
        }

        // MAXまでの素数を求める（直列）
        static void PrimeTest()
        {
            for (int i = 2; i <= MAX; i++)
            {
                JudgePrime(i);
            }
        }

        static void JudgePrime(int target)
        {
            for (int i = 2; i <= target; i++)
            {
                if (i == target)
                {
                    // Console.WriteLine(target);   // 計測時はコメントアウト
                }
                else if (target % i == 0)
                {
                    break;
                }
            }
        }
    }
}