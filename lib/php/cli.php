<?php

declare(strict_types=1);

define('CLI_SUCCESS', 0);
define('CLI_FAILURE', 1);

set_exception_handler(
    static fn(Throwable $ex) => writeln('Error message: ' . $ex->getMessage()) && exit(CLI_FAILURE)
);

function write(string $message): void
{
    print($message);
}

function writeln(string $message): void
{
    write($message . PHP_EOL);
}

function newLine(int $count = 1): void
{
    write(str_repeat(PHP_EOL, $count));
}

function getOption(string $option): ?string
{
    static $mappedOptions = [];

    if (empty($mappedOptions)) {
        $mappedOptions = parseOptions();
    }

    return $mappedOptions[$option] ?? null;
}

function parseOptions(): array
{
    global $argc;
    global $argv;

    $map = [];

    for ($i = 0; $i < $argc; $i++) {
        if (strpos($argv[$i], '-') !== false) {
            $i++;
            $map[str_replace('-', '', $argv[$i - 1])] = $argv[$i];
        }
    }

    return $map;
}

function validateInput(callable $validator, $input, string $message): void
{
    if ($validator($input)) {
        return;
    }

    writeln($message);

    exit(CLI_FAILURE);
}
