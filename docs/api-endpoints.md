# API 端点状态跟踪

> 自动同步用。每行 = 一个 API 端点 + Go 函数 + Go 文件 + 测试状态。
> 规则：如果端点已在 Lifecycle 测试中调用，标注 lifecycle 测试名，不单独写测试。

## Payment Acceptance (pa)

### 支付尝试与争议
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `GET /api/v1/pa/payment_attempts/{id}` | GetPaymentAttempt | payment_attempts.go | TestPaymentAttemptLifecycle | ✅ PASS |
| `GET /api/v1/pa/payment_attempts` | ListPaymentAttempts | payment_attempts.go | TestPaymentAttemptLifecycle | ✅ PASS |
| `GET /api/v1/pa/payment_disputes/{id}` | GetPaymentDispute | payment_disputes.go | TestPaymentDisputeLifecycle | ✅ PASS |
| `GET /api/v1/pa/payment_disputes` | ListPaymentDisputes | payment_disputes.go | TestPaymentDisputeLifecycle | ✅ PASS |

### 配置
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `GET /api/v1/pa/config/payment_method_types` | GetPaymentMethodTypes | config.go | TestConfigLifecycle | ✅ PASS |
| `GET /api/v1/pa/config/banks` | GetBanks | config.go | TestConfigLifecycle | ✅ PASS |
| `GET /api/v1/pa/config/convertible_shopper_currencies` | GetConvertibleShopperCurrencies | config.go | TestConfigLifecycle | ✅ PASS |
| `GET /api/v1/pa/config/reserve_plan` | GetReservePlan | config.go | TestConfigLifecycle | ✅ PASS |
| `GET /api/v1/pa/config/applepay/registered_domains` | GetApplePayDomains | config.go | TestConfigLifecycle | ✅ PASS |
| `POST /api/v1/pa/config/applepay/registered_domains/add_items` | AddApplePayDomains | config.go | TestConfigLifecycle | 🔒 400 |
| `POST /api/v1/pa/config/applepay/registered_domains/remove_items` | RemoveApplePayDomains | config.go | TestConfigLifecycle | ✅ PASS |

### 配置与报价
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/pa/conversion_quotes/create` | CreateConversionQuote | conversion_quotes.go | TestConversionQuoteLifecycle | ✅ PASS |
| `GET /api/v1/pa/conversion_quotes/{id}` | GetConversionQuote | conversion_quotes.go | TestConversionQuoteLifecycle | ✅ PASS |

### 客户管理
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/pa/customers/create` | CreateCustomer | customer.go | TestCustomerLifecycle | ✅ PASS |
| `GET /api/v1/pa/customers/{id}` | GetCustomer | customer.go | TestCustomerLifecycle | ✅ PASS |
| `POST /api/v1/pa/customers/{id}/update` | UpdateCustomer | customer.go | TestCustomerLifecycle | ✅ PASS |
| `GET /api/v1/pa/customers` | ListCustomers | customer.go | TestCustomerLifecycle | ✅ PASS |

### 参考数据
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `GET /api/v1/pa/reference/bin/lookup` | LookupBin | pa_reference_data.go | — | 🚫 |

### 终端 (POS)
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/pa/pos/terminals/create` | CreateTerminal | terminals.go | TestTerminalLifecycle | 🔒 400 |
| `GET /api/v1/pa/pos/terminals/{id}` | GetTerminal | terminals.go | TestTerminalLifecycle | ✅ PASS |
| `POST /api/v1/pa/pos/terminals/{id}/update` | UpdateTerminal | terminals.go | TestTerminalLifecycle | ✅ PASS |
| `GET /api/v1/pa/pos/terminals` | ListTerminals | terminals.go | TestTerminalLifecycle | ✅ PASS |
| `POST /api/v1/pa/pos/terminals/{id}/activate` | ActivateTerminal | terminals.go | TestTerminalLifecycle | 🔒 400 |
| `POST /api/v1/pa/pos/terminals/{id}/deactivate` | DeactivateTerminal | terminals.go | TestTerminalLifecycle | 🔒 400 |
| `POST /api/v1/pa/pos/terminals/{id}/terminate` | TerminateTerminal | terminals.go | TestTerminalLifecycle | 🔒 400 |
| `POST /api/v1/pa/pos/terminals/{id}/reset_password` | ResetTerminalPassword | terminals.go | TestTerminalLifecycle | 🔒 400 |
| `POST /api/v1/pa/pos/terminals/{id}/cancel_current_operation` | CancelCurrentOperation | terminals.go | TestTerminalLifecycle | 🔒 400 |
| `POST /api/v1/pa/pos/terminals/{id}/process_payment_intent` | ProcessPaymentIntentInTerminal | terminals.go | TestTerminalLifecycle | 🔒 400 |

### 支付链接
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/pa/payment_links/create` | CreatePaymentLink | payment_link.go | TestPaymentLinkLifecycle | ✅ PASS |
| `GET /api/v1/pa/payment_links/{id}` | GetPaymentLink | payment_link.go | TestPaymentLinkLifecycle | ✅ PASS |
| `GET /api/v1/pa/payment_links` | ListPaymentLinks | payment_link.go | TestPaymentLinkLifecycle | ✅ PASS |

### 退款与支付方式
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/pa/refunds/create` | CreateRefund | refund.go | TestRefundLifecycle | ✅ PASS |
| `GET /api/v1/pa/refunds/{id}` | GetRefund | refund.go | TestRefundLifecycle | ✅ PASS |
| `GET /api/v1/pa/refunds` | ListRefunds | refund.go | TestRefundLifecycle | ✅ PASS |
| `POST /api/v1/pa/payment_methods/create` | CreatePaymentMethod | payment_method.go | TestPaymentMethodLifecycle | ✅ PASS |
| `GET /api/v1/pa/payment_methods/{id}` | GetPaymentMethod | payment_method.go | TestPaymentMethodLifecycle | ✅ PASS |
| `GET /api/v1/pa/payment_methods` | ListPaymentMethods | payment_method.go | TestPaymentMethodLifecycle | ✅ PASS |
| `POST /api/v1/pa/payment_methods/{id}/disable` | DisablePaymentMethod | payment_method.go | TestPaymentMethodLifecycle | ✅ PASS |
| `POST /api/v1/pa/payment_methods/{id}/update` | UpdatePaymentMethod | payment_method.go | TestPaymentMethodLifecycle | ✅ PASS |

### 核心支付
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/pa/payment_intents/create` | CreatePaymentIntent | payment_intent.go | TestPaymentIntentLifecycle | ✅ PASS |
| `GET /api/v1/pa/payment_intents/{id}` | GetPaymentIntent | payment_intent.go | TestPaymentIntentLifecycle | ✅ PASS |
| `POST /api/v1/pa/payment_intents/{id}/update` | UpdatePaymentIntent | payment_intent.go | TestPaymentIntentLifecycle | ✅ PASS |
| `POST /api/v1/pa/payment_intents/{id}/cancel` | CancelPaymentIntent | payment_intent.go | TestPaymentIntentLifecycle | ✅ PASS |
| `GET /api/v1/pa/payment_intents` | ListPaymentIntents | payment_intent.go | TestPaymentIntentLifecycle | ✅ PASS |
| `POST /api/v1/pa/payment_intents/{id}/confirm` | ConfirmPaymentIntent | payment_intent.go | TestPaymentIntentLifecycle | ✅ PASS |
| `POST /api/v1/pa/payment_intents/{id}/capture` | CapturePaymentIntent | payment_intent.go | TestPaymentIntentLifecycle | ✅ PASS |

