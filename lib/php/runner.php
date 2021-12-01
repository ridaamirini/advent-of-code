<?php

declare(strict_types=1);

require_once 'utils.php';

function run(string $codeFile): void
{
    $binary = getBinaryPathByFilename($codeFile);

    write(
        sprintf(
            'Executing: %s/%s/%s',
            basename(dirname($codeFile, 2)),
            basename(dirname($codeFile)),
            basename($codeFile)
        )
    );
    newLine(2);
    write(
        (string) shell_exec(
            sprintf('%s %s', $binary, $codeFile)
        )
    );
    newLine(2);
}

function runWithExitCode(string $codeFile): int {
    $binary = getBinaryPathByFilename($codeFile);

    write(
        sprintf(
            'Executing: %s/%s/%s',
            basename(dirname($codeFile, 2)),
            basename(dirname($codeFile)),
            basename($codeFile)
        )
    );
    newLine(2);
    $output = $exitCode = null;
    exec(sprintf('%s %s', $binary, $codeFile), $output, $exitCode);
    write(implode("\n", $output));
    newLine(2);

    return $exitCode;
}
