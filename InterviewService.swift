
import Foundation

enum Mode {
    case preprod
    case prod
}

enum TypeService {
    case Network
    case appSetting
    case mock
    case calculate
}

enum Flow {
    case firstFlow, secondFlow, thirdFlow
}

protocol NewServiceInput {
    func getType() -> Any?
    func calculateValue()
}

final class NewService: NSObject {
    var typeService: TypeService
    var mode: Mode
    
    //Get data from database
    let model: [Model] = [Model(id: 1, numberCard: 123, company: "", state: ""),
                          Model(id: 2, numberCard: 123, company: "", state: ""),
                          Model(id: 3, numberCard: 456, company: "", state: ""),
                          Model(id: 1, numberCard: 321, company: "", state: ""),
                          Model(id: 2, numberCard: 123, company: "", state: "")]
    
    init(type: TypeService, mode: Mode) {
        typeService = type
        self.mode = mode
    }
    
    func getFlow() -> Any? {
        if typeService == .appSetting {
            if mode == .preprod {
                return Flow.firstFlow
            } else {
                return Flow.secondFlow
            }
        } else if typeService == .mock {
            return Flow.thirdFlow
        } else {
            return nil
        }
    }
    
    func getState(for accounrId: Int, number: Int) -> String {
        for item in model {
            if item.id == accounrId {
                if item.company == number {
                    return item.state!
                }
            }
        }
    }
}

struct Model {
    let id: Int
    let numberCard: Int
    let company: String?
    let state: String?
}