### 海关申报与资金拆分
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/pa/customs_declarations/create` | CreateCustomsDeclaration | customs_declarations.go | TestCustomsDeclarationLifecycle | 🔒 400 |
| `GET /api/v1/pa/customs_declarations/{id}` | GetCustomsDeclaration | customs_declarations.go | TestCustomsDeclarationLifecycle | 🔒 404 |
| `GET /api/v1/pa/customs_declarations` | ListCustomsDeclarations | customs_declarations.go | TestCustomsDeclarationLifecycle | 🔒 404 |
| `POST /api/v1/pa/customs_declarations/{id}/update` | UpdateCustomsDeclaration | customs_declarations.go | TestCustomsDeclarationLifecycle | 🔒 404 |
| `POST /api/v1/pa/customs_declarations/{id}/redeclare` | RedeclareCustomsDeclaration | customs_declarations.go | TestCustomsDeclarationLifecycle | 🔒 404 |
| `POST /api/v1/pa/funds_splits/create` | CreateFundsSplit | funds_splits.go | TestFundsSplitLifecycle | 🔒 400 |
| `GET /api/v1/pa/funds_splits/{id}` | GetFundsSplit | funds_splits.go | TestFundsSplitLifecycle | 🔒 404 |
| `GET /api/v1/pa/funds_splits` | ListFundsSplits | funds_splits.go | TestFundsSplitLifecycle | 🔒 404 |
| `POST /api/v1/pa/funds_split_reversals/create` | CreateFundsSplitReversal | funds_split_reversals.go | TestFundsSplitReversalLifecycle | 🔒 400 |
| `GET /api/v1/pa/funds_split_reversals/{id}` | GetFundsSplitReversal | funds_split_reversals.go | TestFundsSplitReversalLifecycle | 🔒 404 |
| `GET /api/v1/pa/funds_split_reversals` | ListFundsSplitReversals | funds_split_reversals.go | TestFundsSplitReversalLifecycle | 🔒 404 |

### 支付同意
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/pa/payment_consents/create` | CreatePaymentConsent | payment_consents.go | TestPaymentConsentLifecycle | ✅ PASS |
| `GET /api/v1/pa/payment_consents/{id}` | GetPaymentConsent | payment_consents.go | TestPaymentConsentLifecycle | ✅ PASS |
| `POST /api/v1/pa/payment_consents/{id}/cancel` | CancelPaymentConsent | payment_consents.go | TestPaymentConsentLifecycle | 🔒 404 |
| `GET /api/v1/pa/payment_consents` | ListPaymentConsents | payment_consents.go | TestPaymentConsentLifecycle | ✅ PASS |

### 结算
| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `GET /api/v1/pa/settlement_records/{id}` | GetSettlementRecord | settlement_records.go | TestSettlementRecordLifecycle | ✅ PASS |
| `GET /api/v1/pa/settlement_records` | ListSettlementRecords | settlement_records.go | TestSettlementRecordLifecycle | ✅ PASS |

---

## Core Resources (core)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `GET /api/v1/balances/current` | GetCurrentBalances | balance.go | TestGetCurrentBalances | ✅ PASS |
| `GET /api/v1/balances/history` | GetBalanceHistory | balance.go | TestBalanceHistory | ✅ PASS |
| `POST /api/v1/deposits/create` | CreateDeposit | deposit.go | TestDepositLifecycle | ✅ PASS |
| `GET /api/v1/deposits/{id}` | GetDeposit | deposit.go | TestDepositLifecycle | ✅ PASS |
| `GET /api/v1/deposits` | ListDeposits | deposit.go | TestDepositLifecycle | ✅ PASS |
| `GET /api/v1/direct_debits/{id}` | GetDirectDebit | direct_debits.go | TestDirectDebitsListAndGet | ✅ PASS |
| `GET /api/v1/direct_debits` | ListDirectDebits | direct_debits.go | TestDirectDebitsListAndGet | ✅ PASS |
| `POST /api/v1/direct_debits/{transaction_id}/cancel` | CancelDirectDebit | direct_debits.go | TestDirectDebitsListAndGet | ✅ PASS |
| `POST /api/v1/global_accounts/create` | CreateGlobalAccount | global_account.go | TestGlobalAccountLifecycle | ✅ PASS |
| `POST /api/v1/global_accounts/{id}/close` | DeleteGlobalAccount | global_account.go | TestGlobalAccountLifecycle | ✅ PASS |
| `GET /api/v1/global_accounts/{id}` | GetGlobalAccount | global_account.go | TestGlobalAccountLifecycle | ✅ PASS |
| `GET /api/v1/global_accounts/{id}/bank_details` | GetGlobalAccountBankDetails | global_account.go | TestGlobalAccountLifecycle | ✅ PASS |
| `GET /api/v1/global_accounts/{id}/transactions` | GetGlobalAccountTransactions | global_account.go | TestGlobalAccountLifecycle | ✅ PASS |
| `GET /api/v1/global_accounts` | ListGlobalAccounts | global_account.go | TestGlobalAccountLifecycle | ✅ PASS |
| `POST /api/v1/global_accounts/{id}/update` | UpdateGlobalAccount | global_account.go | TestGlobalAccountLifecycle | ✅ PASS |
| `POST /api/v1/linked_accounts/create` | CreateLinkedAccount | linked_account.go | TestLinkedAccountLifecycle | ✅ PASS |
| `GET /api/v1/linked_accounts/{id}` | GetLinkedAccount | linked_account.go | TestLinkedAccountLifecycle | ✅ PASS |
| `GET /api/v1/linked_accounts` | ListLinkedAccounts | linked_account.go | TestLinkedAccountLifecycle | ✅ PASS |
---

