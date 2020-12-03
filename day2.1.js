const fs = require('fs');

let count = 0;

fs.readFile('./files/day2.txt', (err, data) => {
    const lines = data.toString().split('\n');
    for (const line of lines) {
        const policyValue = line.split(':');
        const numbersLetters = policyValue[0].split(' ');
        
        const numbers = numbersLetters[0].split('-');

        const myPattern = `[${numbersLetters[1]}]`;
        let re = new RegExp(myPattern, "g");
        
        const result = policyValue[1].match(re) || [];


        if (result.length >= +numbers[0] && result.length <= +numbers[1]) {
            count++;
        } 
    }

    
    console.log(count);
})