const fs = require('fs');
const path = require('path');

let data = fs.readFileSync(path.resolve('puzzle.txt'), 'utf8').trim();

function range(size, startAt = 0) {
    return [...Array(size).keys()].map(i => i + startAt);
}

function toCoordPairs(data) {
    return data.split('\n').map(x => x.split(' -> ').map(y => y.split(',').map(y => parseInt(y))));
}

function enumerateLine(coordPairs) {
    let newPairs = [];
    for (var x = 0; x < coordPairs.length; x++) {
        let start = coordPairs[x][0];
        let end = coordPairs[x][1];
        //console.log('enumerating', start, '=>', end);
        let newPoints = [];
        if (start[0] === end[0]) {
            if (start[1] > end[1]) {
                let wait = start;
                start = end;
                end = wait;
            }
            //console.log('horizontal line');
            for (var y = start[1]; y < end[1]; y++) {
                //console.log('adding point', [start[0], y]);
                newPoints.push([start[0], y]);
            }
            newPoints.push(end);
        } else if (start[1] === end[1]) {
            if (start[0] > end[0]) {
                let wait = start;
                start = end;
                end = wait;
            }
            //console.log('vertical line');
            for (var z = start[0]; z < end[0]; z++) {
                //console.log('adding point', [z, start[1]]);
                newPoints.push([z, start[1]]);
            }
            newPoints.push(end);
        } else {
            console.log('diagnonal', start,'=>', end);
            let xs = start[0] > end[0] ? range((start[0] - end[0]) + 1, end[0]) : range((end[0] - start[0]) + 1, start[0]);
            let ys = start[1] > end[1] ? range((start[1] - end[1]) + 1, end[1]) : range((end[1] - start[1]) + 1, start[1]);
            if (start[0] > end[0] && start[1] < end[1]) {
                ys.reverse();
            } else if (start[1] > end[1] && start[0] < end[0]) {
                xs.reverse();
            }
            let zipped = xs.map((x, idx) => [xs[idx], ys[idx]]);
            newPoints.push(...zipped);
        }
        newPairs.push(newPoints);
    }
    return newPairs;
}

// lol what a slow way to do this, I'm sure there is some math that can be done 
// this is an exhaustive search of each coordinated against all other coordinates
// also i'm double counting because i don't consider line a crossing line b any differently than line b crossing line a
function findIntersections(lines) {
    let intersections = [];
    lines.map(line => {
        line.map(coord =>
            lines.map(l => {
                //console.log('looking in line', l, 'for', coord);
                let c = l.find(x => x[0] === coord[0] && x[1] === coord[1]);
                if (c) {
                    if (!intersections[coord.toString()]) {
                        intersections[coord.toString()] = 0;
                    }
                    //console.log('found', coord);
                    intersections[coord.toString()]++;
                }
            }));
    });
    return intersections;
}

let hOrVPairs = toCoordPairs(data);
console.log('converted to pairs');
let enumerated = enumerateLine(hOrVPairs);
console.log('enumerated line');
console.log('finding intersections');
let intersections = findIntersections(enumerated);
console.log('found', Object.keys(intersections).length, 'intersections');
console.log('calculating sum');
console.log(Object.keys(intersections).filter(x => intersections[x] > 1).length);

function drawSegments(segments) {
    let coordPlane = []; 
}