## Payouts (payouts)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/batch_transfers/create` | CreateBatchTransfer | batch_transfer.go | TestBatchTransferLifecycle | ✅ PASS |
| `GET /api/v1/batch_transfers/{id}` | GetBatchTransfer | batch_transfer.go | TestBatchTransferLifecycle | ✅ PASS |
| `GET /api/v1/batch_transfers` | ListBatchTransfers | batch_transfer.go | TestBatchTransferLifecycle | ✅ PASS |
| `POST /api/v1/batch_transfers/{id}/delete` | DeleteBatchTransfer | batch_transfer.go | TestBatchTransferLifecycle | ✅ PASS |
| `POST /api/v1/batch_transfers/{id}/add_items` | AddBatchTransferItems | batch_transfer.go | TestBatchTransferLifecycle | ✅ 200 |
| `POST /api/v1/batch_transfers/{id}/delete_items` | DeleteBatchTransferItems | batch_transfer.go | TestBatchTransferLifecycle | ✅ 200 |
| `GET /api/v1/batch_transfers/{id}/items` | ListBatchTransferItems | batch_transfer.go | TestBatchTransferLifecycle | ✅ PASS |
| `POST /api/v1/batch_transfers/{id}/quote` | QuoteBatchTransfer | batch_transfer.go | TestBatchTransferLifecycle | ✅ 200 |
| `POST /api/v1/batch_transfers/{id}/submit` | SubmitBatchTransfer | batch_transfer.go | TestBatchTransferLifecycle | 🔒 400 |
| `POST /api/v1/beneficiaries/create` | CreateBeneficiary | beneficiary.go | TestBeneficiaryLifecycle | ✅ PASS |
| `POST /api/v1/beneficiaries/{id}/delete` | DeleteBeneficiary | beneficiary.go | TestBeneficiaryLifecycle | ✅ PASS |
| `GET /api/v1/beneficiaries/{id}` | GetBeneficiary | beneficiary.go | TestBeneficiaryLifecycle | ✅ PASS |
| `GET /api/v1/beneficiaries` | ListBeneficiaries | beneficiary.go | TestBeneficiaryLifecycle | ✅ PASS |
| `POST /api/v1/beneficiaries/{id}/update` | UpdateBeneficiary | beneficiary.go | TestBeneficiaryLifecycle | ✅ PASS |
| `POST /api/v1/beneficiaries/validate` | ValidateBeneficiary | beneficiary.go | TestBeneficiaryLifecycle | ✅ PASS |
| `POST /api/v1/beneficiaries/verify_account` | VerifyAccount | beneficiary.go | TestBeneficiaryLifecycle | 🔒 400 |
| `POST /api/v1/beneficiaries/generate_api_schema` | GenerateAPISchema | beneficiary.go | TestBeneficiaryLifecycle | 🔒 405 |
| `POST /api/v1/beneficiaries/generate_form_schema` | GenerateFormSchema | beneficiary.go | TestBeneficiaryLifecycle | 🔒 405 |
| `GET /api/v1/beneficiaries/supported_financial_institutions` | GetSupportedFinancialInstitutions | beneficiary.go | TestBeneficiaryLifecycle | 🔒 400 |
| `POST /api/v1/transfers/{id}/cancel` | CancelTransfer | transfer.go | TestTransferLifecycle | 🔒 400 |
| `POST /api/v1/transfers/create` | CreateTransfer | transfer.go | TestTransferLifecycle | ✅ PASS |
| `GET /api/v1/transfers/{id}` | GetTransfer | transfer.go | TestTransferLifecycle | ✅ PASS |
| `GET /api/v1/transfers` | ListTransfers | transfer.go | TestTransferLifecycle | ✅ PASS |
| `POST /api/v1/transfers/validate` | ValidateTransfer | transfer.go | TestTransferLifecycle | ✅ 200 |
| `POST /api/v1/transfers/{id}/confirm_funding` | ConfirmTransferFunding | transfer.go | TestTransferLifecycle | 🔒 400 |
---

## Transactional FX (fx)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/fx/conversions/create` | CreateConversion | conversion.go | fx_test.go | ✅ PASS |
| `GET /api/v1/fx/conversions/{id}` | GetConversion | conversion.go | fx_test.go | ✅ PASS |
| `GET /api/v1/fx/conversions` | ListConversions | conversion.go | fx_test.go | ✅ PASS |
| `POST /api/v1/fx/conversion_amendments/create` | CreateConversionAmendment | conversion_amendment.go | TestConversionAmendmentLifecycle | ✅ PASS |
| `POST /api/v1/fx/quotes/create` | CreateQuote | quote.go | fx_test.go | ✅ PASS |
| `GET /api/v1/fx/quotes/{id}` | GetQuote | quote.go | fx_test.go | ✅ PASS |
| `GET /api/v1/fx/rates/current` | GetRates | rate.go | fx_test.go | ✅ PASS |
---

