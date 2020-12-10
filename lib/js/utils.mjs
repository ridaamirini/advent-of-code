import url from 'url';
import path from 'path';
import fs from 'fs';
import { performance } from 'perf_hooks';

const readInput = (fileUrl, delimiter = '\n') => {
    const inputFilePath = path.resolve(
        path.dirname(
            url.fileURLToPath(fileUrl)
        ),
        'input.txt'
    );

    return fs.readFileSync(inputFilePath, 'utf-8')
        ?.split(delimiter);
};

const testResults = (expects, results) => {
    if (expects.length === 0) {
        console.log('Skipped test!');

        return;
    }

    for (let i = 0; i < expects.length; i++) {
        const index = i+1;
        const expect = expects[i];
        const result = results[i];

        if (expect === result) {
            console.log('\x1b[32m%s\x1b[0m', `Part ${index}: passed`);
            continue;
        }

        console.log(
            '\x1b[31m%s\x1b[0m', `Part ${index}: failed with Expected: '${expect}' but get '${result}'.`
        );
    }
};

const startExecutionTime = () => performance.now();
const stopExecutionTime = startTime => {
    return (performance.now() - startTime).toFixed(4) + 'ms';
}

const saveBenchmarkTime = (executionTime, fileUrl) => {
    const benchmarkFilePath = path.resolve(
        path.dirname(
            url.fileURLToPath(fileUrl)
        ),
        'benchmark.json'
    );
    let benchmark = [];

    if (fs.existsSync(benchmarkFilePath)) {
        benchmark = JSON.parse(fs.readFileSync(benchmarkFilePath, 'utf-8'));
    }

    benchmark['Javascript'] = executionTime;

    fs.writeFileSync(benchmarkFilePath, JSON.stringify(benchmark));
};

export {
    readInput,
    testResults,
    startExecutionTime,
    stopExecutionTime,
    saveBenchmarkTime,
}