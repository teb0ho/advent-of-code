const fs = require('fs');
const _ = require('lodash');

fs.readFile('./files/day11.txt', (err, data) => {

    if (err) {
        throw err;
    }

    const lines = data.toString().split('\n');

    const columns = lines[0].length;
    const rows = lines.length;

    let arrayMatch = false;
    let seatArea = Array(rows).fill();
    let newSeatArea = Array(rows).fill().map(() => Array(columns).fill()); 
    
    for (let i = 0; i < rows; i++) {
        seatArea[i] = [...lines[i]];
    }

    while (!arrayMatch) {

        for (let i = 0; i < rows; i++) {
            for (let j = 0; j < columns; j++) {
                const seatValue = seatArea[i][j];
                const answer = countSeats(i, j);
    
                if (seatValue === 'L' && answer === 0) {
                    newSeatArea[i][j] = '#';
                }        
    
                if (seatValue === '#' && answer >= 4) {
                    newSeatArea[i][j] = 'L';
                }

                if (seatValue === '.') {
                    newSeatArea[i][j] = '.';
                }
            }
        }
    
        arrayMatch = _.isEqual(seatArea, newSeatArea);
        seatArea = newSeatArea.map(a => a.slice());

    }

    let count = 0;

    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < columns; j++) {
            if (seatArea[i][j] === '#') {
                count++;
            }
        }    
    }

    function countSeats(x, y) {
        let sum = 0;

        for (let i = -1; i < 2; i++) {
            for (let j = -1; j < 2; j++) {
                let column = y + i;
                let row = x + j;            

                if (!(column < 0 || column > columns - 1) && !(row < 0 || row > rows - 1)) {
                    let value = seatArea[row][column];
                    
                    if (value === '#') {
                        sum += 1;
                    }
                }
            } 
        }

        if (seatArea[x][y] === '#') {
            sum -= 1;
        }
        
        return sum;
    }

    console.log(count);

});

