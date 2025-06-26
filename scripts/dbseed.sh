#!/bin/bash

# dbseed.sh - Seed the crm-service with dummy customers

API_URL="http://localhost:8888/v1/customers"

echo "Seeding dummy customers..."

curl -s -X POST "$API_URL" -H 'accept: application/json' -H 'Content-Type: application/json' -d '{
  "firstName": "John",
  "lastName": "Doe",
  "role": "student",
  "email": "john.doe1@example.com"
}'

echo

curl -s -X POST "$API_URL" -H 'accept: application/json' -H 'Content-Type: application/json' -d '{
  "firstName": "Jane",
  "lastName": "Smith",
  "role": "teacher",
  "email": "jane.smith2@example.com"
}'

echo

curl -s -X POST "$API_URL" -H 'accept: application/json' -H 'Content-Type: application/json' -d '{
  "firstName": "Alice",
  "lastName": "Johnson",
  "role": "manager",
  "email": "alice.johnson3@example.com"
}'

echo

curl -s -X POST "$API_URL" -H 'accept: application/json' -H 'Content-Type: application/json' -d '{
  "firstName": "Bob",
  "lastName": "Brown",
  "role": "developer",
  "email": "bob.brown4@example.com"
}'

echo

echo "Seeding complete."
