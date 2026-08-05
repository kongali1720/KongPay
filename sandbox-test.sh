#!/bin/bash

echo "🚀 KONGPAY SANDBOX TEST"
echo "========================"

# 1. Health Check
echo -e "\n📍 1. HEALTH CHECK"
curl -s http://localhost:8080/health | jq

# 2. Create Payment
echo -e "\n📍 2. CREATE PAYMENT"
RESP=$(curl -s -X POST http://localhost:8080/api/v1/payments \
  -H "Content-Type: application/json" \
  -d '{"amount":100000,"currency":"IDR","method":"BANK_TRANSFER"}')
echo $RESP | jq
TX_ID=$(echo $RESP | grep -o '"transaction_id":"[^"]*"' | cut -d'"' -f4)
echo "TX_ID: $TX_ID"

# 3. Process Webhook
echo -e "\n📍 3. PROCESS WEBHOOK"
curl -s -X POST "http://localhost:8080/api/v1/webhooks/payment?provider=bank" \
  -H "Content-Type: application/json" \
  -d "{\"transaction_id\":\"$TX_ID\",\"status\":\"SUCCESS\"}" | jq

# 4. Wait for Settlement
echo -e "\n📍 4. WAITING FOR SETTLEMENT..."
sleep 3

# 5. Check Settlement
echo -e "\n📍 5. CHECK SETTLEMENT"
curl -s "http://localhost:8080/api/v1/settlement/$TX_ID" | jq

# 6. Stats
echo -e "\n📍 6. SETTLEMENT STATS"
curl -s http://localhost:8080/api/v1/settlement/stats | jq

# 7. Database Check
echo -e "\n📍 7. DATABASE CHECK"
psql -h localhost -U kongpay -d kongpay -c "SELECT transaction_id, status FROM settlements ORDER BY created_at DESC LIMIT 3;"

echo -e "\n✅ SANDBOX TEST COMPLETED!"
