<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__);

// Task code
function part01(array $input): int
{
    return array_reduce(
        str_split($input[0]),
        static function ($floor, $item) {
            if ($item === '(') {
                return ++$floor;
            }

            return --$floor;
        }
    );
}

function part02(array $input): int
{
    $commands = str_split($input[0]);
    $floor = 0;

    for ($i = 0, $totalCount = count($commands); $i < $totalCount; $i++) {
        $command = $commands[$i];

        if ($command === '(') {
            $floor++;
        } else {
            $floor--;
        }

        if ($floor === -1) {
            return $i+1;
        }
    }

    return 0;
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
    [232, 1783], // Expected
    [$result01, $result02], // Result
);
