<?php

declare(strict_types=1);

require_once 'cli.php';

function getYearOption(): int
{
    $year = (int) (getOption('year') ?? date('Y'));

    validateInput(
        static fn ($year) => $year >= 2015 && date('Y'),
        $year,
        sprintf('Year: >= 2015 and <= %s!', date('Y'))
    );

    return $year;
}

function getDayOption(): int
{
    $day = (int) getOption('day');

    validateInput(
        static fn ($day) => $day >= 1 && $day <= 25,
        $day,
        'Day: >= 1 and <= 25!'
    );


    return $day;
}
