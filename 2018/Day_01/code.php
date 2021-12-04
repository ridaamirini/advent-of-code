<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__);

// Task code
function part01(array $input): int
{
    return array_reduce(
        $input,
        static function ($carry, $line) {
            return $carry + (int) $line;
        }
    );
}

function part02(array $input): int
{
    $freqList = [0 => 1];
    $freq = 0;

    while (true) {
        foreach ($input as $v) {
            $freq += $v;

            if ($freqList[$freq] ?? 0 === 1) {
                return $freq;
            }

            $freqList[$freq] = 1;
        }
    }
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
    [574, 452], // Expected
    [$result01, $result02], // Result
);
