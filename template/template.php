<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__);

// Task code
function part01(array $input)
{

}

function part02(array $input)
{

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
    [], // Expected
    [$result01, $result02], // Result
);
