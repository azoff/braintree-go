package braintree

import (
	"encoding/xml"
)

type TransactionGateway struct {
	*Braintree
}

// Create initiates a transaction.
func (g *TransactionGateway) Create(tx *Transaction) (*Transaction, error) {
	resp, err := g.execute("POST", "transactions", tx)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// SubmitForSettlement submits the transaction with the specified id for settlement.
// If the amount is omitted, the full amount is settled.
func (g *TransactionGateway) SubmitForSettlement(id string, amount ...float64) (*Transaction, error) {
	var tx *Transaction
	if len(amount) > 0 {
		tx = &Transaction{
			Amount: amount[0],
		}
	}
	resp, err := g.execute("PUT", "transactions/"+id+"/submit_for_settlement", tx)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// Void voids the transaction with the specified id if it has a status of authorized or
// submitted_for_settlement. When the transaction is voided Braintree will do an authorization
// reversal if possible so that the customer won’t have a pending charge on their card
func (g *TransactionGateway) Void(id string) (*Transaction, error) {
	resp, err := g.execute("PUT", "transactions/"+id+"/void", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		fallthrough
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

func (g *TransactionGateway) Refund(id string, amount ...float64) (*Transaction, error) {
	var tx *Transaction
	if len(amount) > 0 {
		tx = &Transaction{
			Amount: amount[0],
		}
	}
	resp, err := g.execute("POST", "transactions/"+id+"/refund", tx)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		fallthrough
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

func (g *TransactionGateway) HoldInEscrow(id string) (*Transaction, error) {
	resp, err := g.execute("PUT", "transactions/"+id+"/hold_in_escrow", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// Find finds the transaction with the specified id.
func (g *TransactionGateway) Find(id string) (*Transaction, error) {
	resp, err := g.execute("GET", "transactions/"+id, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// Release the escrow funds of a transaction.
func (g *TransactionGateway) ReleaseFromEscrow(id string) (*Transaction, error) {
	resp, err := g.execute("PUT", "transactions/"+id+"/release_from_escrow", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// Search finds all transactions matching the search query.
func (g *TransactionGateway) Search(query *SearchQuery) (*TransactionSearchResult, error) {
	resp, err := g.execute("POST", "transactions/advanced_search", query)
	if err != nil {
		return nil, err
	}
	var v TransactionSearchResult
	err = xml.Unmarshal(resp.Body, &v)
	if err != nil {
		return nil, err
	}
	return &v, err
}
