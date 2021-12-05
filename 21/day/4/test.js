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

function play(boards, drawn) {
    let boardStates = [];
    let winningBoards = [];
    // for all drawn numbers
    for (var x = 0; x < drawn.length; x++) {
        let num = drawn[x];
        // mark the number on the board
        for (var y = 0; y < boards.length; y++) {
            if (winningBoards[y]) {
                console.log('board', y, 'has already won before, skipping');
            }
            let board = boards[y];
            let marked = [];
            if (!boardStates[y]) {
                boardStates[y] = [];
            }
            board.map((row, rowIdx) => {
                row.map((boardNumber, numIdx) => {
                    if (boardNumber === num) {
                        // init boardState
                        if (!boardStates[y][rowIdx]) {
                            boardStates[y][rowIdx] = [];
                        }
                        boardStates[y][rowIdx][numIdx] = true;
                        //console.log('marked board', y, 'row', rowIdx, 'column', numIdx, 'number', boardNumber);
                        marked.push([rowIdx, numIdx]);
                    }
                });
            });
            // check to see if we won after marking this board
            for (var m = 0; m < marked.length; m++) {
                let [rowIdx, numIdx] = marked[m];
                // check to see if all rows have been marked
                if (boardStates[y][rowIdx].filter(x => x).length === 5) { // we have 5 numbers in a row marked
                    console.log('we have a full row match on board', y, 'with number', num, 'and row', rowIdx);
                    console.log('marking board', y, 'as a winner');
                    winningBoards[y] = true;
                    if (winningBoards.filter(x => x).length === boards.length) {
                        console.log('board', y, 'is our last winner, returning board state');
                        return [boardStates[y], board, num];
                    }
                }
                // check to see if an entire column has been marked
                let checkedRow = 0; 
                for (var row = 0; row < 5; row++) {
                    if (!boardStates[y][row]) {
                        continue;
                    }
                    if (boardStates[y][row][numIdx]) {
                        checkedRow++;
                    }
                }
                if (checkedRow === 5) {
                    console.log('we have a full col match on board', y, 'with number', num, 'and col', numIdx);
                    console.log('marking board', y, 'as a winner');
                    winningBoards[y] = true;
                    if (winningBoards.filter(x => x).length === boards.length) {
                        console.log('board', y, 'is our last winner, returning board state');
                        return [boardStates[y], board, num];
                    }
                }
            }
        }
    }
    console.log("didn't find a winning board, bug");
}

[boards, drawn] = parseData(data);

let [marks, board, winningNumber] = play(boards, drawn);

let total = 0;
board.map((row, rIdx) => {
    row.map((num, cIdx) => {
        if (!marks[rIdx][cIdx]) {
            total += board[rIdx][cIdx];
        }
    });
});
console.log('total', total, 'multiplied by ', winningNumber, 'is', winningNumber * total);
