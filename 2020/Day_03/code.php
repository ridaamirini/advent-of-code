<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__);

function countSlopeTrees(int $steps, int $down, array $input): int
{
    $mapLength = count($input);
    $lineLength = strlen($input[0]);
    $trees = 0;
    $stepsCount = 0;
    $index = 0;

    foreach ($input as $lineNo => $line) {
        $index += $down;

        if ($index > $mapLength) {
            return $trees;
        }

        $slope = $input[$index] ?? $line;
        $stepsCount = ($stepCount = $stepsCount + $steps) < $lineLength ?
            $stepCount :
            abs($lineLength - $stepCount);

        if ('#' === $slope[$stepsCount]) {
            $trees++;
        }
    }

    return $trees;
}

// Task code
function part01(array $input): int
{
    return countSlopeTrees(3, 1, $input);
}

function part02(array $input): int
{
    $checks = [
        countSlopeTrees(1, 1, $input),
        countSlopeTrees(3, 1, $input),
        countSlopeTrees(5, 1, $input),
        countSlopeTrees(7, 1, $input),
        countSlopeTrees(1, 2, $input),
    ];

    return array_product($checks);
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
    [223, 3517401300], // Expected
    [$result01, $result02], // Result
);
