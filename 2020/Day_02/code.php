<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__);

// Task code
function part01(array $input): int
{
    $validPasswords = 0;

    foreach ($input as $value) {
        $row = explode(' ', $value);
        [$min, $max] = explode('-', $row[0]);
        $character = $row[1][0];
        $password = $row[2];

        $matches = [];
        preg_match_all(
            '/' . $character . '/',
            $password,
            $matches,
            PREG_OFFSET_CAPTURE
        );
        $occurrences = count($matches[0]);


        if ($occurrences >= $min &&
            $occurrences <= $max
        ) {
            $validPasswords++;
        }
    }

    return $validPasswords;
}

function part02(array $input)
{
    $validPasswords = 0;

    foreach ($input as $value) {
        $row = explode(' ', $value);
        [$firstPos, $secondPos] = explode('-', $row[0]);
        $character = $row[1][0];
        $password = $row[2];

        $matches = [];
        preg_match_all(
            '/' . $character . '/',
            $password,
            $matches,
            PREG_OFFSET_CAPTURE
        );
        $occurrences = array_map(
            static fn ($value) => $value[1]+1,
            $matches[0]
        );

        $firstPosInArray = in_array((int) $firstPos, $occurrences, true);
        $secondPosInArray = in_array((int) $secondPos, $occurrences, true);
        if ((true === $firstPosInArray && false === $secondPosInArray) ||
            (true === $secondPosInArray && false === $firstPosInArray)
        ) {
            $validPasswords++;
        }
    }

    return $validPasswords;
}

// Execute
calcExecutionTime();
$result01 = part01($input);
$result02 = part02($input);
$executionTime = calcExecutionTime();

writeln('Solution Part 1: ' . $result01);
writeln('Solution Part 2: ' . $result02);
writeln('Execution time: ' . $executionTime);

saveBenchmarkTime($executionTime, __DIR__);

// Task test
testResults(
    [378, 280], // Expected
    [$result01, $result02], // Result
);