## Issuing (issuing)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `GET /api/v1/issuing/authorizations/{id}` | GetAuthorization | authorization.go | TestAuthorizationLifecycle | 🔒 404 |
| `GET /api/v1/issuing/authorizations` | ListAuthorizations | authorization.go | TestAuthorizationLifecycle | ✅ PASS |
| `POST /api/v1/issuing/cards/create` | CreateCard | card.go | TestCardLifecycle | 🔒 403 |
| `GET /api/v1/issuing/cards/{id}` | GetCard | card.go | TestCardLifecycle | 🔒 403 |
| `GET /api/v1/issuing/cards` | ListCards | card.go | TestCardLifecycle | ✅ PASS |
| `POST /api/v1/issuing/cards/{id}/update` | UpdateCard | card.go | TestCardLifecycle | 🔒 403 |
| `POST /api/v1/issuing/cards/{id}/activate` | ActivateCard | card.go | TestCardLifecycle | 🔒 403 |
| `GET /api/v1/issuing/cards/{id}/details` | GetCardDetails | card.go | TestCardLifecycle | 🔒 403 |
| `GET /api/v1/issuing/cards/{id}/limits` | GetCardLimits | card.go | TestCardLifecycle | 🔒 403 |
| `GET /api/v1/issuing/card_transaction_events/{id}` | GetCardTransactionEvent | card_transaction_events.go | TestCardTransactionEventLifecycle | 🔒 404 |
| `GET /api/v1/issuing/card_transaction_events` | ListCardTransactionEvents | card_transaction_events.go | TestCardTransactionEventLifecycle | ✅ PASS |
| `GET /api/v1/issuing/lifecycles/{id}` | GetLifecycle | card_transaction_lifecycles.go | TestCardTransactionLifecycleLifecycle | 🔒 404 |
| `GET /api/v1/issuing/lifecycles` | ListLifecycles | card_transaction_lifecycles.go | TestCardTransactionLifecycleLifecycle | ✅ PASS |
| `POST /api/v1/issuing/cardholders/create` | CreateCardholder | cardholder.go | TestCardholderLifecycle | ✅ PASS |
| `GET /api/v1/issuing/cardholders/{id}` | GetCardholder | cardholder.go | TestCardholderLifecycle | ✅ PASS |
| `GET /api/v1/issuing/cardholders` | ListCardholders | cardholder.go | TestCardholderLifecycle | ✅ PASS |
| `POST /api/v1/issuing/cardholders/{id}/update` | UpdateCardholder | cardholder.go | TestCardholderLifecycle | 🔒 400 |
| `POST /api/v1/issuing/cardholders/{id}/delete` | DeleteCardholder | cardholder.go | TestDeleteCardholder | ✅ PASS |
| `GET /api/v1/issuing/config` | GetConfig | config.go | TestConfig | ✅ PASS |
| `POST /api/v1/issuing/config/update` | UpdateConfig | config.go | TestConfig | ✅ PASS |
| `GET /api/v1/issuing/digital_wallet_tokens/{id}` | GetDigitalWalletToken | digital_wallet_tokens.go | TestDigitalWalletTokenLifecycle | 🔒 404 |
| `GET /api/v1/issuing/digital_wallet_tokens` | ListDigitalWalletTokens | digital_wallet_tokens.go | TestDigitalWalletTokenLifecycle | ✅ PASS |
| `GET /api/v1/issuing/merchant_brands/{id}` | GetMerchantBrand | merchant_brands.go | TestMerchantBrandLifecycle | ✅ PASS |
| `GET /api/v1/issuing/merchant_brands` | ListMerchantBrands | merchant_brands.go | TestMerchantBrandLifecycle | ✅ PASS |
| `GET /api/v1/issuing/transactions/{id}` | GetTransaction | transaction.go | TestTransactionLifecycle | 🔒 404 |
| `GET /api/v1/issuing/transactions` | ListTransactions | transaction.go | TestTransactionLifecycle | ✅ PASS |
| `POST /api/v1/issuing/transaction_disputes/create` | CreateTransactionDispute | transaction_disputes.go | TestTransactionDisputeLifecycle | 🔒 404 |
| `GET /api/v1/issuing/transaction_disputes/{id}` | GetTransactionDispute | transaction_disputes.go | TestTransactionDisputeLifecycle | 🔒 404 |
| `GET /api/v1/issuing/transaction_disputes` | ListTransactionDisputes | transaction_disputes.go | TestTransactionDisputeLifecycle | ✅ PASS |
| `POST /api/v1/issuing/transaction_disputes/{id}/update` | UpdateTransactionDispute | transaction_disputes.go | TestTransactionDisputeLifecycle | 🔒 404 |
| `POST /api/v1/issuing/transaction_disputes/{id}/cancel` | CancelTransactionDispute | transaction_disputes.go | TestTransactionDisputeLifecycle | 🔒 404 |
| `POST /api/v1/issuing/transaction_disputes/{id}/submit` | SubmitTransactionDispute | transaction_disputes.go | TestTransactionDisputeLifecycle | 🔒 404 |
---

