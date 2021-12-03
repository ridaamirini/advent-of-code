<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__);

// Task code
const UP    = 'U';
const DOWN    = 'D';
const LEFT  = 'L';
const RIGHT = 'R';
const KEY_PAD_PART_1 = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
];
const KEY_PAD_PART_2 = [
    [null, null, 1, null, null],
    [null, 2, 3, 4, null],
    [5, 6, 7, 8, 9],
    [null, 'A', 'B', 'C', null],
    [null, null, 'D', null, null],
];


function part01(array $input): string
{
    $coords = [1, 1];
    $code = '';

    foreach ($input as $moves) {
        $coords = parseMoves(KEY_PAD_PART_1, $coords, $moves);
        [$y, $x] = $coords;
        $code .= KEY_PAD_PART_1[$y][$x];
    }

    return $code;
}

function part02(array $input): string
{
    $coords = [2, 0];
    $code = '';


    foreach ($input as $moves) {
        $coords = parseMoves(KEY_PAD_PART_2, $coords, $moves);
        [$y, $x] = $coords;
        $code .= KEY_PAD_PART_2[$y][$x];
    }

    return $code;
}

function parseMoves(array $keyPad, array $coords, string $moves): array
{
    foreach (str_split($moves) as $move) {
        [$y, $x] = $coords;
        [$nY, $nX] = calculateCoordinatesByMove($y, $x, $move);

        $edge = $keyPad[$nY][$nX] ?? null;
        if ($edge === null) {
            continue;
        }

        $coords = [$nY, $nX];
    }

    return $coords;
}

function calculateCoordinatesByMove(int $y, $x, string $move): array
{
    if ($move === UP) {
        return [--$y, $x];
    }

    if ($move === DOWN) {
        return [++$y, $x];
    }

    if ($move === RIGHT) {
        return [$y, ++$x];
    }

    if ($move === LEFT) {
        return [$y, --$x];
    }

    return [$y, $x];
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
    ['48584', '563B6'], // Expected
    [$result01, $result02], // Result
);
