<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__)[0];

// Task code
function part01(string $input): int
{
    $sum = 0;

    for ($i = 0, $totalCount = strlen($input); $i < $totalCount; $i++) {
        if ($input[$i] === $input[$i - 1]) {
            $sum += $input[$i];
        }
    }

    return $sum;
}

function part02(string $input): int
{
    $sum = 0;

    for ($i = 0, $totalCount = strlen($input); $i < $totalCount; $i++) {
        if ($input[$i] === $input[$i - ($totalCount/2)]) {
            $sum += $input[$i];
        }
    }

    return $sum;
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
    [1223, 1284], // Expected
    [$result01, $result02], // Result
);