## Billing (billing)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/billing_customers/create` | CreateBillingCustomer | billing_customers.go | TestBillingCustomerLifecycle | ✅ PASS |
| `GET /api/v1/billing_customers/{id}` | GetBillingCustomer | billing_customers.go | TestBillingCustomerLifecycle | ✅ PASS |
| `GET /api/v1/billing_customers` | ListBillingCustomers | billing_customers.go | TestBillingCustomerLifecycle | ✅ PASS |
| `POST /api/v1/billing_customers/{id}/update` | UpdateBillingCustomer | billing_customers.go | TestBillingCustomerLifecycle | ✅ PASS |
| `GET /api/v1/billing_transactions/{id}` | GetBillingTransaction | billing_transactions.go | TestBillingTransactionLifecycle | ✅ PASS |
| `GET /api/v1/billing_transactions` | ListBillingTransactions | billing_transactions.go | TestBillingTransactionLifecycle | ✅ PASS |
| `POST /api/v1/billing_checkouts/{id}/cancel` | CancelCheckout | checkouts.go | TestCheckoutOperations | ✅ PASS |
| `POST /api/v1/billing_checkouts/create` | CreateCheckout | checkouts.go | TestCheckoutLifecycle | ✅ PASS |
| `GET /api/v1/billing_checkouts/{id}` | GetCheckout | checkouts.go | TestCheckoutLifecycle | ✅ PASS |
| `GET /api/v1/billing_checkouts` | ListCheckouts | checkouts.go | TestCheckoutLifecycle | ✅ PASS |
| `POST /api/v1/billing_checkouts/{id}/update` | UpdateCheckout | checkouts.go | TestCheckoutOperations | ✅ PASS |
| `POST /api/v1/coupons/create` | CreateCoupon | coupons.go | TestCouponLifecycle | 🔒 401 |
| `GET /api/v1/coupons/{id}` | GetCoupon | coupons.go | TestCouponLifecycle | 🔒 401 |
| `GET /api/v1/coupons` | ListCoupons | coupons.go | TestCouponLifecycle | 🔒 401 |
| `POST /api/v1/coupons/{id}/update` | UpdateCoupon | coupons.go | TestCouponLifecycle | 🔒 401 |
| `POST /api/v1/billing/credit_notes/{id}/add_line_items` | AddCreditNoteLineItems | credit_notes.go | TestCreditNoteOperations | 🔒 401 |
| `POST /api/v1/billing/credit_notes/create` | CreateCreditNote | credit_notes.go | TestCreditNoteLifecycle | 🔒 401 |
| `POST /api/v1/billing/credit_notes/{id}/delete` | DeleteCreditNote | credit_notes.go | TestCreditNoteOperations | 🔒 401 |
| `POST /api/v1/billing/credit_notes/{id}/delete_line_items` | DeleteCreditNoteLineItems | credit_notes.go | TestCreditNoteOperations | 🔒 401 |
| `POST /api/v1/billing/credit_notes/{id}/finalize` | FinalizeCreditNote | credit_notes.go | TestCreditNoteOperations | 🔒 401 |
| `GET /api/v1/billing/credit_notes/{id}` | GetCreditNote | credit_notes.go | TestCreditNoteLifecycle | 🔒 401 |
| `GET /api/v1/billing/credit_notes/{id}/line_items/{id}` | GetCreditNoteLineItem | credit_notes.go | TestCreditNoteOperations | 🔒 401 |
| `GET /api/v1/billing/credit_notes/{id}/line_items` | ListCreditNoteLineItems | credit_notes.go | TestCreditNoteOperations | 🔒 401 |
| `GET /api/v1/billing/credit_notes` | ListCreditNotes | credit_notes.go | TestCreditNoteLifecycle | 🔒 401 |
| `POST /api/v1/billing/credit_notes/preview` | PreviewCreditNote | credit_notes.go | TestCreditNoteOperations | 🔒 401 |
| `POST /api/v1/billing/credit_notes/{id}/update` | UpdateCreditNote | credit_notes.go | TestCreditNoteLifecycle | 🔒 401 |
| `POST /api/v1/billing/credit_notes/{id}/update_line_items` | UpdateCreditNoteLineItems | credit_notes.go | TestCreditNoteOperations | 🔒 401 |
| `POST /api/v1/billing/credit_notes/{id}/void` | VoidCreditNote | credit_notes.go | TestCreditNoteOperations | 🔒 401 |
| `POST /api/v1/invoices/{id}/add_line_items` | AddInvoiceLineItems | invoice.go | TestInvoiceOperations | ✅ PASS |
| `POST /api/v1/invoices/create` | CreateInvoice | invoice.go | TestInvoiceLifecycle | ✅ PASS |
| `POST /api/v1/invoices/{id}/delete` | DeleteInvoice | invoice.go | TestInvoiceOperations | ✅ PASS |
| `POST /api/v1/invoices/{id}/delete_line_items` | DeleteInvoiceLineItems | invoice.go | TestInvoiceOperations | ✅ PASS |
| `POST /api/v1/invoices/{id}/finalize` | FinalizeInvoice | invoice.go | TestInvoiceOperations | ✅ PASS |
| `GET /api/v1/invoices/{id}` | GetInvoice | invoice.go | TestInvoiceLifecycle | ✅ PASS |
| `GET /api/v1/invoices/{id}/line_items/{id}` | GetInvoiceLineItem | invoice.go | TestInvoiceOperations | ✅ PASS |
| `GET /api/v1/invoices/{id}/line_items` | ListInvoiceLineItems | invoice.go | TestInvoiceOperations | ✅ PASS |
| `GET /api/v1/invoices` | ListInvoices | invoice.go | TestInvoiceLifecycle | ✅ PASS |
| `POST /api/v1/invoices/{id}/mark_as_paid` | MarkInvoiceAsPaid | invoice.go | TestInvoiceOperations | ✅ PASS |
| `POST /api/v1/invoices/preview` | PreviewInvoice | invoice.go | TestInvoiceOperations | 🔒 400 |
| `POST /api/v1/invoices/{id}/update` | UpdateInvoice | invoice.go | TestInvoiceLifecycle | ✅ PASS |
| `POST /api/v1/invoices/{id}/update_line_items` | UpdateInvoiceLineItems | invoice.go | TestInvoiceOperations | ✅ PASS |
| `POST /api/v1/invoices/{id}/void` | VoidInvoice | invoice.go | TestInvoiceOperations | 🔒 400 |
| `POST /api/v1/meters/{id}/archive` | ArchiveMeter | meters.go | TestMeterOperations | 🔒 401 |
| `POST /api/v1/meters/create` | CreateMeter | meters.go | TestMeterLifecycle | 🔒 401 |
| `GET /api/v1/meters/{id}` | GetMeter | meters.go | TestMeterLifecycle | 🔒 401 |
| `GET /api/v1/meters/{id}/summaries` | GetMeterSummaries | meters.go | TestMeterOperations | 🔒 401 |
| `GET /api/v1/meters` | ListMeters | meters.go | TestMeterLifecycle | 🔒 401 |
| `POST /api/v1/meters/{id}/restore` | RestoreMeter | meters.go | TestMeterOperations | 🔒 401 |
| `POST /api/v1/meters/{id}/update` | UpdateMeter | meters.go | TestMeterLifecycle | 🔒 401 |
| `POST /api/v1/payment_sources/create` | CreatePaymentSource | payment_sources.go | TestPaymentSourceLifecycle | 🔒 400 |
| `GET /api/v1/payment_sources/{id}` | GetPaymentSource | payment_sources.go | TestPaymentSourceLifecycle | 🔒 400 |
| `GET /api/v1/payment_sources` | ListPaymentSources | payment_sources.go | TestPaymentSourceLifecycle | 🔒 400 |
| `POST /api/v1/prices/create` | CreatePrice | prices.go | TestPriceLifecycle | ✅ PASS |
| `GET /api/v1/prices/{id}` | GetPrice | prices.go | TestPriceLifecycle | ✅ PASS |
| `GET /api/v1/prices` | ListPrices | prices.go | TestPriceLifecycle | ✅ PASS |
| `POST /api/v1/prices/{id}/update` | UpdatePrice | prices.go | TestPriceLifecycle | ✅ PASS |
| `POST /api/v1/products/create` | CreateProduct | product.go | TestProductLifecycle | ✅ PASS |
| `GET /api/v1/products/{id}` | GetProduct | product.go | TestProductLifecycle | ✅ PASS |
| `GET /api/v1/products` | ListProducts | product.go | TestProductLifecycle | ✅ PASS |
| `POST /api/v1/products/{id}/update` | UpdateProduct | product.go | TestProductLifecycle | ✅ PASS |
| `POST /api/v1/subscriptions/{id}/cancel` | CancelSubscription | subscription.go | TestSubscriptionLifecycle | 🔒 400 |
| `POST /api/v1/subscriptions/create` | CreateSubscription | subscription.go | TestSubscriptionLifecycle | ✅ PASS |
| `GET /api/v1/subscriptions/{id}` | GetSubscription | subscription.go | TestSubscriptionLifecycle | ✅ PASS |
| `GET /api/v1/subscriptions/{id}/items/{id}` | GetSubscriptionItem | subscription.go | TestSubscriptionLifecycle | ✅ PASS |
| `GET /api/v1/subscriptions/{id}/items` | ListSubscriptionItems | subscription.go | TestSubscriptionLifecycle | ✅ PASS |
| `GET /api/v1/subscriptions` | ListSubscriptions | subscription.go | TestSubscriptionLifecycle | ✅ PASS |
| `POST /api/v1/subscriptions/{id}/update` | UpdateSubscription | subscription.go | TestSubscriptionLifecycle | ✅ PASS |
| `POST /api/v1/usage_events/batch_ingest` | BatchIngestUsageEvents | usage_event.go | TestUsageEventLifecycle | 🔒 401 |
| `POST /api/v1/usage_events/ingest` | IngestUsageEvent | usage_event.go | TestUsageEventLifecycle | 🔒 401 |
| `POST /api/v1/usage_events/void` | VoidUsageEvent | usage_event.go | TestUsageEventLifecycle | 🔒 401 |
---

