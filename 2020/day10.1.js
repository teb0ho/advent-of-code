const fs = require('fs');

fs.readFile('./files/day10.txt', (err, data) => {
    if (err) {
        throw err;
    }

    const lines = data.toString().split('\n');
    const linesv2 = lines.slice();
    linesv2.sort((a, b) => a - b);

    let difference = 0;
    const finalResults = [[], []];

    for (const _ of lines) {
        for (let i = 0; i < linesv2.length; i++) {
            if (+linesv2[i] - difference === 1) {
                finalResults[0].push(1);
                linesv2.splice(i, 1);
                difference++;
                break;
            }
            else if (+linesv2[i] - difference === 3) {
                finalResults[1].push(1);
                linesv2.splice(i, 1);
                difference += 3;
                break;
            }
        }
    }   

    console.log(finalResults[0].length);
    console.log(finalResults[1].length);
    console.log(finalResults[0].length * (finalResults[1].length + 1));
});