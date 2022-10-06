//
//  ViewController.swift
//  ReactiveRiver
//
//  Created by Andrew Arrow on 7/30/22.
//

import Moya
import ReactiveCocoa
import ReactiveSwift
import UIKit

class RootViewController: UIViewController {
    private let welcomeView: WelcomeView

    private let coinMarketCap = MoyaProvider<CoinMarketCap>(endpointClosure: { (target: CoinMarketCap) -> Endpoint in
        let defaultEndpoint = MoyaProvider.defaultEndpointMapping(for: target)
        let apiKey = Bundle.main.infoDictionary?["COIN_MARKET_CAP_API_KEY"] as? String
        return defaultEndpoint.adding(newHTTPHeaderFields: ["X-CMC_PRO_API_KEY": apiKey ?? ""])
    })

    lazy var getCoinCategoriesAction = makeCoinCategoriesAction()

    private var coinCategories = MutableProperty<[CMC_Category]?>(nil)

    init() {
        welcomeView = WelcomeView()
        super.init(nibName: nil, bundle: nil)
        navigationItem.title = "Web3Captain"
    }

    required init?(coder _: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    override func loadView() {
        view = welcomeView
    }

    func makeCoinCategoriesAction() -> Action<Void, Response, MoyaError> {
        return Action { [weak self] in
            guard let self = self else { return .empty }
            return self.coinMarketCap.reactive.request(.categories)
        }
    }

    override func viewDidLoad() {
        super.viewDidLoad()

        let loadingItem = UIBarButtonItem.loading()
        navigationItem.reactive.rightBarButtonItem <~ getCoinCategoriesAction.isExecuting.map { $0 ? loadingItem : nil }

        getCoinCategoriesAction.errors.observe(on: QueueScheduler.main).observeValues { [weak self] error in
            let alert = UIAlertController(title: "HTTP Error", message: error.localizedDescription, preferredStyle: .alert)
            let defaultAction = UIAlertAction(title: "OK", style: .default)
            alert.addAction(defaultAction)
            self?.present(alert, animated: true, completion: nil)
        }

        let cache = Files.readFile(name: "categories.json")
        if cache.count == 0 {
            getCoinCategoriesAction.apply().startWithResult { [weak self] result in
                switch result {
                case let .success(moyaResponse):
                    let data = moyaResponse.data
                    let jsonString = String(decoding: data, as: UTF8.self)
                    Files.writeFile(name: "categories.json", contents: jsonString)
                    // print("111", jsonString)

                    do {
                        let result = try JSONDecoder().decode(CMC_CategoryResults<CMC_Category>.self, from: data)
                        self?.coinCategories.value = result.data
                    } catch { print(error) }

                default:
                    break
                }
            }
        } else {
            let cacheData = cache.data(using: .utf8)
            do {
                let result = try JSONDecoder().decode(CMC_CategoryResults<CMC_Category>.self, from: cacheData!)
                coinCategories.value = result.data
            } catch { print(error) }
        }

        welcomeView.sections <~ coinCategories.map { [weak self]
            categories in
            var sections: [TableMagicSection] = []
            guard let self = self else { return sections }
            sections = self.makeTableSections(categories: categories)

            return sections
        }
    }

    private func makeTableSections(categories: [CMC_Category]?) -> [TableMagicSection] {
        var sections: [TableMagicSection] = []
        guard let categories = categories else {
            return sections
        }

        var rowsForSection1: [TableMagicRow] = []
        print("12222", categories.count)
        for category in categories {
            rowsForSection1.append(TableMagicRow(topText: category.title, bottomText: category.description))
        }

        let section1 = TableMagicSection(rows: rowsForSection1)
        sections.append(section1)

        return sections
    }
}
