<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__)[0];

// Task code
function part01(string $input): int
{
    $digitList = [];

    for ($i = 0, $totalCount = strlen($input); $i < $totalCount; $i++) {
        if ($input[$i] === $input[$i - 1]) {
            $digitList[] = $input[$i];
        }
    }

    return array_sum($digitList);
}

function part02(string $input): int
{
    $digitList = [];

    for ($i = 0, $totalCount = strlen($input); $i < $totalCount; $i++) {
        if ($input[$i] === $input[$i - ($totalCount/2)]) {
            $digitList[] = $input[$i];
        }
    }

    return array_sum($digitList);
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
    [], // Expected
    [$result01, $result02], // Result
);
