# 項目技術規範 (Go 1.24 & Vue 3 Latest)

## 1. 核心環境 (Environment)
- **後端**: Go 1.24+ (使用最新標準庫特性，如 map 性能優化、泛型類型別名)。
- **前端**: Vue 3.5+ (最新穩定版)。
- **構建工具**: Vite 6+。
- **包管理**: Go Modules (後端), npm (前端)。

## 2. 後端規範 (Go 1.24 Standards)
- **代碼風格**: 遵循官方 `gofmt` 與 `goimports`。
- **語法特性**: 
  - 優先使用 Go 1.24 的泛型新特性（如 Generic Type Aliases）。
  - 對於高性能需求場景，考慮 Go 1.24 的內存分配與 map 遍歷優化。
- **錯誤處理**: 必須使用 `errors.Is` 和 `errors.As` 進行錯誤檢查，禁止使用 `errors.New(fmt.Sprintf(...))`，應使用 `fmt.Errorf("...: %w", err)`。
- **並發**: 優先使用 `context` 控制協程生命週期，避免 goroutine 洩漏。
- **API**: 默認提供 RESTful JSON 接口，結構體標籤使用小駝峰 (json:"userName")。

## 3. 前端規範 (Vue 3 Standards)
- **API 模式**: 必須使用 **Composition API** 搭配 **`<script setup>`** 語法。
- **語言**: 強制使用 **TypeScript**，禁止使用 `any`。
- **狀態管理**: 使用 **Pinia** (替代 Vuex)。
- **組件命名**: 文件名使用 PascalCase (例如 `UserCard.vue`)。
- **CSS**: 優先選用 Tailwind CSS 或 Scoped CSS，禁止全局污染。
- **邏輯抽離**: 複雜邏輯應封裝為 `useSomething.ts` 的 Composable 函數。

## 4. 項目目錄結構
Jules 在生成代碼時請遵循以下結構：
```text
.
├── server/             # Go 1.24 後端代碼
│   ├── main.go         # 入口文件
│   ├── api/            # 接口層
│   ├── internal/       # 內部業務邏輯
│   └── go.mod
├── web/                # Vue 3 前端代碼 (Vite)
│   ├── src/
│   │   ├── components/ # 組件
│   │   ├── composables/# 組合式函數
│   │   ├── stores/     # Pinia stores
│   │   └── App.vue
│   ├── package.json
│   └── vite.config.ts
└── agents.md           # 本指令文件
