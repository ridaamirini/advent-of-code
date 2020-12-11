<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__);

function getRoundModeByChar(string $char): int
{
    switch ($char) {
        case 'F':
        case 'L': return PHP_ROUND_HALF_DOWN;
        case 'B':
        case 'R': return PHP_ROUND_HALF_UP;
        default: throw new Exception('Fizz buzz');
    }
}

function decodeSeatID(string $id): int {
    $charList = str_split($id);
    $yCharList = array_slice($charList, 0, 7);
    $xCharList = array_slice($charList, -3, 3);
    $y = [0, 127];
    $x = [0, 7];

    foreach ($yCharList as $char) {
        $mode = getRoundModeByChar($char);
        $result = round(($y[0] + $y[1]) / 2, 0, $mode);

        if (PHP_ROUND_HALF_DOWN === $mode) {
            $y = [$y[0], $result];

            continue;
        }

        $y = [$result, $y[1]];
    }

    foreach ($xCharList as $char) {
        $mode = getRoundModeByChar($char);
        $result = round(($x[0] + $x[1]) / 2, 0, $mode);

        if (PHP_ROUND_HALF_DOWN === $mode) {
            $x = [$x[0], $result];

            continue;
        }

        $x = [$result, $x[1]];
    }

    return current($y) * 8 + current($x);
}

// Task code
function part01(array $input): int
{
    $ids = array_map('decodeSeatID', $input);

    return max($ids);
}

function part02(array $input): ?int
{
    $ids = array_map('decodeSeatID', $input);

    for($i = min($ids) + 1; $i < max($ids) - 1; $i++ ) {
        if(!in_array($i, $ids, true)) {
            return $i;
        }
    }

    return null;
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
    [892, 625], // Expected
    [$result01, $result02], // Result
);
