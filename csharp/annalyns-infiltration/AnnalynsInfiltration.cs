using System;

static class QuestLogic
{
    public static bool CanFastAttack(bool knightIsAwake)
    {
        if (!knightIsAwake)
        {
            return true;
        }
        return false;
    }

    public static bool CanSpy(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake)
    {
        if (knightIsAwake || archerIsAwake || prisonerIsAwake)
        {
            return true;
        }
        return false;
    }

    public static bool CanSignalPrisoner(bool archerIsAwake, bool prisonerIsAwake)
    {
        if (prisonerIsAwake && !archerIsAwake)
        {
            return true;
        }
        return false;
    }

    public static bool CanFreePrisoner(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake, bool petDogIsPresent)
    {
        if (prisonerIsAwake && (!knightIsAwake && !archerIsAwake))
        {
            return true;
        }
        if (petDogIsPresent && !archerIsAwake)
        {
            return true;
        }
        return false;
    }
}
