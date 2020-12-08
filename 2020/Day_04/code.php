<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__, PHP_EOL . PHP_EOL);

function normalizePassport(array $input): array
{
    $passportList = [];
    foreach ($input as $line) {
        $values = explode(
            ':',
            str_replace([' ', PHP_EOL], ':', $line)
        );

        $passport = [];
        for ($i = 1, $iMax = count($values); $i < $iMax; ($i+=2)) {
            $passport[$values[$i-1]] = $values[$i];
        }

        $passportList[] = $passport;
    }

    return $passportList;
}

// Task code
function part01(array $input): int
{
    $requiredFields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'];
    $valid = 0;

    foreach (normalizePassport($input) as $passport) {
        foreach ($requiredFields as $key) {
            if (!array_key_exists($key, $passport)) {
                continue 2;
            }
        }

        $valid++;
    }

    return $valid;
}

function part02(array $input): int
{
    $validators = [
        'byr' => static fn ($v) => $v >= 1920 && $v <= 2002,
        'iyr' => static fn ($v) => $v >= 2010 && $v <= 2020,
        'eyr' => static fn ($v) => $v >= 2020 && $v <= 2030,
        'hgt' => static function ($v) {
            $unit = substr($v, -2);
            $value = substr($v, 0, -2);

            switch ($unit) {
                case 'cm': return $value >= 150 && $value <= 193;
                case 'in': return $value >= 59 && $value <= 76;
                default: return false;
            }
        },
        'hcl' => static fn ($v) => (bool) preg_match('/#[0-9a-f]{6}/', $v),
        'ecl' => static fn ($v) => in_array($v, ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']),
        'pid' => static fn ($v) => strlen($v) === 9,
    ];
    $valid = 0;

    foreach (normalizePassport($input) as $passport) {
        foreach ($validators as $key => $validator) {
            if (!array_key_exists($key, $passport) ||
                !$validator($passport[$key])
            ) {
                continue 2;
            }
        }

        $valid++;
    }

    return $valid;
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
    [216, 150], // Expected
    [$result01, $result02], // Result
);
