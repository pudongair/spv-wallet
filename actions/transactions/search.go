package transactions

import (
	"net/http"

	"github.com/bitcoin-sv/spv-wallet/actions/common"
	"github.com/bitcoin-sv/spv-wallet/engine/spverrors"
	"github.com/bitcoin-sv/spv-wallet/internal/query"
	"github.com/bitcoin-sv/spv-wallet/mappings"
	"github.com/bitcoin-sv/spv-wallet/models/filter"
	"github.com/bitcoin-sv/spv-wallet/models/response"
	"github.com/bitcoin-sv/spv-wallet/server/reqctx"
	"github.com/gin-gonic/gin"
)

// transactions will fetch a list of transactions filtered on conditions and metadata
// Get transactions godoc
// @Summary		Get transactions
// @Description	Get transactions
// @Tags		Transactions
// @Produce		json
// @Param		SwaggerCommonParams query swagger.CommonFilteringQueryParams false "Supports options for pagination and sorting to streamline data exploration and analysis"
// @Param		TransactionParams query filter.TransactionFilter false "Supports targeted resource searches with filters"
// @Success		200 {object} response.PageModel[response.Transaction] "Page of transactions"
// @Failure		400	"Bad request - Error while parsing SearchTransactions from request body"
// @Failure 	500	"Internal server error - Error while searching for transactions"
// @Router		/api/v1/transactions [get]
// @Security	x-auth-xpub
func transactions(c *gin.Context, userContext *reqctx.UserContext) {
	reqXPubID := userContext.GetXPubID()

	searchParams, err := query.ParseSearchParams[filter.TransactionFilter](c)
	if err != nil {
		spverrors.ErrorResponse(c, spverrors.ErrCannotParseQueryParams, reqctx.Logger(c))
		return
	}

	conditions := searchParams.Conditions.ToDbConditions()
	metadata := mappings.MapToMetadata(searchParams.Metadata)
	pageOptions := mappings.MapToDbQueryParams(&searchParams.Page)

	// Record a new transaction (get the hex from parameters)
	transactions, err := reqctx.Engine(c).GetTransactionsByXpubID(
		c.Request.Context(),
		reqXPubID,
		metadata,
		conditions,
		pageOptions,
	)
	if err != nil {
		spverrors.ErrorResponse(c, err, reqctx.Logger(c))
		return
	}

	contracts := make([]*response.Transaction, 0)
	for _, transaction := range transactions {
		contracts = append(contracts, mappings.MapToTransactionContract(transaction))
	}

	count, err := reqctx.Engine(c).GetTransactionsByXpubIDCount(
		c.Request.Context(),
		reqXPubID,
		metadata,
		conditions,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	result := response.PageModel[response.Transaction]{
		Content: contracts,
		Page:    common.GetPageDescriptionFromSearchParams(pageOptions, count),
	}
	c.JSON(http.StatusOK, result)
}
