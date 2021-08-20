const fs = require('fs');
       
fs.readFile('./files/day6.txt', (err, data) => {

    if (err) {
        throw err;
    }

    const lines = data.toString().split('\n\n');
    const answeredYes = [];

    for (const line of lines) {
        let answers = line.split('\n');
        const uniqueAnswers = [];
        let count = 0;

        for (const answer of answers) {
            for (let i = 0; i < answer.length; i++) {
                if(!uniqueAnswers.includes(answer[i])) {
                    uniqueAnswers.push(answer[i]);
                }       
            }
        }
        
        uniqueAnswers.map(a => {
            if (answers.every(p => p.includes(a))) {
                count++;
            }
        });
        answeredYes.push(count);
    }

    const finalAnswer = answeredYes.reduce((a, b) => a + b);
    console.log(finalAnswer); 
});
