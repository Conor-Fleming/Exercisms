class Lasagna
{
    public int ExpectedMinutesInOven()
    {
        //Cookbook calls for 40 minutes in over
        return 40;
    }

    public int RemainingMinutesInOven(int elapsedTimeInOven)
    {
        //Expected time less elapsed time in oven
        return ExpectedMinutesInOven() - elapsedTimeInOven;
    }

    public int PreparationTimeInMinutes(int layers)
    {
        //Each layer takes about 2 minutes
        return layers * 2;
    }

    public int ElapsedTimeInMinutes(int layers, int elapsedTimeInOven)
    {
        //Prep time plus elapsed time in over gives total elapsed time
        return PreparationTimeInMinutes(layers) + elapsedTimeInOven;
    }
}
