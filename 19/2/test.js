
// 1 => adds numbers read from 2 pos, stores result in 3rd
// 3 ints after op code

// 2 => multiplies inputs from positions, same as above mostly

// 99 => halt
let opcodes = [1, 2, 99];

//Once you're done processing an opcode, move to the next one by stepping forward 4 positions.

const fs = require('fs');
const path = require('path');

let program = fs.readFileSync(path.resolve('program.txt'), 'utf8').trim().split(',').map(x => parseInt(x));

function runProgram(prog, x, y, debug = false) {
    let cursor = 0;
    let fear = 0;
    prog[1] = x;
    prog[2] = y;
    do {
	let command = prog[cursor];
	      fear++;
	if (fear > prog.length) {
	    if (debug) console.log('error processing program');
	    return;
	}
	switch (command) {
	case 1:
	    var parameters = prog.slice(cursor, cursor + 4);
	    var sum = prog[parameters[1]] + prog[parameters[2]];
	    if (debug) console.log('add:', prog[parameters[1]], prog[parameters[2]], 'set:', parameters[3]);
	    prog[parameters[3]] = sum;
	    cursor = cursor + 4;
	    continue;
	case 2:
	    var parameters = prog.slice(cursor, cursor + 4);
	    var product = prog[parameters[1]] * prog[parameters[2]];
	    if (debug) console.log('multiply::', prog[parameters[1]], prog[parameters[2]], 'set:', parameters[3]);
	    prog[parameters[3]] = product;
	    cursor = cursor + 4;
	    continue;
	case 99:
	    return prog;
	default:
	    console.log('invalid command:', command);
	    return;
	}
    } while (true)
}

console.log(process.argv);
let sol = runProgram([...program], 10, 2)[0];

let needle = 19690720;

console.log(needle - sol);
