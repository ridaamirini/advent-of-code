<?php

require_once __DIR__ . '/../../lib/php/utils.php';

$input = readInput(__DIR__);

// Task code
function part01(array $input): int
{
    $bagGraph = transpose(
        createBagGraph($input)
    );

    return traverse($bagGraph, 'shiny gold') - 1;
}

function part02(array $input): int
{
    $graph = createBagGraph($input);

    return countBags($graph, 'shiny gold')-1;
}

function createBagGraph(array $input): array
{
    $bagGraph = [];

    foreach ($input as $line) {
        $matches = null;
        $matched = preg_match('/(.*) bags contain (.*)/', $line, $matches);

        if (!$matched) {
            continue;
        }

        [, $outerBag, $innerBags] = $matches;
        $bagGraph[$outerBag] = [];

        $innerBagList = explode(', ', $innerBags);
        foreach ($innerBagList as $innerBag) {
            $matches = null;
            $matched = preg_match('/(\d+) (.*) (bags|bag)/', $innerBag, $matches);

            if (!$matched) {
                continue;
            }

            [, $weight, $bag] = $matches;
            $bagGraph[$outerBag][$bag] = $weight;
        }
    }

    return $bagGraph;
}

function transpose(array $graph): array
{
    $result = [];

    foreach ($graph as $outerBag => $innerBags) {
        foreach ($innerBags as $innerBag => $weight) {
            if (!array_key_exists($innerBag, $result)) {
                $result[$innerBag] = [];
            }

            $result[$innerBag][] = $outerBag;
        }
    }

    return $result;
}

function traverse(array $graph, string $node, array &$hits = []): int
{
    if (in_array($node, $hits, true)) {
        return 0;
    }

    $hits[] = $node;

    $count = 1;
    foreach ($graph[$node] ?? [] as $parent) {
        $count += traverse($graph, $parent, $hits);
    }

    return $count;
}

function countBags(array $graph, string $node): int
{
    $innerBags = $graph[$node];
    $count = 1;

    foreach ($innerBags as $innerBag => $weight) {
        $count += countBags($graph, $innerBag) * $weight;
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
    [326, 5635], // Expected
    [$result01, $result02], // Result
);
