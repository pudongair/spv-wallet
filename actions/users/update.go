package users

import (
	"net/http"

	"github.com/bitcoin-sv/spv-wallet/engine"
	"github.com/bitcoin-sv/spv-wallet/engine/spverrors"
	"github.com/bitcoin-sv/spv-wallet/mappings"
	"github.com/bitcoin-sv/spv-wallet/server/reqctx"
	"github.com/gin-gonic/gin"
)

// update will update an existing model
// Update current user information godoc
// @Summary		Update current user information
// @Description	Update current user information
// @Tags		Users
// @Produce		json
// @Param		Metadata body engine.Metadata false " "
// @Success		200 {object} response.Xpub "Updated xPub"
// @Failure		400	"Bad request - Error while parsing Metadata from request body"
// @Failure 	500	"Internal Server Error - Error while updating xPub"
// @Router		/api/v1/users/current [patch]
// @Security	x-auth-xpub
func update(c *gin.Context, userContext *reqctx.UserContext) {
	logger := reqctx.Logger(c)

	var requestBody engine.Metadata
	if err := c.Bind(&requestBody); err != nil {
		spverrors.ErrorResponse(c, spverrors.ErrCannotBindRequest, logger)
		return
	}

	// Get an xPub
	var xPub *engine.Xpub
	var err error
	xPub, err = reqctx.Engine(c).UpdateXpubMetadata(
		c.Request.Context(), userContext.GetXPubID(), requestBody,
	)
	if err != nil {
		spverrors.ErrorResponse(c, err, logger)
		return
	}

	if userContext.GetAuthType() == reqctx.AuthTypeAccessKey {
		xPub.RemovePrivateData()
	}

	contract := mappings.MapToXpubContract(xPub)
	c.JSON(http.StatusOK, contract)
}
