<?php

declare(strict_types=1);

require_once 'utils.php';

define('API', 'https://adventofcode.com');

function aocGet(string $endpoint): string
{
    $sessionId = getConfig()['sessionId'];
    $context = stream_context_create([
        'http' => [
            'header' => "Cookie: session=$sessionId\r\n",
        ]
    ]);

    return (string) file_get_contents(
        sprintf('%s/%s', API, $endpoint),
        false,
        $context
    );
}

function fetchTaskDescription(int $year, int $day): string
{
    $rawDescription = aocGet(
        sprintf('%d/day/%d', $year, $day)
    );

    return strip_tags(
        getStringBetween(
            $rawDescription,
            '<article class="day-desc">',
            '</article>'
        )
    );
}

function fetchTaskInput(int $year, int $day): string
{
    return trim(
        aocGet(
            sprintf('%d/day/%d/input', $year, $day)
        )
    );
}
