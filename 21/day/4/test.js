const fs = require('fs');
const path = require('path');

let data = fs.readFileSync(path.resolve('puzzle.txt'), 'utf8').trim();

function parseData(data) {
    let drawn = data.split('\n')[0].split(',').map(x => parseInt(x));
    let rest = data.split('\n').slice(2);
    let boards = [];
    let board = [];
    for (var x = 0; x < rest.length; x++) {
        if (rest[x] === "") {
            boards.push(board.map(row => row.split(/(\d+)/).filter(num => !isNaN(parseInt(num))).map(num => parseInt(num))));
            board = [];
            continue;
        }
        board.push(rest[x]);
    }
    //boards.push(board.map(row => row.split(/(\d+)/).filter(num => !isNaN(parseInt(num))).map(parseInt)));
    boards.push(board.map(row => row.split(/(\d+)/).filter(num => !isNaN(parseInt(num))).map(num => parseInt(num))));
    console.log(boards);
    return [boards, drawn];
}

function playSparse(boards, drawn) {
    let winningBoards = [];
    let marked = [];
    // for all drawn numbers
    for (var x = 0; x < drawn.length; x++) {
        let num = drawn[x];
        // mark the number on the board
        for (var y = 0; y < boards.length; y++) {
            if (winningBoards[y]) {
                console.log('board', y, 'has already won before, skipping');
            }
            let board = boards[y];
            if (!marked[y]) {
                marked[y] = [];
            }
            board.map((row, rowIdx) => {
                row.map((boardNumber, numIdx) => {
                    if (boardNumber === num) {
                        marked[y].push([rowIdx, numIdx]);
                    }
                });
            });
            for (var m = 0; m < marked[y].length; m++) {
                let [rowIdx, numIdx] = marked[y][m];
                // check to see if we won after marking this board
                if (marked[y].filter(x => x[0] === rowIdx).length === 5) {
                    winningBoards[y] = true;
                    if (winningBoards.filter(x => x).length === boards.length) {
                        console.log('board', y, 'is our last winner, returning board state');
                        return [marked[y], board, num];
                    }
                }

                if (marked[y].filter(x => x[1] === numIdx).length === 5) {
                    winningBoards[y] = true;
                    if (winningBoards.filter(x => x).length === boards.length) {
                        console.log('board', y, 'is our last winner, returning board state');
                        return [marked[y], board, num];
                    }
                }
            }
        }
    }
    console.log("didn't find a winning board, bug");
}

[boards, drawn] = parseData(data);

//let [marks, board, winningNumber] = play(boards, drawn);
let [marks, board, winningNumber] = playSparse(boards, drawn);

console.log(marks);
let total = 0;
board.map((row, rIdx) => {
    row.map((num, cIdx) => {
        if (!marks.find(x => x[0] === rIdx && x[1] === cIdx)) {
            total += board[rIdx][cIdx];
        } else {
            console.log('found mark', [rIdx, cIdx]);
        }
    });
});
console.log('total', total, 'multiplied by ', winningNumber, 'is', winningNumber * total);
