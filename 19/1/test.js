/**
Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.

For example:

For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
For a mass of 1969, the fuel required is 654.
For a mass of 100756, the fuel required is 33583.
The Fuel Counter-Upper needs to know the total fuel requirement. To find it, individually calculate the fuel needed for the mass of each module (your puzzle input), then add together all the fuel values.

What is the sum of the fuel requirements for all of the modules on your spacecraft?
**/

const fs = require('fs')
const path = require('path')

// part I

function round(x) {
    return x >= 10 ? x - (x % 10) : x
}

function calculateFuelRequirement(mass) {
    return Math.trunc(mass / 3) - 2
}

let data = fs.readFileSync(path.resolve('input.txt'), 'utf8')

let masses = data.split('\n')

let totalFuel = masses.map(calculateFuelRequirement).reduce((x, y) => x + y)

console.log("Total Fuel Required: %d", totalFuel)

// part II
function recursivelyCalculateFuel(mass, total) {
    let fuel = calculateFuelRequirement(mass)
    if (fuel <= 0) {
	return total
    }
    return recursivelyCalculateFuel(fuel, total + fuel)
}

let totalFuelPartII = masses.map(x => recursivelyCalculateFuel(x, 0)).reduce((x, y) => x + y)

console.log("Total Fuel Required, accounting for fuel weight: %d", totalFuelPartII)
