Example usages:

```go
package main

import (
	"fmt"
	"log"

	"github.com/267H/orderedquery"
)

func main() {
	testOrderedQuery()
	testSpecialQueryCases()
	fmt.Println("\nOrderedQuery tests passed successfully!")
}

func testOrderedQuery() {
	fmt.Println("\nTesting OrderedQuery:")

	query := orderedquery.NewQuery(5)
	query.Add("page", "1")
	query.Add("sort", "desc")
	query.Add("filter", "active")
	query.Add("view", "detailed")
	query.Add("limit", "50")

	queryString := query.Encode()
	fmt.Printf("Basic Query string: %s\n", queryString)
	expectedQuery := "page=1&sort=desc&filter=active&view=detailed&limit=50"
	assertEqual(queryString, expectedQuery, "Basic Query")

	complexQuery := orderedquery.NewQuery(4)
	complexQuery.Add("search", "test product")
	complexQuery.Add("category", "electronics/phones")
	complexQuery.Add("price_range", "100-500")
	complexQuery.Add("brand[]", "apple,samsung,google")

	complexQueryString := complexQuery.Encode()
	fmt.Printf("Complex Query string: %s\n", complexQueryString)
	expectedComplexQuery := "search=test+product&category=electronics%2Fphones&price_range=100-500&brand%5B%5D=apple%2Csamsung%2Cgoogle"
	assertEqual(complexQueryString, expectedComplexQuery, "Complex Query")
	fmt.Println("✓ Query strings verified")
}

func testSpecialQueryCases() {
	fmt.Println("\nTesting Special Query Cases:")

	emptyQuery, err := orderedquery.ParseQuery("")
	if err != nil {
		log.Fatalf("Failed to parse empty query: %v", err)
	}
	assertEqual(emptyQuery.Encode(), "", "Empty Query")

	prefixQuery, err := orderedquery.ParseQuery("?key=value&special=a/b/c")
	if err != nil {
		log.Fatalf("Failed to parse query with prefix: %v", err)
	}
	assertEqual(prefixQuery.Encode(), "key=value&special=a%2Fb%2Fc", "Prefixed Query")
	fmt.Println("✓ Special query cases handled correctly")
}

func assertEqual(got, expected, testName string) {
	if got != expected {
		log.Fatalf("\n%s test failed.\nExpected: %s\nGot: %s", testName, expected, got)
	}
}
