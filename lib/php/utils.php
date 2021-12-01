<?php

declare(strict_types=1);

require_once 'cli.php';

function getRootPath(): string
{
    return __DIR__ . '/../..';
}

function getWorkingPath(int $year, ?int $day = null): string
{
    $parent = sprintf('%s/%d', getRootPath(), $year);

    if (null === $day) {
        return $parent;
    }

    return sprintf(
        '%s/Day_%02d',
        $parent,
        $day
    );
}

function getConfig(): array
{
    return include getRootPath() . '/config.php';
}

function mkdirIfNotExists(string $path): bool
{
    if (file_exists($path)) {
        return true;
    }

    $concurrentDirectory = dirname($path);
    if (!file_exists($concurrentDirectory) &&
        !mkdir($concurrentDirectory) &&
        !is_dir($concurrentDirectory)
    ) {
        throw new \RuntimeException(sprintf('Directory "%s" was not created', $concurrentDirectory));
    }

    return \mkdir($path);
}

function copyFileIfNotExists(string $resource, string $destination): void
{
    if (file_exists($destination)) {
        return;
    }

    copy($resource, $destination);
}

function getStringBetween(string $haystack, string $start, string $end): string
{
    $startLength = strlen($start);
    $startPosition = strpos($haystack, $start, $startLength) + $startLength;
    $endPosition = strpos($haystack, $end, $startPosition) - $startPosition;

    return trim(substr($haystack, $startPosition, $endPosition));
}

function locateBinaryPath(string $name): string
{
    $nameEscaped = escapeshellarg($name);

    $path = shell_exec("command -v $nameEscaped || which $nameEscaped || type -p $nameEscaped");
    if (empty($path)) {
        throw new \RuntimeException(
            sprintf('Can\'t locate [%s] - neither of [command|which|type] commands are available', $nameEscaped)
        );
    }

    return trim(str_replace("$name is", "", $path));
}

function getBinaryPathByFilename(string $fileName): string
{
    static $binaryExtMap = null;
    if (null === $binaryExtMap) {
        $binaryExtMap = include __DIR__ . '/../../binary-ext.map.php';
    }

    $extension = '.' . pathinfo($fileName, PATHINFO_EXTENSION);
    $binaryPath = null;

    foreach ($binaryExtMap as $binary => $values) {
        $extensionList = $values['extensions'] ?? $values;
        if (in_array($extension, $extensionList, true)) {
            $binaryPath = locateBinaryPath($binary);

            if (isset($values['options'])) {
                $binaryPath .= ' ' . implode(' ', $values['options']);
            }
        }
    }

    if (null === $binaryPath) {
        throw new \RuntimeException(
            sprintf('No binary found in "binary-ext.map" for "%s"', $extension)
        );
    }

    return $binaryPath;
}

function getDayPadded(int $day): string
{
    return str_pad((string) $day, 2, '0', STR_PAD_LEFT);
}

function arrayFlatten(array $items): array
{
    return array_reduce(
        $items,
        static fn ($carry, $item) => is_array($item)
            ? [...$carry, ...arrayFlatten($item)]
            : [...$carry, $item],
        []
    );
}

function readInput(string $currentDir, string $delimiter = PHP_EOL): array
{
    return explode($delimiter, file_get_contents($currentDir . '/input.txt'));
}

function testResults(array $expects, array $results): void
{
    if (empty($expects)) {
        writeln('Skipped tests!');

        return;
    }

    foreach ($expects as $key => $value) {
        $index = $key + 1;
        $result = $results[$key];

        if ($value === $result) {
            writeln(
                sprintf("\033[32mPart %d: passed\033[37m", $index)
            );
            continue;
        }

        writeln(
            sprintf(
                "\033[31mPart %d: failed with Expected: '%s' but get '%s'. \033[37m",
                $index,
                $value,
                $result
            )
        );
    }
}

function calcExecutionTime(): ?string
{
    static $startTime = null;
    if (null === $startTime) {
        $startTime = microtime(true);

        return null;
    }

    $result = (microtime(true) - $startTime) * 1000;
    $startTime = null;

    return sprintf('%.4f', $result) . 'ms';
}

function saveBenchmarkTime(string $executionTime, string $currentDir): void
{
    $benchmarkFilePath = $currentDir . '/benchmark.json';
    $benchmark = [];

    if (file_exists($benchmarkFilePath)) {
        $benchmark = json_decode(
            file_get_contents($benchmarkFilePath),
            true,
            512,
            JSON_THROW_ON_ERROR
        );
    }

    $benchmark['PHP'] = $executionTime;

    file_put_contents($benchmarkFilePath, json_encode($benchmark, JSON_PRETTY_PRINT));
}
