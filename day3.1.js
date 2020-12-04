const fs = require('fs');

let ex = [];

fs.readFile('./files/day3.txt', (err, data) => {
    if (err) {
        throw err;
    }

    let index = 3;
    const lines = data.toString().split('\n');

    for (let i = 1; i < lines.length; i++) {
        
        if (lines[i][index] == '#') {
            ex.push(1);
        } 
        
        for (let j = 0; j < 3; j++) {
            index++;
            if (index == lines[i].length) {
                index = 0;
            }
        }
    }

    console.log(ex.length);
});

