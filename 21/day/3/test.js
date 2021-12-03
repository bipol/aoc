const fs = require('fs');
const path = require('path');

let data = fs.readFileSync(path.resolve('puzzle.txt'), 'utf8').trim().split('\n');

let gamma_rate = [];
let epsilon_rate = [];

let oxygen_generator_rating = [];
let c02_scrubber_rating = [];
//To find oxygen generator rating, determine the most common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 1 in the position being considered.
//    To find CO2 scrubber rating, determine the least common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 0 in the position being considered.
let bit_criteria = 0;

// gamma_rate * epsilon_rate = power consumption
let power_consumption = 0;

function toMatrix(data) {
    let matrix = [];
    return data.map((x, idx) => matrix[idx] = x.split(''));
}

function getCol(idx, matrix) {
    return matrix.map(x => x[idx]);
}

function getMCB(col) {
    let sum = col.reduce((x, y) => parseInt(x) + parseInt(y));
    return sum >= col.length - sum ? 1 : 0;
}

function findOxygenGeneratorRating(matrix, idx) {
    let newMatrix = [];

    if (matrix.length === 1) {
        return matrix[0];
    }
    let lcb = getMCB(getCol(idx, matrix));
    newMatrix = matrix.filter(row => parseInt(row[idx]) === lcb);
    return findOxygenGeneratorRating(newMatrix, idx+1);
}

function findC02ScrubberRating(matrix, idx) {
    let newMatrix = [];

    if (matrix.length === 1) {
        return matrix[0];
    }
    let lcb = getMCB(getCol(idx, matrix)) === 0 ? 1 : 0;
    newMatrix = matrix.filter(row => parseInt(row[idx]) === lcb);
    return findC02ScrubberRating(newMatrix, idx+1);
}

let matrix = toMatrix(data);
for (let x = 0; x < matrix[0].length; x++) {
    gamma_rate[x] = getMCB(getCol(x, matrix));
    epsilon_rate[x] = getMCB(getCol(x, matrix)) === 0 ? 1 : 0;
}

console.log('gamma_rate', parseInt(gamma_rate.join(''), 2), 'epsilon_rate', parseInt(epsilon_rate.join(''),2));
console.log('power_consumption', parseInt(gamma_rate.join(''), 2) *  parseInt(epsilon_rate.join(''),2));
console.log('oxygen_generator_rating', parseInt(findOxygenGeneratorRating(matrix, 0).join(''), 2));
console.log('c02_scrubber_rating', parseInt(findC02ScrubberRating(matrix, 0).join(''), 2));
console.log('life_support_rating', parseInt(findC02ScrubberRating(matrix, 0).join(''), 2) * parseInt(findOxygenGeneratorRating(matrix, 0).join(''), 2));
