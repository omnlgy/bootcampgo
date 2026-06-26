#!/usr/bin/env bash
set -e

API_BASE="http://localhost:8080/api"

echo "=========================================="
echo "  AUTH ENDPOINT TESTS"
echo "=========================================="

echo ""
echo "--- 1. Login with valid credentials ---"
curl -s -X POST "$API_BASE/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"alice@example.com","password":"password"}' > /tmp/login_resp.json
python3 -m json.tool /tmp/login_resp.json
TOKEN=*** /tmp/login_resp.json)
echo "  Token received: ${TOKEN:0:20}..."
echo ""

echo "--- 2. Login with wrong password ---"
curl -s -X POST "$API_BASE/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"alice@example.com","password":"wrong"}' | python3 -m json.tool
echo ""

echo "--- 3. Login with empty body ---"
curl -s -X POST "$API_BASE/login" \
  -H "Content-Type: application/json" \
  -d '{}' | python3 -m json.tool
echo ""

echo "--- 4. Login with non-existent email ---"
curl -s -X POST "$API_BASE/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"nobody@test.com","password":"password"}' | python3 -m json.tool
echo ""

echo "--- 5. Logout with valid token ---"
TOKEN=*** /tmp/login_resp.json)
curl -s -X POST "$API_BASE/logout" \
  -H "Authorization: Bearer *** | python3 -m json.tool
echo ""

echo "--- 6. Logout without auth header ---"
curl -s -X POST "$API_BASE/logout" | python3 -m json.tool
echo ""

echo "--- 7. Logout with invalid auth format ---"
curl -s -X POST "$API_BASE/logout" \
  -H "Authorization: *** | python3 -m json.tool
echo ""

echo "=========================================="
echo "  DONE"
echo "=========================================="
