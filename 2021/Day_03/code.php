<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__);

// Task code
function part01(array $input): int
{
    $gamma = '';
    $epsilon = '';

    for ($pos = 0, $byteLength = strlen($input[0]); $pos < $byteLength; $pos++) {
        $mostCommonIndex = commonBit($input, $pos);
        $leastCommonIndex = $mostCommonIndex === 0 ? 1 : 0;

        $gamma .= $mostCommonIndex;
        $epsilon .= $leastCommonIndex;
    }

    return bindec($gamma) * bindec($epsilon);
}

function part02(array $input): int
{
    $oxygen = filterBytes($input, true);
    $carbon = filterBytes($input, false);

    return bindec($oxygen) * bindec($carbon);
}

function commonBit(array $bytes, int $position): int
{
    $occurrences = [0, 0];

    foreach ($bytes as $byte) {
        $bit = (int) $byte[$position];
        $occurrences[$bit]++;
    }

    if ($occurrences[0] === $occurrences[1]) {
        return -1;
    }

    return array_search(max($occurrences), $occurrences, true);
}

function filterBytes(array $input, bool $mostCommon): string
{
    $bytes = $input;

    for ($pos = 0, $byteLength = strlen($bytes[0]); $pos < $byteLength; $pos++) {
        if (count($bytes) === 1) {
            break;
        }

        $mostCommonIndex = commonBit($bytes, $pos);
        $bytes = array_filter(
            $bytes,
            static function($byte) use ($mostCommonIndex, $mostCommon, $pos){
                if ($mostCommonIndex === -1) {
                    if ($mostCommon) {
                        return $byte[$pos] === '1';
                    }

                    return $byte[$pos] === '0';
                }

                if (!$mostCommon) {
                    return $byte[$pos] === ($mostCommonIndex === 0 ? '1' : '0');
                }

                return $byte[$pos] === (string) $mostCommonIndex;
            }
        );
    }

    return current($bytes);
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
    [845186, 4636702], // Expected
    [$result01, $result02], // Result
);
