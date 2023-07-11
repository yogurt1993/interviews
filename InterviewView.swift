
import UIKit

class ViewController: UIViewController {
    var output: ProfileViewOutput?
    let resultButton = UILabel()
    
    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    
    override func viewWillAppear(_ animated: Bool) {
        resultButton.topAnchor.constraint(equalTo: view.safeAreaLayoutGuide.topAnchor).isActive = true
        resultButton.centerXAnchor.constraint(equalTo: view.centerXAnchor).isActive = true
        resultButton.centerYAnchor.constraint(equalTo: view.centerYAnchor).isActive = true
        resultButton.widthAnchor.constraint(equalToConstant: 100).isActive = true
        resultButton.heightAnchor.constraint(equalToConstant: 60).isActive = true
        view.addSubview(resultButton)
    }
    
    // MARK: Life cycle
    override func viewDidLoad() {
        super.viewDidLoad()
        output?.viewIsReady()
    }
}

// MARK: ProfileViewInput
extension ViewController: ProfileViewInput {
    func setResult(value: String) {
        resultButton.font = UIFont.systemFont(ofSize: 15)
        resultButton.textColor = .brown
        resultButton.layer.cornerRadius = 15
        resultButton.layer.borderColor = UIColor.black.cgColor
        resultButton.layer.borderWidth = 1
        resultButton.text = value
    }
}

