const fs = require('fs');
       
fs.readFile('./files/day7.txt', (err, data) => {

    if (err) {
        throw err;
    }

    const lines = data.toString().split('\n');
    const answer = [];
    let bagsContainingGold = [];
    let found = true;
    let bagCount = 0;

    for (const line of lines) {
        
        if (line.match(/^shiny gold/g)) {
            let bagsInfo = line.split(' contain ')[1];
            bagCount += +bagsInfo.slice(0, bagsInfo.indexOf(' '));
            answer.push(bagCount);
            let bagColor = bagsInfo.slice(bagsInfo.indexOf(' ') + 1, 
            bagsInfo.indexOf(' bag'));
            bagsContainingGold.push({[bagColor]: bagCount});
        }
    }

    while (found) {
        found = false;
        const newArray = [];
        let total = 0;
        for (const bagName of bagsContainingGold) {
            const keyName = Object.keys(bagName)[0];
            for (const line of lines) {
                let re = new RegExp(`^${Object.keys(bagName)[0]}`, 'g');
                if (line.match(re) || [].length) {
                    found = true;
                    if (!line.includes('no other bags')) {
                        const bagsFound = line.split(' contain ')[1].split(', ');
                        for (const bagFound of bagsFound) {
                            const bagCount = bagName[keyName] * +bagFound.slice(0, bagFound.indexOf(' ', 0));
                            const bagColor = bagFound.slice(bagFound.indexOf(' ') + 1,
                            bagFound.indexOf(' bag'));

                            newArray.push({[bagColor] : bagCount});
                            total += bagCount;
                        }
                    }
                }
            }
        }
        answer.push(total);
        bagsContainingGold = newArray;
    } 
    console.log(answer.reduce((a, b) => a + b));
});

