const fs = require('fs');

let results = [];

fs.readFile('./files/day2.txt', (err, data) => {

    if (err) {
        throw err;
    }
    
    const lines = data.toString().split('\n');
    for (let i = 0; i < lines.length; i++) {
        let count = 0;
        const policyValue = lines[i].split(':');
        const numbersLetters = policyValue[0].split(' ');
        const numbers = numbersLetters[0].split('-');

        policyValue[1] = policyValue[1].trim();
        
        const a = policyValue[1][numbers[0] - 1];
        const b = numbersLetters[1];

        const c = policyValue[1][numbers[1] - 1];

        if (a == b) {
            count++;
        }

        if (c == b) {
            count++;
        }

        if (count == 1) {
            results.push(i);
        }
    }

    
    console.log(results.length);
})