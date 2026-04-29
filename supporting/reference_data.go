package supporting

// --- Reference Data ---
//
// 旧版 CreateReferenceData / GetReferenceData / UpdateReferenceData / ListReferenceData 已移除。
// 对应的端点 POST/GET /api/v1/reference_data/* 在官方 API 文档中不存在。
//
// 官方 Reference Data API 位于 /api/v1/reference/ 下，包含以下端点：
//   GET  /api/v1/reference/industry_categories       — 行业类别
//   GET  /api/v1/reference/invalid_conversion_dates   — 无效兑换日期
//   GET  /api/v1/reference/order_items/expiring_quota  — 即将过期的限额
//   GET  /api/v1/reference/order_items/quota          — 转账/存款限额
//   POST /api/v1/reference/order_items/upload         — 提高限额
//   GET  /api/v1/reference/supported_currencies       — 支持的货币
//
// 官方文档: https://www.airwallex.com/docs/api/supporting_services/reference_data.md
//
// 如需实现以上端点，请按 spend 包的模式逐端点获取文档并生成对应 struct 和 service 函数。
