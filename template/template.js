import {
    readInput,
    saveBenchmarkTime, startExecutionTime, stopExecutionTime,
    testResults
} from '../../lib/js/utils.mjs';

const input = readInput(import.meta.url);

// Task code
const part01 = (input) => {

};

const part02 = (input) => {

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
    [], // Expected
    [result01, result02] // Results
);