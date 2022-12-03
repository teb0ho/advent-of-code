//
//  day2.2.swift
//  2022
//
//  Created by Teboho Mphure on 2022/12/04.
//

import Foundation

class Day2_2 : Day2_1 {
    var outcome = ["X": 0, "Y": 3, "Z": 6]
    func GetAmendedStrategyValue() {
        let lines = FileUtils.ReadFile(directory: "repos/advent-of-code/2022/2022/files/day2.txt")
        
        
        for line in lines {
            let pair = line.components(separatedBy: " ")
            
            if pair.count == 2 {
                var answer = 0
                let endResult: Int = outcome[pair[1]]!
                
                for key in myPlays.keys {
                    answer = EvaluateOptions(first: pair[0], second: key)
                    if endResult == answer {
                        total += answer + myPlays[key]!
                        break
                    }
                }
            }
        }
        
        print(total)
    }
}
