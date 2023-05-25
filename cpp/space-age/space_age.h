#if !defined(SPACE_AGE_H)
#define SPACE_AGE_H

namespace space_age
{
    class space_age
    {
        double _seconds;
        double earth_year = 31557600;

    public:
        space_age(double seconds)
        {
            _seconds = seconds;
        }

        double seconds() const
        {
            return _seconds;
        }

        double on_earth() const
        {
            return _seconds / earth_year;
        }

        double on_mercury() const
        {
            return on_earth() / .2408467;
        }

        double on_venus() const
        {
            return on_earth() / .61519726;
        }

        double on_mars() const
        {
            return on_earth() / 1.8808158;
        }

        double on_jupiter() const
        {
            return on_earth() / 11.862615;
        }

        double on_saturn() const
        {
            return on_earth() / 29.447498;
        }

        double on_uranus() const
        {
            return on_earth() / 84.016846;
        }

        double on_neptune() const
        {
            return on_earth() / 164.79132;
        }
    };
} // namespace space_age

#endif // SPACE_AGE_H