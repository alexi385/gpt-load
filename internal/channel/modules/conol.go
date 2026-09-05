package modules

import (
	"gpt-load/internal/channel/spec"
	"gpt-load/internal/execution"
	"gpt-load/internal/protocol"
)

// Conol declares the API key-backed Conol.ai gateway channel.
func Conol() spec.Module {
	return spec.Module{
		Definition: spec.Definition{
			ID:          spec.Conol,
			Name:        "Conol",
			Mark:        "CO",
			Icon:        "conol",
			SearchTerms: []string{"conol", "gateway", "proxy", "deepseek", "claude", "gemini"},
			Description: "Conol.ai unified AI gateway - supports DeepSeek, GPT, Claude, Gemini, Qwen, GLM, Kimi",
			Connection: spec.Connection{
				Type:            spec.ConnectionAPIKey,
				CredentialInput: "batch_text",
			},
			Params: []spec.Field{},
			Credentials: []spec.Field{
				{
					Key:        "token",
					Label:      "Session Token",
					InputKind:  spec.InputSecret,
					Required:   true,
					Sensitive:  true,
					Normalizer: spec.NormalizeNonEmpty,
				},
				{
					Key:        "passkey",
					Label:      "Passkey",
					InputKind:  spec.InputSecret,
					Required:   true,
					Sensitive:  true,
					Normalizer: spec.NormalizeNonEmpty,
				},
				{
					Key:        "email",
					Label:      "Email",
					InputKind:  spec.InputText,
					Required:   true,
					Sensitive:  false,
					Normalizer: spec.NormalizeNonEmpty,
				},
			},
			Provider: spec.ProviderBinding{
				ProviderKind:      spec.ProviderOpenAICompatible,
				CatalogProviderID: "",
				FixedBaseURL:      "https://conol.ai",
				EndpointPolicy:    spec.EndpointFixedWithOverride,
			},
			Routes: []spec.Route{
				spec.NewRoute(protocol.OpenAICompletions, execution.OperationChatCompletion, execution.RouteNative),
				spec.NewRoute(protocol.OpenAICompletions, execution.OperationListModels, execution.RouteNative),
				spec.NewRoute(protocol.OpenAICompletions, execution.OperationProbe, execution.RouteNative),
				spec.NewResponsesCreateRoute(execution.RouteConverted, spec.ResponsesStoreHandlingStateless),
				spec.NewRoute(protocol.OpenAIResponses, execution.OperationProbe, execution.RouteConverted),
				spec.NewRoute(protocol.Anthropic, execution.OperationChatCompletion, execution.RouteConverted),
				spec.NewRoute(protocol.Anthropic, execution.OperationListModels, execution.RouteConverted),
				spec.NewRoute(protocol.Gemini, execution.OperationChatCompletion, execution.RouteConverted),
				spec.NewRoute(protocol.Gemini, execution.OperationListModels, execution.RouteConverted),
			},
		},
	}
}
