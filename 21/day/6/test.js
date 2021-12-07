const fs = require('fs');
const path = require('path');
const util = require('util');
const Stream = require('stream');

let data = fs.readFileSync(path.resolve('puzzle.txt'), 'utf8').trim();

let seed = data.split(',').map(x => parseInt(x));
let simulation = [];

seed.map(x => {
    if (simulation[parseInt(x)]) {
        simulation[x]++;
    } else {
        simulation[parseInt(x)] = 1;
    }
})

function runSimArray(arr, days) {
    let prevSeed = arr;
    for (var day = 0; day < days; day++) {
        console.log('calculating day', day);
        let newSeed = [];
        prevSeed.map((val, key) => {
            switch(key) {
            case 0:
                newSeed[6] = prevSeed[0];
                newSeed[8] = prevSeed[0];
                break;
            default:
                newSeed[key-1] = newSeed[key-1] ? newSeed[key-1] + prevSeed[key] : prevSeed[key];
                break;
            }
        });
        prevSeed = newSeed;
    }
    return prevSeed;
}

console.log(runSimArray(simulation, 256).reduce((x, y) => x + y));
