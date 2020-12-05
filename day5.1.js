const fs = require('fs');
       
fs.readFile('./files/day5.txt', (err, data) => {

    if (err) {
        throw err;
    }

    const lines = data.toString().split('\n');

    const seatIds = [];

    console.log(lines.length);

    for (boardingPass of lines) {

        let currentRowRange = 128;
        let availableRows = [0 , 128];
        let availableColumns = [0, 8];
        let currentColumnRange = 8;

        for (let i = 0; i < boardingPass.length; i++) {
        
            if (i < 7) {
                currentRowRange /= 2;

                if (boardingPass[i] == 'F') {
                    availableRows[1] = availableRows[0] + currentRowRange  - 1;
                } else {

                    availableRows[0] = availableRows[0] + currentRowRange;
                    availableRows[1] = availableRows[0] + currentRowRange - 1; 
                }

            }

            if (i >= 7) {
                currentColumnRange /= 2;

                if (boardingPass[i] == 'L') {
                    availableColumns[1] = availableColumns[0] + currentColumnRange  - 1;
                } else {

                    availableColumns[0] = availableColumns[0] + currentColumnRange;
                    availableColumns[1] = availableColumns[0] + currentColumnRange - 1; 
                }
            }
        }

        
        const answer = availableRows[1] * 8 + availableColumns[1];
        seatIds.push(answer);
    }

    let max = seatIds.reduce((a, b) => Math.max(a, b));
    console.log(max);
    
});
