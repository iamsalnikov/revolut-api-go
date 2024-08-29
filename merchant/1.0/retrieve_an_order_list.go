package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type RetrieveAnOrderListFilter struct {
	FromCreateDate      time.Time
	ToCreatedDate       time.Time
	CustomerId          string
	Email               string
	MerchantOrderExtRef string
	State               []string
}

type RetrieveAnOrderListPagination struct {
	CreatedBefore time.Time
	Limit         int
}

func (c *Merchant) RetrieveAnOrderList(
	filter RetrieveAnOrderListFilter,
	pagination RetrieveAnOrderListPagination,
) (RetrieveAnOrderListResponse, error) {
	params := url.Values{}
	if !filter.FromCreateDate.IsZero() {
		params.Add("from_create_date", filter.FromCreateDate.Format(time.RFC3339))
	}
	if !filter.ToCreatedDate.IsZero() {
		params.Add("to_create_date", filter.ToCreatedDate.Format(time.RFC3339))
	}
	if filter.CustomerId != "" {
		params.Add("customer_id", filter.CustomerId)
	}
	if filter.Email != "" {
		params.Add("email", filter.Email)
	}
	if filter.MerchantOrderExtRef != "" {
		params.Add("merchant_order_ext_ref", filter.MerchantOrderExtRef)
	}
	if filter.State != nil {
		params.Add("state", strings.Join(filter.State, ","))
	}
	if pagination.Limit != 0 {
		params.Add("limit", strconv.Itoa(pagination.Limit))
	}
	if !pagination.CreatedBefore.IsZero() {
		params.Add("created_by_create_date", pagination.CreatedBefore.Format(time.RFC3339))
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "orders"},
		http.MethodGet,
		make([]byte, 0),
		params,
	)
	if err != nil {
		return nil, err
	}

	response := make([]Order, 0)
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
