<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__);

// Task code
function part01(array $input): int
{
    $result = 0;

    foreach ($input as $firstValue) {
        foreach ($input as $secondValue) {
            $sum = $firstValue + $secondValue;
            if ($sum === 2020) {
                $result = $firstValue * $secondValue;
            }
        }
    }

    return $result;
}

function part02(array $input): int
{
    $result = 0;

    foreach ($input as $firstValue) {
        foreach ($input as $secondValue) {
            foreach ($input as $thirdValue) {
                $sum = $firstValue + $secondValue + $thirdValue;

                if ($sum === 2020) {
                    $result = $firstValue * $secondValue * $thirdValue;
                }
            }
        }
    }

    return $result;
}

// Execute
calcExecutionTime();
$result01 = part01($input);
writeln('Solution Part 1: ' . $result01);
$result02 = part02($input);
writeln('Solution Part 2: ' . $result02);
$executionTime = calcExecutionTime();

echo("Execution time: $executionTime" . PHP_EOL);
saveBenchmarkTime($executionTime, __DIR__);

// Task test
testResults(
    [935419, 49880012], // Expected
    [$result01, $result02], // Result
);
