package admin

import (
	"net/http"

	"github.com/bitcoin-sv/spv-wallet/engine/spverrors"
	"github.com/bitcoin-sv/spv-wallet/mappings"
	"github.com/bitcoin-sv/spv-wallet/models"
	"github.com/bitcoin-sv/spv-wallet/server/reqctx"
	"github.com/gin-gonic/gin"
)

// subscribeWebhook will subscribe to a webhook to receive notifications
// @Summary		Subscribe to a webhook
// @Description	Subscribe to a webhook to receive notifications
// @Tags		Admin
// @Produce		json
// @Param		SubscribeRequestBody body models.SubscribeRequestBody false "URL to subscribe to and optional token header and value"
// @Success		200 {boolean} bool "Success response"
// @Failure 	500	"Internal server error - Error while subscribing to the webhook"
// @Router		/api/v1/admin/webhooks/subscriptions [post]
// @Security	x-auth-xpub
func subscribeWebhook(c *gin.Context, _ *reqctx.AdminContext) {
	logger := reqctx.Logger(c)
	requestBody := models.SubscribeRequestBody{}
	if err := c.Bind(&requestBody); err != nil {
		spverrors.ErrorResponse(c, spverrors.ErrCannotBindRequest.WithTrace(err), logger)
		return
	}

	err := reqctx.Engine(c).SubscribeWebhook(c.Request.Context(), requestBody.URL, requestBody.TokenHeader, requestBody.TokenValue)
	if err != nil {
		spverrors.ErrorResponse(c, spverrors.ErrWebhookSubscriptionFailed.WithTrace(err), logger)
		return
	}

	c.JSON(http.StatusOK, true)
}

// unsubscribeWebhook will unsubscribe to a webhook to receive notifications
// @Summary		Unsubscribe to a webhook
// @Description	Unsubscribe to a webhook to stop receiving notifications
// @Tags		Admin
// @Produce		json
// @Param		UnsubscribeRequestBody body models.UnsubscribeRequestBody false "URL to unsubscribe from"
// @Success		200
// @Failure 	500	"Internal server error - Error while unsubscribing to the webhook"
// @Router		/api/v1/admin/webhooks/subscriptions [delete]
// @Security	x-auth-xpub
func unsubscribeWebhook(c *gin.Context, _ *reqctx.AdminContext) {
	logger := reqctx.Logger(c)
	requestModel := models.UnsubscribeRequestBody{}
	if err := c.Bind(&requestModel); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := reqctx.Engine(c).UnsubscribeWebhook(c.Request.Context(), requestModel.URL); err != nil {
		spverrors.ErrorResponse(c, spverrors.ErrWebhookUnsubscriptionFailed.WithTrace(err), logger)
		return
	}

	c.Status(http.StatusOK)
}

// getAllWebhooks will return all the stored webhooks
// @Summary		Get All Webhooks
// @Description	Get All Webhooks currently subscribed to
// @Tags		Admin
// @Produce		json
// @Success		200 {object} []models.Webhook "List of webhooks"
// @Failure 	500	"Internal server error - Error while getting all webhooks"
// @Router		/api/v1/admin/webhooks/subscriptions [get]
// @Security	x-auth-xpub
func getAllWebhooks(c *gin.Context, _ *reqctx.AdminContext) {
	wh, err := reqctx.Engine(c).GetWebhooks(c.Request.Context())
	if err != nil {
		spverrors.ErrorResponse(c, err, reqctx.Logger(c))
		return
	}

	webhookDTOs := make([]*models.Webhook, len(wh))
	for i, w := range wh {
		webhookDTOs[i] = mappings.MapToWebhookContract(w)
	}

	c.JSON(http.StatusOK, webhookDTOs)
}