## Finance (finance)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/finance/financial_reports/create` | CreateFinancialReport | financial_reports.go | TestFinancialReportsLifecycle | ✅ PASS |
| `GET /api/v1/finance/financial_reports/{id}` | GetFinancialReport | financial_reports.go | TestFinancialReportsLifecycle | ✅ PASS |
| `GET /api/v1/finance/financial_reports/{id}/content` | GetFinancialReportContent | financial_reports.go | TestFinancialReportsLifecycle | ✅ PASS |
| `GET /api/v1/finance/financial_reports` | ListFinancialReports | financial_reports.go | TestFinancialReportsLifecycle | ✅ PASS |
| `GET /api/v1/financial_transactions/{id}` | GetFinancialTransaction | financial_transaction.go | TestListFinancialTransactions | ✅ PASS |
| `GET /api/v1/financial_transactions` | ListFinancialTransactions | financial_transaction.go | TestListFinancialTransactions | ✅ PASS |
| `GET /api/v1/pa/financial/settlements/{id}` | GetSettlement | settlement.go | TestSettlementLifecycle | ✅ PASS |
| `GET /api/v1/pa/financial/settlements` | ListSettlements | settlement.go | TestSettlementLifecycle | ✅ PASS |
| `GET /api/v1/pa/financial/settlements/{id}/report` | GetSettlementReport | settlement.go | TestSettlementLifecycle | ✅ PASS |
---

## Scale (scale)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/accounts/create` | CreateAccount | account.go | scale_test.go | 🔒 400 |
| `GET /api/v1/accounts/{id}` | GetAccount | account.go | scale_test.go | ✅ PASS |
| `GET /api/v1/accounts` | ListAccounts | account.go | scale_test.go | ✅ PASS |
| `POST /api/v1/charges/create` | CreateCharge | charges.go | TestChargeLifecycle | 🔒 403 |
| `GET /api/v1/charges/{id}` | GetCharge | charges.go | TestChargeLifecycle | ✅ PASS |
| `GET /api/v1/charges` | ListCharges | charges.go | TestChargeLifecycle | ✅ PASS |
| `POST /api/v1/connected_account_transfers/create` | CreateConnectedAccountTransfer | connected_account_transfers.go | TestConnectedAccountTransferLifecycle | 🔒 403 |
| `GET /api/v1/connected_account_transfers/{id}` | GetConnectedAccountTransfer | connected_account_transfers.go | TestConnectedAccountTransferLifecycle | ✅ PASS |
| `GET /api/v1/connected_account_transfers` | ListConnectedAccountTransfers | connected_account_transfers.go | TestConnectedAccountTransferLifecycle | ✅ PASS |
| `POST /api/v1/hosted_flows/{id}/authorize` | AuthorizeHostedFlow | hosted_flows.go | TestHostedFlowLifecycle | ✅ PASS |
| `POST /api/v1/hosted_flows/create` | CreateHostedFlow | hosted_flows.go | TestHostedFlowLifecycle | 🔒 400 |
| `GET /api/v1/hosted_flows/{id}` | GetHostedFlow | hosted_flows.go | TestHostedFlowLifecycle | ✅ PASS |
| `POST /api/v1/accounts/invitation_links/create` | CreateInvitationLink | invitation_links.go | TestInvitationLinkLifecycle | 🔒 400 |
| `GET /api/v1/accounts/invitation_links/{id}` | GetInvitationLink | invitation_links.go | TestInvitationLinkLifecycle | ✅ PASS |
| `POST /api/v1/platform_liquidity_programs/{id}/deposit` | DepositFunds | platform_liquidity_programs.go | TestProgramLifecycle | ✅ PASS |
| `GET /api/v1/platform_liquidity_programs/{id}` | GetProgram | platform_liquidity_programs.go | TestProgramLifecycle | ✅ PASS |
| `GET /api/v1/platform_liquidity_programs/{id}/program_spending_accounts` | ListProgramSpendingAccounts | platform_liquidity_programs.go | TestProgramLifecycle | ✅ PASS |
| `GET /api/v1/platform_liquidity_programs/{id}/transactions` | ListProgramTransactions | platform_liquidity_programs.go | TestProgramLifecycle | ✅ PASS |
| `POST /api/v1/platform_liquidity_programs/{id}/withdraw` | WithdrawFunds | platform_liquidity_programs.go | TestProgramLifecycle | ✅ PASS |
| `POST /api/v1/platform_reports/create` | CreatePlatformReport | platform_reports.go | TestPlatformReportLifecycle | 🔒 400 |
| `GET /api/v1/platform_reports/{id}` | GetPlatformReport | platform_reports.go | TestPlatformReportLifecycle | ✅ PASS |
| `GET /api/v1/psp_settlement_deposits` | ListPSPSettlementDeposits | psp_settlement_deposits.go | TestPSPSettlementDepositLifecycle | ✅ PASS |
| `POST /api/v1/psp_settlement_intents/create` | CreatePSPSettlementIntent | psp_settlement_intents.go | TestPSPSettlementIntentLifecycle | 🔒 400 |
| `GET /api/v1/psp_settlement_intents/{id}` | GetPSPSettlementIntent | psp_settlement_intents.go | TestPSPSettlementIntentLifecycle | ✅ PASS |
| `GET /api/v1/psp_settlement_intents` | ListPSPSettlementIntents | psp_settlement_intents.go | TestPSPSettlementIntentLifecycle | ✅ PASS |
| `POST /api/v1/psp_settlement_splits/{id}/cancel` | CancelPSPSettlementSplit | psp_settlement_splits.go | TestPSPSettlementSplitLifecycle | ✅ PASS |
| `GET /api/v1/psp_settlement_splits/{id}` | GetPSPSettlementSplit | psp_settlement_splits.go | TestPSPSettlementSplitLifecycle | ✅ PASS |
| `GET /api/v1/psp_settlement_splits` | ListPSPSettlementSplits | psp_settlement_splits.go | TestPSPSettlementSplitLifecycle | ✅ PASS |
| `POST /api/v1/psp_settlement_splits/{id}/release` | ReleasePSPSettlementSplit | psp_settlement_splits.go | TestPSPSettlementSplitLifecycle | ✅ PASS |
| `POST /api/v1/psp_settlement_intents/{id}/split` | SplitPSPSettlementIntent | psp_settlement_splits.go | TestPSPSettlementSplitLifecycle | ✅ PASS |
---

