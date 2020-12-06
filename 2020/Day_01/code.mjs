import {
    readInput,
    saveBenchmarkTime, startExecutionTime, stopExecutionTime,
    testResults
} from '../../lib/js/utils.mjs';

const input = readInput(import.meta.url)
    .map(value => parseInt(value));

// Task code
const part01 = (input) => {
    for (let i = 0; i < input.length; i++) {
        for (let j = 0; j < input.length; j++) {
            const sum = input[i] + input[j];

            if (2020 === sum) {
                return input[i] * input[j];
            }
        }
    }
};

const part02 = (input) => {
    for (let i = 0; i < input.length; i++) {
        for (let j = 0; j < input.length; j++) {
            for (let k = 0; k < input.length; k++) {
                const sum = input[i] + input[j] + input[k];

                if (2020 === sum) {
                    return input[i] * input[j] * input[k];
                }
            }
        }
    }
};

// Execute
const startTime = startExecutionTime();
const result01 = part01(input);
const result02 = part02(input);
const executionTime = stopExecutionTime(startTime);

console.log('Solution Part 1: ' + result01);
console.log('Solution Part 2: ' + result02);
console.log('Execution time: ' + executionTime);

saveBenchmarkTime(executionTime, import.meta.url);

// Task test
testResults(
[935419, 49880012],
[result01, result02]
);