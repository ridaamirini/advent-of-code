<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__);

// Task code
function part01(array $input): int
{
    $iLength = count($input);
    $count = 0;

    for ($i = 1; $i < $iLength; $i++) {
        if ($input[$i] > $input[$i - 1]) {
            $count++;
        }
    }

    return $count;
}

function part02(array $input): int
{
    $iLen = count($input);
	$count = 0;
	$depth = $input[0] + $input[1] + $input[2];

	for ($i = 1; $i < $iLen-2; $i++) {
        $cDepth = $input[$i] + $input[$i + 1] + $input[$i + 2];

        if ($cDepth > $depth) {
            $count++;
		}

        $depth = $cDepth;
    }

    return $count;
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
    [1713, 1734], // Expected
    [$result01, $result02], // Result
);
