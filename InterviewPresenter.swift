
import UIKit
import Foundation

protocol ProfileViewInput: AnyObject {
    var output: ProfileViewOutput? { get set }
    
    func setResult(value: String)
}

protocol ProfileViewOutput {
    func viewIsReady()
}

enum TypeCalculate: String {
    case sum
    case difference
}

class Presenter: ProfileViewOutput {
    weak var view: ProfileViewInput!
    
    var a = Int()
    let b: Int
    var c = 0
    let typeCalculate: TypeCalculate
    
    init(koef a: Int, b: Int, c: Int, typeCalculate: TypeCalculate) {
        a = a
        b = b
        c = c
        typeCalculate = typeCalculate
    }
    
    private func calculate() -> Int {
        switch typeCalculate {
        case .sum:
            return a + b + c
        case .difference:
            return a - b - c
        }
    }
    
    func resultString() -> String {
        "My result \(calculate()) for type \(typeCalculate.rawValue)"
    }
    
    func viewIsReady() {
        view.setResult(value: resultString())
    }
}
