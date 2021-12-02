const fs = require('fs')
const path = require('path')

let hpos = 0;
let depth = 0;
let aim = 0;

let data = fs.readFileSync(path.resolve('puzzle.txt'), 'utf8').trim().split('\n')

function handleComm(command, depth, hpos)
{
    console.log(command)
    switch(command[0]) {
    case 'up':
	return [depth - command[1], hpos]
    case 'down':
	return [depth + command[1], hpos]
    case 'forward':
	return [depth, hpos + command[1]]
    }
}

function handleAim(command, depth, hpos, aim)
{
    console.log(command)
    switch(command[0]) {
    case 'up':
	return [depth, hpos, aim - command[1]]
    case 'down':
	return [depth, hpos, aim + command[1]]
    case 'forward':
	return [depth + (aim * command[1]), hpos + command[1], aim]
    }
}

function parseCourse(data, depth, hpos, aim) {
    let tup = [depth, hpos, aim]
    return data.map(x => x.split(' '))
	.map(x => [x[0], parseInt(x[1])])
	.reduce((x, y) => handleAim(y, x[0], x[1], x[2]), tup)
}

let [newD, newHPOS, newAim] = parseCourse(data, depth, hpos, aim)
console.log(newD, newHPOS, newAim)
console.log('depth * hpos', newD * newHPOS)
