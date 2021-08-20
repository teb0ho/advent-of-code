const fs = require('fs');

let results = [];
const rules = [{x: 1, y: 1}, {x: 3, y: 1}, {x: 5, y: 1}, {x: 7, y: 1}, {x: 1, y: 2}];


fs.readFile('./files/day3.txt', (err, data) => {

    if (err) {
        throw err;
    }
    
    const lines = data.toString().split('\n');
    
    for (const rule of rules) {
        let index = 0;
        let i = rule.y;
        let count = 0;

        while (i < lines.length) {
            
            for (let j = 0; j < rule.x; j++) {
                index++;
                if (index == lines[i].length) {
                    index = 0;
                }
            }
            
            if (lines[i][index] == '#') {
                count++;
            }
            
            i += rule.y;
            
        }

        results.push(count);
        
    }

    const finalResult = results.reduce((i, j) => i * j);

    console.log(results);

    console.log(finalResult);
});
