
const fs = require('fs')
const path = require('path')

let data = fs.readFileSync(path.resolve('puzzle.txt'), 'utf8').trim().split('\n').map(x => parseInt(x))

let increased = 0

function incrFunc(x, idx) {
    if (idx+1 < data.length) {
	if (parseInt(data[idx+1]) > parseInt(x)) {
	    increased++
	}
    }
}

data.map(incrFunc)

console.log(increased)

let slidingWindow = [];

function createWindows(x, idx) {
    let window = data.slice(idx, idx+3);
    slidingWindow.push(window);
}

data.map(createWindows)
//console.log(slidingWindow);

let sums = slidingWindow.map((x, idx) => x.reduce((x, y) => x + y))
//console.log(sums);

let numIncreased = 0;
function cmp(x, idx) {
    if (idx+1 < sums.length) {
	if (sums[idx+1] > x) {
	    numIncreased++
	}
    }
}

sums.map(cmp)
console.log('prewindowed', numIncreased)

function diff(data, agg) {
    for (x = 0; x < data.length-1; x++) {
	let a = data.slice(x, x+3).reduce((x, y) => x + y);
	let b = data.slice(x+1, x+4).reduce((x, y) => x + y);
	if (b > a) {
	    agg++
	}
    }
    return agg
}

console.log('ll', diff(data, 0))

function inPlace(data, agg) {
    for (x = 0; x < data.length-1; x++) {
	let a = data.slice(x, x+3).reduce((x, y) => x + y);
	let b = data.slice(x+1, x+4).reduce((x, y) => x + y);
	if (b > a) {
	    agg++
	}
    }
    return agg
}