## Simulation (simulation)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/simulation/account/amendments/{id}/approve` | SimulateAccountAmendmentApprove | accounts.go | TestSimulateAccountAmendmentApprove_InvalidID | 🔒 400 |
| `POST /api/v1/simulation/account/amendments/{id}/reject` | SimulateAccountAmendmentReject | accounts.go | TestSimulateAccountAmendmentReject_InvalidID | 🔒 400 |
| `POST /api/v1/simulation/accounts/{id}/update_status` | SimulateAccountUpdateStatus | accounts.go | TestSimulateAccountUpdateStatus | 🔒 401 |
| `POST /api/v1/simulation/billing/payment_sources/{id}/fail_next_autocharge` | SimulateBillingFailNextAutocharge | billing.go | TestSimulateBillingFailNextAutocharge | 🔒 400 |
| `POST /api/v1/simulation/deposits/{id}/reject` | SimulateDirectDebitReject | deposits.go | TestSimulateDirectDebitOperations | ✅ PASS |
| `POST /api/v1/simulation/deposits/{id}/reverse` | SimulateDirectDebitReverse | deposits.go | TestSimulateDirectDebitOperations | ✅ PASS |
| `POST /api/v1/simulation/deposits/{id}/settle` | SimulateDirectDebitSettle | deposits.go | TestSimulateDirectDebitOperations | ✅ PASS |
| `POST /api/v1/simulation/deposit/create` | SimulateGlobalAccountDeposit | deposits.go | TestSimulateGlobalAccountDeposit | ✅ PASS |
| `POST /api/v1/simulation/issuing/cardholders/{id}/pass_review` | SimulateIssuingCardholderPassReview | issuing.go | TestSimulateIssuingCardholderPassReview | 🔒 400 |
| `POST /api/v1/simulation/issuing/threeds/notify` | SimulateIssuingThreedsNotify | issuing.go | TestSimulateIssuingThreedsNotify | 🔒 500 |
| `POST /api/v1/simulation/issuing/{id}/capture` | SimulateIssuingTransactionCapture | issuing.go | TestSimulateIssuingTransactionLifecycle | ✅ PASS |
| `POST /api/v1/simulation/issuing/create` | SimulateIssuingTransactionCreate | issuing.go | TestSimulateIssuingTransactionLifecycle | 🔒 500 |
| `POST /api/v1/simulation/issuing/refund` | SimulateIssuingTransactionRefund | issuing.go | TestSimulateIssuingTransactionLifecycle | ✅ PASS |
| `POST /api/v1/simulation/issuing/{id}/reverse` | SimulateIssuingTransactionReverse | issuing.go | TestSimulateIssuingTransactionLifecycle | ✅ PASS |
| `POST /api/v1/simulation/linked_accounts/{id}/fail_microdeposits` | SimulateLinkedAccountFailMicrodeposits | linked_accounts.go | TestSimulateLinkedAccountOperations_InvalidID | 🔒 400 |
| `POST /api/v1/simulation/linked_accounts/{id}/mandate/accept` | SimulateLinkedAccountMandateAccept | linked_accounts.go | TestSimulateLinkedAccountOperations_InvalidID | 🔒 400 |
| `POST /api/v1/simulation/linked_accounts/{id}/mandate/cancel` | SimulateLinkedAccountMandateCancel | linked_accounts.go | TestSimulateLinkedAccountOperations_InvalidID | 🔒 400 |
| `POST /api/v1/simulation/linked_accounts/{id}/mandate/reject` | SimulateLinkedAccountMandateReject | linked_accounts.go | TestSimulateLinkedAccountOperations_InvalidID | 🔒 400 |
| `POST /api/v1/simulation/pa/payment_disputes/create` | SimulatePaymentDisputeCreate | payments.go | TestSimulatePaymentDisputeLifecycle | 🔒 500 |
| `POST /api/v1/simulation/pa/payment_disputes/{id}/escalate` | SimulatePaymentDisputeEscalate | payments.go | TestSimulatePaymentDisputeLifecycle | ✅ PASS |
| `POST /api/v1/simulation/pa/payment_disputes/{id}/resolve` | SimulatePaymentDisputeResolve | payments.go | TestSimulatePaymentDisputeLifecycle | ✅ PASS |
| `POST /api/v1/simulation/pa/shopper_actions/{id}` | SimulateShopperAction | payments.go | TestSimulateShopperAction | ✅ PASS |
| `POST /api/v1/simulation/payments/{id}/transition` | SimulatePayoutPaymentTransition | payouts.go | TestSimulatePayoutPaymentTransition_InvalidID | 🔒 400 |
| `POST /api/v1/simulation/rfis/{id}/close` | SimulateRFIClose | rfis.go | TestSimulateRFIOperations_InvalidRequests | 🔒 404 |
| `POST /api/v1/simulation/rfis/create` | SimulateRFICreate | rfis.go | TestSimulateRFIOperations_InvalidRequests | 🔒 400 |
| `POST /api/v1/simulation/rfis/{id}/follow_up` | SimulateRFIFollowUp | rfis.go | TestSimulateRFIOperations_InvalidRequests | 🔒 400 |
| `POST /api/v1/simulation/pa/pos/terminals/confirm_payment_intent` | SimulateTerminalConfirmPaymentIntent | terminals.go | TestSimulateTerminalLifecycle | 🔒 400 |
| `POST /api/v1/simulation/pa/pos/terminals/generate_activation_code` | SimulateTerminalGenerateActivationCode | terminals.go | TestSimulateTerminalLifecycle | 🔒 400 |
| `GET /api/v1/simulation/pa/pos/terminals/payment_scenarios` | SimulateTerminalPaymentScenarios | terminals.go | TestSimulateTerminalLifecycle | ✅ PASS |
| `POST /api/v1/simulation/pa/pos/terminals/turn_off` | SimulateTerminalTurnOff | terminals.go | TestSimulateTerminalLifecycle | 🔒 400 |
| `POST /api/v1/simulation/pa/pos/terminals/turn_on` | SimulateTerminalTurnOn | terminals.go | TestSimulateTerminalLifecycle | 🔒 400 |
| `POST /api/v1/simulation/transfers/{id}/transition` | SimulateTransferTransition | transfers.go | TestSimulateTransferTransition | 🔒 500 |
---

## Spend (spend)

