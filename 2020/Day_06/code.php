<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__, PHP_EOL . PHP_EOL);

// Task code
function part01(array $input): int
{
    $sum = 0;

    foreach ($input as $value) {
        $sum += count(
            array_unique(
                str_split(
                    str_replace(PHP_EOL, '', $value)
                )
            )
        );
    }

    return $sum;
}

function part02(array $input): int
{
    $sum = 0;

    foreach ($input as $line) {
        $group = explode(PHP_EOL, $line);
        $groupCount = count($group);
        $questionList = [];

        foreach ($group as $person) {
            foreach (str_split($person) as $question) {
                $questionList[$question] = ($questionList[$question] ?? 0) + 1;
            }
        }

        $sum += count(array_filter($questionList, static fn ($answers) => $answers === $groupCount));
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
    [6748, 3445], // Expected
    [$result01, $result02], // Result
);
