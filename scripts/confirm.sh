#!/bin/bash
# Usage: ./scripts/confirm.sh KONG-1785893172544559141

TX_ID=$1

curl -X POST "http://localhost:8080/api/v1/webhooks/payment?provider=manual" \
  -H "Content-Type: application/json" \
  -d "{\"transaction_id\":\"$TX_ID\",\"status\":\"SUCCESS\"}"

echo "✅ Transaksi $TX_ID dikonfirmasi!"