> Bills/Expenses 不支持 Create/Update（官方 API 无此端点）。所有 Spend API 在 Sandbox 返回 401（权限不足）。

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `GET /api/v1/spend/bills/{id}` | GetBill | bill.go | TestBillLifecycle | 🔒 401 |
| `GET /api/v1/spend/bills` | ListBills | bill.go | TestBillLifecycle | 🔒 401 |
| `POST /api/v1/spend/bills/{id}/sync` | SyncBill | bill.go | TestBillLifecycle | 🔒 401 |
| `GET /api/v1/spend/expenses/{id}` | GetExpense | expense.go | TestExpenseLifecycle | 🔒 401 |
| `GET /api/v1/spend/expenses` | ListExpenses | expense.go | TestExpenseLifecycle | 🔒 401 |
| `POST /api/v1/spend/expenses/{id}/sync` | SyncExpense | expense.go | TestExpenseLifecycle | 🔒 401 |
| `POST /api/v1/spend/purchase_orders/create` | CreatePurchaseOrder | purchase_order.go | TestPurchaseOrderLifecycle | 🔒 401 |
| `GET /api/v1/spend/purchase_orders/{id}` | GetPurchaseOrder | purchase_order.go | TestPurchaseOrderLifecycle | 🔒 401 |
| `GET /api/v1/spend/purchase_orders` | ListPurchaseOrders | purchase_order.go | TestPurchaseOrderLifecycle | 🔒 401 |
| `POST /api/v1/spend/purchase_orders/{id}/sync` | SyncPurchaseOrder | purchase_order.go | TestPurchaseOrderLifecycle | 🔒 401 |
| `GET /api/v1/spend/reimbursement_reports/{id}` | GetReimbursementReport | reimbursement.go | TestReimbursementLifecycle | 🔒 401 |
| `GET /api/v1/spend/reimbursement_reports` | ListReimbursementReports | reimbursement.go | TestReimbursementLifecycle | 🔒 401 |
| `POST /api/v1/spend/reimbursement_reports/{id}/mark_as_paid` | MarkReimbursementReportAsPaid | reimbursement.go | TestReimbursementLifecycle | 🔒 401 |
| `POST /api/v1/spend/vendors/create` | CreateVendor | vendor.go | TestVendorLifecycle | 🔒 401 |
| `GET /api/v1/spend/vendors/{id}` | GetVendor | vendor.go | TestVendorLifecycle | 🔒 401 |
| `GET /api/v1/spend/vendors` | ListVendors | vendor.go | TestVendorLifecycle | 🔒 401 |
| `POST /api/v1/spend/vendors/{id}/sync` | SyncVendor | vendor.go | TestVendorLifecycle | 🔒 401 |
---

## Risk (risk)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/risk/fraud_feedback/create` | CreateFraudFeedback | fraud_feedback.go | TestFraudFeedbackLifecycle | ✅ PASS |
| `GET /api/v1/risk/fraud_feedback/{id}` | GetFraudFeedback | fraud_feedback.go | TestFraudFeedbackLifecycle | ✅ PASS |
| `GET /api/v1/risk/fraud_feedback` | ListFraudFeedback | fraud_feedback.go | TestFraudFeedbackLifecycle | ✅ PASS |
| `POST /api/v1/rfis/create` | CreateRFI | rfi.go | TestRFILifecycle | 🔒 405 |
| `GET /api/v1/rfis/{id}` | GetRFI | rfi.go | TestRFILifecycle | ✅ PASS |
| `GET /api/v1/rfis` | ListRFIs | rfi.go | TestRFILifecycle | ✅ PASS |
| `POST /api/v1/rfis/{id}/respond` | RespondRFI | rfi.go | TestRFILifecycle | ✅ PASS |
| `POST /api/v1/sellers/create` | CreateSeller | seller.go | TestSellerLifecycle | 🔒 403 |
| `GET /api/v1/sellers/{id}` | GetSeller | seller.go | TestSellerLifecycle | ✅ PASS |
| `GET /api/v1/sellers` | ListSellers | seller.go | TestSellerLifecycle | ✅ PASS |
| `POST /api/v1/risk/watchlist/create` | CreateWatchlistEntry | watchlist.go | TestWatchlistLifecycle | ✅ PASS |
| `GET /api/v1/risk/watchlist/{id}` | GetWatchlistEntry | watchlist.go | TestWatchlistLifecycle | ✅ PASS |
| `GET /api/v1/risk/watchlist` | ListWatchlistEntries | watchlist.go | TestWatchlistLifecycle | ✅ PASS |
| `POST /api/v1/risk/watchlist/{id}/update` | UpdateWatchlistEntry | watchlist.go | TestWatchlistLifecycle | ✅ PASS |
---

## Supporting (supporting)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `GET /api/v1/ecosystem/connected_stores/{id}` | GetConnectedStore | connected_store.go | TestConnectedStoreLifecycle | ✅ PASS |
| `GET /api/v1/ecosystem/connected_stores` | ListConnectedStores | connected_store.go | TestConnectedStoreLifecycle | ✅ PASS |
| `POST /api/v1/files/download_links` | GetDownloadLinks | file.go | TestFileLifecycle | ✅ PASS |
| `POST /api/v1/files/upload` | UploadFile | file.go | TestFileLifecycle | ✅ PASS |
---

## Webhook (webhook)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| ``webhook.VerifySignature(...)`` | — (本地计算) | verify.go | webhook_test.go | ✅ PASS |
| `POST /api/v1/webhooks/create` | CreateWebhook | endpoint.go | TestWebhookLifecycle | 🔒 401 |
| `POST /api/v1/webhooks/{id}/delete` | DeleteWebhook | endpoint.go | TestWebhookLifecycle | 🔒 401 |
| `GET /api/v1/webhooks/{id}` | GetWebhook | endpoint.go | TestWebhookLifecycle | 🔒 401 |
| `GET /api/v1/webhooks` | ListWebhooks | endpoint.go | TestWebhookLifecycle | 🔒 401 |
| `POST /api/v1/webhooks/{id}/update` | UpdateWebhook | endpoint.go | TestWebhookLifecycle | 🔒 401 |
---

## Account Capability (capability)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/account_capabilities/{id}/enable` | EnableAccountCapability | account_capability.go | TestEnableAccountCapability | ✅ PASS |
| `GET /api/v1/account_capabilities/{id}` | GetAccountCapability | account_capability.go | TestGetAccountCapability | ✅ PASS |
| `GET /api/v1/account_capabilities/funding_limits` | GetFundingLimits | account_capability.go | TestGetFundingLimits | ✅ PASS |
---

## Confirmation Letter (confirmation)

| 端点 | 函数 | Go 文件 | 测试 | 状态 |
|------|------|------|------|------|
| `POST /api/v1/confirmation_letters/create` | CreateConfirmationLetter | confirmation_letter.go | TestConfirmationLetterCreate | ✅ PASS |
---

