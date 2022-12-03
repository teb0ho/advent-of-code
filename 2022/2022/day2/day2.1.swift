//
//  day2.1.swift
//  2022
//
//  Created by Teboho Mphure on 2022/12/02.
//

import Foundation


class Day2_1 {
    var myPlays = ["X": 1, "Y": 2, "Z": 3]
    var total = 0
    
    func GetStrategyValue() {
        let lines = FileUtils.ReadFile(directory: "repos/advent-of-code/2022/2022/files/day2.txt")
        
        for line in lines {
            let pair = line.components(separatedBy: " ")
            
            if pair.count == 2 {
                var answer = 0
                answer = EvaluateOptions(first: pair[0], second: pair[1])
                total += answer + myPlays[pair[1]]!
//                switch(pair[0]) {
//                case "A":
//                    answer = EvaluateOptions(first: pair[0], second: "Y")
//                    total += answer + myPlays["Y"]!
//                case "B":
//                    answer = EvaluateOptions(first: pair[0], second: "X")
//                    total += answer + myPlays["X"]!
//                default:
//                    answer = EvaluateOptions(first: pair[0], second: "Z")
//                    total += answer + myPlays["Z"]!
//                }
        
                
            }
        }
        
        print(total)
    }
    
    func EvaluateOptions(first: String, second: String) -> Int {
        var result = 0
        
        if first == "A" { // rock (opp)
            switch second {
            case "X": // rock
                result = 3
            case "Y": // paper
                result = 6
            default: // scissors
                result = 0
            }
        } else if first == "B" { // paper (opp)
            switch second {
            case "X": // rock
                result = 6
            case "Y": // paper
                result = 3
            default: // scissors
                result = 0
            }
        } else { // scissor
            switch second {
            case "X": // rock
                result = 6
            case "Y": // paper
                result = 0
            default: // scissor
                result = 3
            }
        }
        
        return result
    }
}
